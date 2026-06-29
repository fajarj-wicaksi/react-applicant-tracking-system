import { useEffect, useState } from 'react';
import { adminApi, TenantAdmin, TenantStats } from '@/features/admin/api/admin-api';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Building2, Users, Briefcase, FileText, HardDrive } from 'lucide-react';

function formatBytes(bytes: number): string {
  if (bytes === 0) return '0 B';
  const k = 1024;
  const sizes = ['B', 'KB', 'MB', 'GB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return parseFloat((bytes / Math.pow(k, i)).toFixed(1)) + ' ' + sizes[i];
}

export function TenantsPage() {
  const [tenants, setTenants] = useState<TenantAdmin[]>([]);
  const [selectedStats, setSelectedStats] = useState<TenantStats | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchTenants = async () => {
      try {
        const data = await adminApi.listTenants();
        setTenants(data);
      } catch (err) {
        console.error('Failed to fetch tenants:', err);
      } finally {
        setLoading(false);
      }
    };
    fetchTenants();
  }, []);

  const handleViewStats = async (tenantId: string) => {
    try {
      const stats = await adminApi.getTenantStats(tenantId);
      setSelectedStats(stats);
    } catch (err) {
      console.error('Failed to fetch stats:', err);
    }
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center h-64">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary" />
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Tenant Management</h1>
        <p className="text-muted-foreground mt-1">View and manage all registered organizations</p>
      </div>

      <div className="grid gap-6 lg:grid-cols-3">
        {/* Tenant List */}
        <div className="lg:col-span-2">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Building2 className="h-5 w-5 text-primary" />
                All Tenants
              </CardTitle>
              <CardDescription>{tenants.length} organizations registered</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-3">
                {tenants.map(tenant => (
                  <button
                    key={tenant.id}
                    onClick={() => handleViewStats(tenant.id)}
                    className={`w-full text-left rounded-xl border p-4 transition-all duration-200 hover:shadow-md hover:border-primary/30 ${
                      selectedStats?.tenantId === tenant.id ? 'border-primary bg-primary/5' : 'border-border/50'
                    }`}
                  >
                    <div className="flex items-center justify-between">
                      <div>
                        <h3 className="font-semibold">{tenant.name}</h3>
                        <p className="text-sm text-muted-foreground">{tenant.domain || 'No domain set'}</p>
                      </div>
                      <div className="flex items-center gap-3 text-sm">
                        <span className="flex items-center gap-1 text-muted-foreground">
                          <Users className="h-3.5 w-3.5" /> {tenant.userCount}
                        </span>
                        <span className="flex items-center gap-1 text-muted-foreground">
                          <Briefcase className="h-3.5 w-3.5" /> {tenant.positionCount}
                        </span>
                        <span className={`rounded-full px-2.5 py-0.5 text-xs font-medium ${
                          tenant.subscription?.status === 'active'
                            ? 'bg-green-100 text-green-700 dark:bg-green-900/30 dark:text-green-400'
                            : 'bg-zinc-100 text-zinc-600 dark:bg-zinc-800 dark:text-zinc-400'
                        }`}>
                          {tenant.subscription?.plan?.name ?? 'Free'}
                        </span>
                      </div>
                    </div>
                  </button>
                ))}
                {tenants.length === 0 && (
                  <p className="text-center text-muted-foreground py-8">No tenants found</p>
                )}
              </div>
            </CardContent>
          </Card>
        </div>

        {/* Stats Panel */}
        <div>
          <Card className="sticky top-6">
            <CardHeader>
              <CardTitle className="text-lg">Tenant Details</CardTitle>
              <CardDescription>
                {selectedStats ? selectedStats.tenantName : 'Select a tenant to view details'}
              </CardDescription>
            </CardHeader>
            <CardContent>
              {selectedStats ? (
                <div className="space-y-4">
                  <div className="flex items-center justify-between rounded-lg bg-muted/50 p-3">
                    <div className="flex items-center gap-2 text-sm">
                      <Users className="h-4 w-4 text-blue-500" />
                      <span>Users</span>
                    </div>
                    <span className="font-bold">{selectedStats.userCount}</span>
                  </div>
                  <div className="flex items-center justify-between rounded-lg bg-muted/50 p-3">
                    <div className="flex items-center gap-2 text-sm">
                      <Briefcase className="h-4 w-4 text-purple-500" />
                      <span>Positions</span>
                    </div>
                    <span className="font-bold">{selectedStats.positionCount}</span>
                  </div>
                  <div className="flex items-center justify-between rounded-lg bg-muted/50 p-3">
                    <div className="flex items-center gap-2 text-sm">
                      <FileText className="h-4 w-4 text-amber-500" />
                      <span>Applications</span>
                    </div>
                    <span className="font-bold">{selectedStats.applicationCount}</span>
                  </div>
                  <div className="flex items-center justify-between rounded-lg bg-muted/50 p-3">
                    <div className="flex items-center gap-2 text-sm">
                      <HardDrive className="h-4 w-4 text-emerald-500" />
                      <span>Storage</span>
                    </div>
                    <span className="font-bold">{formatBytes(selectedStats.storageUsage)}</span>
                  </div>
                </div>
              ) : (
                <p className="text-sm text-muted-foreground text-center py-8">
                  Click on a tenant from the list to see its usage statistics.
                </p>
              )}
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}
