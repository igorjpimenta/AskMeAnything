package api

import (
	"context"
	"log/slog"
	"net/http"
	"sync"

	"github.com/igorjpimenta/AskMeAnything/internal/store/pgstore"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/gorilla/websocket"
)

type apiHandler struct {
	q           *pgstore.Queries
	r           *chi.Mux
	upgrader    websocket.Upgrader
	subscribers map[string]map[*websocket.Conn]context.CancelFunc
	mu          *sync.Mutex
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	h := apiHandler{
		q:           q,
		upgrader:    websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }},
		subscribers: make(map[string]map[*websocket.Conn]context.CancelFunc),
		mu:          &sync.Mutex{},
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/subscribe/{room_id}", h.handleSubscribe)

	r.Route("/api", func(r chi.Router) {
		r.Route("/rooms", func(r chi.Router) {
			r.Post("/", h.handleCreateRoom)
			r.Get("/", h.handleGetRooms)

			r.Route("/{room_id}/messages", func(r chi.Router) {
				r.Post("/", h.handleCreateRoomMessage)
				r.Get("/", h.handleGetRoomMessages)

				r.Route("/{message_id}", func(r chi.Router) {
					r.Get("/", h.handleGetRoomMessage)
					r.Patch("/react", h.handleReactToMessage)
					r.Delete("/react", h.handleRemoveReactFromMessage)
					r.Patch("/answer", h.handleMaskMessageAsAnswered)
				})
			})
		})
	})

	h.r = r
	return h
}

const (
	MessageKindMessageCreated = "message_created"
)

type MessageMessageCreated struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type Message struct {
	Kind   string `json:"kind"`
	Value  any    `json:"value"`
	RoomID string `json:"-"`
}

func (h apiHandler) handleSubscribe(w http.ResponseWriter, r *http.Request) {
	roomID, ok := h.getRoom(w, r)
	if !ok {
		return
	}

	c, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.handleError(w, "failed to upgrade to ws connection", err, http.StatusBadRequest)
		return
	}

	defer c.Close()

	ctx, cancel := context.WithCancel(r.Context())
	h.addSubscriber(roomID.String(), c, cancel)
	defer h.removeSubscriber(roomID.String(), c)

	<-ctx.Done()
}

func (h *apiHandler) addSubscriber(roomID string, conn *websocket.Conn, cancel context.CancelFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.subscribers[roomID]; !ok {
		h.subscribers[roomID] = make(map[*websocket.Conn]context.CancelFunc)
	}
	h.subscribers[roomID][conn] = cancel
	slog.Info("new client connected", "room_id", roomID, "client_ip", conn.RemoteAddr().String())
}

func (h *apiHandler) removeSubscriber(roomID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.subscribers[roomID], conn)
}

func (h apiHandler) handleCreateRoom(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Theme string `json:"theme"`
	}
	if err := h.decodeJSONBody(w, r, &body); err != nil {
		return
	}

	roomID, err := h.q.InsertRoom(r.Context(), body.Theme)
	if err != nil {
		h.handleError(w, "failed to insert room", err, http.StatusInternalServerError)
		return
	}

	type response struct {
		ID string `json:"id"`
	}

	h.respondWithJSON(w, http.StatusOK, response{ID: roomID.String()})
}

func (h apiHandler) handleGetRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.q.GetRooms(r.Context())
	if err != nil {
		h.handleError(w, "failed to get rooms", err, http.StatusInternalServerError)
		return
	}

	if rooms == nil {
		rooms = []pgstore.Room{}
	}

	h.respondWithJSON(w, http.StatusOK, rooms)
}

func (h apiHandler) handleCreateRoomMessage(w http.ResponseWriter, r *http.Request) {
	roomID, ok := h.getRoom(w, r)
	if !ok {
		return
	}

	var body struct {
		Message string `json:"message"`
	}
	if err := h.decodeJSONBody(w, r, &body); err != nil {
		return
	}

	messageID, err := h.q.InsertMessage(r.Context(), pgstore.InsertMessageParams{RoomID: roomID, Message: body.Message})
	if err != nil {
		h.handleError(w, "failed to insert message", err, http.StatusInternalServerError)
		return
	}

	type response struct {
		ID string `json:"id"`
	}

	h.respondWithJSON(w, http.StatusOK, response{ID: messageID.String()})

	go h.notifyClients(Message{
		Kind:   MessageKindMessageCreated,
		RoomID: roomID.String(),
		Value: MessageMessageCreated{
			ID:      messageID.String(),
			Message: body.Message,
		},
	})
}
func (h apiHandler) handleGetRoomMessages(w http.ResponseWriter, r *http.Request) {
	roomID, ok := h.getRoom(w, r)
	if !ok {
		return
	}

	messages, err := h.q.GetRoomMessages(r.Context(), roomID)
	if err != nil {
		h.handleError(w, "failed to get messages from the room", err, http.StatusInternalServerError)
		return
	}

	if messages == nil {
		messages = []pgstore.Message{}
	}

	h.respondWithJSON(w, http.StatusOK, messages)
}
func (h apiHandler) handleGetRoomMessage(w http.ResponseWriter, r *http.Request)         {}
func (h apiHandler) handleReactToMessage(w http.ResponseWriter, r *http.Request)         {}
func (h apiHandler) handleRemoveReactFromMessage(w http.ResponseWriter, r *http.Request) {}
func (h apiHandler) handleMaskMessageAsAnswered(w http.ResponseWriter, r *http.Request)  {}
