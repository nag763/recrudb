import { Bot } from "lucide-react";
import AppSidebarGroup from "../AppSidebarGroup";
import { useQuery } from "@tanstack/react-query";
import { getCoordinatorSessions } from "@/interactions/agents";
import { SidebarMenuItem } from "../ui/sidebar";
import { cn } from "@/lib/utils";
import { useCallback, useMemo } from "react";
import { useNavigate } from "@tanstack/react-router";

export default function ChatSidebar() {

    const navigate = useNavigate();

    const { data, isPending } = useQuery({
        queryKey: ['profile-sessions'],
        queryFn: async () => getCoordinatorSessions(),
    });

    const createNewChatCallback = useCallback(() => {
        navigate({to: "/chat"})
    }, [navigate])

    const sessions = useMemo(() => data?.data, [data]);
    return <AppSidebarGroup groupLabel={"Chats"} GroupIcon={Bot}>
        <SidebarMenuItem className="font-semibold" onClick={createNewChatCallback}>Create new</SidebarMenuItem>
        {isPending && <p>Loading ...</p>}
        {!isPending && sessions.map((e) => <SidebarMenuItem
            className={cn(
                ' p-0.5 text-truncate',
            )}>
            {e.name}
        </SidebarMenuItem>)}
    </AppSidebarGroup>
}