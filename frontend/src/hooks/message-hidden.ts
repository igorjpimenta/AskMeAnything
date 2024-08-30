import { GetRoomMessagesResponse } from "../http/get-room-messages"

import { QueryClient } from "@tanstack/react-query"

export interface MessageHidden {
    id: string
}

export function handleMessageHidden(
    queryClient: QueryClient,
    roomId: string,
    value: MessageHidden
) {
    queryClient.setQueryData<GetRoomMessagesResponse>(['messages', roomId], state => {
        if (!state) {
            return undefined
        }

        return {
            messages: state.messages.map(item => {
                if (item.id === value.id) {
                    return { ...item, hidden: true }
                }
    
                return item
            }),
        }
    })
}
