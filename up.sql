CREATE TABLE MessageRow (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    Data BLOB,
    Timestamp INTEGER,
    Topic TEXT,
    Partition INTEGER
);