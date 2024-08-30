import { GetRoomMessagesResponse } from "../http/get-room-messages"

import { QueryClient } from "@tanstack/react-query"

export interface MessageUnanswered {
    id: string
}

export function handleMessageUnanswered(
    queryClient: QueryClient,
    roomId: string,
    value: MessageUnanswered
) {
    queryClient.setQueryData<GetRoomMessagesResponse>(['messages', roomId], state => {
        if (!state) {
            return undefined
        }

        return {
            messages: state.messages.map(item => {
                if (item.id === value.id) {
                    return { ...item, answered: false }
                }
    
                return item
            }),
        }
    })
}
