CREATE TABLE if not exists clients
(
    id          BIGSERIAL PRIMARY KEY,
    first_name  TEXT        NOT NULL,
    last_name   TEXT        NOT NULL,
    middle_name TEXT,
    login       TEXT UNIQUE NOT NULL,
    password    TEXT        NOT NULL,
    e_mail      TEXT,
    phone       TEXT
);

INSERT INTO clients(first_name, last_name, middle_name, login, password, e_mail, phone)
VALUES ('ADMIN', 'ADMINISTRATOR', 'ADM', 'admin', 'admin', 'admin@mail.com', '1111')
ON CONFLICT DO NOTHING;

INSERT INTO clients(first_name, last_name, login, password)
values (?, ?, ? , ?);