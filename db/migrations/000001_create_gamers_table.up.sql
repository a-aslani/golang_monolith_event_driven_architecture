CREATE TABLE IF NOT EXISTS gamers
(
    id          TEXT NOT NULL,
    first_name  TEXT NULL,
    last_name   TEXT NOT NULL,
    email       TEXT NOT NULL,
    password    TEXT NOT NULL,
    is_approved bool        NOT NULL DEFAULT FALSE,
    created_at  timestamptz NOT NULL DEFAULT NOW(),
    updated_at  timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);