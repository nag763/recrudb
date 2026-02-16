import { LetterText, Trash2 } from "lucide-react";
import { ContextMenu, ContextMenuContent, ContextMenuItem, ContextMenuTrigger } from "../ui/context-menu";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { deleteOffer } from "@/interactions/offers";
import { toast } from "sonner";
import { coverLetter } from "@/interactions/agents";

export default function OfferContextMenu({ children, offerId }: Readonly<{ children: ReactNode, offerId: string }>) {

    const queryClient = useQueryClient();

    const { mutate: remove } = useMutation({
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



    const { mutateAsync: writeCover } = useMutation({
        mutationFn: coverLetter,
    });

    const writeCoverToast = () => {
        toast.promise(writeCover(offerId), {
            loading: "Preparing the letter...",
            success: (data) => {
                if (data?.data?.content) {
                    navigator.clipboard.writeText(data.data.content);
                    return "Cover letter copied to clipboard!";
                }
                return "Please try again ..."
            },
            error: "An error happened :("
        })
    }


    return <ContextMenu>
        <ContextMenuTrigger
            asChild
        >
            {children}
        </ContextMenuTrigger>
        <ContextMenuContent>
            <ContextMenuItem onClick={writeCoverToast}>
                <LetterText />
                Write a cover
            </ContextMenuItem>
            <ContextMenuItem onClick={() => remove(offerId)}>
                <Trash2 />
                Delete
            </ContextMenuItem>
        </ContextMenuContent>
    </ContextMenu >
}