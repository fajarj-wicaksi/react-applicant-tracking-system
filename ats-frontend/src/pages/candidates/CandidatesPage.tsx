import { useEffect, useState, useCallback } from 'react';
import { pipelineApi, Candidate } from '@/features/pipeline/api/pipeline-api';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Modal } from '@/components/ui/modal';
import { CandidateForm } from '@/features/pipeline/ui/CandidateForm';
import { Plus, Search, Edit, Trash2, Mail, Phone, ExternalLink, AlertTriangle } from 'lucide-react';

export function CandidatesPage() {
  const [candidates, setCandidates] = useState<Candidate[]>([]);
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState('');
  const [showForm, setShowForm] = useState(false);
  const [editTarget, setEditTarget] = useState<Candidate | null>(null);
  const [deleteTarget, setDeleteTarget] = useState<Candidate | null>(null);
  const [deleting, setDeleting] = useState(false);

  const fetchCandidates = useCallback(async () => {
    try {
      const data = await pipelineApi.listCandidates();
      setCandidates(data);
    } catch (err) {
      console.error('Failed to fetch candidates:', err);
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => { fetchCandidates(); }, [fetchCandidates]);

  const handleFormSuccess = () => {
    setShowForm(false);
    setEditTarget(null);
    fetchCandidates();
  };

  const handleDelete = async () => {
    if (!deleteTarget) return;
    setDeleting(true);
    try {
      await pipelineApi.deleteCandidate(deleteTarget.id);
      setDeleteTarget(null);
      fetchCandidates();
    } catch (err) {
      console.error('Failed to delete candidate:', err);
    } finally {
      setDeleting(false);
    }
  };

  const filteredCandidates = candidates.filter(c =>
    `${c.firstName} ${c.lastName}`.toLowerCase().includes(search.toLowerCase()) ||
    c.email.toLowerCase().includes(search.toLowerCase())
  );

  return (
    <div className="space-y-6">
      <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Candidates</h1>
          <p className="text-muted-foreground mt-1">Manage and view all applicants</p>
        </div>
        <Button onClick={() => { setEditTarget(null); setShowForm(true); }}>
          <Plus className="mr-2 h-4 w-4" /> Add Candidate
        </Button>
      </div>

      <div className="flex items-center gap-4 bg-card p-4 rounded-xl border shadow-sm">
        <div className="relative flex-1 max-w-md">
          <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
          <Input
            placeholder="Search by name or email..."
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
        <div className="rounded-xl border bg-card overflow-hidden shadow-sm">
          <div className="overflow-x-auto">
            <table className="w-full text-sm text-left">
              <thead className="bg-muted/50 text-muted-foreground">
                <tr>
                  <th className="px-6 py-4 font-medium">Candidate</th>
                  <th className="px-6 py-4 font-medium">Contact</th>
                  <th className="px-6 py-4 font-medium">Source</th>
                  <th className="px-6 py-4 font-medium">Applied Date</th>
                  <th className="px-6 py-4 font-medium text-right">Actions</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-border/50">
                {filteredCandidates.map(c => (
                  <tr key={c.id} className="hover:bg-muted/30 transition-colors">
                    <td className="px-6 py-4">
                      <div className="flex items-center gap-3">
                        <div className="w-10 h-10 rounded-full bg-linear-to-br from-primary to-blue-600 flex items-center justify-center text-white font-bold shrink-0">
                          {c.firstName[0]}{c.lastName?.[0] ?? ''}
                        </div>
                        <div>
                          <div className="font-medium text-foreground">{c.firstName} {c.lastName}</div>
                          {c.resumeUrl && (
                            <a href={c.resumeUrl} target="_blank" rel="noreferrer" className="text-xs text-primary hover:underline flex items-center gap-1 mt-0.5">
                              <ExternalLink className="h-3 w-3" /> View Resume
                            </a>
                          )}
                        </div>
                      </div>
                    </td>
                    <td className="px-6 py-4">
                      <div className="space-y-1 text-muted-foreground">
                        <div className="flex items-center gap-2"><Mail className="h-3 w-3" /> {c.email}</div>
                        {c.phone && <div className="flex items-center gap-2"><Phone className="h-3 w-3" /> {c.phone}</div>}
                      </div>
                    </td>
                    <td className="px-6 py-4">
                      {c.source
                        ? <span className="inline-flex items-center rounded-md bg-secondary px-2 py-1 text-xs font-medium">{c.source}</span>
                        : <span className="text-muted-foreground">—</span>}
                    </td>
                    <td className="px-6 py-4 text-muted-foreground">
                      {new Date(c.createdAt).toLocaleDateString('en-US', { year: 'numeric', month: 'short', day: 'numeric' })}
                    </td>
                    <td className="px-6 py-4 text-right">
                      <div className="flex items-center justify-end gap-2">
                        <Button variant="ghost" size="icon" className="h-8 w-8 text-muted-foreground hover:text-foreground"
                          onClick={() => { setEditTarget(c); setShowForm(true); }}>
                          <Edit className="h-4 w-4" />
                        </Button>
                        <Button variant="ghost" size="icon" className="h-8 w-8 text-red-500 hover:bg-red-500/10"
                          onClick={() => setDeleteTarget(c)}>
                          <Trash2 className="h-4 w-4" />
                        </Button>
                      </div>
                    </td>
                  </tr>
                ))}
                {filteredCandidates.length === 0 && (
                  <tr><td colSpan={5} className="px-6 py-12 text-center text-muted-foreground">No candidates found.</td></tr>
                )}
              </tbody>
            </table>
          </div>
        </div>
      )}

      {/* Create / Edit Modal */}
      <Modal
        isOpen={showForm}
        onClose={() => { setShowForm(false); setEditTarget(null); }}
        title={editTarget ? 'Edit Candidate' : 'Add New Candidate'}
        size="lg"
      >
        <CandidateForm
          initial={editTarget ?? undefined}
          onSuccess={handleFormSuccess}
          onCancel={() => { setShowForm(false); setEditTarget(null); }}
        />
      </Modal>

      {/* Delete Confirm Modal */}
      <Modal isOpen={!!deleteTarget} onClose={() => setDeleteTarget(null)} title="Delete Candidate" size="sm">
        <div className="space-y-4">
          <div className="flex items-start gap-3 p-3 rounded-lg bg-red-500/10 border border-red-500/20">
            <AlertTriangle className="h-5 w-5 text-red-500 shrink-0 mt-0.5" />
            <p className="text-sm text-red-600 dark:text-red-400">
              Delete <strong>"{deleteTarget?.firstName} {deleteTarget?.lastName}"</strong>? This cannot be undone.
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
