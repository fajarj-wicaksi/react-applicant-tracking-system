import { KanbanBoard } from '@/features/pipeline/ui/KanbanBoard';

export function PipelinePage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Pipeline</h1>
        <p className="text-muted-foreground mt-1">Drag and drop candidates between stages to manage your recruitment pipeline.</p>
      </div>
      <KanbanBoard />
    </div>
  );
}
