import { GetRoomMessagesResponse } from "../http/get-room-messages"
import { sortMessages } from "../utils/sort-messages"

import { QueryClient } from "@tanstack/react-query"

export interface MessageCreated {
    id: string,
    message: string
}

export function handleMessageCreated(
    queryClient: QueryClient,
    roomId: string,
    value: MessageCreated
) {
    queryClient.setQueryData<GetRoomMessagesResponse>(['messages', roomId], state => {
        const messages = [
            ...(state?.messages ?? []),
            {
                id: value.id,
                text: value.message,
                amountOfReactions: 0,
                answered: false,
            },
        ]

        return {
            messages: sortMessages(messages),
        }
    })
}
