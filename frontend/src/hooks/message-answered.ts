import { GetRoomMessagesResponse } from "../http/get-room-messages"
import { sortMessages } from "../utils/sort-messages"

import { QueryClient } from "@tanstack/react-query"

export interface MessageAnswered {
    id: string
}

export function handleMessageAnswered(
    queryClient: QueryClient,
    roomId: string,
    value: MessageAnswered
) {
    queryClient.setQueryData<GetRoomMessagesResponse>(['messages', roomId], state => {
        if (!state) {
            return undefined
        }

        const messages = state.messages.map(item => {
            if (item.id === value.id) {
                return { ...item, answered: true }
            }

            return item
        })

        return {
            messages: sortMessages(messages),
        }
    })
}
