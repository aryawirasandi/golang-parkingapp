CREATE TABLE visitors (
    id INT NOT NULL,
    vehicle_number VARCHAR(10),
    clock_in DATE NOT NULL,
    clock_out DATE NOT NULL,
    type_vehicle ENUM('motocycle', 'car', 'truct') NOT NULL,
    createdAt DATE NOT NULL,
    updatedAt DATE NOT NULL,
    PRIMARY KEY (id),
    UNIQUE(vehicle_number)
)