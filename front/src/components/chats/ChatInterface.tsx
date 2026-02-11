import { getCoordinatorSession } from "@/interactions/agents";
import { useQuery } from "@tanstack/react-query";
import { useMemo } from "react";
import ChatEvent from "./ChatEvent";

export default function ChatInterface({sessionId} : Readonly<{sessionId: string}>) {

    const { data, isPending } = useQuery({
        queryKey: [`profile-session-${sessionId}`],
        queryFn: async () => getCoordinatorSession(sessionId),
    });

    const session = useMemo(() => data?.data.Session, [data]);

    if(isPending) {
        return <div className="h-full text-sm">Waiting ...</div>
    }

    return <div className="flex-1 min-h-0 overflow-auto flex flex-col-reverse px-[15%] py-12 gap-4">
        {session && session.eventWrapper.Events.filter((e) => e.Content.parts.every((p) => p.text != null)).map((e) => <ChatEvent key={e.ID} event={e}/>).reverse()}
    </div>
}