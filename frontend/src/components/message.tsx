import { reactToMessage } from "../http/react-to-message"
import { removeReactFromMessage } from "../http/remove-react-from-message"
import { markMessageAnswered } from "../http/mark-message-answered"
import { markMessageUnanswered } from "../http/mark-message-unanswered"

import { ArrowUp, CheckCircle, CircleSlash } from "lucide-react"
import { useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import { toast } from "sonner"
import { Tooltip } from "react-tooltip"

interface MessageProps {
    id: string
    text: string
    amountOfReactions: number
    answered: boolean
    isOwner: boolean
}

export function Message({
    id: messageId,
    text,
    amountOfReactions,
    answered=false,
    isOwner,
}: MessageProps) {
    const [hasReacted, setHasReacted] = useState(false)
    const [hasAnswered, setHasAnswered] = useState(answered)
    const { roomId } = useParams()

    if (!roomId) {
        throw new Error('Message components must be used within room page')
    }

    useEffect(() => {
        const storedReaction = localStorage.getItem(`reacted-${roomId}-${messageId}`)
        if (storedReaction) {
            setHasReacted(true)
        }
    }, [roomId, messageId])

    async function handleReactToMessage() {
        if (!roomId || !messageId) {
            return
        }

        try {
            await reactToMessage({ roomId, messageId })
            localStorage.setItem(`reacted-${roomId}-${messageId}`, 'true')
            setHasReacted(true)

        } catch {
            toast.error('Error reacting to message!')
        }
    }

    async function handleRemoveReactFromMessage() {
        if (!roomId || !messageId) {
            return
        }

        try {
            await removeReactFromMessage({ roomId, messageId })
            localStorage.removeItem(`reacted-${roomId}-${messageId}`)
            setHasReacted(false)

        } catch {
            toast.error('Error removing react from message!')
        }
    }

    async function handleMarkAsAnswered() {
        if (!roomId || !messageId || !isOwner) {
            return
        }

        const ownerToken = localStorage.getItem(`owner_token-${roomId}`)
        if (!ownerToken) {
            return
        }

        try {
            await markMessageAnswered({ roomId, messageId, ownerToken })
            setHasAnswered(true)

        } catch {
            toast.error('Error marking message as answered!')
        }
    }

    async function handleMarkAsUnanswered() {
        if (!roomId || !messageId || !isOwner) {
            return
        }

        const ownerToken = localStorage.getItem(`owner_token-${roomId}`)
        if (!ownerToken) {
            return
        }

        try {
            await markMessageUnanswered({ roomId, messageId, ownerToken })
            setHasAnswered(false)

        } catch {
            toast.error('Error marking message as unanswered!')
        }
    }

    return (
        <div className="flex justify-between items-center">
            <li
                data-answered={hasAnswered}
                className="relative ml-4 leading-relaxed text-zinc-100 data-[answered=true]:opacity-50 data-[answered=true]:pointer-events-none"
            >
                {text}

                {<button
                    data-reacted={hasReacted}
                    onClick={hasReacted ? handleRemoveReactFromMessage : handleReactToMessage}
                    type="button"
                    className="mt-3 flex items-center gap-2 select-none text-sm font-medium data-[reacted=true]:text-orange-400 data-[reacted=true]:hover:text-orange-500 data-[reacted=false]:text-zinc-400 data-[reacted=false]:hover:text-zinc-300"
                >
                    <ArrowUp className="size-4" />
                    Like question ({amountOfReactions})
                </button>}
            </li>

            {isOwner && (
                <div>
                    <button
                        data-answered={hasAnswered}
                        onClick={!hasAnswered ? handleMarkAsAnswered : handleMarkAsUnanswered}
                        type="button"
                        className="gap-2 text-sm font-medium pointer-events-auto opacity-100 data-[answered=false]:text-green-500 data-[answered=false]:hover:text-green-600 data-[answered=true]:text-yellow-500 data-[answered=true]:hover:text-yellow-600"
                        data-tooltip-content={`Mark as ${hasAnswered ? "Unanswered" : "Answered"}`}
                        data-tooltip-id={`tooltip-change-answered-state-${messageId}`}
                    >
                        {!hasAnswered ? <CheckCircle className="size-4" /> : <CircleSlash className="size-4" />}
                    </button>
                    
                    <Tooltip id={`tooltip-change-answered-state-${messageId}`} place="top" />
                </div>
            )}
        </div>
    )
}