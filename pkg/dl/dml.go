package dl

const CheckId = `SELECT id FROM clients WHERE id = $1;`

const CheckLogin = `SELECT login from clients WHERE login = $1;`

const CheckPhone = `SELECT phone from clients WHERE phone = $1;`

const CheckPass = `SELECT password from clients WHERE id = $1;`

const CheckPassAndLogin = `SELECT password from clients WHERE login = $1;`

const ClientDML = `INSERT INTO clients(first_name, last_name, middle_name, login, password, e_mail, avatar, phone)
VALUES ('ADMIN', 'ADMINISTRATOR', 'ADM', 'admin', '$2y$12$qbxGqj3HFqrkWdXxRSI1q.t3YQf3pNcxbsdIGtZzOKENvBspnq9jq', 'admin@mail.com', 'NoAvatar', '1111')
ON CONFLICT DO NOTHING`

const SignIn = `SELECT first_name, last_name, middle_name, e_mail, avatar, phone FROM clients WHERE login = $1;`

const ClientNew = `INSERT INTO clients(first_name, last_name, middle_name, login, password, e_mail, avatar, phone)
values ($1, $2, $3 , $4, $5, $6, $7, $8);`

const ClientUpdatePass = `UPDATE clients SET password = $2 WHERE id = $1;`

const ClientUpdateAvatar = `UPDATE clients SET avatar = $2 WHERE id = $1;`

const ClientUpdateData = `UPDATE clients SET first_name = $2, last_name = $3, middle_name = $4, e_mail = $5 WHERE id = $1;`
