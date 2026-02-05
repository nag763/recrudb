import { cn } from "@/lib/utils"
import { Bot, User } from "lucide-react"
import { useMemo } from "react"
import ReactMarkdown from "react-markdown"

export default function ChatEvent({ event }) {
    if (!event) {
        return null
    }

    const isUser = event.Author === "user"
    const text = useMemo(() => event.Content.parts.reduce((acc, val) => acc.concat(val.text), ""), [event])

    return (
        <div className={cn(
            "flex w-full gap-4 p-6 transition-colors",
            isUser ? "bg-white flex-row-reverse" : "bg-slate-50/50"
        )}>
            <div className={cn("flex gap-2 items-center", isUser && "flex-row-reverse")}>
                <div className={cn(
                    "flex h-8 w-8 shrink-0 select-none items-center justify-center rounded-md border shadow-sm",
                    isUser ? "bg-white text-slate-900" : "bg-indigo-600 text-white"
                )}>
                    {isUser ? <User size={18} /> : <Bot size={18} />}
                </div>
                <div className="flex-1 space-y-2 overflow-hidden px-1">
                    <div className="prose prose-slate max-w-none break-words dark:prose-invert">
                        <ReactMarkdown>{text}</ReactMarkdown>
                    </div>
                </div>
            </div>
        </div>
    )
}