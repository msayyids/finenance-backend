CREATE TABLE verification_keys (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    key TEXT NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    verifed_at TIMESTAMP,
    is_used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);