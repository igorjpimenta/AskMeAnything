import { reactToMessage } from "../http/react-to-message"
import { removeReactFromMessage } from "../http/remove-react-from-message"

import { ArrowUp } from "lucide-react"
import { useState } from "react"
import { useParams } from "react-router-dom"
import { toast } from "sonner"

interface MessageProps {
    id: string
    text: string
    amountOfReactions: number
    answered?: boolean
}

export function Message({
    id: messageId,
    text,
    amountOfReactions,
    answered=false
}: MessageProps) {
    const [hasReacted, setHasReacted] = useState(false)
    const { roomId } = useParams()

    async function handleReactToMessage() {
        if (!roomId) {
            throw new Error('Message components must be used within room page')
        }

        if (!messageId) {
            return
        }

        try {
            await reactToMessage({ roomId, messageId })
            setHasReacted(true)

        } catch {
            toast.error('Error reacting to message!')
        }
    }

    async function handleRemoveReactFromMessage() {
        if (!roomId) {
            return
        }

        if (!messageId) {
            return
        }

        try {
            await removeReactFromMessage({ roomId, messageId })
            setHasReacted(false)

        } catch {
            toast.error('Error removing react from message!')
        }
    }

    return (
        <li data-answered={answered} className="ml-4 leading-relaxed text-zinc-100 data-[answered=true]:opacity-50 data-[answered=true]:pointer-events-none">
            {text}

            {hasReacted ? (
                <button
                    onClick={handleRemoveReactFromMessage}
                    type="button"
                    className="mt-3 flex items-center gap-2 select-none text-orange-400 text-sm font-medium hover:text-orange-500"
                >
                    <ArrowUp className="size-4" />
                    Like question ({amountOfReactions})
                </button>
            ) : (
                <button
                    onClick={handleReactToMessage}
                    type="button"
                    className="mt-3 flex items-center gap-2 select-none text-zinc-400 text-sm font-medium hover:text-zinc-300"
                >
                    <ArrowUp className="size-4" />
                    Like question ({amountOfReactions})
                </button>
            )}
        </li>
    )
}