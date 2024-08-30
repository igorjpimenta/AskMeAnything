import { config } from "../../config"
import { handleMessageCreated, MessageCreated } from "./message-created"
import { handleMessageReacted, MessageReacted } from "./message-reacted"
import { handleMessageAnswered, MessageAnswered } from "./message-answered"
import { handleMessageUnanswered, MessageUnanswered } from "./message-unanswered"
import { handleMessageUnhidden, MessageUnhidden } from "./message-unhidden"
import { handleMessageHidden, MessageHidden } from "./message-hidden"

import { useQueryClient } from "@tanstack/react-query"
import { useEffect } from "react"

type WebSocketMessage =
    | { kind: 'message_created'; value: MessageCreated }
    | { kind: 'message_reaction_increased'; value: MessageReacted }
    | { kind: 'message_reaction_decreased'; value: MessageReacted }
    | { kind: 'message_answered'; value: MessageAnswered }
    | { kind: 'message_unanswered'; value: MessageUnanswered}
    | { kind: 'message_hidden'; value: MessageHidden }
    | { kind: 'message_unhidden'; value: MessageUnhidden}

export function useWebSocketMessages(roomId: string) {
    const queryClient = useQueryClient()

    useEffect(() => {
        const ws = new WebSocket(`${config.WS_URL}/subscribe/${roomId}`)
    
        ws.onopen = () => {
            console.log('Websocket connected!')
        }
    
        ws.onmessage = (event) => {
            const data: WebSocketMessage = JSON.parse(event.data)
    
            switch(data.kind) {
                case 'message_created':
                    handleMessageCreated(queryClient, roomId, data.value)
                    break
                
                case 'message_reaction_increased':
                case 'message_reaction_decreased':
                    handleMessageReacted(queryClient, roomId, data.value)
                    break
    
                case 'message_answered':
                    handleMessageAnswered(queryClient, roomId, data.value)
                    break
    
                case 'message_unanswered':
                    handleMessageUnanswered(queryClient, roomId, data.value)
                    break
    
                case 'message_hidden':
                    handleMessageHidden(queryClient, roomId, data.value)
                    break
    
                case 'message_unhidden':
                    handleMessageUnhidden(queryClient, roomId, data.value)
                    break
            }
        }
    
        ws.onclose = () => {
            console.log('Websocket connection closed!')
        }

        return () => {
            ws.close()
        }
    }, [roomId, queryClient])
}