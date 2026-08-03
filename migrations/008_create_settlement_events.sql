CREATE TABLE IF NOT EXISTS settlement_events (

    id UUID PRIMARY KEY,

    settlement_id UUID NOT NULL,

    event_type VARCHAR(50) NOT NULL,

    old_status VARCHAR(30),

    new_status VARCHAR(30),

    created_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_settlement_event
        FOREIGN KEY(settlement_id)
        REFERENCES settlements(id)
        ON DELETE CASCADE
);
