create table IF NOT EXISTS users (
    id SERIAL primary key,
    name varchar(255) not null,
    email varchar(255) unique not null ,
    password varchar(255) not null,
    is_verifed boolean default false,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);