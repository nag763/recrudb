import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useCallback } from 'react';
import { Field } from '../ui/field';
import { Input } from '../ui/input';
import { analyzeCv } from '@/interactions/agents';
import LoadingCoffee from './LoadingCoffee';

export default function UploadCv() {
  const queryClient = useQueryClient();

  const { mutate, isPending: cvSubmitIsPending } = useMutation({
    mutationFn: (newFile: File) => analyzeCv(newFile),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['profile'] });
    },
    onError: (error) => {
      console.error('Upload failed', error);
    },
  });

  const onFileChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      mutate(e.target.files[0]);
    }
  }, []);

  if (cvSubmitIsPending) {
    return <LoadingCoffee/>;
  }

  return (
    <div className='h-full flex w-full items-center justify-between'>
      <div className='flex flex-col gap-8'>
        <h1 className='text-4xl'>Start by uploading your CV</h1>
        <form>
          <Field>
            <Input
              disabled={cvSubmitIsPending}
              onChange={onFileChange}
              required={true}
              id="picture"
              type="file"
            />
          </Field>
          <p>To create your profile, start by submitting your CV </p>
        </form>
      </div>
      <img src='/upload_cv.svg'></img>
    </div>
  );
}
