import { apiClient } from '@/shared/api/axios';
import { z } from 'zod';

export const loginSchema = z.object({
  email: z.string().email('Alamat email tidak valid'),
  password: z.string().min(6, 'Password minimal 6 karakter'),
});

export type LoginCredentials = z.infer<typeof loginSchema>;

export interface AuthUser {
  id: string;
  email: string;
  firstName: string;
  lastName: string;
  role: string;
  tenantId: string;
}

export interface AuthResponse {
  accessToken: string;
  refreshToken: string;
  user: AuthUser;
}

export const authApi = {
  login: async (credentials: LoginCredentials): Promise<AuthResponse> => {
    const response = await apiClient.post<AuthResponse>('/auth/login', credentials);
    return response.data;
  },
};
