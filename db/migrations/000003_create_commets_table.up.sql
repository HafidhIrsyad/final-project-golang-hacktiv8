begin;

create table if not exists comments(
    id serial primary key ,
    user_id int references users(id),
    photo_id int not null references photos(id),
    message text not null ,
    created_at timestamp,
    updated_at timestamp
);


commit;
