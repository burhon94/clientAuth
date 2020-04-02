package dl

const ClientDML = `INSERT INTO clients(first_name, last_name, middle_name, login, password, e_mail, phone)
VALUES ('ADMIN', 'ADMINISTRATOR', 'ADM', 'admin', 'admin', 'admin@mail.com', '1111')
ON CONFLICT DO NOTHING;`

const ClientNew = `INSERT INTO clients(first_name, last_name, login, password)
values ($1, $2, $3 , $4);`
