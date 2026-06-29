import { apiClient } from '@/shared/api/axios';
import { Candidate, Position } from '@/features/pipeline/api/pipeline-api';

export type InterviewStatus = 'Scheduled' | 'Completed' | 'Cancelled';

export interface Interview {
  id: string;
  candidateId: string;
  positionId: string;
  title: string;
  scheduledAt: string;
  duration: number;
  status: InterviewStatus;
  location: string;
  notes: string;
  candidate?: Candidate;
  position?: Position;
}

export interface CreateInterviewRequest {
  candidateId: string;
  positionId: string;
  title: string;
  scheduledAt: string;
  duration: number;
  location?: string;
  notes?: string;
}

export const interviewApi = {
  listInterviews: async (): Promise<Interview[]> => {
    const response = await apiClient.get<Interview[]>('/interviews');
    return response.data;
  },

  createInterview: async (data: CreateInterviewRequest): Promise<Interview> => {
    const response = await apiClient.post<Interview>('/interviews', data);
    return response.data;
  },

  updateInterview: async (id: string, data: Partial<CreateInterviewRequest & { status: InterviewStatus }>): Promise<Interview> => {
    const response = await apiClient.put<Interview>(`/interviews/${id}`, data);
    return response.data;
  },

  deleteInterview: async (id: string): Promise<void> => {
    await apiClient.delete(`/interviews/${id}`);
  },
};
