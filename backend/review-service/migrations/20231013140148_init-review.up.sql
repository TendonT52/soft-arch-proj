CREATE TABLE IF NOT EXISTS reviews(
    rid BIGSERIAL PRIMARY KEY,
    uid BIGSERIAL NOT NULL,
    cid BIGSERIAL NOT NULL,
    title VARCHAR(255) NOT NULL,
    description json NOT NULL,
    rating integer NOT NULL CHECK (rating >= 1 AND rating <= 5),
    anonymous BOOLEAN NOT NULL DEFAULT FALSE,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
)