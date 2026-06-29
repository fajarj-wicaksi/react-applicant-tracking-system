import { apiClient } from '@/shared/api/axios';

export interface User {
  id: string;
  firstName: string;
  lastName: string;
  email: string;
  roleName: string;
  createdAt: string;
}

export interface CreateUserRequest {
  firstName: string;
  lastName: string;
  email: string;
  password?: string;
  roleId: string;
}

export interface UpdateUserRequest {
  firstName?: string;
  lastName?: string;
  email?: string;
  password?: string;
  roleId?: string;
}

export const userApi = {
  listUsers: async (): Promise<User[]> => {
    const response = await apiClient.get<User[]>('/users');
    return response.data;
  },

  getUser: async (id: string): Promise<User> => {
    const response = await apiClient.get<User>(`/users/${id}`);
    return response.data;
  },

  createUser: async (data: CreateUserRequest): Promise<User> => {
    const response = await apiClient.post<User>('/users', data);
    return response.data;
  },

  updateUser: async (id: string, data: UpdateUserRequest): Promise<void> => {
    await apiClient.put(`/users/${id}`, data);
  },

  deleteUser: async (id: string): Promise<void> => {
    await apiClient.delete(`/users/${id}`);
  },
};
