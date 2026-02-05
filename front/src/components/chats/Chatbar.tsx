import { SendIcon } from "lucide-react";
import { Button } from "../ui/button";
import { Label } from "../ui/label";
import { Textarea } from "../ui/textarea";
import { useCallback, useMemo, useState } from "react";

export default function Chatbar() {

    const [input, setInput] = useState("");

    const updateInput = useCallback((e) => {
        setInput(e.target.value);
    }, [setInput]);

    const inputEmpty = useMemo(() => input.length == 0, [input]);

    return <div className="w-full h-full flex flex-col items-center justify-center border-t-2 p-4">
        <div className="w-1/3 flex flex-col gap-2">
            <Label className="text-2xl">Start typing below</Label>
            <Textarea value={input} onChange={updateInput} className=" shadow-2xl h-80 resize-none" placeholder="Start typing here ..." />
            <Button disabled={inputEmpty}>Send<SendIcon/></Button>
        </div>
    </div>
}