CREATE TABLE IF NOT EXISTS inventory (
    id UUID NOT NULL,
    wallet_id UUID NOT NULL,
    external_reference UUID NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS item (
    id UUID NOT NULL,
    inventory_id UUID NOT NULL,
    external_reference UUID NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS item_event (
    id UUID NOT NULL,
    occurred_on BIGINT NOT NULL,
    type TEXT NOT NULL,
    third_party_inventory_id UUID,
    trade_reference UUID,
    value BIGINT,
    test BOOLEAN NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS slot (
    id UUID NOT NULL,
    item_id UUID NOT NULL,
    inventory_id UUID NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS wallet (
    id UUID NOT NULL,
    last_wallet_event_id UUID NOT NULL,
    value BIGINT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS wallet_event (
    id UUID NOT NULL,
    wallet_id UUID NOT NULL,
    type TEXT NOT NULL,
    third_party_wallet_id UUID,
    trade_reference UUID,
    value BIGINT NOT NULL,
    test BOOLEAN NOT NULL,
    PRIMARY KEY (id)
);
