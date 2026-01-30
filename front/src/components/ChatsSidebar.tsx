import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuItem,
} from "@/components/ui/sidebar"
import { BriefcaseBusiness, PlusCircle } from "lucide-react"
import { Button } from "./ui/button"
import { Dialog, DialogTrigger } from "@/components/ui/dialog"
import { useQuery } from "@tanstack/react-query"
import { getOffers } from "@/interactions/offers"
import SubmitOfferDialog from "./SubmitOfferDialog"

export function ChatSidebar() {

  const { data: offers, isPending } = useQuery(
    {
      queryKey: ['profile-offers'],
      queryFn: async () => getOffers(),
    }
  )

  return (
    <Sidebar>
      <SidebarHeader className="flex flex-row border-b-4 items-center align-middle">
        <BriefcaseBusiness /> Offers
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupContent>
            <SidebarMenu>
              {isPending && <p>Loading ...</p>}
              {!isPending && offers && offers?.data.map(e => (
                <SidebarMenuItem className="border-b-2 p-0.5">
                  <div className="flex flex-col gap-0.5">
                    <p>{e.name}</p>
                    <small className="flex w-full justify-end">{new Date(e.createdAt).toDateString()}</small>
                  </div>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>

          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <Dialog>
          <DialogTrigger asChild>
            <Button>
              <PlusCircle /> Add a new offer
            </Button>
          </DialogTrigger>
          <SubmitOfferDialog />
        </Dialog>
      </SidebarFooter>
    </Sidebar>
  )
}