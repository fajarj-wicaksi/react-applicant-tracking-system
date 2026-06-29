import { useState, useEffect, useCallback } from 'react';
import {
  DndContext,
  DragEndEvent,
  DragOverEvent,
  DragOverlay,
  DragStartEvent,
  PointerSensor,
  useSensor,
  useSensors,
  closestCorners,
} from '@dnd-kit/core';
import { Application, ApplicationStage, pipelineApi } from '../api/pipeline-api';
import { KanbanColumn } from './KanbanColumn';
import { KanbanCard } from './KanbanCard';

const STAGES: ApplicationStage[] = ['Applied', 'Screening', 'Interview', 'Offer', 'Hired', 'Rejected'];

export function KanbanBoard() {
  const [columns, setColumns] = useState<Record<ApplicationStage, Application[]>>({
    Applied: [],
    Screening: [],
    Interview: [],
    Offer: [],
    Hired: [],
    Rejected: [],
  });
  const [activeApp, setActiveApp] = useState<Application | null>(null);
  const [loading, setLoading] = useState(true);

  const sensors = useSensors(
    useSensor(PointerSensor, { activationConstraint: { distance: 5 } })
  );

  const fetchPipeline = useCallback(async () => {
    try {
      const data = await pipelineApi.getPipeline();
      if (data.stages) {
        // Ensure all stages exist
        const merged: Record<ApplicationStage, Application[]> = {
          Applied: [],
          Screening: [],
          Interview: [],
          Offer: [],
          Hired: [],
          Rejected: [],
        };
        for (const stage of STAGES) {
          if (data.stages[stage]) {
            merged[stage] = data.stages[stage];
          }
        }
        setColumns(merged);
      }
    } catch (err) {
      console.error('Failed to fetch pipeline:', err);
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    fetchPipeline();
  }, [fetchPipeline]);

  const findApplication = (id: string): { app: Application; stage: ApplicationStage } | null => {
    for (const stage of STAGES) {
      const app = columns[stage].find(a => a.id === id);
      if (app) return { app, stage };
    }
    return null;
  };

  const handleDragStart = (event: DragStartEvent) => {
    const found = findApplication(event.active.id as string);
    if (found) setActiveApp(found.app);
  };

  const handleDragOver = (event: DragOverEvent) => {
    const { active, over } = event;
    if (!over) return;

    const activeId = active.id as string;
    const overId = over.id as string;

    const activeResult = findApplication(activeId);
    if (!activeResult) return;

    // Determine target stage
    let targetStage: ApplicationStage | null = null;
    if (STAGES.includes(overId as ApplicationStage)) {
      targetStage = overId as ApplicationStage;
    } else {
      const overResult = findApplication(overId);
      if (overResult) targetStage = overResult.stage;
    }

    if (!targetStage || targetStage === activeResult.stage) return;

    // Move between columns
    setColumns(prev => {
      const sourceItems = prev[activeResult.stage].filter(a => a.id !== activeId);
      const destItems = [...prev[targetStage!], { ...activeResult.app, stage: targetStage! }];
      return { ...prev, [activeResult.stage]: sourceItems, [targetStage!]: destItems };
    });
  };

  const handleDragEnd = async (event: DragEndEvent) => {
    setActiveApp(null);
    const { active } = event;
    const activeId = active.id as string;

    const result = findApplication(activeId);
    if (!result) return;

    // Persist stage change to backend
    try {
      await pipelineApi.updateStage(activeId, result.stage, result.app.stageOrder);
    } catch (err) {
      console.error('Failed to update stage:', err);
      // Re-fetch to reset on error
      fetchPipeline();
    }
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center h-64">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary"></div>
      </div>
    );
  }

  return (
    <DndContext
      sensors={sensors}
      collisionDetection={closestCorners}
      onDragStart={handleDragStart}
      onDragOver={handleDragOver}
      onDragEnd={handleDragEnd}
    >
      <div className="flex gap-4 overflow-x-auto pb-4 px-1">
        {STAGES.map(stage => (
          <KanbanColumn key={stage} stage={stage} applications={columns[stage]} />
        ))}
      </div>

      <DragOverlay>
        {activeApp ? <KanbanCard application={activeApp} /> : null}
      </DragOverlay>
    </DndContext>
  );
}
