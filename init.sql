DROP TABLE IF EXISTS address; 
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS account; 

CREATE TABLE account (
    ID UUID                             NOT NULL PRIMARY KEY,
    User     VARCHAR(10)                NOT NULL UNIQUE,
    Pas      VARCHAR(100)               NULL,
    Salt     VARCHAR(100)               NULL
);


CREATE TABLE users (
    ACCOUNT_ID  UUID references account(id)  NOT NULL,
    ID          UUID            NOT NULL PRIMARY KEY,
    NAME        VARCHAR(50)     NOT NULL,
    LAST_NAME    VARCHAR(100)   NULL,
    AGE         SMALLINT        NOT NULL
);
-- para falar a verdade, o address poderia ser só um []json na tabela de usuários, mas seguimos com tabela a parte-- 


CREATE TABLE address (
    ID UUID                             NOT NULL PRIMARY KEY,
    USER_ID  UUID references users(id)  NOT NULL,
    ZIP_CODE VARCHAR(10)                NOT NULL,
    DETAILS  VARCHAR(100)               NULL,
    STATE    VARCHAR(100)               NULL,
    COUNTRY  VARCHAR(100)               NULL,
    CITY     VARCHAR(100)               NULL
);

