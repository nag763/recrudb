import Chatbar from '@/components/chats/Chatbar'
import ChatInterface from '@/components/chats/ChatInterface'
import { createFileRoute, useSearch } from '@tanstack/react-router'

export const Route = createFileRoute('/chat/')({
  component: RouteComponent,
  validateSearch: (search: Record<string, unknown>) => {
    return {
      id: (search.id as string) || undefined,
    }
  },
})

function RouteComponent() {
  const { id } = useSearch({from: "/chat/"})
  return <div className='w-full h-dvh flex flex-col overflow-hidden'>
    {id && 
      <ChatInterface sessionId={id}/>
    }
    <Chatbar messagesShown={!!id} />
    
  </div>
}
