ALTER TABLE visitors DROP COLUMN clock_in;
ALTER TABLE visitors DROP COLUMN clock_out;
ALTER TABLE visitors DROP COLUMN createdAt;
ALTER TABLE visitors DROP COLUMN updatedAt;
ALTER TABLE visitors ADD clock_in DATE NULL;
ALTER TABLE visitors ADD clock_out DATE NULL;
ALTER TABLE visitors ADD createdAt DATE NULL;
ALTER TABLE visitors ADD updatedAt DATE NULL;