CREATE TABLE IF NOT EXISTS posts (
    pid BIGSERIAL PRIMARY KEY,
    uid BIGSERIAL NOT NULL,
    topic VARCHAR(255) NOT NULL,
    description json NOT NULL,
    period VARCHAR(255) NOT NULL,
    how_to json NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS open_positions (
    oid BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS posts_open_positions (
    pid BIGINT NOT NULL,
    oid BIGINT NOT NULL,
    PRIMARY KEY (pid, oid),
    FOREIGN KEY (pid) REFERENCES posts(pid),
    FOREIGN KEY (oid) REFERENCES open_positions(oid),
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS required_skills (
    sid BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS posts_required_skills (
    pid BIGINT NOT NULL,
    sid BIGINT NOT NULL,
    PRIMARY KEY (pid, sid),
    FOREIGN KEY (pid) REFERENCES posts(pid),
    FOREIGN KEY (sid) REFERENCES required_skills(sid),
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS benefits (
    bid BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS posts_benefits (
    pid BIGINT NOT NULL,
    bid BIGINT NOT NULL,
    PRIMARY KEY (pid, bid),
    FOREIGN KEY (pid) REFERENCES posts(pid),
    FOREIGN KEY (bid) REFERENCES benefits(bid),
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);