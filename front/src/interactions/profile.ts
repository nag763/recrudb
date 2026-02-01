import type { ApiResponse } from './type'
import type {Profile} from '@/types/profile';
import apiClient from '@/lib/client'
import {  ProfileSchema } from '@/types/profile'

export const getProfile = (): ApiResponse<Profile> =>
  apiClient.get('/profile/me', { schema: ProfileSchema })
export const deleteProfile = () => apiClient.delete('/profile/me')
