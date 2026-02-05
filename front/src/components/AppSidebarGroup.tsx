import { ChevronDown, PlusIcon } from "lucide-react";
import { SidebarGroup, SidebarGroupAction, SidebarGroupContent, SidebarGroupLabel, SidebarMenu, useSidebar } from "./ui/sidebar";
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from "./ui/collapsible";

export default function AppSidebarGroup({ groupLabel, GroupIcon, children }) {

    const { open } = useSidebar();

    return <Collapsible defaultOpen className="group/collapsible">
        <SidebarGroup>
            <SidebarGroupLabel asChild>
                <div className="flex gap-2 text-md items-center"><GroupIcon size={16} />{groupLabel}</div>

            </SidebarGroupLabel>
            <SidebarGroupAction>
                <PlusIcon></PlusIcon>
                <CollapsibleTrigger asChild>
                    <ChevronDown className="cursor-pointer transition-transform -rotate-90 group-data-[state=open]/collapsible:rotate-0" />
                </CollapsibleTrigger>
            </SidebarGroupAction>
            <SidebarGroupContent>
                <CollapsibleContent>
                    <SidebarMenu hidden={!open} className="border-l ml-2 p-2 max-h-96 overflow-auto">
                        {children}
                    </SidebarMenu>
                </CollapsibleContent>
            </SidebarGroupContent>
        </SidebarGroup>
    </Collapsible>
}