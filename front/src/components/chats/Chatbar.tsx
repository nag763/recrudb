import { SendIcon } from "lucide-react";
import { Button } from "../ui/button";
import { Label } from "../ui/label";
import { Textarea } from "../ui/textarea";
import { useCallback, useMemo, useState } from "react";
import { cn } from "@/lib/utils";

export default function Chatbar({ messagesShown }) {

    const [input, setInput] = useState("");

    const updateInput = useCallback((e) => {
        setInput(e.target.value);
    }, [setInput]);

    const inputEmpty = useMemo(() => input.length == 0, [input]);

    return <div className={cn("w-full h-full flex flex-col items-center justify-center border-t-2 p-4", messagesShown && 'flex flex-row h-auto')}>
        {!messagesShown && <Label className="text-2xl">Start typing below</Label>}
        <Textarea value={input} onChange={updateInput} className={cn("shadow-2xl resize-none ", !messagesShown && "h-80")} placeholder="Start typing here ..." />
        <Button disabled={inputEmpty}>Send<SendIcon /></Button>
    </div>
}