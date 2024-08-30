package api

const (
	MessageKindMessageCreated           = "message_created"
	MessageKindMessageReactionIncreased = "message_reaction_increased"
	MessageKindMessageReactionDecreased = "message_reaction_decreased"
	MessageKindMessageAnswered          = "message_answered"
	MessageKindMessageUnanswered        = "message_unanswered"
)

type MessageMessageUnanswered struct {
	ID string `json:"id"`
}

type MessageMessageAnswered struct {
	ID string `json:"id"`
}

type MessageMessageReactionDecreased struct {
	ID    string `json:"id"`
	Count int64  `json:"count"`
}

type MessageMessageReactionIncreased struct {
	ID    string `json:"id"`
	Count int64  `json:"count"`
}

type MessageMessageCreated struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type Message struct {
	Kind   string `json:"kind"`
	Value  any    `json:"value"`
	RoomID string `json:"-"`
}
