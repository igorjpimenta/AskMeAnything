// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package pgstore

import (
	"context"

	"github.com/google/uuid"
)

const getMessage = `-- name: GetMessage :one
select
    "id", "room_id", "message", "reaction_count", "answered", "hidden"
from messages
where
    id = $1
`

func (q *Queries) GetMessage(ctx context.Context, id uuid.UUID) (Message, error) {
	row := q.db.QueryRow(ctx, getMessage, id)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.RoomID,
		&i.Message,
		&i.ReactionCount,
		&i.Answered,
		&i.Hidden,
	)
	return i, err
}

const getRoom = `-- name: GetRoom :one
select
    "id", "theme", "owner_token"
from rooms
where id = $1
`

func (q *Queries) GetRoom(ctx context.Context, id uuid.UUID) (Room, error) {
	row := q.db.QueryRow(ctx, getRoom, id)
	var i Room
	err := row.Scan(&i.ID, &i.Theme, &i.OwnerToken)
	return i, err
}

const getRoomMessages = `-- name: GetRoomMessages :many
select
    "id", "room_id", "message", "reaction_count", "answered", "hidden"
from messages
where
    room_id = $1
`

func (q *Queries) GetRoomMessages(ctx context.Context, roomID uuid.UUID) ([]Message, error) {
	rows, err := q.db.Query(ctx, getRoomMessages, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.RoomID,
			&i.Message,
			&i.ReactionCount,
			&i.Answered,
			&i.Hidden,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRooms = `-- name: GetRooms :many
select
    "id", "theme", "owner_token"
from rooms
`

func (q *Queries) GetRooms(ctx context.Context) ([]Room, error) {
	rows, err := q.db.Query(ctx, getRooms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Room
	for rows.Next() {
		var i Room
		if err := rows.Scan(&i.ID, &i.Theme, &i.OwnerToken); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const hideMessage = `-- name: HideMessage :exec
update messages
set
    hidden = true
where
    id = $1
`

func (q *Queries) HideMessage(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, hideMessage, id)
	return err
}

const insertMessage = `-- name: InsertMessage :one
insert into messages("room_id", "message")
    values
        ($1, $2)
returning "id"
`

type InsertMessageParams struct {
	RoomID  uuid.UUID `db:"room_id" json:"room_id"`
	Message string    `db:"message" json:"message"`
}

func (q *Queries) InsertMessage(ctx context.Context, arg InsertMessageParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, insertMessage, arg.RoomID, arg.Message)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const insertRoom = `-- name: InsertRoom :one
insert into rooms("theme", "owner_token")
    values
        ($1, $2)
returning "id"
`

type InsertRoomParams struct {
	Theme      string    `db:"theme" json:"theme"`
	OwnerToken uuid.UUID `db:"owner_token" json:"owner_token"`
}

func (q *Queries) InsertRoom(ctx context.Context, arg InsertRoomParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, insertRoom, arg.Theme, arg.OwnerToken)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const markMessageAnswered = `-- name: MarkMessageAnswered :exec
update messages
set
    answered = true
where
    id = $1
`

func (q *Queries) MarkMessageAnswered(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, markMessageAnswered, id)
	return err
}

const markMessageUnanswered = `-- name: MarkMessageUnanswered :exec
update messages
set
    answered = false
where
    id = $1
`

func (q *Queries) MarkMessageUnanswered(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, markMessageUnanswered, id)
	return err
}

const reactToMessage = `-- name: ReactToMessage :one
update messages
set
    reaction_count = reaction_count + 1
where
    id = $1
returning reaction_count
`

func (q *Queries) ReactToMessage(ctx context.Context, id uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, reactToMessage, id)
	var reaction_count int64
	err := row.Scan(&reaction_count)
	return reaction_count, err
}

const removeReactionFromMessage = `-- name: RemoveReactionFromMessage :one
update messages
set
    reaction_count = reaction_count - 1
where
    id = $1
returning reaction_count
`

func (q *Queries) RemoveReactionFromMessage(ctx context.Context, id uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, removeReactionFromMessage, id)
	var reaction_count int64
	err := row.Scan(&reaction_count)
	return reaction_count, err
}

const unhideMessage = `-- name: UnhideMessage :exec
update messages
set
    hidden = false
where
    id = $1
`

func (q *Queries) UnhideMessage(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, unhideMessage, id)
	return err
}
