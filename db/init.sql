create table players
(
    id      serial       PRIMARY KEY,
    name    varchar(255) not null,
    club    varchar(255) not null,
    games   int          not null,
    goals   int          not null,
    assists int          not null
);

INSERT INTO public.players (name, club, games, goals, assists)
VALUES
    ('Messi', 'PSG', 546, 480, 197),
    ('Cristiano Ronaldo', 'MU', 615, 494, 128),
    ('Salah', 'Liverpool', 357, 175, 74);

INSERT INTO public.players(name, club, games, goals, assists)
SELECT
    left(md5(random()::text), 10),
    left(md5(random()::text), 10),
    (random()*300)::int,
    (random()*200)::int,
    (random()*50)::int
from generate_series(1,2000);

create index players_games_goals_assists_index
    on players (games, goals, assists);

