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

const dqlSelectKinship = `
SELECT 
	COALESCE(relationship_main.person_id, '') AS person_id,
    COALESCE(relationship_main.type, relationship_member.type, 'undefined') AS "type", 
	COALESCE(relationship_member.relationship_id, '') AS relationship_id
FROM relationship AS relationship_main
INNER JOIN relationship relationship_member 
ON relationship_member.relationship_id = relationship_main.id
WHERE TRUE
`