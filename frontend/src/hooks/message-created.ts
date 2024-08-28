import { GetRoomMessagesResponse } from "../http/get-room-messages"

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
        return {
            messages: [
                ...(state?.messages ?? []),
                {
                    id: value.id,
                    text: value.message,
                    amountOfReactions: 0,
                    answered: false,
                },
            ],
        }
    })
}
