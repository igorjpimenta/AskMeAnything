import { config } from "../../config";

interface ReactToMessageRequest {
    roomId: string
    messageId: string
}

export async function reactToMessage({ roomId, messageId }: ReactToMessageRequest) {
    const response = await fetch(`${config.API_URL}/api/rooms/${roomId}/messages/${messageId}/react`, {
        method: 'PATCH',
    })

    const data: { count: number } = await response.json()

    return { amountOfReactions: data.count }
}