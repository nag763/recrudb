import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useCallback } from 'react';
import { Field } from './ui/field';
import { Input } from './ui/input';
import { analyzeCv } from '@/interactions/agents';

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
    return <p>Your CV is being processed...</p>;
  }

  return (
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
  );
}
