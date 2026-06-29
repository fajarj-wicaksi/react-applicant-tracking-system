import { useEffect, useState, useCallback } from 'react';
import { userApi, User } from '@/features/users/api/user-api';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Modal } from '@/components/ui/modal';
import { UserForm } from '@/features/users/ui/UserForm';
import { Search, Plus, Edit, Trash2, Shield, Mail, AlertTriangle } from 'lucide-react';

export function UsersPage() {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState('');
  const [showForm, setShowForm] = useState(false);
  const [editTarget, setEditTarget] = useState<User | null>(null);
  const [deleteTarget, setDeleteTarget] = useState<User | null>(null);
  const [deleting, setDeleting] = useState(false);

  const fetchUsers = useCallback(async () => {
    try {
      const data = await userApi.listUsers();
      setUsers(data);
    } catch (err) {
      console.error('Failed to fetch users:', err);
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => { fetchUsers(); }, [fetchUsers]);

  const handleFormSuccess = () => {
    setShowForm(false);
    setEditTarget(null);
    fetchUsers();
  };

  const handleDelete = async () => {
    if (!deleteTarget) return;
    setDeleting(true);
    try {
      await userApi.deleteUser(deleteTarget.id);
      setDeleteTarget(null);
      fetchUsers();
    } catch (err) {
      console.error('Failed to delete user:', err);
    } finally {
      setDeleting(false);
    }
  };

  const filteredUsers = users.filter(u =>
    `${u.firstName} ${u.lastName}`.toLowerCase().includes(search.toLowerCase()) ||
    u.email.toLowerCase().includes(search.toLowerCase())
  );

  return (
    <div className="space-y-6">
      <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Team Members</h1>
          <p className="text-muted-foreground mt-1">Manage users and their roles in your organization</p>
        </div>
        <Button onClick={() => { setEditTarget(null); setShowForm(true); }}>
          <Plus className="mr-2 h-4 w-4" /> Add User
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
          <table className="w-full text-sm text-left">
            <thead className="bg-muted/50 text-muted-foreground">
              <tr>
                <th className="px-6 py-4 font-medium">User</th>
                <th className="px-6 py-4 font-medium">Role</th>
                <th className="px-6 py-4 font-medium">Joined Date</th>
                <th className="px-6 py-4 font-medium text-right">Actions</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-border/50">
              {filteredUsers.map(user => (
                <tr key={user.id} className="hover:bg-muted/30 transition-colors">
                  <td className="px-6 py-4">
                    <div className="flex items-center gap-3">
                      <div className="w-10 h-10 rounded-full bg-linear-to-br from-indigo-500 to-purple-600 flex items-center justify-center text-white font-bold shrink-0">
                        {user.firstName[0]}{user.lastName?.[0] ?? ''}
                      </div>
                      <div>
                        <div className="font-medium text-foreground">{user.firstName} {user.lastName}</div>
                        <div className="flex items-center gap-1 text-xs text-muted-foreground mt-0.5">
                          <Mail className="h-3 w-3" /> {user.email}
                        </div>
                      </div>
                    </div>
                  </td>
                  <td className="px-6 py-4">
                    <span className="inline-flex items-center gap-1 rounded-md bg-primary/10 px-2 py-1 text-xs font-medium text-primary">
                      <Shield className="h-3 w-3" /> {user.roleName || 'Member'}
                    </span>
                  </td>
                  <td className="px-6 py-4 text-muted-foreground">
                    {new Date(user.createdAt).toLocaleDateString()}
                  </td>
                  <td className="px-6 py-4 text-right">
                    <div className="flex items-center justify-end gap-2">
                      <Button variant="ghost" size="icon" className="h-8 w-8 text-muted-foreground hover:text-foreground"
                        onClick={() => { setEditTarget(user); setShowForm(true); }}>
                        <Edit className="h-4 w-4" />
                      </Button>
                      <Button variant="ghost" size="icon" className="h-8 w-8 text-red-500 hover:bg-red-500/10"
                        onClick={() => setDeleteTarget(user)}>
                        <Trash2 className="h-4 w-4" />
                      </Button>
                    </div>
                  </td>
                </tr>
              ))}
              {filteredUsers.length === 0 && (
                <tr><td colSpan={4} className="px-6 py-12 text-center text-muted-foreground">No users found.</td></tr>
              )}
            </tbody>
          </table>
        </div>
      )}

      {/* Create / Edit Modal */}
      <Modal
        isOpen={showForm}
        onClose={() => { setShowForm(false); setEditTarget(null); }}
        title={editTarget ? 'Edit User' : 'Add New User'}
        size="md"
      >
        <UserForm
          initial={editTarget ?? undefined}
          onSuccess={handleFormSuccess}
          onCancel={() => { setShowForm(false); setEditTarget(null); }}
        />
      </Modal>

      {/* Delete Confirm Modal */}
      <Modal isOpen={!!deleteTarget} onClose={() => setDeleteTarget(null)} title="Remove User" size="sm">
        <div className="space-y-4">
          <div className="flex items-start gap-3 p-3 rounded-lg bg-red-500/10 border border-red-500/20">
            <AlertTriangle className="h-5 w-5 text-red-500 shrink-0 mt-0.5" />
            <p className="text-sm text-red-600 dark:text-red-400">
              Remove <strong>"{deleteTarget?.firstName} {deleteTarget?.lastName}"</strong> from the organization? This cannot be undone.
            </p>
          </div>
          <div className="flex justify-end gap-3">
            <Button variant="outline" onClick={() => setDeleteTarget(null)}>Cancel</Button>
            <Button variant="destructive" onClick={handleDelete} disabled={deleting}>
              {deleting ? 'Removing...' : 'Yes, Remove'}
            </Button>
          </div>
        </div>
      </Modal>
    </div>
  );
}
