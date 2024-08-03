CREATE TABLE users (
    id int AUTO_INCREMENT,
    username VARCHAR(25),
    password VARCHAR(100),
    role ENUM('staff', 'super') NOT NULL DEFAULT 'staff',
    createdAt DATE,
    updatedAt DATE,
    PRIMARY KEY (id)
)