import { apiClient } from '@/shared/api/axios';

// --- Types ---

export interface Candidate {
  id: string;
  firstName: string;
  lastName: string;
  email: string;
  phone: string;
  resumeUrl: string;
  source: string;
  createdAt: string;
}

export interface Position {
  id: string;
  title: string;
  department: string;
  location: string;
  type: string;
  description: string;
  isOpen: boolean;
  createdAt: string;
}

export type ApplicationStage = 'Applied' | 'Screening' | 'Interview' | 'Offer' | 'Hired' | 'Rejected';

export interface Application {
  id: string;
  candidateId: string;
  positionId: string;
  stage: ApplicationStage;
  stageOrder: number;
  notes: string;
  appliedAt: string;
  candidate?: Candidate;
  position?: Position;
}

export interface PipelineResponse {
  stages: Record<ApplicationStage, Application[]>;
}

// --- API ---

export const pipelineApi = {
  getPipeline: async (): Promise<PipelineResponse> => {
    const response = await apiClient.get<PipelineResponse>('/pipeline');
    return response.data;
  },

  updateStage: async (applicationId: string, stage: ApplicationStage, stageOrder: number): Promise<void> => {
    await apiClient.patch(`/applications/${applicationId}/stage`, { stage, stageOrder });
  },

  createCandidate: async (data: Omit<Candidate, 'id' | 'createdAt'>): Promise<Candidate> => {
    const response = await apiClient.post<Candidate>('/candidates', data);
    return response.data;
  },

  listCandidates: async (): Promise<Candidate[]> => {
    const response = await apiClient.get<Candidate[]>('/candidates');
    return response.data;
  },

  createPosition: async (data: Omit<Position, 'id' | 'createdAt' | 'isOpen'>): Promise<Position> => {
    const response = await apiClient.post<Position>('/positions', data);
    return response.data;
  },

  listPositions: async (): Promise<Position[]> => {
    const response = await apiClient.get<Position[]>('/positions');
    return response.data;
  },

  createApplication: async (data: { candidateId: string; positionId: string; notes?: string }): Promise<Application> => {
    const response = await apiClient.post<Application>('/applications', data);
    return response.data;
  },

  updatePosition: async (id: string, data: Partial<Omit<Position, 'id' | 'createdAt'>>): Promise<Position> => {
    const response = await apiClient.put<Position>(`/positions/${id}`, data);
    return response.data;
  },

  deletePosition: async (id: string): Promise<void> => {
    await apiClient.delete(`/positions/${id}`);
  },

  updateCandidate: async (id: string, data: Partial<Omit<Candidate, 'id' | 'createdAt'>>): Promise<Candidate> => {
    const response = await apiClient.put<Candidate>(`/candidates/${id}`, data);
    return response.data;
  },

  deleteCandidate: async (id: string): Promise<void> => {
    await apiClient.delete(`/candidates/${id}`);
  },
};

