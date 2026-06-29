import { useEffect, useState, useCallback } from 'react';
import { taskApi, Task } from '@/features/tasks/api/task-api';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Modal } from '@/components/ui/modal';
import { TaskForm } from '@/features/tasks/ui/TaskForm';
import { Plus, Clock, CheckCircle2, Circle, Edit, Trash2, AlertTriangle, Calendar } from 'lucide-react';

export function TasksPage() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [editTarget, setEditTarget] = useState<Task | null>(null);
  const [deleteTarget, setDeleteTarget] = useState<Task | null>(null);
  const [deleting, setDeleting] = useState(false);

  const fetchTasks = useCallback(async () => {
    const data = await taskApi.listTasks();
    setTasks(data);
    setLoading(false);
  }, []);

  useEffect(() => { fetchTasks(); }, [fetchTasks]);

  const handleFormSuccess = () => {
    setShowForm(false);
    setEditTarget(null);
    fetchTasks();
  };

  const toggleComplete = async (task: Task) => {
    const newStatus = task.status === 'Completed' ? 'Pending' : 'Completed';
    try {
      await taskApi.updateTask(task.id, { status: newStatus });
      fetchTasks();
    } catch (err) {
      console.error('Failed to update task:', err);
    }
  };

  const handleDelete = async () => {
    if (!deleteTarget) return;
    setDeleting(true);
    try {
      await taskApi.deleteTask(deleteTarget.id);
      setDeleteTarget(null);
      fetchTasks();
    } catch (err) {
      console.error('Failed to delete task:', err);
    } finally {
      setDeleting(false);
    }
  };

  const pendingTasks = tasks.filter(t => t.status !== 'Completed');
  const completedTasks = tasks.filter(t => t.status === 'Completed');

  const TaskRow = ({ task }: { task: Task }) => (
    <div className="flex items-start gap-3 p-3 rounded-lg border hover:bg-muted/50 transition-colors group">
      <button onClick={() => toggleComplete(task)}
        className={`mt-0.5 transition-colors shrink-0 ${task.status === 'Completed' ? 'text-green-500' : 'text-muted-foreground hover:text-green-500'}`}>
        {task.status === 'Completed' ? <CheckCircle2 className="h-5 w-5" /> : <Circle className="h-5 w-5" />}
      </button>
      <div className="flex-1 min-w-0">
        <p className={`text-sm font-medium ${task.status === 'Completed' ? 'line-through text-muted-foreground' : ''}`}>{task.title}</p>
        {task.description && <p className="text-xs text-muted-foreground mt-0.5 line-clamp-1">{task.description}</p>}
        {task.dueDate && (
          <div className="flex items-center gap-1 text-xs text-muted-foreground mt-1">
            <Calendar className="h-3 w-3" /> Due: {new Date(task.dueDate).toLocaleDateString()}
          </div>
        )}
      </div>
      <div className="flex gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
        <Button variant="ghost" size="icon" className="h-7 w-7 text-muted-foreground hover:text-foreground"
          onClick={() => { setEditTarget(task); setShowForm(true); }}>
          <Edit className="h-3.5 w-3.5" />
        </Button>
        <Button variant="ghost" size="icon" className="h-7 w-7 text-red-500 hover:bg-red-500/10"
          onClick={() => setDeleteTarget(task)}>
          <Trash2 className="h-3.5 w-3.5" />
        </Button>
      </div>
    </div>
  );

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Tasks</h1>
          <p className="text-muted-foreground mt-1">Manage your to-dos and recruitment tasks</p>
        </div>
        <Button onClick={() => { setEditTarget(null); setShowForm(true); }}>
          <Plus className="mr-2 h-4 w-4" /> Add Task
        </Button>
      </div>

      {loading ? (
        <div className="flex justify-center p-8">
          <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary" />
        </div>
      ) : (
        <div className="grid gap-6 md:grid-cols-2">
          <Card>
            <CardHeader>
              <CardTitle className="text-lg flex items-center gap-2">
                <Clock className="h-5 w-5 text-amber-500" />
                Active Tasks
                <span className="ml-auto text-sm font-normal text-muted-foreground">{pendingTasks.length} tasks</span>
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-3">
              {pendingTasks.map(task => <TaskRow key={task.id} task={task} />)}
              {pendingTasks.length === 0 && (
                <p className="text-sm text-muted-foreground text-center py-8">You're all caught up! 🎉</p>
              )}
            </CardContent>
          </Card>

          <Card className="opacity-80 hover:opacity-100 transition-opacity">
            <CardHeader>
              <CardTitle className="text-lg flex items-center gap-2">
                <CheckCircle2 className="h-5 w-5 text-green-500" />
                Completed
                <span className="ml-auto text-sm font-normal text-muted-foreground">{completedTasks.length} tasks</span>
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-3">
              {completedTasks.map(task => <TaskRow key={task.id} task={task} />)}
              {completedTasks.length === 0 && (
                <p className="text-sm text-muted-foreground text-center py-8">No completed tasks yet.</p>
              )}
            </CardContent>
          </Card>
        </div>
      )}

      {/* Create / Edit Modal */}
      <Modal
        isOpen={showForm}
        onClose={() => { setShowForm(false); setEditTarget(null); }}
        title={editTarget ? 'Edit Task' : 'Create New Task'}
        size="md"
      >
        <TaskForm
          initial={editTarget ?? undefined}
          onSuccess={handleFormSuccess}
          onCancel={() => { setShowForm(false); setEditTarget(null); }}
        />
      </Modal>

      {/* Delete Confirm Modal */}
      <Modal isOpen={!!deleteTarget} onClose={() => setDeleteTarget(null)} title="Delete Task" size="sm">
        <div className="space-y-4">
          <div className="flex items-start gap-3 p-3 rounded-lg bg-red-500/10 border border-red-500/20">
            <AlertTriangle className="h-5 w-5 text-red-500 shrink-0 mt-0.5" />
            <p className="text-sm text-red-600 dark:text-red-400">
              Delete <strong>"{deleteTarget?.title}"</strong>? This cannot be undone.
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
