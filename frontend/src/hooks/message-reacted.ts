import { GetRoomMessagesResponse } from "../http/get-room-messages"
import { sortMessages } from "../utils/sort-messages"

import { QueryClient } from "@tanstack/react-query"

export interface MessageReacted {
    id: string,
    count: number
}

export function handleMessageReacted(
    queryClient: QueryClient,
    roomId: string,
    value: MessageReacted
) {
    queryClient.setQueryData<GetRoomMessagesResponse>(['messages', roomId], state => {
        if (!state) {
            return undefined
        }

        const messages = state.messages.map(item => {
            if (item.id === value.id) {
                return { ...item, amountOfReactions: value.count }
            }

            return item
        })

        return {
            messages: sortMessages(messages),
        }
    })
}
