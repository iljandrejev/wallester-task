create type gender as ENUM ('Male', 'Female');

create table customers
(
    id serial constraint customers_pkey primary key,
    firstname  varchar(100) not null,
    lastname   varchar(100) not null,
    gender     gender       not null,
    birthday   date         not null,
    email      varchar(255) not null,
    address    varchar(200),
    hash       varchar(255) not null,
    created_at timestamp    default now() not null,
    updated_at timestamp    default now() not null,
);
