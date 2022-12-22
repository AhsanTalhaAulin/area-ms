CREATE DATABASE IF NOT EXISTS area_db;
USE area_db;

CREATE TABLE areas(
    areaID int AUTO_INCREMENT,
    gym boolean,
    school boolean,
    store boolean,
    lab boolean,
    office boolean,
    PRIMARY KEY (areaID)
);

CREATE TABLE dist_areas(
    areaID int AUTO_INCREMENT,
    gym int,
    school int,
    store int,
    lab int,
    office int,
    PRIMARY KEY (areaID)
);

CREATE USER 'area_user'@'%' IDENTIFIED BY '12345678';
GRANT ALL PRIVILEGES ON *.* TO 'area_user'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;