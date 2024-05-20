-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS data_schema;

CREATE TABLE IF NOT EXISTS data_schema.data(
    OrderUID VARCHAR(255) PRIMARY KEY,
    Data JSONB
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA data_schema CASCADE;
-- +goose StatementEnd
