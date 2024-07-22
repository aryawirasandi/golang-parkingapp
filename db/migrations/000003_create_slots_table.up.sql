CREATE TABLE slots (
    id INT NOT NULL PRIMARY KEY,
    visitor_id INT NOT NULL,
    slot_name VARCHAR(50) NOT NULL UNIQUE,
    FOREIGN KEY (visitor_id) REFERENCES visitors (id)
);