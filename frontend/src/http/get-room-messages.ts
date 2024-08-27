import { config } from "../../config";

interface GetRoomMessagesRequest {
    roomId: string
}

export async function getRoomMessages({ roomId }: GetRoomMessagesRequest) {
    const response = await fetch(`${config.API_URL}/api/rooms/${roomId}/messages`)

    const data: Array<{
        ID: string
        RoomID: string
        Message: string
        ReactionCount: number
        Answered: boolean
    }> = await response.json()

    return {
        messages: data
            .map(item => ({
                id: item.ID,
                text: item.Message,
                amountOfReactions: item.ReactionCount,
                answered: item.Answered,
            }))
            .sort((a, b) => {
                if (b.answered !== a.answered) {
                    return a.answered ? 1 : -1
                }

                return b.amountOfReactions - a.amountOfReactions
            })
    }
}