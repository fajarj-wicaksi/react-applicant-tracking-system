import { useState } from 'react';
import { Label } from '@/components/ui/label';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { Loader2 } from 'lucide-react';
import { userApi, User, CreateUserRequest, UpdateUserRequest } from '@/features/users/api/user-api';

interface UserFormProps {
  onSuccess: () => void;
  onCancel: () => void;
  initial?: User;
}

export function UserForm({ onSuccess, onCancel, initial }: UserFormProps) {
  const [form, setForm] = useState({
    firstName: initial?.firstName ?? '',
    lastName: initial?.lastName ?? '',
    email: initial?.email ?? '',
    password: '',
    roleId: '00000000-0000-0000-0000-000000000002', // default: user role
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const set = (key: keyof typeof form) => (e: React.ChangeEvent<HTMLInputElement>) =>
    setForm(f => ({ ...f, [key]: e.target.value }));

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!form.firstName || !form.email) { setError('First name and email are required'); return; }
    if (!initial && !form.password) { setError('Password is required for new users'); return; }
    setLoading(true);
    setError('');
    try {
      if (initial?.id) {
        const data: UpdateUserRequest = {
          firstName: form.firstName,
          lastName: form.lastName,
          email: form.email,
          roleId: form.roleId,
        };
        if (form.password) data.password = form.password;
        await userApi.updateUser(initial.id, data);
      } else {
        const data: CreateUserRequest = {
          firstName: form.firstName,
          lastName: form.lastName,
          email: form.email,
          password: form.password,
          roleId: form.roleId,
        };
        await userApi.createUser(data);
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
          <Input id="email" type="email" placeholder="john@company.com" value={form.email} onChange={set('email')} />
        </div>
        <div className="col-span-2 space-y-1.5">
          <Label htmlFor="password">{initial ? 'New Password (leave blank to keep current)' : 'Password *'}</Label>
          <Input id="password" type="password" placeholder="Min. 6 characters" value={form.password} onChange={set('password')} />
        </div>
      </div>

      {error && <p className="text-sm text-red-500">{error}</p>}

      <div className="flex justify-end gap-3 pt-2">
        <Button type="button" variant="outline" onClick={onCancel}>Cancel</Button>
        <Button type="submit" disabled={loading}>
          {loading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
          {initial?.id ? 'Save Changes' : 'Create User'}
        </Button>
      </div>
    </form>
  );
}
