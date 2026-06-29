import { useEffect, useState } from 'react';
import { Calendar, dateFnsLocalizer } from 'react-big-calendar';
import { format, parse, startOfWeek, getDay } from 'date-fns';
import { id } from 'date-fns/locale';
import 'react-big-calendar/lib/css/react-big-calendar.css';

import { interviewApi, Interview } from '@/features/interviews/api/interview-api';
import { Button } from '@/components/ui/button';
import { Plus } from 'lucide-react';

const locales = {
  'id': id,
};

const localizer = dateFnsLocalizer({
  format,
  parse,
  startOfWeek,
  getDay,
  locales,
});

export function InterviewsPage() {
  const [interviews, setInterviews] = useState<Interview[]>([]);

  useEffect(() => {
    const fetchInterviews = async () => {
      const data = await interviewApi.listInterviews();
      setInterviews(data);
    };
    fetchInterviews();
  }, []);

  const events = interviews.map(inv => {
    const start = new Date(inv.scheduledAt);
    const end = new Date(start.getTime() + inv.duration * 60000);
    return {
      id: inv.id,
      title: inv.title,
      start,
      end,
      resource: inv,
    };
  });

  return (
    <div className="space-y-6 h-[calc(100vh-6rem)] flex flex-col">
      <div className="flex justify-between items-center shrink-0">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Interviews</h1>
          <p className="text-muted-foreground mt-1">Schedule and manage candidate interviews</p>
        </div>
        <Button className="bg-primary text-primary-foreground">
          <Plus className="mr-2 h-4 w-4" /> Schedule Interview
        </Button>
      </div>

      <div className="flex-1 bg-card rounded-xl border shadow-sm p-4 min-h-0">
        <style>{`
          .rbc-calendar { font-family: inherit; }
          .rbc-header { padding: 8px; font-weight: 600; text-transform: uppercase; font-size: 0.75rem; color: hsl(var(--muted-foreground)); }
          .rbc-today { background-color: hsl(var(--primary)/0.05); }
          .rbc-event { background-color: hsl(var(--primary)); border-radius: 6px; }
        `}</style>
        <Calendar
          localizer={localizer}
          events={events}
          startAccessor="start"
          endAccessor="end"
          style={{ height: '100%' }}
          views={['month', 'week', 'day', 'agenda']}
          defaultView="week"
          tooltipAccessor={(e) => `${e.title}\nCandidate: ${(e.resource as Interview).candidate?.firstName || 'Unknown'}`}
        />
      </div>
    </div>
  );
}
