-- name: GetRoom :one
select
    "id", "theme"
from tb_rooms
where id = $1;

-- name: GetRooms :many
select
    "id", "theme"
from tb_rooms;

-- name: InsertRoom :one
insert into tb_rooms("theme")
    values
        ($1)
returning "id";

-- name: GetMessage :one
select
    "id", "room_id", "message", "reaction_count", "answered"
from tb_messages
where
    id = $1;

-- name: GetRoomMessages :many
select
    "id", "room_id", "message", "reaction_count", "answered"
from tb_messages
where
    room_id = $1;

-- name: InsertMessage :one
insert into tb_messages("room_id", "message")
    values
        ($1, $2)
returning "id";

-- name: ReactToMessage :one
update tb_messages
set
    reaction_count = rection_count + 1
where
    id = $1
returning reaction_count;

-- name: RemoveReactionFromMessage :one
update tb_messages
set
    reaction_count = rection_count - 1
where
    id = $1
returning reaction_count;

-- name: MarkMessageAnswered :exec
update tb_messages
set
    answered = true
where
    id = $1;