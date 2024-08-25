create table if not exists tb_messages(
    id                uuid            primary key     not null    default gen_random_uuid(),
    room_id           uuid                            not null,
    message           varchar(255)                    not null,
    reaction_count    bigint                          not null    default 0,
    answered          boolean                         not null    default false,
    foreign key (room_id) references tb_rooms(id)
);

---- create above / drop bellow ----

-- drop table if exists tb_messages;