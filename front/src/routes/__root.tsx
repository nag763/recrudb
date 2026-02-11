import {
  Outlet,
  createRootRouteWithContext,
} from '@tanstack/react-router';
import { TanStackRouterDevtoolsPanel } from '@tanstack/react-router-devtools';
import { TanStackDevtools } from '@tanstack/react-devtools';


import TanStackQueryDevtools from '../integrations/tanstack-query/devtools';

import type { QueryClient } from '@tanstack/react-query';
import AuthProvider from '@/components/AuthProvider';
import AuthGuard from '@/components/AuthGuard';
import { Toaster } from '@/components/ui/sonner';
import { SidebarProvider } from '@/components/ui/sidebar';
import { AppSidebar } from '@/components/AppSidebar';

interface MyRouterContext {
  queryClient: QueryClient;
}

export const Route = createRootRouteWithContext<MyRouterContext>()({
  component: () => (
    <AuthProvider>
      <AuthGuard>
        <SidebarProvider>
          <AppSidebar />
          {/* <Header /> */}
          <Outlet />
          <TanStackDevtools
            config={{
              position: 'bottom-right',
            }}
            plugins={[
              {
                name: 'Tanstack Router',
                render: <TanStackRouterDevtoolsPanel />,
              },
              TanStackQueryDevtools,
            ]}
          />
          <Toaster richColors />

        </SidebarProvider>
      </AuthGuard>
    </AuthProvider>
  ),
});
