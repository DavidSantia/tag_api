--
-- Create groups table
--

CREATE TABLE IF NOT EXISTS groups (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(50) DEFAULT NULL,
  sess_seconds int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET utf8;

--
-- Create images table
--

CREATE TABLE IF NOT EXISTS images (
  id int(11) NOT NULL AUTO_INCREMENT,
  width int(11) DEFAULT NULL,
  height int(11) DEFAULT NULL,
  url varchar(255) DEFAULT NULL,
  title varchar(100) DEFAULT NULL,
  artist varchar(100) DEFAULT NULL,
  gallery varchar(100) DEFAULT NULL,
  organization varchar(100) DEFAULT NULL,
  media varchar(25) DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET utf8;

--
-- Create images_groups table
--

CREATE TABLE IF NOT EXISTS images_groups (
  id int(11) NOT NULL AUTO_INCREMENT,
  group_id int(11) NOT NULL,
  image_id int(11) NOT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET utf8;

--
-- Create users table
--

CREATE TABLE IF NOT EXISTS users (
  id int(11) NOT NULL AUTO_INCREMENT,
  group_id int(11) DEFAULT NULL,
  guid varchar(40) DEFAULT NULL,
  first_name varchar(50) DEFAULT NULL,
  middle_init varchar(2) DEFAULT NULL,
  last_name varchar(50) DEFAULT NULL,
  email varchar(50) DEFAULT NULL,
  addr varchar(50) DEFAULT NULL,
  city varchar(50) DEFAULT NULL,
  state varchar(2) DEFAULT NULL,
  zip varchar(10) DEFAULT NULL,
  gender varchar(1) DEFAULT NULL,
  status boolean DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET utf8;
