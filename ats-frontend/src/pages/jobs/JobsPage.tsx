import { useEffect, useState, useCallback } from 'react';
import { pipelineApi, Position } from '@/features/pipeline/api/pipeline-api';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Modal } from '@/components/ui/modal';
import { PositionForm } from '@/features/pipeline/ui/PositionForm';
import { Briefcase, MapPin, Plus, Building, Search, Edit, Trash2, AlertTriangle } from 'lucide-react';

export function JobsPage() {
  const [positions, setPositions] = useState<Position[]>([]);
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState('');
  const [showForm, setShowForm] = useState(false);
  const [editTarget, setEditTarget] = useState<Position | null>(null);
  const [deleteTarget, setDeleteTarget] = useState<Position | null>(null);
  const [deleting, setDeleting] = useState(false);

  const fetchPositions = useCallback(async () => {
    try {
      const data = await pipelineApi.listPositions();
      setPositions(data);
    } catch (err) {
      console.error('Failed to fetch positions:', err);
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => { fetchPositions(); }, [fetchPositions]);

  const handleFormSuccess = () => {
    setShowForm(false);
    setEditTarget(null);
    fetchPositions();
  };

  const handleDelete = async () => {
    if (!deleteTarget) return;
    setDeleting(true);
    try {
      await pipelineApi.deletePosition(deleteTarget.id);
      setDeleteTarget(null);
      fetchPositions();
    } catch (err) {
      console.error('Failed to delete position:', err);
    } finally {
      setDeleting(false);
    }
  };

  const filteredPositions = positions.filter(pos =>
    pos.title.toLowerCase().includes(search.toLowerCase()) ||
    pos.department?.toLowerCase().includes(search.toLowerCase())
  );

  return (
    <div className="space-y-6">
      <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Jobs</h1>
          <p className="text-muted-foreground mt-1">Manage open positions and job postings</p>
        </div>
        <Button onClick={() => { setEditTarget(null); setShowForm(true); }}>
          <Plus className="mr-2 h-4 w-4" /> Add Position
        </Button>
      </div>

      <div className="flex items-center gap-4 bg-card p-4 rounded-xl border shadow-sm">
        <div className="relative flex-1">
          <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
          <Input
            placeholder="Search by title or department..."
            className="pl-9 bg-muted/50 border-transparent focus-visible:ring-1"
            value={search}
            onChange={e => setSearch(e.target.value)}
          />
        </div>
      </div>

      {loading ? (
        <div className="flex justify-center p-8">
          <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary" />
        </div>
      ) : (
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          {filteredPositions.map(pos => (
            <Card key={pos.id} className="group hover:shadow-md transition-all duration-200 border-border/60 hover:border-primary/30">
              <CardHeader className="pb-3 flex flex-row items-start justify-between">
                <div>
                  <CardTitle className="text-lg group-hover:text-primary transition-colors">{pos.title}</CardTitle>
                  <CardDescription className="flex items-center gap-1 mt-1">
                    <Building className="h-3 w-3" /> {pos.department}
                  </CardDescription>
                </div>
                <div className="flex gap-1">
                  <Button
                    variant="ghost" size="icon"
                    className="h-8 w-8 text-muted-foreground hover:text-foreground"
                    onClick={() => { setEditTarget(pos); setShowForm(true); }}
                  >
                    <Edit className="h-4 w-4" />
                  </Button>
                  <Button
                    variant="ghost" size="icon"
                    className="h-8 w-8 text-red-500 hover:bg-red-500/10"
                    onClick={() => setDeleteTarget(pos)}
                  >
                    <Trash2 className="h-4 w-4" />
                  </Button>
                </div>
              </CardHeader>
              <CardContent>
                <div className="flex flex-wrap gap-2 mb-4">
                  <span className="inline-flex items-center gap-1 rounded-md bg-blue-500/10 px-2 py-1 text-xs font-medium text-blue-600 dark:text-blue-400">
                    <MapPin className="h-3 w-3" /> {pos.location}
                  </span>
                  <span className="inline-flex items-center gap-1 rounded-md bg-purple-500/10 px-2 py-1 text-xs font-medium text-purple-600 dark:text-purple-400">
                    <Briefcase className="h-3 w-3" /> {pos.type}
                  </span>
                  <span className={`inline-flex items-center rounded-md px-2 py-1 text-xs font-medium ${pos.isOpen ? 'bg-green-500/10 text-green-600 dark:text-green-400' : 'bg-zinc-500/10 text-zinc-600 dark:text-zinc-400'}`}>
                    {pos.isOpen ? 'Open' : 'Closed'}
                  </span>
                </div>
                <p className="text-sm text-muted-foreground line-clamp-2">{pos.description}</p>
                <div className="mt-4 pt-4 border-t text-xs text-muted-foreground">
                  Created {new Date(pos.createdAt).toLocaleDateString()}
                </div>
              </CardContent>
            </Card>
          ))}
          {filteredPositions.length === 0 && (
            <div className="col-span-full py-12 text-center text-muted-foreground border-2 border-dashed rounded-xl">
              No positions found. Create one to get started.
            </div>
          )}
        </div>
      )}

      {/* Create / Edit Modal */}
      <Modal
        isOpen={showForm}
        onClose={() => { setShowForm(false); setEditTarget(null); }}
        title={editTarget ? 'Edit Position' : 'Create New Position'}
        size="lg"
      >
        <PositionForm
          initial={editTarget ?? undefined}
          onSuccess={handleFormSuccess}
          onCancel={() => { setShowForm(false); setEditTarget(null); }}
        />
      </Modal>

      {/* Delete Confirm Modal */}
      <Modal isOpen={!!deleteTarget} onClose={() => setDeleteTarget(null)} title="Delete Position" size="sm">
        <div className="space-y-4">
          <div className="flex items-start gap-3 p-3 rounded-lg bg-red-500/10 border border-red-500/20">
            <AlertTriangle className="h-5 w-5 text-red-500 shrink-0 mt-0.5" />
            <p className="text-sm text-red-600 dark:text-red-400">
              Are you sure you want to delete <strong>"{deleteTarget?.title}"</strong>? This action cannot be undone.
            </p>
          </div>
          <div className="flex justify-end gap-3">
            <Button variant="outline" onClick={() => setDeleteTarget(null)}>Cancel</Button>
            <Button variant="destructive" onClick={handleDelete} disabled={deleting}>
              {deleting ? 'Deleting...' : 'Yes, Delete'}
            </Button>
          </div>
        </div>
      </Modal>
    </div>
  );
}
