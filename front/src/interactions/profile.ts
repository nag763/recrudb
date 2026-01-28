import apiClient from "@/lib/client";

export const getProfile = () => (apiClient.get('/profile/me'));
export const deleteProfile = () => (apiClient.delete('/profile/me'));