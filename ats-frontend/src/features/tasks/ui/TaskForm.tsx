import { useState } from 'react';
import { Label } from '@/components/ui/label';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { Loader2 } from 'lucide-react';
import { taskApi, Task } from '@/features/tasks/api/task-api';

interface TaskFormProps {
  onSuccess: () => void;
  onCancel: () => void;
  initial?: Task;
}

const STATUSES: Task['status'][] = ['Pending', 'In Progress', 'Completed'];

export function TaskForm({ onSuccess, onCancel, initial }: TaskFormProps) {
  const [form, setForm] = useState({
    title: initial?.title ?? '',
    description: initial?.description ?? '',
    dueDate: initial?.dueDate ? initial.dueDate.slice(0, 10) : '',
    status: initial?.status ?? 'Pending' as Task['status'],
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!form.title.trim()) { setError('Title is required'); return; }
    setLoading(true);
    setError('');
    try {
      const payload = {
        title: form.title,
        description: form.description,
        status: form.status,
        dueDate: form.dueDate || undefined,
      };
      if (initial?.id) {
        await taskApi.updateTask(initial.id, payload);
      } else {
        await taskApi.createTask(payload);
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
      <div className="space-y-1.5">
        <Label htmlFor="taskTitle">Task Title *</Label>
        <Input id="taskTitle" placeholder="e.g. Follow up with candidate" value={form.title}
          onChange={e => setForm(f => ({ ...f, title: e.target.value }))} />
      </div>
      <div className="space-y-1.5">
        <Label htmlFor="taskDesc">Description</Label>
        <textarea id="taskDesc" placeholder="Task details..." value={form.description}
          onChange={e => setForm(f => ({ ...f, description: e.target.value }))}
          className="flex min-h-[80px] w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring resize-none" />
      </div>
      <div className="grid grid-cols-2 gap-4">
        <div className="space-y-1.5">
          <Label htmlFor="dueDate">Due Date</Label>
          <Input id="dueDate" type="date" value={form.dueDate}
            onChange={e => setForm(f => ({ ...f, dueDate: e.target.value }))} />
        </div>
        <div className="space-y-1.5">
          <Label htmlFor="taskStatus">Status</Label>
          <select id="taskStatus" value={form.status} onChange={e => setForm(f => ({ ...f, status: e.target.value as Task['status'] }))}
            className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring">
            {STATUSES.map(s => <option key={s} value={s}>{s}</option>)}
          </select>
        </div>
      </div>

      {error && <p className="text-sm text-red-500">{error}</p>}

      <div className="flex justify-end gap-3 pt-2">
        <Button type="button" variant="outline" onClick={onCancel}>Cancel</Button>
        <Button type="submit" disabled={loading}>
          {loading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
          {initial?.id ? 'Save Changes' : 'Create Task'}
        </Button>
      </div>
    </form>
  );
}
