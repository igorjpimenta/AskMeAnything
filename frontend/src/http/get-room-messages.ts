import { config } from "../../config"
import { sortMessages } from "../utils/sort-messages"

interface GetRoomMessagesRequest {
    roomId: string
}

export interface GetRoomMessagesResponse {
    messages: {
        id: string
        text: string
        amountOfReactions: number
        answered: boolean
    }[]
}

export async function getRoomMessages({ roomId }: GetRoomMessagesRequest): Promise<GetRoomMessagesResponse> {
    const response = await fetch(`${config.API_URL}/api/rooms/${roomId}/messages`)

    const data: Array<{
        ID: string
        RoomID: string
        Message: string
        ReactionCount: number
        Answered: boolean
    }> = await response.json()

    const messages = data.map(item => ({
        id: item.ID,
        text: item.Message,
        amountOfReactions: item.ReactionCount,
        answered: item.Answered,
    }))

    return {
        messages: sortMessages(messages)
    }
}