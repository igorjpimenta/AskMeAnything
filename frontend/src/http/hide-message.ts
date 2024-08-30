import { config } from "../../config";

interface HideMessageRequest {
    roomId: string
    messageId: string
    ownerToken: string
}

export async function hideMessage({ roomId, messageId, ownerToken }: HideMessageRequest) {
    await fetch(`${config.API_URL}/api/rooms/${roomId}/messages/${messageId}/hide`, {
        method: 'PATCH',
        headers: { 'Owner-Token': ownerToken },
    })
}