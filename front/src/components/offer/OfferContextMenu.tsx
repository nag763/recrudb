import { Trash2 } from "lucide-react";
import { ContextMenu, ContextMenuContent, ContextMenuItem, ContextMenuTrigger } from "../ui/context-menu";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { deleteOffer } from "@/interactions/offers";
import { toast } from "sonner";
import type { ReactNode } from "react";

export default function OfferContextMenu({ children, offerId }: Readonly<{ children: ReactNode, offerId: string }>) {


    const queryClient = useQueryClient();

    const { mutate } = useMutation({
        mutationFn: deleteOffer,
        onSuccess: () => {
            queryClient
                .invalidateQueries({ queryKey: ['profile-offers'] })
                .then(() => toast.success("The offer has been deleted with success"));
        },
        onError: (error) => {
            console.error('Delete failed', error);
            toast.error("An errror happened while processing your offer. Try again!")
        },
    });
    return <ContextMenu>
        <ContextMenuTrigger
            asChild
        >
            {children}
        </ContextMenuTrigger>
        <ContextMenuContent>
            <ContextMenuItem onClick={() => mutate(offerId)}>
                <Trash2 />
                Delete
            </ContextMenuItem>
        </ContextMenuContent>
    </ContextMenu>
}