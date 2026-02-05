import { createFileRoute, useNavigate } from '@tanstack/react-router';
import { useProfile } from '@/hooks/useProfile';

export const Route = createFileRoute('/')({
  component: App,
});

function App() {
  const { data, isLoading } = useProfile();

  const navigate = useNavigate();


  if (isLoading) {
    return <p>Loading...</p>;
  }

  if (data?.status == 204) {
    navigate({ to: '/profile' });
  }

  return (
    <div className='w-full h-full   min-h-screen  flex flex-col gap-12 items-center align-middle justify-center col-span-full'>
      <img src="/daydream.svg" height={400} width={400} />
      <h1 className='text-4xl'>I will be waiting....</h1>
      <h2>Use the panel on the left to start using the app.</h2>
    </div>
  );
}
