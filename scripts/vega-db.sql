CREATE TABLE IF NOT EXISTS users
(
    id              INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, --id
    email           TEXT                              NOT NULL, --email
    hashed_password TEXT                              NOT NULL, --hashed password
    avatar          TEXT                              NOT NULL, --avatar
    name            TEXT                              NOT NULL, --name
    gender          TEXT,                                       --gender
    birth_date      TEXT,                                       --birth date
    location        TEXT,                                       --location
    bio             TEXT,                                       --bio
    created_at      TEXT                              NOT NULL, --created at
    updated_at      TEXT,                                       --updated at
    deleted_at      TEXT                                        --deleted at
); --users

CREATE TABLE IF NOT EXISTS movies
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, --id
    cover         TEXT                              NOT NULL, --cover
    title         TEXT                              NOT NULL, --title
    release_date  TEXT,                                       --release date
    location      TEXT,                                       --location
    director      TEXT,                                       --director
    language      TEXT,                                       --language
    runtime       TEXT,                                       --runtime
    storyline     TEXT,                                       --storyline
    critic_rating REAL                              NOT NULL, --critic rating
    user_rating   REAL                              NOT NULL, --user rating
    created_at    TEXT                              NOT NULL, --created at
    updated_at    TEXT,                                       --updated at
    deleted_at    TEXT                                        --deleted at
); --movies

CREATE TABLE IF NOT EXISTS reviews
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, --id
    user_id    INTEGER                           NOT NULL, --user id
    movie_id   INTEGER                           NOT NULL, --movie id
    rating     REAL                              NOT NULL, --rating
    content    TEXT                              NOT NULL, --content
    created_at TEXT                              NOT NULL, --created at
    updated_at TEXT,                                       --updated at
    deleted_at TEXT,                                       --deleted at
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (movie_id) REFERENCES movies (id)
); --reviews

CREATE TABLE IF NOT EXISTS genres
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, --id
    name       TEXT                              NOT NULL, --name
    created_at TEXT                              NOT NULL, --created at
    updated_at TEXT,                                       --updated at
    deleted_at TEXT                                        --deleted at
); --genres;

CREATE TABLE IF NOT EXISTS stars
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, --id
    name       TEXT                              NOT NULL, --name
    created_at TEXT                              NOT NULL, --created at
    updated_at TEXT,                                       --updated at
    deleted_at TEXT                                        --deleted at
); --stars;

CREATE TABLE IF NOT EXISTS movie_genres
(
    movie_id   INTEGER NOT NULL, --movie id
    genre_id   INTEGER NOT NULL, --genre id
    created_at TEXT    NOT NULL, --created at
    updated_at TEXT,             --updated at
    deleted_at TEXT,             --deleted at
    PRIMARY KEY (movie_id, genre_id),
    FOREIGN KEY (movie_id) REFERENCES movies (id),
    FOREIGN KEY (genre_id) REFERENCES genres (id)
); --movie_genres;

CREATE TABLE IF NOT EXISTS movie_stars
(
    movie_id   INTEGER NOT NULL, --movie id
    star_id    INTEGER NOT NULL, --star id
    created_at TEXT    NOT NULL, --created at
    updated_at TEXT,             --updated at
    deleted_at TEXT,             --deleted at
    PRIMARY KEY (movie_id, star_id),
    FOREIGN KEY (movie_id) REFERENCES movies (id),
    FOREIGN KEY (star_id) REFERENCES stars (id)
); --movie_stars;
