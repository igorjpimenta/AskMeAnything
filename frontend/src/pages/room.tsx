import amaLogo from '../assets/ama-logo.svg'
import { Messages } from '../components/messages'
import { CreateMessageForm } from '../components/message-form'
import { getRoom } from '../http/get-room'

import { useNavigate, useParams } from "react-router-dom"
import { Eye, EyeOff, Share2 } from "lucide-react"
import { toast } from "sonner"
import { Suspense, useEffect, useState } from "react"
import { Tooltip } from 'react-tooltip'

export function Room() {
    const navigate = useNavigate()
    const { roomId } = useParams()
    const [isOwner, setIsOwner] = useState(false)
    const [showHiddenMessages, setShowHiddenMessages] = useState(true)

    useEffect(() => {
        async function fetchRoomDetails() {
            if (!roomId) {
                navigate(`/`)
                return
            }

            const ownerToken = localStorage.getItem(`owner_token-${roomId}`)
            const { ownership } = await getRoom({ roomId, ownerToken })
            setIsOwner(ownership)
        }

        fetchRoomDetails()
    }, [roomId, navigate])

    function handleShareRoom() {
        const url = window.location.href.toString()

        if (navigator.share != undefined && navigator.canShare()) {
            navigator.share({ url })
            
        } else {
            navigator.clipboard.writeText(url)

            toast.info('The room URL was copied to your clipboard!')
        }
    }
    
    function handleToggleHiddenMessages() {
        setShowHiddenMessages(prev => !prev)
    }

    return (
        <div className="mx-auto max-w-[640px] flex flex-col gap-6 py-10 px-4">
            <div className="flex items-center justify-between px-3">
                <div className="flex gap-3">
                    <img src={amaLogo} alt="Ask Me Anything Logo" className="h-5" />

                    <span className="text-sm text-zinc-500 truncate">
                        Room code: <span className="text-zinc-300">{roomId}</span>
                    </span>
                </div>

                <div className="flex space-x-3">
                    {isOwner && (
                        <>
                            <button
                                data-tooltip-id="tooltip-toggle-hidden-messages"
                                data-tooltip-content={`${showHiddenMessages ? "Hide" : "Show"} hidden messages`}
                                type="button"
                                onClick={handleToggleHiddenMessages}
                                className="ml-3 bg-zinc-800 text-zinc-300 px-3 py-1.5 gap-1.5 flex items-center rounded-lg font-medium text-sm transition-colors hover:bg-zinc-700"
                            >
                                {showHiddenMessages ? <Eye className="size-4"/> : <EyeOff className="size-4"/>}
                            </button>

                            <Tooltip id={"tooltip-toggle-hidden-messages"} place="top" />
                        </>
                    )}
                        
                    <button
                        data-tooltip-id="tooltip-room-sharing"
                        data-tooltip-content="Share room"
                        type="button"
                        onClick={handleShareRoom}
                        className="ml-auto bg-zinc-800 text-zinc-300 px-3 py-1.5 gap-1.5 flex items-center rounded-lg font-medium text-sm transition-colors hover:bg-zinc-700"
                    >
                        {!isOwner && (<span>Share room</span>)}

                        <Share2 className="size-4"/>
                    </button>

                    {isOwner && (<Tooltip id={"tooltip-room-sharing"} place="top" />)}
                    
                </div>
            </div>

            <div className="h-px w-full bg-zinc-900"></div>
            
            <CreateMessageForm />
            
            <Suspense fallback={<p>Loading...</p>}>
                <Messages
                    isOwner={isOwner}
                    showHiddenMessages={showHiddenMessages}
                />
            </Suspense>
        </div>
    )
}
