import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useCallback, useState } from 'react';
import { toast } from 'sonner';
import { DialogFooter, DialogHeader } from '../ui/dialog';
import { Button } from '../ui/button';
import { Textarea } from '../ui/textarea';
import type { FormEvent } from 'react';
import { postOffer } from '@/interactions/agents';
import {
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogTitle,
} from '@/components/ui/dialog';

export default function SubmitOfferDialog(): React.ReactNode {
  const [data, setData] = useState<string>('');

  const queryClient = useQueryClient();

  const mutation = useMutation({
    mutationFn: postOffer,
    onSuccess: () => {
      setData('');
      queryClient.invalidateQueries({ queryKey: ['profile-offers'] });
      toast.success('Your offer has been registered');
    },
  });

  const onFormSubmitCallback = useCallback(
    (e: FormEvent) => {
      e.preventDefault();
      mutation.mutate(data);
    },
    [mutation],
  );

  return (
    <DialogContent className="sm:max-w-md flex flex-col gap-8">
      <form onSubmit={onFormSubmitCallback}>
        <DialogHeader>
          <DialogTitle>Copy paste your offer below</DialogTitle>
          <DialogDescription>
            A match score will be calculated and the critical informations will
            be displayed.
          </DialogDescription>
        </DialogHeader>
        <div className="flex items-center gap-2">
          <div className="grid flex-1 gap-2">
            <Textarea
              className="max-h-40"
              name="data"
              minLength={10}
              value={data}
              onChange={(e) => setData(e.target.value)}
            />
          </div>
        </div>
        <DialogFooter className="mt-2">
          <DialogClose className="flex w-full justify-end space-x-2">
            <Button variant={'default'} type="button">
              Close
            </Button>
            <Button variant={'secondary'}>Submit</Button>
          </DialogClose>
        </DialogFooter>
      </form>
    </DialogContent>
  );
}
