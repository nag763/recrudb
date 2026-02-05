import {
  BriefcaseBusiness,
  CircleUserRound,
  LogOutIcon,
  PlusCircle,
  Settings2Icon,
} from 'lucide-react';
import { Link } from '@tanstack/react-router';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from './ui/dropdown-menu';
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from './ui/tooltip';
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarRail,
  useSidebar,
} from '@/components/ui/sidebar';
import OfferSidebar from './offer/OfferSidebar';
import { Dialog, DialogTrigger } from './ui/dialog';
import { Button } from './ui/button';
import SubmitOfferDialog from './offer/SubmitOfferDialog';
import ChatSidebar from './chats/ChatSidebar';

export function AppSidebar() {
  const { open, setOpen } = useSidebar();

  return (
    <Sidebar variant="inset" className="border-r-2" collapsible="icon">
      <SidebarHeader className="flex flex-row border-b-4 w-full justify-between">
        <div className="flex flex-row items-center align-middle gap-2">
          <Tooltip>
            <TooltipTrigger asChild>
              <img src='/recrudb.webp' className='cursor-pointer' height={32} width={32} onClick={() => setOpen(!open)} />
            </TooltipTrigger>
            <TooltipContent hidden={open} side="right">
              Menu
            </TooltipContent>
          </Tooltip>{' '}
          <span hidden={!open}>recrudb</span>
        </div>
        <DropdownMenu >
          <DropdownMenuTrigger hidden={!open} asChild>
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
      <SidebarContent className='gap-0'>
        <ChatSidebar/>
        <OfferSidebar />
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
      <SidebarRail />
    </Sidebar>
  );
}
