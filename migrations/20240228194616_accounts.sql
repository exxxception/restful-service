-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS accounts (
 id SERIAL PRIMARY KEY,
 balance DECIMAL(15, 2) NOT NULL DEFAULT 0,
 created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS accounts;
-- +goose StatementEnd
