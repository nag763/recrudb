import {
  BriefcaseBusiness,
  CircleUserRound,
  LogOutIcon,
  PlusCircle,
  Settings2Icon,
  Trash2,
} from 'lucide-react'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { useCallback, useContext } from 'react'
import { Link } from '@tanstack/react-router'
import { Button } from './ui/button'
import SubmitOfferDialog from './SubmitOfferDialog'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from './ui/dropdown-menu'
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuItem,
} from '@/components/ui/sidebar'
import { Dialog, DialogTrigger } from '@/components/ui/dialog'
import { deleteOffer, getOffers } from '@/interactions/offers'
import {
  ContextMenu,
  ContextMenuContent,
  ContextMenuItem,
  ContextMenuTrigger,
} from '@/components/ui/context-menu'
import { OfferContext, OfferContextDispatcher } from '@/context/OfferContext'
import { cn } from '@/lib/utils'

export function ChatSidebar() {
  const { data: offers, isPending } = useQuery({
    queryKey: ['profile-offers'],
    queryFn: async () => getOffers(),
  })

  const selectedOffer = useContext(OfferContext)
  const dispatch = useContext(OfferContextDispatcher)

  const queryClient = useQueryClient()

  const { mutate } = useMutation({
    mutationFn: deleteOffer,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['profile-offers'] })
    },
    onError: (error) => {
      console.error('Delete failed', error)
    },
  })

  const itemClicked = useCallback(
    (id: string) => {
      if (selectedOffer == id) {
        return
      }
      dispatch &&
        dispatch({
          type: 'selected',
          id: id,
        })
    },
    [dispatch, selectedOffer],
  )

  return (
    <Sidebar variant="inset" className="border-r-2">
      <SidebarHeader className="flex flex-row border-b-4 w-full justify-between">
        <div className="flex flex-row items-center align-middle gap-2">
          <BriefcaseBusiness /> Offers
        </div>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <CircleUserRound className="cursor-pointer"></CircleUserRound>
          </DropdownMenuTrigger>
          <DropdownMenuContent>
            <DropdownMenuItem asChild>
              <Link to={'/profile'}>
                {' '}
                <Settings2Icon />
                Settings
              </Link>
            </DropdownMenuItem>
            <DropdownMenuItem asChild>
              <Link to={'/logout'}>
                <LogOutIcon />
                Logout
              </Link>
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupContent>
            <SidebarMenu>
              {isPending && <p>Loading ...</p>}
              {!isPending &&
                offers &&
                offers?.data?.map((e) => (
                  <SidebarMenuItem
                    className={cn(
                      'border-b-2 p-0.5',
                      selectedOffer == e.id && 'font-bold',
                    )}
                  >
                    <ContextMenu>
                      <ContextMenuTrigger
                        onClick={() => itemClicked(e.id)}
                        asChild
                      >
                        <div className="flex flex-col gap-0.5 cursor-pointer">
                          <p>{e.name}</p>
                          <small className="flex w-full justify-end">
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
