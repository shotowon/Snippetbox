-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_snippets_created ON snippets(created);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_snippets_created;
-- +goose StatementEnd
