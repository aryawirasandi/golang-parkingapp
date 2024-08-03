CREATE TABLE slots (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    visitor_id INT NULL,
    slot_name VARCHAR(50) NOT NULL UNIQUE,
    FOREIGN KEY (visitor_id) REFERENCES visitors (id)
);