CREATE TABLE IF NOT EXISTS Person(
                                     id SERIAL NOT NULL PRIMARY KEY,
                                     name TEXT NOT NULL,
                                     age INTEGER NOT NULL,
                                     license BOOL NOT NULL
);