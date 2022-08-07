package dbrelationship

const dqlSelect = `
SELECT
	COALESCE(id, '') AS id,
	COALESCE(person_id, '') AS person_id,
	COALESCE("type", 'undefined') AS "type",
	COALESCE(relationship_id, '') AS relationship_id
FROM relationship
WHERE TRUE
`

const dmlInsertRelationship = `
INSERT INTO relationship(
	id, 
	person_id, 
	"type", 
	relationship_id
)VALUES(
	?, /* id, */
	?, /* person_id, */
	?, /* "type", */
	? /* relationship_id */
)
`

const dmlDeleteRelationship = `
DELETE FROM relationship
WHERE id = ? 
`
