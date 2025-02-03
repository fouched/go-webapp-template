create table customer (
    id bigserial primary key,
    customer_name varchar(255) unique not null,
    tel varchar(64) not null,
    email varchar(128) not null,
    address_1 varchar(255) not null default(''),
    address_2 varchar(255) not null default(''),
    address_3 varchar(255) not null default(''),
    post_code varchar(32) not null default(''),
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);