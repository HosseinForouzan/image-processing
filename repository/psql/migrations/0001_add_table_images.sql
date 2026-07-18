-- +migrate Up

CREATE TYPE image_status AS ENUM (
    'processing',
    'completed',
    'failed'
);

CREATE TABLE images (
    id SERIAL PRIMARY KEY,

    original_name VARCHAR(255) NOT NULL,

    original_key TEXT NOT NULL,

    thumbnail_key TEXT,

    content_type VARCHAR(100) NOT NULL,

    size BIGINT NOT NULL CHECK (size >= 0),

    status image_status NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- -migrate Down

DROP TABLE images;