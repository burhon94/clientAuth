CREATE TABLE if not exists clients
(
    id          BIGSERIAL PRIMARY KEY,
    first_name  TEXT        NOT NULL,
    last_name   TEXT        NOT NULL,
    middle_name TEXT,
    login       TEXT UNIQUE NOT NULL,
    password    TEXT        NOT NULL,
    e_mail      TEXT,
    avatar      TEXT,
    phone       TEXT UNIQUE NOT NULL
);

INSERT INTO clients(first_name, last_name, middle_name, login, password, e_mail, avatar, phone)
VALUES ('ADMIN', 'ADMINISTRATOR', 'ADM', 'admin', 'admin', 'admin@mail.com', 'NoAvatar', '1111')
ON CONFLICT DO NOTHING;

INSERT INTO clients(first_name, last_name, middle_name, login, password, e_mail, avatar, phone)
values (?, ?, ?, ?, ?, ?, ?, ?);

SELECT first_name, last_name, middle_name, e_mail, avatar, phone, password FROM clients WHERE login = ?;