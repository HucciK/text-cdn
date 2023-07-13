CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    email TEXT,
    password_hash TEXT,
    name TEXT NOT NULL,
    coin_balance INTEGER,
    crystall_balance INTEGER
);

CREATE TABLE dailybox_rewards(
    id SERIAL PRIMARY KEY,
    day SMALLINT,
    reward_type SMALLINT,
    reward_value INTEGER,
);

CREATE TABLE dailybox_progress(

);