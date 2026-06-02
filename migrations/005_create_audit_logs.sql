CREATE TABLE IF NOT EXISTS audit_logs (
    id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    actor_id       UUID REFERENCES users(id),
    action         VARCHAR(50) NOT NULL,
    token          VARCHAR(100),
    field_type     VARCHAR(50),
    access_level   VARCHAR(20),
    ip_address     VARCHAR(50),
    success        BOOLEAN NOT NULL,
    failure_reason TEXT,
    created_at     TIMESTAMPTZ DEFAULT NOW()
);
