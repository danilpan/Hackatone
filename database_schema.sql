CREATE TABLE users
(
    iin       CHAR(12) NOT NULL
        CONSTRAINT users_pk
            PRIMARY KEY,
    full_name VARCHAR  NOT NULL
);

ALTER TABLE users
    OWNER TO admin;

CREATE TABLE establishment_types
(
    id   SERIAL
        CONSTRAINT establishment_types_pk
            PRIMARY KEY,
    name VARCHAR NOT NULL
);

ALTER TABLE establishment_types
    OWNER TO admin;

CREATE TABLE establishments
(
    id            SERIAL
        CONSTRAINT establishments_pk
            PRIMARY KEY,
    name          VARCHAR NOT NULL,
    address       VARCHAR NOT NULL,
    type_id       INTEGER NOT NULL
        CONSTRAINT establishments_establishment_types_id_fk
            REFERENCES establishment_types,
    average_check INTEGER NOT NULL,
    rating        INTEGER NOT NULL
        CONSTRAINT check_name
            CHECK ((1 <= rating) AND (rating <= 5))
);

ALTER TABLE establishments
    OWNER TO admin;

CREATE TABLE tables
(
    id               SERIAL
        CONSTRAINT tables_pk
            PRIMARY KEY,
    establishment_id INTEGER NOT NULL
        CONSTRAINT tables_establishments_id_fk
            REFERENCES establishments,
    number           INTEGER NOT NULL
);

ALTER TABLE tables
    OWNER TO admin;

CREATE TABLE reservations
(
    id        INTEGER   NOT NULL
        CONSTRAINT reservations_pk
            PRIMARY KEY,
    table_id  INTEGER   NOT NULL
        CONSTRAINT reservations_tables_id_fk
            REFERENCES tables,
    user_iin  CHAR(12)  NOT NULL
        CONSTRAINT reservations_users_iin_fk
            REFERENCES users,
    time_from TIMESTAMP NOT NULL,
    time_to   TIMESTAMP NOT NULL
);

ALTER TABLE reservations
    OWNER TO admin;

