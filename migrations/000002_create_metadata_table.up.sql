CREATE TABLE IF NOT EXISTS metadata (
   id serial PRIMARY KEY,
   title VARCHAR(300),
   last_modified timestamp without time zone,
   version VARCHAR(20),
   oscal_version VARCHAR(20)
);
