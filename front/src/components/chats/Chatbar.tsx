import { Loader2, SendIcon } from "lucide-react";
import { Button } from "../ui/button";
import { Label } from "../ui/label";
import { Textarea } from "../ui/textarea";
import { useCallback, useMemo, useRef, useState } from "react";
import { cn } from "@/lib/utils";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { sendToCoordinator } from "@/interactions/agents";
import { useNavigate, useSearch } from "@tanstack/react-router";
import { toast } from "sonner";

export default function Chatbar({ messagesShown, className }) {

    const queryClient = useQueryClient();

    const navigate = useNavigate()

    const { id: sessionId } = useSearch({ from: "/chat/" })

    const [input, setInput] = useState("");
    const textareaRef = useRef<HTMLTextAreaElement>(null);

    const updateInput = useCallback((e) => {
        setInput(e.target.value);
    }, [setInput]);

    const { mutate, isPending: submitIsPending } = useMutation({
        mutationFn: async (content: string) => sendToCoordinator({ content, sessionId }),
        onSuccess: (resp) => {
            if (!sessionId) {
                navigate({ to: "/chat", search: { id: resp?.data.sessionId } })
                queryClient.invalidateQueries({ queryKey: [`profile-sessions`] });
            }
            setInput("");
            queryClient.invalidateQueries({ queryKey: [`profile-session-${sessionId}`] });
            textareaRef.current?.focus();
        },
        onError: (error) => {
            console.error(error);
            toast.error("The message couldn't be sent");
        },
    });


    const inputEmpty = useMemo(() => input.length == 0, [input]);

    const handleSubmit = useCallback((e) => {
        e.preventDefault();
        if (!submitIsPending) {
            mutate(input);
            setInput("")
        }
    }, [input, setInput, mutate, submitIsPending]);

    const placeholder = useMemo(() => submitIsPending ? "Your request is being processed ... " : "Start typing here...", [submitIsPending])

    const handleKeyDown = useCallback((e: React.KeyboardEvent) => {
        if (e.key === 'Enter' && !e.shiftKey) {
            e.preventDefault();
            handleSubmit(e);
        }
    }, [handleSubmit]);

    return <form onSubmit={handleSubmit} className={cn(!messagesShown && "flex-1 flex flex-col justify-center")}>
        <div className={cn("w-full flex flex-col items-start justify-center border-t-2 p-4 max-h-full", messagesShown && 'flex flex-row items-center px-20 gap-4', !messagesShown && "max-w-3xl mx-auto border-t-0", className)}>
            {!messagesShown && (
                <div className="text-left mb-8 space-y-2">
                    <Label className="text-4xl font-bold">How can I help you?</Label>
                    <p className="text-muted-foreground">Ask me anything about your job offers, requirements, or application status.</p>
                </div>
            )}
            <div className="relative w-full">
                <Textarea ref={textareaRef} disabled={submitIsPending} minLength={1} value={input} onChange={updateInput} onKeyDown={handleKeyDown} className={cn("shadow-2xl resize-none max-h-30", !messagesShown && "h-40")} placeholder={placeholder} />
                {submitIsPending && (
                    <div className="absolute right-4 bottom-4 animate-spin text-indigo-600">
                        <Loader2 size={20} />
                    </div>
                )}
            </div>
            <Button hidden={true} disabled={inputEmpty || submitIsPending}>Send<SendIcon /></Button>
        </div>
    </form>
}