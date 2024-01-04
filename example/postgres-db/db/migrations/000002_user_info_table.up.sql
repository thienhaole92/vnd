CREATE TABLE IF NOT EXISTS user_info (
    id VARCHAR PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL
);

ALTER TABLE
    user_info
ADD
    CONSTRAINT idx_user_info_id_unique_constraint UNIQUE (id);