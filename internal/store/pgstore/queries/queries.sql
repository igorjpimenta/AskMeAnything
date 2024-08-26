-- name: GetRoom :one
select
    "id", "theme"
from rooms
where id = $1;

-- name: GetRooms :many
select
    "id", "theme"
from rooms;

-- name: InsertRoom :one
insert into rooms("theme")
    values
        ($1)
returning "id";

-- name: GetMessage :one
select
    "id", "room_id", "message", "reaction_count", "answered"
from messages
where
    id = $1;

-- name: GetRoomMessages :many
select
    "id", "room_id", "message", "reaction_count", "answered"
from messages
where
    room_id = $1;

-- name: InsertMessage :one
insert into messages("room_id", "message")
    values
        ($1, $2)
returning "id";

-- name: ReactToMessage :one
update messages
set
    reaction_count = reaction_count + 1
where
    id = $1
returning reaction_count;

-- name: RemoveReactionFromMessage :one
update messages
set
    reaction_count = reaction_count - 1
where
    id = $1
returning reaction_count;

-- name: MarkMessageAnswered :exec
update messages
set
    answered = true
where
    id = $1;