CREATE TABLE IF NOT EXISTS person (
	id TEXT NOT NULL,
    first_name TEXT,
    last_name TEXT,
    gender TEXT,
	CONSTRAINT Person_ID_UN UNIQUE (id)
);

CREATE TABLE IF NOT EXISTS relationship (
	id TEXT NOT NULL,
	person_id TEXT NOT NULL,
	"type" TEXT,
	relationship_id TEXT,
	CONSTRAINT relationship_id_un UNIQUE (id)
);
