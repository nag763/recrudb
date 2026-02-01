import apiClient from "@/lib/client";
import type { ApiResponse } from "./type";
import { ProfileSchema, type Profile } from "@/types/profile";

export const getProfile = () : ApiResponse<Profile> => (apiClient.get('/profile/me', {schema: ProfileSchema}));
export const deleteProfile = () => (apiClient.delete('/profile/me'));