-- +goose Up
-- +goose StatementBegin
ALTER TABLE todos ADD user_id int;
ALTER TABLE users ADD email varchar(255) UNIQUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE todos DROP COLUMN user_id;
ALTER TABLE users DROP COLUMN email;

-- +goose StatementEnd
