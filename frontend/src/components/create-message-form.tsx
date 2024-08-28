import { createMessage } from "../http/create-message";

import { ArrowRight } from "lucide-react";
import { useParams } from "react-router-dom";
import { toast } from "sonner";

export function CreateMessageForm() {
    const { roomId } = useParams()

    async function handleCreateMessage(data: FormData) {
        const message = data.get('message')?.toString()

        if (!roomId) {
            return
        }

        if (!message) {
            toast.error('Write a message.')
            return
        }

        try {
            await createMessage({ roomId, message })

            toast.success('Message created!')

        } catch {
            toast.error('Error posting a message!')
        }
    }

    return (
        <form
        action={handleCreateMessage}
        className="flex items-center gap-2 bg-zinc-900 p-2 rounded-xl border border-zinc-800 ring-orange-400 ring-offset-2 ring-offset-zinc-950 focus-within:ring-1"
    >
        <input
            type="text"
            name="message"
            placeholder="What do you want to ask?"
            autoComplete="off"
            className="flex-1 text-sm bg-transparent mx-2 outline-none text-zinc-100 placeholder:text-zinc-500"
        />
        
        <button
            type="submit"
            className="bg-orange-400 text-orange-950 px-3 py-1.5 gap-1.5 flex items-center rounded-lg font-medium text-sm transition-colors hover:bg-orange-500"
        >
            Make a question
            <ArrowRight className="size-4"/>
        </button>
    </form>
    )
}