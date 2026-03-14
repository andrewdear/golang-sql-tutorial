-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id              BIGSERIAL PRIMARY KEY,
    username        varchar(255)
);

CREATE TABLE todos
(
    id              BIGSERIAL PRIMARY KEY,
    content         text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE todos;
-- +goose StatementEnd
