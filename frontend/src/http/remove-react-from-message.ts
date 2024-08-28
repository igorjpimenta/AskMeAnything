import { config } from "../../config";

interface RemoveReactFromMessageRequest {
    roomId: string
    messageId: string
}

export async function removeReactFromMessage({ roomId, messageId }: RemoveReactFromMessageRequest) {
    const response = await fetch(`${config.API_URL}/api/rooms/${roomId}/messages/${messageId}/react`, {
        method: 'DELETE',
    })

    const data: { count: number } = await response.json()

    return { amountOfReactions: data.count }
}