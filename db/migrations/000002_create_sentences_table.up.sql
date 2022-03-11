CREATE TABLE IF NOT EXISTS sentences 
(
    id              serial  PRIMARY KEY,
    story_id    integer REFERENCES stories (id),
    sentence        text        NOT NULL,
    created_at      timestamptz NOT NULL DEFAULT NOW(),
    updated_at      timestamptz NOT NULL DEFAULT NOW(),
    is_deleted      boolean     NOT NULL DEFAULT FALSE
)