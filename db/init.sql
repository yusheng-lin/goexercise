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
        INSERT INTO users (acct, pwd, fullname, created_at)
            VALUES ('alan','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','alan pu','2022-01-03'),
                   ('ben','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','ben yang','2022-01-02'),
                   ('jane','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','jane lin','2022-01-01'),
                   ('ted','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','ted lui','2022-01-04'),
                   ('caesar','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','caesar chou','2022-01-05'),
                   ('penny','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','penny tsai','2022-01-07'),
                   ('scott','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','scott chang','2022-01-06'),
                   ('miles','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','miles liao','2022-01-09'),
                   ('winnie','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','winnie huang','2022-01-08'),
                   ('luis','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','luis weng','2022-01-10'),
                   ('lily','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','lily chang','2022-01-19'),
                   ('alice','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','alice chang','2022-01-11'),
                   ('james','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','james wang','2022-01-12'),
                   ('johnny','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','johnny lai','2022-01-13'),
                   ('iris','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','iris lai','2022-01-16'),
                   ('eureka','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','eureka ye','2022-01-14'),
                   ('howard','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','howard liao','2022-01-15'),
                   ('chris','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','chris wang','2022-01-17'),
                   ('knuck','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','knuck wang','2022-01-18'),
                   ('daniel','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','daniel shih','2022-01-20'),
                   ('terence','41e622a7d7993bcfc8eddd57edb0ee9a5c25b09d219bc805dcb22b08f8a852d6','terence shih','2022-01-21');
    END IF;
END
$do$