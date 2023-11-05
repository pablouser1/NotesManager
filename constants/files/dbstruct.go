package files

const DB_STRUCT string = `
-- subjects definition
CREATE TABLE subjects (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	slug TEXT
);


-- units definition
CREATE TABLE units (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	num INTEGER,
	name TEXT,
	subject_id INTEGER,
	CONSTRAINT units_FK FOREIGN KEY (subject_id) REFERENCES subjects(id) ON DELETE CASCADE
);

-- variants definition
CREATE TABLE variants (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	slug TEXT
);
`
