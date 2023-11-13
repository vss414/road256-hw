-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.players
(
    id      serial       PRIMARY KEY,
    name    varchar(255) not null,
    club    varchar(255) not null,
    games   int          not null,
    goals   int          not null,
    assists int          not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.players;
-- +goose StatementEnd
