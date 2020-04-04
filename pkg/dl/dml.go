package dl

const ClientDML = `INSERT INTO clients(first_name, last_name, middle_name, login, password, e_mail, avatar, phone)
VALUES ('ADMIN', 'ADMINISTRATOR', 'ADM', 'admin', 'admin', 'admin@mail.com', 'NoAvatar', '1111')
ON CONFLICT DO NOTHING`

const SignIn = `SELECT first_name, last_name, middle_name, e_mail, avatar, phone FROM clients WHERE login = $1;`

const ClientNew = `INSERT INTO clients(first_name, last_name, middle_name, login, password, e_mail, avatar, phone)
values ($1, $2, $3 , $4, $5, $6, $7, $8);`

const ClientUpdatePass = `UPDATE clients SET password = $2 WHERE id = $1;`

const ClientUpdateAvatar = `UPDATE clients SET avatar = $2 WHERE id = $1;`