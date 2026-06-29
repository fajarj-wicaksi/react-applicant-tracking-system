import { useDroppable } from '@dnd-kit/core';
import { SortableContext, verticalListSortingStrategy } from '@dnd-kit/sortable';
import { Application, ApplicationStage } from '../api/pipeline-api';
import { KanbanCard } from './KanbanCard';

const stageConfig: Record<ApplicationStage, { label: string; color: string; bgColor: string }> = {
  Applied: { label: 'Applied', color: 'text-blue-600', bgColor: 'bg-blue-500/10 border-blue-500/20' },
  Screening: { label: 'Screening', color: 'text-amber-600', bgColor: 'bg-amber-500/10 border-amber-500/20' },
  Interview: { label: 'Interview', color: 'text-purple-600', bgColor: 'bg-purple-500/10 border-purple-500/20' },
  Offer: { label: 'Offer', color: 'text-emerald-600', bgColor: 'bg-emerald-500/10 border-emerald-500/20' },
  Hired: { label: 'Hired', color: 'text-green-600', bgColor: 'bg-green-500/10 border-green-500/20' },
  Rejected: { label: 'Rejected', color: 'text-red-600', bgColor: 'bg-red-500/10 border-red-500/20' },
};

interface KanbanColumnProps {
  stage: ApplicationStage;
  applications: Application[];
}

export function KanbanColumn({ stage, applications }: KanbanColumnProps) {
  const { setNodeRef, isOver } = useDroppable({ id: stage });
  const config = stageConfig[stage];

  return (
    <div
      ref={setNodeRef}
      className={`flex flex-col min-w-[280px] w-[300px] rounded-xl border transition-all duration-200 ${
        isOver ? 'border-primary/50 bg-primary/5 shadow-lg' : 'border-border/50 bg-muted/30'
      }`}
    >
      {/* Column Header */}
      <div className={`flex items-center justify-between rounded-t-xl px-4 py-3 border-b ${config.bgColor}`}>
        <div className="flex items-center gap-2">
          <span className={`text-sm font-bold ${config.color}`}>{config.label}</span>
        </div>
        <span className={`inline-flex items-center justify-center rounded-full w-6 h-6 text-xs font-bold ${config.color} bg-white/80 dark:bg-zinc-900/60`}>
          {applications.length}
        </span>
      </div>

      {/* Cards */}
      <div className="flex-1 p-2 space-y-2 overflow-y-auto max-h-[calc(100vh-220px)]">
        <SortableContext items={applications.map(a => a.id)} strategy={verticalListSortingStrategy}>
          {applications.map(app => (
            <KanbanCard key={app.id} application={app} />
          ))}
        </SortableContext>

        {applications.length === 0 && (
          <div className="flex items-center justify-center h-24 rounded-lg border-2 border-dashed border-border/40 text-xs text-muted-foreground">
            Drop candidate here
          </div>
        )}
      </div>
    </div>
  );
}
