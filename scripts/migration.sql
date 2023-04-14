CREATE TABLE IF NOT EXISTS wallet (
    id UUID NOT NULL,
    value BIGINT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS wallet_event_store (
    id UUID NOT NULL,
    wallet_id UUID NOT NULL,
    type TEXT NOT NULL,
    third_party_wallet_id UUID,
    value BIGINT NOT NULL,
    test BOOLEAN NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (wallet_id) REFERENCES wallet(id),
    FOREIGN KEY (third_party_wallet_id) REFERENCES wallet(id),
);

CREATE TABLE IF NOT EXISTS inventory (
    id UUID NOT NULL,
    wallet_id UUID NOT NULL,
    external_reference UUID NOT NULL UNIQUE,
    PRIMARY KEY (id),
    FOREIGN KEY (wallet_id) REFERENCES wallet(id)
);

CREATE TABLE IF NOT EXISTS item (
    id UUID NOT NULL,
    inventory_id UUID NOT NULL,
    external_reference UUID NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (inventory_id) REFERENCES inventory(id)
);

CREATE TABLE IF NOT EXISTS item_event_store (
    id UUID NOT NULL,
    occurred_on BIGINT NOT NULL,
    type TEXT NOT NULL,
    third_party_inventory_id UUID,
    wallet_event_store_id UUID,
    external_reference UUID NOT NULL,
    test BOOLEAN NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (third_party_inventory_id) REFERENCES inventory(id),
    FOREIGN KEY (wallet_event_store_id) REFERENCES wallet_event_store(id)
);