package dbrelationship

import (
	"database/sql"

	"github.com/darllantissei/genealogy-tree/application/model"
)

func (db RelationshipDB) getFieldsRelationship(rtshp model.Relationship) []interface{} {
	return []interface{}{
		rtshp.PersonID,
		sql.NullString{String: "", Valid: false},
		sql.NullString{String: "", Valid: false},
	}
}

func (db RelationshipDB) getFieldsMemberRelationship(member model.RelationshipMember) []interface{} {
	return []interface{}{
		db.commonDB.GetUUID(),
		member.PersonID,
		member.Type,
		member.RelationshipID,
	}
}
