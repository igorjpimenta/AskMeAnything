import { config } from "../../config";

interface UnhideMessageRequest {
    roomId: string
    messageId: string
    ownerToken: string
}

export async function unhideMessage({ roomId, messageId, ownerToken }: UnhideMessageRequest) {
    await fetch(`${config.API_URL}/api/rooms/${roomId}/messages/${messageId}/hide`, {
        method: 'DELETE',
        headers: { 'Owner-Token': ownerToken },
    })
}