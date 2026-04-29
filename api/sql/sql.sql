CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users (
  id int auto_increment primary key,
  name varchar(50) not null,
  nick varchar(50) not null unique,
  email varchar(50) not null unique,
  password varchar(100) not null,
  created_at timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS followers (
  user_id int not null,
  follower_id int not null,
  primary key (user_id, follower_id),
  foreign key (user_id) references users(id) on delete cascade,
  foreign key (follower_id) references users(id) on delete cascade
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS posts (
  id int auto_increment primary key,
  title varchar(50) not null,
  content varchar(300) not null,
  author_id int not null,
  foreign key (author_id) references users(id) on delete cascade,
  likes int default 0,
  created_at timestamp default current_timestamp()
) ENGINE=INNODB;
