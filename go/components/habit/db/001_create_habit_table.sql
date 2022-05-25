CREATE TYPE activity_type AS ENUM ('behaviour', 'task');

CREATE TABLE activity (
  id UUID PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  created_time TIMESTAMP NOT NULL,
  modified_time TIMESTAMP NOT NULL,
  type activity_type NOT NULL,
  is_archived boolean NOT NULL DEFAULT FALSE,
  is_deleted boolean NOT NULL DEFAULT FALSE,
  is_paused boolean NOT NULL DEFAULT FALSE
);

---- create above / drop below ----
DROP TYPE activity_type;
DROP TABLE activity;
