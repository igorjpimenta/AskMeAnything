import { Message } from "./message";
import { getRoomMessages } from "../http/get-room-messages"

import { useParams } from "react-router-dom"
import { useSuspenseQuery } from "@tanstack/react-query"
import { useWebSocketMessages } from "../hooks/use-websocket-messages"

interface MessagesProps {
    isOwner: boolean
}

export function Messages({ isOwner }: MessagesProps) {
    const { roomId } = useParams()

    if (!roomId) {
        throw new Error('Messages components must be used within room page')
    }

    const { data } = useSuspenseQuery({
        queryKey: ['messages', roomId],
        queryFn: () => getRoomMessages({ roomId })
    })

    useWebSocketMessages(roomId)

    const sortedMessages = data.messages.sort((a, b) => {
        if (b.answered !== a.answered) {
            return a.answered ? 1 : -1
        }
        return b.amountOfReactions - a.amountOfReactions
    })


    return (
        <ol className="list-decimal list-outside px-3 space-y-8">
            {sortedMessages.map(message => {
                return (
                    <Message
                        key={message.id}
                        id={message.id}
                        text={message.text}
                        amountOfReactions={message.amountOfReactions}
                        answered={message.answered}
                        isOwner={isOwner}
                    />
                )
            })}
        </ol>
   ) 
}