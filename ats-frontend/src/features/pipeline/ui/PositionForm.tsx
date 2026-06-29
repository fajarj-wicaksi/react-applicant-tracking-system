import { useState } from 'react';
import { Label } from '@/components/ui/label';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { Loader2 } from 'lucide-react';
import { pipelineApi } from '@/features/pipeline/api/pipeline-api';

interface PositionFormProps {
  onSuccess: () => void;
  onCancel: () => void;
  initial?: {
    id: string;
    title: string;
    department: string;
    location: string;
    type: string;
    description: string;
    isOpen: boolean;
  };
}

const JOB_TYPES = ['Full-time', 'Part-time', 'Contract', 'Internship', 'Remote'];

export function PositionForm({ onSuccess, onCancel, initial }: PositionFormProps) {
  const [form, setForm] = useState({
    title: initial?.title ?? '',
    department: initial?.department ?? '',
    location: initial?.location ?? '',
    type: initial?.type ?? 'Full-time',
    description: initial?.description ?? '',
    isOpen: initial?.isOpen ?? true,
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!form.title.trim()) { setError('Title is required'); return; }
    setLoading(true);
    setError('');
    try {
      if (initial?.id) {
        await pipelineApi.updatePosition(initial.id, form);
      } else {
        await pipelineApi.createPosition(form);
      }
      onSuccess();
    } catch (err: unknown) {
      setError((err as { message?: string }).message ?? 'An error occurred');
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div className="grid grid-cols-2 gap-4">
        <div className="col-span-2 space-y-1.5">
          <Label htmlFor="title">Job Title *</Label>
          <Input id="title" placeholder="e.g. Senior Software Engineer" value={form.title} onChange={e => setForm(f => ({ ...f, title: e.target.value }))} />
        </div>
        <div className="space-y-1.5">
          <Label htmlFor="department">Department</Label>
          <Input id="department" placeholder="e.g. Engineering" value={form.department} onChange={e => setForm(f => ({ ...f, department: e.target.value }))} />
        </div>
        <div className="space-y-1.5">
          <Label htmlFor="location">Location</Label>
          <Input id="location" placeholder="e.g. Jakarta, Remote" value={form.location} onChange={e => setForm(f => ({ ...f, location: e.target.value }))} />
        </div>
        <div className="space-y-1.5">
          <Label htmlFor="type">Type</Label>
          <select id="type" value={form.type} onChange={e => setForm(f => ({ ...f, type: e.target.value }))}
            className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring">
            {JOB_TYPES.map(t => <option key={t} value={t}>{t}</option>)}
          </select>
        </div>
        <div className="flex items-center gap-3 pt-6">
          <input type="checkbox" id="isOpen" checked={form.isOpen} onChange={e => setForm(f => ({ ...f, isOpen: e.target.checked }))} className="h-4 w-4 rounded accent-primary" />
          <Label htmlFor="isOpen">Open Position</Label>
        </div>
        <div className="col-span-2 space-y-1.5">
          <Label htmlFor="description">Description</Label>
          <textarea id="description" placeholder="Job description..." value={form.description}
            onChange={e => setForm(f => ({ ...f, description: e.target.value }))}
            className="flex min-h-[100px] w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring resize-none" />
        </div>
      </div>

      {error && <p className="text-sm text-red-500">{error}</p>}

      <div className="flex justify-end gap-3 pt-2">
        <Button type="button" variant="outline" onClick={onCancel}>Cancel</Button>
        <Button type="submit" disabled={loading}>
          {loading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
          {initial?.id ? 'Save Changes' : 'Create Position'}
        </Button>
      </div>
    </form>
  );
}
