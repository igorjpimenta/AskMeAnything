import { config } from "../../config"

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
        id: string
        room_id: string
        message: string
        reaction_count: number
        answered: boolean
    }> = await response.json()

    return {
        messages: data.map(item => ({
            id: item.id,
            text: item.message,
            amountOfReactions: item.reaction_count,
            answered: item.answered,
        })),
    }
}