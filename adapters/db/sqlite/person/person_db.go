package dbperson

import (
	"context"
	"errors"
	"strings"

	commondb "github.com/darllantissei/genealogy-tree/adapters/db/sqlite/common"
	"github.com/darllantissei/genealogy-tree/application/model"
)

type PersonDB struct {
	commonDB commondb.CommonDB
}

func NewPersonDB(dbcommon commondb.CommonDB) *PersonDB {
	return &PersonDB{
		commonDB: dbcommon,
	}

}

func (db *PersonDB) Insert(ctx context.Context, prsn model.Person) (model.Person, error) {

	prsn.ID = db.commonDB.GetUUID()

	fields := []interface{}{
		prsn.ID,
	}

	fields = append(fields, db.getFields(prsn)...)

	result, err := db.commonDB.ExecStatement(dmlInsert, fields...)

	if err != nil {

		err = db.commonDB.TreatError(err, "fail to persist person on database")

		return model.Person{}, err
	}

	row, err := result.RowsAffected()

	if err != nil {

		err = db.commonDB.TreatError(err, "fail to get records affected in database")

		return model.Person{}, err
	}

	if row <= 0 {
		return model.Person{}, errors.New("the data of person wasn't persisted")
	}

	return prsn, nil
}
func (db *PersonDB) Update(ctx context.Context, prsn model.Person) error {

	if strings.EqualFold(prsn.ID, "") {
		return errors.New("the ID of person undefined to update")
	}

	fields := db.getFields(prsn)

	fields = append(fields, prsn.ID)

	result, err := db.commonDB.ExecStatement(dmlUpdate, fields...)

	if err != nil {

		err = db.commonDB.TreatError(err, "fail to update person on database")

		return err
	}

	row, err := result.RowsAffected()

	if err != nil {

		err = db.commonDB.TreatError(err, "fail to get records affected on update in database")

		return err
	}

	if row <= 0 {
		return errors.New("the data of person wasn't updated")
	}

	return nil
}
func (db *PersonDB) Get(ctx context.Context, prsn model.Person) (model.Person, error) {

	query := dqlSelect

	args := []interface{}{}

	if !strings.EqualFold(prsn.ID, "") {

		query += ` AND person.id = ? `

		args = append(args, prsn.ID)
	}

	if len(args) <= 0 {
		query += ` AND FALSE `
	}

	rows, err := db.commonDB.Query(query, args...)

	if err != nil {

		err = db.commonDB.TreatError(err, "Fail to get person in database")

		return model.Person{}, err
	}

	defer rows.Close()

	prsn = model.Person{}

	for rows.Next() {

		err := rows.Scan(
			&prsn.ID,
			&prsn.FirstName,
			&prsn.LastName,
			&prsn.Gender,
		)

		if err != nil {

			err = db.commonDB.TreatError(err, "Fail to mapping fields in the struct person")

			return model.Person{}, err
		}
	}

	return prsn, nil
}

func (db *PersonDB) List(ctx context.Context) ([]model.Person, error) {

	rows, err := db.commonDB.Query(dqlSelect)

	if err != nil {

		err = db.commonDB.TreatError(err, "Fail to get all person in database")

		return []model.Person{}, err
	}

	defer rows.Close()

	prsn := model.Person{}
	persons := []model.Person{}

	for rows.Next() {

		err := rows.Scan(
			&prsn.ID,
			&prsn.FirstName,
			&prsn.LastName,
			&prsn.Gender,
		)

		if err != nil {

			err = db.commonDB.TreatError(err, "Fail to mapping fields in the struct person when get all persons")

			return []model.Person{}, err
		}

		persons = append(persons, prsn)
	}

	return persons, nil
}
