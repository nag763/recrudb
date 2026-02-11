import { Trash2 } from "lucide-react";
import { ContextMenu, ContextMenuContent, ContextMenuItem, ContextMenuTrigger } from "../ui/context-menu";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";
import type { ReactNode } from "react";
import { deleteSession } from "@/interactions/agents";

export default function ChatContextMenu({ children, sessionId }: Readonly<{ children: ReactNode, sessionId: string }>) {


    const queryClient = useQueryClient();

    const { mutate } = useMutation({
        mutationFn: deleteSession,
        onSuccess: () => {
            queryClient
                .invalidateQueries({ queryKey: ['profile-sessions'] })
                .then(() => toast.success("Your chat has been deleted with success"));
        },
        onError: (error) => {
            console.error('Delete failed', error);
            toast.error("An errror happened while deleting the session. Try again!")
        },
    });
    return <ContextMenu>
        <ContextMenuTrigger
            asChild
        >
            {children}
        </ContextMenuTrigger>
        <ContextMenuContent>
            <ContextMenuItem onClick={() => mutate(sessionId)}>
                <Trash2 />
                Delete
            </ContextMenuItem>
        </ContextMenuContent>
    </ContextMenu>
}