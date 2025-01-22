DROP DATABASE IF EXISTS vega_db;
CREATE DATABASE vega_db;
USE vega_db;

DROP TABLE IF EXISTS users;
CREATE TABLE users(
    `id` INT AUTO_INCREMENT COMMENT 'id' ,
    `email` VARCHAR(255) NOT NULL  COMMENT 'email' ,
    `hashed_password` VARCHAR(255) NOT NULL  COMMENT 'hashed password' ,
    `avatar` VARCHAR(255) NOT NULL  COMMENT 'avatar' ,
    `name` VARCHAR(255) NOT NULL  COMMENT 'name' ,
    `gender` VARCHAR(255)   COMMENT 'gender' ,
    `birth_date` VARCHAR(255)   COMMENT 'birth date' ,
    `location` VARCHAR(255)   COMMENT 'location' ,
    `bio` VARCHAR(255)   COMMENT 'bio' ,
    `created_at` DATETIME NOT NULL  COMMENT 'created at' ,
    `updated_at` DATETIME   COMMENT 'updated at' ,
    `deleted_at` DATETIME   COMMENT 'deleted at' ,
    PRIMARY KEY (id)
)  COMMENT = 'users';

DROP TABLE IF EXISTS movies;
CREATE TABLE movies(
    `id` INT AUTO_INCREMENT COMMENT 'id' ,
    `cover` VARCHAR(255) NOT NULL  COMMENT 'cover' ,
    `title` VARCHAR(255) NOT NULL  COMMENT 'title' ,
    `release_date` VARCHAR(255) NOT NULL  COMMENT 'release date' ,
    `location` VARCHAR(255) NOT NULL  COMMENT 'location' ,
    `director` VARCHAR(255) NOT NULL  COMMENT 'director' ,
    `runtime` VARCHAR(255) NOT NULL  COMMENT 'runtime' ,
    `language` VARCHAR(255) NOT NULL  COMMENT 'language' ,
    `storyline` VARCHAR(255) NOT NULL  COMMENT 'storyline' ,
    `critic_rating` DECIMAL(24,6) NOT NULL  COMMENT 'critic rating' ,
    `user_rating` DECIMAL(24,6) NOT NULL  COMMENT 'user rating' ,
    `created_at` DATETIME NOT NULL  COMMENT 'created at' ,
    `updated_at` DATETIME   COMMENT 'updated at' ,
    `deleted_at` DATETIME   COMMENT 'deleted at' ,
    PRIMARY KEY (id)
)  COMMENT = 'movies';

DROP TABLE IF EXISTS reviews;
CREATE TABLE reviews(
    `id` INT AUTO_INCREMENT COMMENT 'id' ,
    `user_id` INT NOT NULL  COMMENT 'user id' ,
    `movie_id` INT NOT NULL  COMMENT 'movie id' ,
    `rating` DECIMAL(24,6) NOT NULL  COMMENT 'rating' ,
    `content` VARCHAR(255) NOT NULL  COMMENT 'content' ,
    `created_at` DATETIME NOT NULL  COMMENT 'created at' ,
    `updated_at` DATETIME   COMMENT 'updated at' ,
    `deleted_at` DATETIME   COMMENT 'deleted at' ,
    PRIMARY KEY (id)
)  COMMENT = 'reviews';

DROP TABLE IF EXISTS genres;
CREATE TABLE genres(
    `id` INT AUTO_INCREMENT COMMENT 'id' ,
    `name` VARCHAR(255) NOT NULL  COMMENT 'name' ,
    `created_at` DATETIME NOT NULL  COMMENT 'created at' ,
    `updated_at` DATETIME   COMMENT 'updated at' ,
    `deleted_at` DATETIME   COMMENT 'deleted at' ,
    PRIMARY KEY (id)
)  COMMENT = 'genres;';

DROP TABLE IF EXISTS stars;
CREATE TABLE stars(
    `id` INT AUTO_INCREMENT COMMENT 'id' ,
    `name` VARCHAR(255) NOT NULL  COMMENT 'name' ,
    `created_at` DATETIME NOT NULL  COMMENT 'created at' ,
    `updated_at` DATETIME   COMMENT 'updated at' ,
    `deleted_at` DATETIME   COMMENT 'deleted at' ,
    PRIMARY KEY (id)
)  COMMENT = 'stars;';

DROP TABLE IF EXISTS movie_genres;
CREATE TABLE movie_genres(
    `movie_id` INT NOT NULL  COMMENT 'movie id' ,
    `genre_id` INT NOT NULL  COMMENT 'genre id' ,
    `created_at` DATETIME NOT NULL  COMMENT 'created at' ,
    `updated_at` DATETIME   COMMENT 'updated at' ,
    `deleted_at` DATETIME   COMMENT 'deleted at' ,
    PRIMARY KEY (movie_id,genre_id)
)  COMMENT = 'movie_genres;';

DROP TABLE IF EXISTS movie_stars;
CREATE TABLE movie_stars(
    `movie_id` INT NOT NULL  COMMENT 'movie id' ,
    `star_id` INT NOT NULL  COMMENT 'star id' ,
    `created_at` DATETIME NOT NULL  COMMENT 'created at' ,
    `updated_at` DATETIME   COMMENT 'updated at' ,
    `deleted_at` DATETIME   COMMENT 'deleted at' ,
    PRIMARY KEY (movie_id,star_id)
)  COMMENT = 'movie_stars;';

alter table reviews add constraint FK_reviews1 foreign key (user_id)
      references users (id) on delete restrict on update restrict;
      
alter table reviews add constraint FK_reviews2 foreign key (movie_id)
      references movies (id) on delete restrict on update restrict;
      
alter table movie_genres add constraint FK_movie_genres1 foreign key (movie_id)
      references movies (id) on delete restrict on update restrict;
      
alter table movie_genres add constraint FK_movie_genres2 foreign key (genre_id)
      references genres (id) on delete restrict on update restrict;
      
alter table movie_stars add constraint FK_movie_stars1 foreign key (movie_id)
      references movies (id) on delete restrict on update restrict;
      
alter table movie_stars add constraint FK_movie_stars2 foreign key (star_id)
      references stars (id) on delete restrict on update restrict;

