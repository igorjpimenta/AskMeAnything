package api

import "github.com/go-chi/chi/v5"

func routes(h *apiHandler) chi.Router {
	r := chi.NewRouter()

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

	return r
}
