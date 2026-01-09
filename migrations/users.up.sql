CREATE TABLE IF NOT EXISTS Users (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    email VARCHAR(320) NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE UNIQUE Index users_emails_unique
ON Users (LOWER(email));