import { useEffect, useState } from 'react';
import { adminApi, SystemMonitoring, TenantAdmin } from '@/features/admin/api/admin-api';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Building2, Users, Briefcase, FileText, CreditCard, Activity } from 'lucide-react';

function StatCard({ icon: Icon, label, value, description, color }: {
  icon: React.ElementType;
  label: string;
  value: string | number;
  description?: string;
  color: string;
}) {
  return (
    <Card className="relative overflow-hidden border-border/50 hover:shadow-lg transition-all duration-300 group">
      <div className={`absolute inset-0 ${color} opacity-0 group-hover:opacity-100 transition-opacity duration-300`} />
      <CardHeader className="flex flex-row items-center justify-between pb-2 relative">
        <CardTitle className="text-sm font-medium text-muted-foreground">{label}</CardTitle>
        <div className={`p-2 rounded-lg ${color}`}>
          <Icon className="h-4 w-4 text-white" />
        </div>
      </CardHeader>
      <CardContent className="relative">
        <div className="text-3xl font-bold">{value}</div>
        {description && <p className="text-xs text-muted-foreground mt-1">{description}</p>}
      </CardContent>
    </Card>
  );
}

export function AdminDashboard() {
  const [monitoring, setMonitoring] = useState<SystemMonitoring | null>(null);
  const [tenants, setTenants] = useState<TenantAdmin[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const [monitoringData, tenantsData] = await Promise.all([
          adminApi.getSystemMonitoring(),
          adminApi.listTenants(),
        ]);
        setMonitoring(monitoringData);
        setTenants(tenantsData);
      } catch (err) {
        console.error('Failed to fetch admin data:', err);
      } finally {
        setLoading(false);
      }
    };
    fetchData();
  }, []);

  if (loading) {
    return (
      <div className="flex items-center justify-center h-64">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary" />
      </div>
    );
  }

  return (
    <div className="space-y-8">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Admin Dashboard</h1>
        <p className="text-muted-foreground mt-1">System overview and monitoring</p>
      </div>

      {/* Stats Grid */}
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-5">
        <StatCard
          icon={Building2}
          label="Total Tenants"
          value={monitoring?.totalTenants ?? 0}
          description="Active organizations"
          color="bg-blue-600"
        />
        <StatCard
          icon={Users}
          label="Total Users"
          value={monitoring?.totalUsers ?? 0}
          description="Across all tenants"
          color="bg-emerald-600"
        />
        <StatCard
          icon={Briefcase}
          label="Open Positions"
          value={monitoring?.totalPositions ?? 0}
          description="All job openings"
          color="bg-purple-600"
        />
        <StatCard
          icon={FileText}
          label="Applications"
          value={monitoring?.totalApplications ?? 0}
          description="Total submissions"
          color="bg-amber-600"
        />
        <StatCard
          icon={CreditCard}
          label="Active Subscriptions"
          value={monitoring?.activeSubscriptions ?? 0}
          description="Paid plans"
          color="bg-pink-600"
        />
      </div>

      {/* Recent Tenants Table */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Activity className="h-5 w-5 text-primary" />
            Recent Tenants
          </CardTitle>
          <CardDescription>Overview of all registered organizations</CardDescription>
        </CardHeader>
        <CardContent>
          <div className="overflow-x-auto">
            <table className="w-full text-sm">
              <thead>
                <tr className="border-b border-border/50">
                  <th className="text-left py-3 px-4 font-semibold text-muted-foreground">Tenant</th>
                  <th className="text-left py-3 px-4 font-semibold text-muted-foreground">Domain</th>
                  <th className="text-center py-3 px-4 font-semibold text-muted-foreground">Users</th>
                  <th className="text-center py-3 px-4 font-semibold text-muted-foreground">Positions</th>
                  <th className="text-left py-3 px-4 font-semibold text-muted-foreground">Plan</th>
                  <th className="text-left py-3 px-4 font-semibold text-muted-foreground">Status</th>
                </tr>
              </thead>
              <tbody>
                {tenants.map(tenant => (
                  <tr key={tenant.id} className="border-b border-border/30 hover:bg-muted/50 transition-colors">
                    <td className="py-3 px-4 font-medium">{tenant.name}</td>
                    <td className="py-3 px-4 text-muted-foreground">{tenant.domain || '—'}</td>
                    <td className="py-3 px-4 text-center">{tenant.userCount}</td>
                    <td className="py-3 px-4 text-center">{tenant.positionCount}</td>
                    <td className="py-3 px-4">
                      <span className="inline-flex items-center rounded-full bg-primary/10 px-2.5 py-0.5 text-xs font-medium text-primary">
                        {tenant.subscription?.plan?.name ?? 'Free'}
                      </span>
                    </td>
                    <td className="py-3 px-4">
                      <span className={`inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium ${
                        tenant.subscription?.status === 'active'
                          ? 'bg-green-100 text-green-700 dark:bg-green-900/30 dark:text-green-400'
                          : 'bg-zinc-100 text-zinc-600 dark:bg-zinc-800 dark:text-zinc-400'
                      }`}>
                        {tenant.subscription?.status ?? 'No subscription'}
                      </span>
                    </td>
                  </tr>
                ))}
                {tenants.length === 0 && (
                  <tr>
                    <td colSpan={6} className="py-8 text-center text-muted-foreground">
                      No tenants found
                    </td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
