package dl

const ClientDDL = `CREATE TABLE if not exists clients
(
    id BIGSERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    middle_name TEXT,
    login TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    e_mail TEXT,
    phone TEXT
);`
