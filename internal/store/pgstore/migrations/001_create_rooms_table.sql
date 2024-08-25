create table if not exists tb_rooms (
    id    uuid            primary key     not null    default gen_random_uuid(),
    theme varchar(255)                    not null
);

---- create above / drop bellow ----

-- drop table if exists tb_rooms;