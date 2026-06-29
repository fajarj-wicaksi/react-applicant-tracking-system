import { useSortable } from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';
import { Application } from '../api/pipeline-api';

interface KanbanCardProps {
  application: Application;
}

export function KanbanCard({ application }: KanbanCardProps) {
  const { attributes, listeners, setNodeRef, transform, transition, isDragging } = useSortable({
    id: application.id,
    data: { application },
  });

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
    opacity: isDragging ? 0.5 : 1,
  };

  const candidate = application.candidate;
  const position = application.position;

  return (
    <div
      ref={setNodeRef}
      style={style}
      {...attributes}
      {...listeners}
      className="group relative rounded-xl border border-border/60 bg-card p-4 shadow-sm 
                 hover:shadow-md hover:border-primary/30 transition-all duration-200 cursor-grab active:cursor-grabbing"
    >
      {/* Candidate Name */}
      <div className="flex items-center gap-2 mb-2">
        <div className="w-8 h-8 rounded-full bg-linear-to-br from-primary to-blue-600 flex items-center justify-center text-white text-xs font-bold shrink-0">
          {candidate ? `${candidate.firstName[0]}${candidate.lastName[0]}` : '??'}
        </div>
        <div className="min-w-0">
          <p className="text-sm font-semibold truncate">
            {candidate ? `${candidate.firstName} ${candidate.lastName}` : 'Unknown'}
          </p>
          <p className="text-xs text-muted-foreground truncate">
            {candidate?.email}
          </p>
        </div>
      </div>

      {/* Position */}
      {position && (
        <div className="mt-2 flex items-center gap-1.5">
          <span className="inline-flex items-center rounded-md bg-primary/10 px-2 py-0.5 text-xs font-medium text-primary">
            {position.title}
          </span>
        </div>
      )}

      {/* Source & Date */}
      <div className="mt-3 flex items-center justify-between text-xs text-muted-foreground">
        {candidate?.source && (
          <span className="bg-muted rounded px-1.5 py-0.5">{candidate.source}</span>
        )}
        <span>{new Date(application.appliedAt).toLocaleDateString('id-ID', { day: 'numeric', month: 'short' })}</span>
      </div>
    </div>
  );
}
