import { config } from "../../config";

interface MarkMessageUnansweredRequest {
    roomId: string
    messageId: string
    ownerToken: string
}

export async function markMessageUnanswered({ roomId, messageId, ownerToken }: MarkMessageUnansweredRequest) {
    await fetch(`${config.API_URL}/api/rooms/${roomId}/messages/${messageId}/answer`, {
        method: 'DELETE',
        headers: { 'Owner-Token': ownerToken },
    })
}
