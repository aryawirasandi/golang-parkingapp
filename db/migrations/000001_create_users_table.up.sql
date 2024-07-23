CREATE TABLE users (
    id int AUTO_INCREMENT,
    username VARCHAR(25),
    password VARCHAR(100),
    token VARCHAR(100) DEFAULT '',
    role ENUM('staff', 'super') NOT NULL DEFAULT 'staff',
    createdAt DATE,
    updatedAt DATE,
    PRIMARY KEY (id)
)