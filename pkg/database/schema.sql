--DROP TABLE IF EXISTS `games`;
CREATE TABLE IF NOT EXISTS `games` (
	`id`	integer not null primary key autoincrement,
	`name`	varchar(255) not null,
	`description`	text
);
--DROP TABLE IF EXISTS `players`;
CREATE TABLE IF NOT EXISTS `players` (
	`id`	integer not null primary key autoincrement,
	`name`	varchar(255) not null
);
--DROP TABLE IF EXISTS `players_games`;
CREATE TABLE IF NOT EXISTS `players_games` (
	`player_id`	integer,
	`game_id`	integer,
    foreign key(`player_id`) references `players`(`id`),
    foreign key(`game_id`) references `games`(`id`) on delete CASCADE
);
--DROP TABLE IF EXISTS `stats`;
CREATE TABLE IF NOT EXISTS `stats` (
	`player_id`	integer,
	`game_id`	integer,
	`date`	datetime not null,
    foreign key(`player_id`) references `players`(`id`),
    foreign key(`game_id`) references `games`(`id`) on delete CASCADE
);
--DROP INDEX IF EXISTS `players_name_unique`;
CREATE UNIQUE INDEX IF NOT EXISTS `players_name_unique` on `players` (`name`);