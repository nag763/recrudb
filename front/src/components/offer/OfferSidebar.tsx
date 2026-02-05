import { Briefcase, Trash2 } from "lucide-react";
import { ContextMenu, ContextMenuContent, ContextMenuItem, ContextMenuTrigger } from "../ui/context-menu";
import { SidebarMenuItem } from "../ui/sidebar";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { deleteOffer, getOffers } from "@/interactions/offers";
import useSelectOffer from "@/hooks/useSelectOffer";
import { cn } from "@/lib/utils";
import { toast } from "sonner";
import AppSidebarGroup from "../AppSidebarGroup";

export default function OfferSidebar() {

    const { data: offers, isPending } = useQuery({
        queryKey: ['profile-offers'],
        queryFn: async () => getOffers(),
    });

    const queryClient = useQueryClient();

    const { mutate } = useMutation({
        mutationFn: deleteOffer,
        onSuccess: () => {
            queryClient
                .invalidateQueries({ queryKey: ['profile-offers'] })
                .then(() => toast.success("Your offer is now available!"));
        },
        onError: (error) => {
            console.error('Delete failed', error);
            toast.error("An errror happened while processing your offer. Try again!")
        },
    });

    const { selectedOffer, onOfferSelected: itemClicked } = useSelectOffer();

    return <AppSidebarGroup groupLabel={"Offers"} GroupIcon={Briefcase}>
        {isPending && <p>Loading ...</p>}
        {!isPending &&
            offers?.data?.slice(0, 5).map((e) => (
                <SidebarMenuItem
                    className={cn(
                        ' p-0.5',
                        selectedOffer?.id == e.id && 'font-bold',
                    )}>
                    <ContextMenu>
                        <ContextMenuTrigger
                            onClick={() => itemClicked(e.id)}
                            asChild
                        >
                            <div className="flex flex-col gap-0.5 cursor-pointer">
                                <p>{e.name}</p>
                                <small
                                    className="flex w-full justify-end"
                                >
                                    {e.createdAt.toDateString()}
                                </small>
                            </div>
                        </ContextMenuTrigger>
                        <ContextMenuContent>
                            <ContextMenuItem onClick={() => mutate(e.id)}>
                                <Trash2 />
                                Delete
                            </ContextMenuItem>
                        </ContextMenuContent>
                    </ContextMenu>
                </SidebarMenuItem>
            ))}
            <SidebarMenuItem className="font-semibold">
                See all
            </SidebarMenuItem>
    </AppSidebarGroup>
}