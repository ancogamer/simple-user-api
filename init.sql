DROP TABLE IF EXISTS address; 
DROP TABLE IF EXISTS users;


CREATE TABLE users (
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
