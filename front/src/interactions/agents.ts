import apiClient from '@/lib/client'

export const analyzeCv = (file: File) => {
  const formData = new FormData()
  formData.append('file', file)
  return apiClient.postForm('/agents/cvAnalyzer', formData)
}

export const postOffer = (offerText: string) =>
  apiClient.post('/agents/jobOffer', { data: offerText })
