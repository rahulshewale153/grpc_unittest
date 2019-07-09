
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `articles` (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `nickname` varchar(250) NOT NULL,
  `title` varchar(250) NOT NULL,
  `article_creation_date` date NOT NULL,
  `content` text NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime NULL,
 PRIMARY KEY (`id`)) ENGINE=InnoDB;

CREATE TABLE `comments` (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `nickname` varchar(250) NOT NULL,
  `article_id` int(10) UNSIGNED NOT NULL,
  `content` text NOT NULL,
  `comment_creation_date` date NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime  NULL,
 PRIMARY KEY (`id`),
 KEY `article_id` (`article_id`),
 CONSTRAINT `comment_articles` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`)
 ON DELETE RESTRICT ON UPDATE CASCADE
    ) ENGINE=InnoDB ;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
Drop TABLE  IF EXISTS `comments`; 
Drop TABLE  IF EXISTS `articles`; 