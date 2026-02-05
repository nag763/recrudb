import { getCoordinatorSession } from "@/interactions/agents";
import { useQuery } from "@tanstack/react-query";
import { useMemo } from "react";
import ChatEvent from "./ChatEvent";

export default function ChatInterface({sessionId, className} : Readonly<{sessionId: string}>) {

    const { data, isPending } = useQuery({
        queryKey: [`profile-session-${sessionId}`],
        queryFn: async () => getCoordinatorSession(sessionId),
    });

    const session = useMemo(() => data?.data.Session, [data]);

    if(isPending) {
        return <p>Waiting ...</p>
    }

    return <div className="h-screen flex-1 min-h-0 overflow-auto flex flex-col-reverse">
        {session && session.eventWrapper.Events.map((e) => <ChatEvent key={e.ID} event={e}/>)}
    </div>
}