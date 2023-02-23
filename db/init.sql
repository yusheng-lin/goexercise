CREATE TABLE IF NOT EXISTS users( 
    acct varchar(20) NOT NULL, 
    pwd varchar(200) NOT NULL, 
    fullname varchar(50) NOT NULL, 
    created_at timestamptz NOT NULL DEFAULT (now()), 
    updated_at timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT PK_users PRIMARY KEY (acct)
);

DO
$do$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM users) THEN
        INSERT INTO users
            VALUES ('ben11','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','ben yang'),
                   ('jane22','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','jane lin'),
                   ('john33','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','john chen');
    END IF;
END
$do$