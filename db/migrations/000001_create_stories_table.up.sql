CREATE TABLE IF NOT EXISTS stories 
(
    id          serial  PRIMARY KEY,
    title       text        NOT NULL,
    created_at  timestamptz NOT NULL DEFAULT NOW(),
    updated_at  timestamptz NOT NULL DEFAULT NOW(),
    is_deleted  boolean     NOT NULL DEFAULT FALSE
);