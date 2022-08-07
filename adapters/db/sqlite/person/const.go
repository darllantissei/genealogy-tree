package dbperson

const dmlInsert = `
INSERT INTO person (
	id, 
	first_name, 
	last_name, 
	gender
) VALUES(
	?, /* id,  */
	?, /* first_name,  */
	?, /* last_name,  */
	? /* gender */
)
`

const dmlUpdate = `
UPDATE person SET 
	first_name=?,
	last_name=?, 
	gender=?
WHERE id=?
`

const dqlSelect = `
SELECT 
	COALESCE(id, '') AS id, 
	COALESCE(first_name, '') AS first_name, 
	COALESCE(last_name, '') AS last_name, 
	COALESCE(gender, '') AS gender
FROM person
WHERE TRUE
`
