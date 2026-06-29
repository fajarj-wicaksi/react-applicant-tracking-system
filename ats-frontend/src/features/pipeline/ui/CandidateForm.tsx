import { useState } from 'react';
import { Label } from '@/components/ui/label';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { Loader2 } from 'lucide-react';
import { pipelineApi, Candidate } from '@/features/pipeline/api/pipeline-api';

interface CandidateFormProps {
  onSuccess: () => void;
  onCancel: () => void;
  initial?: Candidate;
}

const SOURCES = ['LinkedIn', 'Indeed', 'Referral', 'Company Website', 'Job Fair', 'Other'];

export function CandidateForm({ onSuccess, onCancel, initial }: CandidateFormProps) {
  const [form, setForm] = useState({
    firstName: initial?.firstName ?? '',
    lastName: initial?.lastName ?? '',
    email: initial?.email ?? '',
    phone: initial?.phone ?? '',
    resumeUrl: initial?.resumeUrl ?? '',
    source: initial?.source ?? '',
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const set = (key: keyof typeof form) => (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) =>
    setForm(f => ({ ...f, [key]: e.target.value }));

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!form.firstName || !form.email) { setError('First name and email are required'); return; }
    setLoading(true);
    setError('');
    try {
      if (initial?.id) {
        await pipelineApi.updateCandidate(initial.id, form);
      } else {
        await pipelineApi.createCandidate({ ...form, lastName: form.lastName });
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
        <div className="space-y-1.5">
          <Label htmlFor="firstName">First Name *</Label>
          <Input id="firstName" placeholder="John" value={form.firstName} onChange={set('firstName')} />
        </div>
        <div className="space-y-1.5">
          <Label htmlFor="lastName">Last Name</Label>
          <Input id="lastName" placeholder="Doe" value={form.lastName} onChange={set('lastName')} />
        </div>
        <div className="col-span-2 space-y-1.5">
          <Label htmlFor="email">Email *</Label>
          <Input id="email" type="email" placeholder="john.doe@example.com" value={form.email} onChange={set('email')} />
        </div>
        <div className="space-y-1.5">
          <Label htmlFor="phone">Phone</Label>
          <Input id="phone" placeholder="+62 812 3456 7890" value={form.phone} onChange={set('phone')} />
        </div>
        <div className="space-y-1.5">
          <Label htmlFor="source">Source</Label>
          <select id="source" value={form.source} onChange={set('source')}
            className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring">
            <option value="">-- Select source --</option>
            {SOURCES.map(s => <option key={s} value={s}>{s}</option>)}
          </select>
        </div>
        <div className="col-span-2 space-y-1.5">
          <Label htmlFor="resumeUrl">Resume URL</Label>
          <Input id="resumeUrl" placeholder="https://drive.google.com/..." value={form.resumeUrl} onChange={set('resumeUrl')} />
        </div>
      </div>

      {error && <p className="text-sm text-red-500">{error}</p>}

      <div className="flex justify-end gap-3 pt-2">
        <Button type="button" variant="outline" onClick={onCancel}>Cancel</Button>
        <Button type="submit" disabled={loading}>
          {loading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
          {initial?.id ? 'Save Changes' : 'Add Candidate'}
        </Button>
      </div>
    </form>
  );
}
