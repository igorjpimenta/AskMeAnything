import { GetRoomMessagesResponse } from "../http/get-room-messages"

import { QueryClient } from "@tanstack/react-query"

export interface MessageUnhidden {
    id: string
}

export function handleMessageUnhidden(
    queryClient: QueryClient,
    roomId: string,
    value: MessageUnhidden
) {
    queryClient.setQueryData<GetRoomMessagesResponse>(['messages', roomId], state => {
        if (!state) {
            return undefined
        }

        return {
            messages: state.messages.map(item => {
                if (item.id === value.id) {
                    return { ...item, hidden: false }
                }
    
                return item
            }),
        }
    })
}
