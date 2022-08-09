package dbrelationship

import (
	"context"
	"errors"
	"strings"

	commondb "github.com/darllantissei/genealogy-tree/adapters/db/sqlite/common"
	"github.com/darllantissei/genealogy-tree/application/model"
)

type RelationshipDB struct {
	commonDB commondb.CommonDB
}

func NewRelationshipDB(dbcommon commondb.CommonDB) *RelationshipDB {
	return &RelationshipDB{
		commonDB: dbcommon,
	}

}

func (db *RelationshipDB) Record(ctx context.Context, rtshp model.Relationship) (model.Relationship, error) {

	rtshp.ID = db.commonDB.GetUUID()

	fields := []interface{}{
		rtshp.ID,
	}

	fields = append(fields, db.getFieldsRelationship(rtshp)...)

	result, err := db.commonDB.ExecStatement(dmlInsertRelationship, fields...)

	if err != nil {

		err = db.commonDB.TreatError(err, "fail to persist relationship on database")

		return model.Relationship{}, err
	}

	row, err := result.RowsAffected()

	if err != nil {

		err = db.commonDB.TreatError(err, "fail to get records affected in database when record relationship")

		return model.Relationship{}, err
	}

	if row <= 0 {
		return model.Relationship{}, errors.New("the data of relationship wasn't persisted")
	}

	err = db.insertMembers(ctx, rtshp)

	if err != nil {
		db.delete(ctx, rtshp)

		return model.Relationship{}, err
	}

	return rtshp, nil
}
func (db *RelationshipDB) Get(ctx context.Context, rtshp model.Relationship) (model.Relationship, error) {
	query := dqlSelect

	args := []interface{}{}

	if !strings.EqualFold(rtshp.ID, "") {

		query += ` AND relationship.id = ? `

		args = append(args, rtshp.ID)
	}

	if !strings.EqualFold(rtshp.PersonID, "") {

		query += ` AND relationship.person_id = ? `

		query += ` AND relationship.relationship_id IS NULL `

		args = append(args, rtshp.PersonID)
	}

	if len(args) <= 0 {
		query += ` AND FALSE `
	}

	rows, err := db.commonDB.Query(query, args...)

	if err != nil {

		err = db.commonDB.TreatError(err, "Fail to get relationship in database")

		return model.Relationship{}, err
	}

	defer rows.Close()

	rtshp = model.Relationship{}

	ghostFieldType := ""
	ghostFieldRtshpID := ""

	for rows.Next() {

		err := rows.Scan(
			&rtshp.ID,
			&rtshp.PersonID,
			&ghostFieldType,
			&ghostFieldRtshpID,
		)

		if err != nil {

			err = db.commonDB.TreatError(err, "Fail to mapping fields in the struct relationship")

			return model.Relationship{}, err
		}
	}

	members, _ := db.getMembers(ctx, rtshp)

	if len(members) > 0 {
		rtshp.Members = append(rtshp.Members, members...)
	}

	return rtshp, nil
}

func (db *RelationshipDB) getMembers(ctx context.Context, rtshp model.Relationship) ([]model.RelationshipMember, error) {
	query := dqlSelect

	args := []interface{}{}

	if !strings.EqualFold(rtshp.ID, "") {

		query += ` AND relationship.relationship_id = ? `

		args = append(args, rtshp.ID)
	}

	if len(args) <= 0 {
		query += ` AND FALSE `
	}

	rows, err := db.commonDB.Query(query, args...)

	if err != nil {

		err = db.commonDB.TreatError(err, "Fail to get members of relationship in database")

		return []model.RelationshipMember{}, err
	}

	defer rows.Close()

	members := []model.RelationshipMember{}
	member := model.RelationshipMember{}

	ghostFieldRtshpID := ""

	for rows.Next() {

		err := rows.Scan(
			&member.RelationshipID,
			&member.PersonID,
			&member.Type,
			&ghostFieldRtshpID,
		)

		if err != nil {

			err = db.commonDB.TreatError(err, "Fail to mapping fields in the struct members of relationship")

			return []model.RelationshipMember{}, err
		}

		members = append(members, member)
	}

	return members, nil
}

func (db *RelationshipDB) insertMembers(ctx context.Context, rtshp model.Relationship) error {

	if strings.EqualFold(rtshp.ID, "") {
		return nil
	}

	for _, member := range rtshp.Members {

		member.RelationshipID = rtshp.ID

		fields := db.getFieldsMemberRelationship(member)

		result, err := db.commonDB.ExecStatement(dmlInsertRelationship, fields...)

		if err != nil {

			err = db.commonDB.TreatError(err, "fail to persist members of relationship on database")

			return err
		}

		row, err := result.RowsAffected()

		if err != nil {

			err = db.commonDB.TreatError(err, "fail to get records affected in database when record members of relationship")

			return err
		}

		if row <= 0 {
			return errors.New("the data of members of relationship wasn't persisted")
		}
	}

	return nil
}

func (db *RelationshipDB) delete(ctx context.Context, rtshp model.Relationship) error {

	if strings.EqualFold(rtshp.ID, "") {
		return nil
	}

	args := []interface{}{
		rtshp.ID,
	}

	result, err := db.commonDB.ExecStatement(dmlDeleteRelationship, args...)

	if err != nil {

		err = db.commonDB.TreatError(err, "fail to delete relationship on database")

		return err
	}

	row, err := result.RowsAffected()

	if err != nil {

		err = db.commonDB.TreatError(err, "fail to get records affected in database when delete relationship")

		return err
	}

	if row <= 0 {
		return errors.New("the relationship wasn't deleted")
	}

	return nil
}

func (db *RelationshipDB) GetKinship(ctx context.Context, member model.RelationshipMember) ([]model.RelationshipMember, error) {
	query := dqlSelectKinship

	args := []interface{}{}

	if !strings.EqualFold(member.PersonID, "") {

		args = append(args, member.PersonID)
	}

	if len(args) <= 0 {
		query += ` AND FALSE `
	}

	rows, err := db.commonDB.Query(query, args...)

	if err != nil {

		err = db.commonDB.TreatError(err, "Fail to get member relationship in database")

		return []model.RelationshipMember{}, err
	}

	defer rows.Close()

	member = model.RelationshipMember{}

	members := []model.RelationshipMember{}

	for rows.Next() {

		err := rows.Scan(
			&member.PersonID,
			&member.Type,
			&member.RelationshipID,
		)

		if err != nil {

			err = db.commonDB.TreatError(err, "Fail to mapping fields in the struct member relationship")

			return []model.RelationshipMember{}, err
		}

		members = append(members, member)
	}

	return members, nil
}
