import { cn } from "@/lib/utils"
import { Bot } from "lucide-react"
import { useMemo } from "react"
import ReactMarkdown from "react-markdown"
import Pre from "../Pre"


export default function ChatEvent({ event }) {
    if (!event) {
        return null
    }

    const isUser = event.Author === "user"
    const text = useMemo(() => event.Content.parts.reduce((acc, val) => acc.concat(val.text), ""), [event])
    if (text == "" || text == undefined) {
        return <p>No text</p>
    }
    return (
        <div className={cn(
            "flex w-full gap-4 p-2 transition-colors",
            isUser && "flex-row-reverse"

        )}>
            <div className={cn("flex gap-2", isUser && "flex-row-reverse")}>
                {!isUser && <div className={cn(
                    "flex h-8 w-8 shrink-0 select-none items-center justify-center rounded-md border shadow-sm mt-2",
                    isUser ? "bg-white text-slate-900" : "bg-indigo-600 text-white"
                )}>
                    {<Bot size={18} />}
                </div>
                }
                <div className={cn("flex-1 space-y-2 overflow-auto p-4 rounded-3xl rounded-tr-lg max-w-full ", isUser && "bg-gray-50/50 flex-row-reverse  border")}>
                    <div className="prose prose-slate dark:prose-invert max-w-full overflow-auto">
                        <ReactMarkdown components={{pre: Pre}} >{text}</ReactMarkdown>
                    </div>
                </div>
            </div>
        </div>
    )
}