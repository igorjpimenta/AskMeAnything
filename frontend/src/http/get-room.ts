import { config } from "../../config"

interface GetRoomRequest {
    roomId: string
    ownerToken?: string | null
}

export async function getRoom({ roomId, ownerToken }: GetRoomRequest) {
    if (!ownerToken) {
        ownerToken = ''
    }

    const response = await fetch(`${config.API_URL}/api/rooms/${roomId}`, {
        method: 'GET',
        headers: { 'Owner-Token': ownerToken },
    })

    const data: {
        id: string
        theme: string
        ownership: boolean
    } = await response.json()

    return {
        theme: data.theme,
        ownership: data.ownership,
    }
}