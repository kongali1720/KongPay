CREATE TABLE IF NOT EXISTS reconciliation_events (

    id UUID PRIMARY KEY,

    settlement_id UUID NOT NULL,

    reconciliation_id UUID NOT NULL,

    event_type VARCHAR(100) NOT NULL,

    old_status VARCHAR(50),

    new_status VARCHAR(50),

    created_at TIMESTAMP DEFAULT NOW()

);
