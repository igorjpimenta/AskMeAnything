package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

func (h apiHandler) handleSubscribe(w http.ResponseWriter, r *http.Request) {
	_, roomID, ok := h.getRoom(w, r)
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

func (h apiHandler) notifyClients(msg Message) {
	h.mu.Lock()
	defer h.mu.Unlock()

	subscribers, ok := h.subscribers[msg.RoomID]
	if !ok || len(subscribers) == 0 {
		return
	}

	for conn, cancel := range subscribers {
		if err := conn.WriteJSON(msg); err != nil {
			slog.Error("failed to send message to client", "error", err)
			cancel()
		}
	}
}
