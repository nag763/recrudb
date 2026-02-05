import Chatbar from '@/components/chats/Chatbar'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/chat/')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div className='h-full w-full flex-col'>
    <Chatbar></Chatbar>
  </div>
}
