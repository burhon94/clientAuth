package dl

const ClientDML = `INSERT INTO clients(first_name, last_name, middle_name, login, password, e_mail, avatar, phone)
VALUES ('ADMIN', 'ADMINISTRATOR', 'ADM', 'admin', 'admin', 'admin@mail.com', 'NoAvatar', '1111')
ON CONFLICT DO NOTHING`

const ClientNew = `INSERT INTO clients(first_name, last_name, middle_name, login, password, e_mail, avatar, phone)
values ($1, $2, $3 , $4, $5, $6, $7, $8);`
