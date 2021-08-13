CREATE TABLE user
(
    id   BIGINT auto_increment
        PRIMARY KEY,
    name LONGTEXT NULL,
    sn   BIGINT DEFAULT 0
) ENGINE = InnoDB
    DEFAULT CHARSET utf8;

CREATE TABLE company
(
    id   BIGINT auto_increment
        PRIMARY KEY,
    sn   BIGINT DEFAULT 0,
    name LONGTEXT NULL
) ENGINE = InnoDB
    DEFAULT CHARSET utf8;