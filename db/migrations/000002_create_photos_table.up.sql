begin;

create table if not exists photos(
    id serial primary key ,
    title varchar(200) not null ,
    caption varchar(200) not null ,
    photo_url text not null ,
    user_id int not null references users(id),
    updated_at timestamp,
    created_at timestamp
);

commit;