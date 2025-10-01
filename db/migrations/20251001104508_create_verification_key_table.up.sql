CREATE TABLE verification_keys (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    key TEXT NOT NULL,       -- simpan hash dari token
    expired_at TIMESTAMP NOT NULL,  -- masa berlaku
    is_used BOOLEAN DEFAULT FALSE,  -- apakah sudah dipakai
    created_at TIMESTAMP DEFAULT NOW()
);