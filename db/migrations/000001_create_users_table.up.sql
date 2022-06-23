begin;

create table if not exists users(
    id serial primary key,
    username varchar(200) not null ,
    email varchar(200) unique not null ,
    password varchar(200) not null ,
    age int not null ,
    updated_at timestamp,
    created_at timestamp
);

commit;