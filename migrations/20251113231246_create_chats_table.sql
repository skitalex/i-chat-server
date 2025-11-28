-- +goose Up
CREATE TABLE IF NOT EXISTS chats (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_chats_group ON chats ((id >> 62)) WHERE (id >> 62) = 1;

CREATE TABLE IF NOT EXISTS id_sequences (
    name TEXT PRIMARY KEY,
    last_value BIGINT NOT NULL
);

-- Инициализация (выполните один раз)
INSERT INTO id_sequences (name, last_value)
VALUES ('chat_id', 0)
ON CONFLICT DO NOTHING;
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
