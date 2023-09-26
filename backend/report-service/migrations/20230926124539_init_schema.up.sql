CREATE TYPE "reporttype" AS ENUM ('Scam And Fraudulent Listing', 'Fake Review', 'Suspicious User', 'Website Bug', 'Suggestion');

CREATE TABLE IF NOT EXISTS reports (
    id BIGSERIAL PRIMARY KEY,
    uid BIGSERIAL NOT NULL,
    topic VARCHAR(255) NOT NULL,
    type reporttype NOT NULL,
    description VARCHAR(1000) NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);