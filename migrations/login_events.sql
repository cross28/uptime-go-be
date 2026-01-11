CREATE TABLE IF NOT EXISTS login_events (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    user_id UUID NOT NULL,
    provider VARCHAR(32) NOT NULL,
    ip_address VARCHAR(15) NOT NULL,
    user_agent VARCHAR(300) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);