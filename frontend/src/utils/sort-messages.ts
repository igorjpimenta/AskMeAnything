import { GetRoomMessagesResponse } from "../http/get-room-messages"

export function sortMessages(messages: GetRoomMessagesResponse['messages']): GetRoomMessagesResponse['messages'] {
    return messages.sort((a, b) => {
        if (b.answered !== a.answered) {
            return a.answered ? 1 : -1
        }
        return b.amountOfReactions - a.amountOfReactions
    })
}
