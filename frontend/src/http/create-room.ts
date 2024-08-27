import { config } from "../../config";

interface CreateRoomRequest {
    theme: string
}

export async function createRoom({ theme }: CreateRoomRequest) {
    const response = await fetch(`${config.API_URL}/api/rooms`, {
        method: 'POST',
        body: JSON.stringify({
            theme
        })
    })

    const data: { id: string } = await response.json()

    return { roomId: data.id }
}