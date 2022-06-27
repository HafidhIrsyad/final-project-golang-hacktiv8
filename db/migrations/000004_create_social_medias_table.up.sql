begin;

create table if not exists social_medias(
    id serial primary key ,
    name varchar(200) not null ,
    social_media_url text not null ,
    user_id int not null references users(id)
    created_at timestamp
    updated_at timestamp
);

commit;