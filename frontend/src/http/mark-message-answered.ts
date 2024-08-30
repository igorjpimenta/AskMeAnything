import { config } from "../../config";

interface MarkMessageAnsweredRequest {
    roomId: string
    messageId: string
    ownerToken: string
}

export async function markMessageAnswered({ roomId, messageId, ownerToken }: MarkMessageAnsweredRequest) {
    await fetch(`${config.API_URL}/api/rooms/${roomId}/messages/${messageId}/answer`, {
        method: 'PATCH',
        headers: { 'Owner-Token': ownerToken },
    })
}