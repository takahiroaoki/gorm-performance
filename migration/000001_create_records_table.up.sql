DROP TABLE IF EXISTS db.records;

CREATE TABLE db.records(
    id INT NOT NULL AUTO_INCREMENT,
    test VARCHAR(10) NOT NULL,

    PRIMARY KEY(id)
) ENGINE = innoDB DEFAULT CHARSET = utf8;