import { LayoutDashboard, Users, Briefcase, Calendar, CheckSquare, Settings } from "lucide-react";
import { Link, useLocation } from "react-router-dom";
import { cn } from "@/lib/utils";

const navItems = [
  { icon: LayoutDashboard, label: "Dashboard", href: "/" },
  { icon: Briefcase, label: "Jobs", href: "/jobs" },
  { icon: Users, label: "Candidates", href: "/candidates" },
  { icon: Calendar, label: "Interviews", href: "/interviews" },
  { icon: CheckSquare, label: "Tasks", href: "/tasks" },
];

export function Sidebar() {
  const location = useLocation();

  return (
    <aside className="hidden w-64 flex-col border-r bg-card/50 md:flex h-screen sticky top-0 backdrop-blur-sm">
      <div className="flex h-16 items-center px-6 border-b">
        <span className="text-2xl font-bold bg-linear-to-r from-primary to-blue-600 bg-clip-text text-transparent">TalentFlow</span>
      </div>
      
      <nav className="flex-1 space-y-1 p-4 overflow-y-auto">
        {navItems.map((item) => (
          <Link
            key={item.href}
            to={item.href}
            className={cn(
              "flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium transition-all duration-200",
              location.pathname === item.href 
                ? "bg-primary text-primary-foreground shadow-md" 
                : "text-muted-foreground hover:bg-muted hover:text-foreground"
            )}
          >
            <item.icon className="h-5 w-5" />
            {item.label}
          </Link>
        ))}
      </nav>

      <div className="p-4 border-t mt-auto">
        <Link
          to="/settings"
          className="flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium text-muted-foreground hover:bg-muted hover:text-foreground transition-all"
        >
          <Settings className="h-5 w-5" />
          Settings
        </Link>
      </div>
    </aside>
  );
}
