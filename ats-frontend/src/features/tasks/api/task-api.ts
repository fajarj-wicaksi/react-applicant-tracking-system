import { apiClient } from '@/shared/api/axios';
import { User } from '@/features/users/api/user-api';

export type TaskStatus = 'Pending' | 'In Progress' | 'Completed';

export interface Task {
  id: string;
  title: string;
  description: string;
  dueDate?: string;
  status: TaskStatus;
  assignedTo?: string;
  assignee?: User;
  createdAt: string;
}

export interface CreateTaskRequest {
  title: string;
  description?: string;
  dueDate?: string;
  status?: TaskStatus;
  assignedTo?: string;
}

export const taskApi = {
  listTasks: async (): Promise<Task[]> => {
    const response = await apiClient.get<Task[]>('/tasks');
    return response.data;
  },

  createTask: async (data: CreateTaskRequest): Promise<Task> => {
    const response = await apiClient.post<Task>('/tasks', data);
    return response.data;
  },

  updateTask: async (id: string, data: Partial<CreateTaskRequest>): Promise<Task> => {
    const response = await apiClient.put<Task>(`/tasks/${id}`, data);
    return response.data;
  },

  deleteTask: async (id: string): Promise<void> => {
    await apiClient.delete(`/tasks/${id}`);
  },
};
