import { SendIcon } from "lucide-react";
import { Button } from "../ui/button";
import { Label } from "../ui/label";
import { Textarea } from "../ui/textarea";
import { useCallback, useMemo, useState } from "react";
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

    const placeholder = useMemo(() => submitIsPending ? "Your request is being processed ... ": "Start typing here...", [submitIsPending])


    return <form onSubmit={handleSubmit}><div className={cn("w-full h-full flex flex-col items-center justify-center border-t-2 p-4", messagesShown && 'flex flex-row h-auto', className)}>
        {!messagesShown && <Label className="text-2xl">Start typing below</Label>}
        <Textarea minLength={1} value={input} onChange={updateInput} className={cn("shadow-2xl resize-none ", !messagesShown && "h-80")} placeholder={placeholder}/>
        <Button disabled={inputEmpty || submitIsPending}>Send<SendIcon /></Button>
    </div></form>
}