import { apiClient } from '@/shared/api/axios';

// --- Types ---

export interface BillingPlan {
  id: string;
  name: string;
  description: string;
  priceMonthly: number;
  maxUsers: number;
  maxStorage: number;
  maxPositions: number;
}

export interface Subscription {
  id: string;
  status: string;
  startDate: string;
  endDate?: string;
  plan?: BillingPlan;
}

export interface TenantAdmin {
  id: string;
  name: string;
  domain: string;
  userCount: number;
  positionCount: number;
  subscription?: Subscription;
  createdAt: string;
}

export interface TenantStats {
  tenantId: string;
  tenantName: string;
  userCount: number;
  positionCount: number;
  applicationCount: number;
  storageUsage: number;
}

export interface SystemMonitoring {
  totalTenants: number;
  totalUsers: number;
  totalPositions: number;
  totalApplications: number;
  activeSubscriptions: number;
}

// --- API ---

export const adminApi = {
  listTenants: async (): Promise<TenantAdmin[]> => {
    const response = await apiClient.get<TenantAdmin[]>('/admin/tenants');
    return response.data;
  },

  getTenantStats: async (tenantId: string): Promise<TenantStats> => {
    const response = await apiClient.get<TenantStats>(`/admin/tenants/${tenantId}/stats`);
    return response.data;
  },

  getSystemMonitoring: async (): Promise<SystemMonitoring> => {
    const response = await apiClient.get<SystemMonitoring>('/admin/monitoring');
    return response.data;
  },
};
