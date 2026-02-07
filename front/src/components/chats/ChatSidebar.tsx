import { Bot } from "lucide-react";
import AppSidebarGroup from "../AppSidebarGroup";
import { useQuery } from "@tanstack/react-query";
import { getCoordinatorSessions } from "@/interactions/agents";
import { SidebarMenuItem } from "../ui/sidebar";
import { cn } from "@/lib/utils";
import { useMemo } from "react";
import { Link } from "@tanstack/react-router";

export default function ChatSidebar() {

    const { data, isPending } = useQuery({
        queryKey: ['profile-sessions'],
        queryFn: async () => getCoordinatorSessions(),
    });


    const sessions = useMemo(() => data?.data, [data]);
    return <AppSidebarGroup groupLabel={"Chats"} GroupIcon={Bot}>

        <Link to="/chat" search={{ id: undefined }}>
            <SidebarMenuItem className="font-semibold">Create new</SidebarMenuItem>
        </Link>

        {isPending && <p>Loading ...</p>}
        {!isPending && sessions.map((e) => <Link to="/chat" search={{ id: e.id }}>
            <SidebarMenuItem
                className={cn(
                    ' p-0.5 text-truncate',
                )}>
                {60 < e.name.length ? `${e.name.slice(0, 60)}...` : e.name}
            </SidebarMenuItem>
        </Link>
        )}
    </AppSidebarGroup>
}