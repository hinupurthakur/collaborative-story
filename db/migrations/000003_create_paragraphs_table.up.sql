CREATE TABLE IF NOT EXISTS paragraphs 
(
    id          serial  PRIMARY KEY,
    sentences   text[]        NOT NULL,
    story_id    integer REFERENCES stories (id),
    created_at  timestamptz NOT NULL DEFAULT NOW(),
    updated_at  timestamptz NOT NULL DEFAULT NOW(),
    is_deleted  boolean     NOT NULL DEFAULT FALSE
);