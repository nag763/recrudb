import {
  BriefcaseBusiness,
  CircleUserRound,
  LogOutIcon,
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
  SidebarHeader,
  SidebarRail,
  useSidebar,
} from '@/components/ui/sidebar';
import OfferSidebar from './offer/OfferSidebar';

export function ChatSidebar() {
  const { open, setOpen } = useSidebar();

  return (
    <Sidebar variant="inset" className="border-r-2" collapsible="icon">
      <SidebarHeader className="flex flex-row border-b-4 w-full justify-between">
        <div className="flex flex-row items-center align-middle gap-2">
          <Tooltip>
            <TooltipTrigger asChild>
              <BriefcaseBusiness onClick={() => setOpen(!open)} />
            </TooltipTrigger>
            <TooltipContent hidden={open} side="right">
              Menu
            </TooltipContent>
          </Tooltip>{' '}
          <span hidden={!open}> Offers</span>
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
      <OfferSidebar />
      <SidebarRail />
    </Sidebar>
  );
}
