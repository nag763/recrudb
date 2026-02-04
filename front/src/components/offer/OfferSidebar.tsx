import { PlusCircle, Trash2 } from "lucide-react";
import { ContextMenu, ContextMenuContent, ContextMenuItem, ContextMenuTrigger } from "../ui/context-menu";
import { SidebarContent, SidebarFooter, SidebarGroup, SidebarGroupContent, SidebarMenu, SidebarMenuItem, useSidebar } from "../ui/sidebar";
import { Tooltip, TooltipContent, TooltipTrigger } from "../ui/tooltip";
import { Dialog, DialogTrigger } from "../ui/dialog";
import { Button } from "../ui/button";
import SubmitOfferDialog from "./SubmitOfferDialog";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { deleteOffer, getOffers } from "@/interactions/offers";
import useSelectOffer from "@/hooks/useSelectOffer";
import { cn } from "@/lib/utils";
import { toast } from "sonner";

export default function OfferSidebar() {

    const { open } = useSidebar();

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

    return <>
        <SidebarContent>
            <SidebarGroup>
                <SidebarGroupContent>
                    <SidebarMenu>
                        {isPending && <p>Loading ...</p>}
                        {!isPending &&
                            offers?.data?.map((e) => (
                                <SidebarMenuItem
                                    className={cn(
                                        'border-b-2 p-0.5',
                                        selectedOffer?.id == e.id && 'font-bold',
                                        !open && 'text-xs truncate',
                                    )}>
                                    <ContextMenu>
                                        <ContextMenuTrigger
                                            onClick={() => itemClicked(e.id)}
                                            asChild
                                        >
                                            <div className="flex flex-col gap-0.5 cursor-pointer">
                                                <p>{e.name}</p>
                                                <small
                                                    hidden={!open}
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
                    </SidebarMenu>
                </SidebarGroupContent>
            </SidebarGroup>
        </SidebarContent>
        <SidebarFooter>
            <Dialog>
                <Tooltip>
                    <TooltipTrigger asChild className="flex gap-2 items-center">
                        <DialogTrigger asChild>
                            <Button>
                                <PlusCircle />
                                <span hidden={!open}>Add a new offer</span>
                            </Button>
                        </DialogTrigger>
                    </TooltipTrigger>
                    <TooltipContent hidden={open} side="right">
                        Add a new offer
                    </TooltipContent>
                </Tooltip>
                <SubmitOfferDialog />
            </Dialog>
        </SidebarFooter>
    </>
}