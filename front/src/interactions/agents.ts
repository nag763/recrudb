import apiClient from '@/lib/client';

export const analyzeCv = (file: File) => {
  const formData = new FormData();
  formData.append('file', file);
  return apiClient.postForm('/agents/cvAnalyzer', formData);
};

export const postOffer = (offerText: string) =>
  apiClient.post('/agents/jobOffer', { data: offerText });

export const getCoordinatorSessions = () => apiClient.get('/agents/coordinator/sessions')
export const getCoordinatorSession = (id: string) => apiClient.get(`/agents/coordinator/sessions/${id}`);
export const sendToCoordinator = ({sessionId, content}) : {sessionId?: string, data: string} => apiClient.post('/agents/coordinator', {data: content, sessionId});
export const deleteSession = (id: string) => apiClient.delete(`/agents/coordinator/sessions/${id}`);