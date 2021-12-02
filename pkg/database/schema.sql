--DROP TABLE IF EXISTS games CASCADE;
CREATE TABLE IF NOT EXISTS games (
	id	serial not null primary key,
	name	varchar(255) not null,
	description	text
);
--DROP TABLE IF EXISTS players CASCADE;
CREATE TABLE IF NOT EXISTS players (
	id	serial not null primary key,
	name	varchar(255) not null
);
--DROP TABLE IF EXISTS players_games CASCADE;
CREATE TABLE IF NOT EXISTS players_games (
	player_id	integer,
	game_id	integer,
    foreign key(player_id) references players(id),
    foreign key(game_id) references games(id) on delete CASCADE
);
--DROP TABLE IF EXISTS stats CASCADE;
CREATE TABLE IF NOT EXISTS stats (
	player_id	integer,
	game_id	integer,
	date	timestamptz not null,
    foreign key(player_id) references players(id),
    foreign key(game_id) references games(id) on delete CASCADE
);
--DROP INDEX IF EXISTS players_name_unique;
CREATE UNIQUE INDEX IF NOT EXISTS players_name_unique on players (name);