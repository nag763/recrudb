import { Briefcase } from "lucide-react";
import { SidebarMenuItem } from "../ui/sidebar";
import { useQuery } from "@tanstack/react-query";
import { getOffers } from "@/interactions/offers";
import AppSidebarGroup from "../AppSidebarGroup";
import { Link } from "@tanstack/react-router";
import OfferContextMenu from "./OfferContextMenu";

export default function OfferSidebar() {



    const { data: offers, isPending } = useQuery({
        queryKey: ['profile-offers'],
        queryFn: async () => getOffers(),
    });


    return <AppSidebarGroup groupLabel={"Offers"} GroupIcon={Briefcase}>
        {isPending && <p>Loading ...</p>}
        {!isPending &&
            offers?.data?.slice(0, 5).map((e) => (
                <SidebarMenuItem
                    className={"p-0.5"}>
                    <OfferContextMenu offerId={e.id}>
                        <Link to="/offer" search={{ id: e.id }}>

                            <div className="flex flex-col gap-0.5 cursor-pointer">
                                <p>{e.name}</p>
                                <small
                                    className="flex w-full justify-end"
                                >
                                    {e.createdAt.toDateString()}
                                </small>
                            </div>
                        </Link>
                    </OfferContextMenu>
                </SidebarMenuItem>
            ))}
        <Link to="/offer" search={{ id: undefined }}><SidebarMenuItem className="font-semibold cursor-pointer">
            See all
        </SidebarMenuItem>
        </Link>
    </AppSidebarGroup>
}