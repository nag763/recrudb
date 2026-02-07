import {  createFileRoute } from '@tanstack/react-router';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { deleteProfile, getProfile } from '@/interactions/profile';
import UploadCv from '@/components/profile/UploadCv';
import { Button } from '@/components/ui/button';
import { Trash2Icon } from 'lucide-react';
import { Tooltip, TooltipContent, TooltipTrigger } from '@/components/ui/tooltip';

export const Route = createFileRoute('/profile/')({
  component: RouteComponent,
});

function RouteComponent() {
  const queryClient = useQueryClient();

  const { data, isPending: profileIsPending } = useQuery({
    queryKey: ['profile'],
    queryFn: () => getProfile(),
  });

  const { mutate: deleteProfileMutate } = useMutation({
    mutationFn: () => deleteProfile(),
    onSuccess: () => {
      queryClient.invalidateQueries(['profile']);
    },
  });

  if (profileIsPending) {
    return <p>Loading ...</p>;
  }

  return (
    <div className="px-80 py-12">
      {data?.status == 204 && <UploadCv />}

      {data?.status == 200 && (
        <div className="space-y-8">
          <div className='w-full flex justify-between'>
            <div>
              <h1 className="text-3xl font-bold">
                {data.data.firstName} {data.data.lastName}
              </h1>
              <div className="text-gray-600 mt-2 flex gap-3 text-sm">
                {data.data.contact?.email && (
                  <span>{data.data.contact.email}</span>
                )}
                {data.data.contact?.number && (
                  <span>{data.data.contact.number}</span>
                )}
                {data.data.location?.city && (
                  <span>{data.data.location?.city}</span>
                )}
              </div>
            </div>
            <Tooltip>
              <TooltipTrigger asChild>
                <Button variant="ghost" onClick={deleteProfileMutate}>
                  <Trash2Icon className='text-red-500' />
                </Button>
              </TooltipTrigger>
              <TooltipContent>
                Delete profile
              </TooltipContent>
            </Tooltip>
          </div>

          {data.data.lookingFor && 
            <p className='bg-gray-100 p-4 italic rounded-3xl'>{data.data.lookingFor}</p>
          }

          {data.data.skills?.length > 0 && (
            <div>
              <h2 className="text-xl font-semibold mb-3">Skills</h2>
              <div className="flex flex-wrap gap-2">
                {data.data.skills.map((skill: string, i: number) => (
                  <span
                    key={i}
                    className="bg-gray-100 px-3 py-1 rounded-full text-sm"
                  >
                    {skill}
                  </span>
                ))}
              </div>
            </div>
          )}

          {data.data.experiences?.length > 0 && (
            <div>
              <h2 className="text-xl font-semibold mb-4">Experience</h2>
              <div className="space-y-8">
                {data.data.experiences.map((exp, i: number) => (
                  <div
                    key={i}
                    className="relative pl-4 border-l-2 border-gray-200"
                  >
                    <div className="mb-2">
                      <h3 className="font-bold text-lg">{exp.role}</h3>
                      <div className="text-gray-600 text-sm">
                        <span className="font-medium">{exp.company}</span>
                        <span className="mx-2">•</span>
                        <span>
                          {exp?.startDate?.toLocaleDateString()} -
                          {exp?.endDate?.toLocaleDateString()}
                        </span>
                      </div>
                    </div>
                    <p className="text-gray-700 text-sm whitespace-pre-line mb-3">
                      {exp.mission}
                    </p>
                    {exp.skills?.length > 0 && (
                      <div className="flex flex-wrap gap-1">
                        {exp.skills.map((skill: string, j: number) => (
                          <span
                            key={j}
                            className="text-xs text-gray-500 bg-gray-100 px-2 py-1 rounded"
                          >
                            {skill}
                          </span>
                        ))}
                      </div>
                    )}
                  </div>
                ))}
              </div>
            </div>
          )}
        </div>
      )}
    </div>
  );
}
