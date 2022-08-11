create type gender as ENUM ('Male', 'Female');

create table customers
(
    created_at timestamp    default now()                  not null,
    email      varchar(255)                                not null,
    address    varchar(200),
    hash       varchar(255) default 'a'::character varying not null,
    updated_at timestamp    default now()                  not null,
    gender     gender                                      not null,
    birthday   date                                        not null,
    lastname   varchar(100)                                not null,
    firstname  varchar(100)                                not null,
    id         serial
        constraint customers_pkey
            primary key
);
