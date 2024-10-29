CREATE TABLE auth (
    id INT NOT NULL AUTO_INCREMENT,
    uuid VARCHAR(40),
    email VARCHAR(255) NOT NULL, 
    password VARCHAR(255) NOT NULL,
    createdAt timestamp NOT NULL DEFAULT current_timestamp,
    updatedAt timestamp NOT NULL DEFAULT current_timestamp,
    primary key (id)
);

CREATE TRIGGER before_insert_auth
    BEFORE INSERT ON auth
    FOR EACH ROW
        SET NEW.`password` = SHA1( NEW.`password` );

