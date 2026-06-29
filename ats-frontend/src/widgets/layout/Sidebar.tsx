import { LayoutDashboard, Users, Briefcase, Calendar, CheckSquare, Settings, Kanban, Shield, Building2, CreditCard } from "lucide-react";
import { Link, useLocation } from "react-router-dom";
import { cn } from "@/lib/utils";
import { useAuth } from "@/app/providers/AuthProvider";

const mainNavItems = [
  { icon: LayoutDashboard, label: "Dashboard", href: "/" },
  { icon: Briefcase, label: "Jobs", href: "/jobs" },
  { icon: Users, label: "Candidates", href: "/candidates" },
  { icon: Kanban, label: "Pipeline", href: "/pipeline" },
  { icon: Calendar, label: "Interviews", href: "/interviews" },
  { icon: CheckSquare, label: "Tasks", href: "/tasks" },
];

const adminNavItems = [
  { icon: Shield, label: "Admin Dashboard", href: "/admin" },
  { icon: Building2, label: "Tenants", href: "/admin/tenants" },
  { icon: CreditCard, label: "Billing", href: "/admin/billing" },
];

export function Sidebar() {
  const location = useLocation();
  const { user, logout } = useAuth();

  return (
    <aside className="hidden w-64 flex-col border-r bg-card/50 md:flex h-screen sticky top-0 backdrop-blur-sm">
      <div className="flex h-16 items-center px-6 border-b">
        <span className="text-2xl font-bold bg-linear-to-r from-primary to-blue-600 bg-clip-text text-transparent">TalentFlow</span>
      </div>
      
      <nav className="flex-1 space-y-1 p-4 overflow-y-auto">
        {/* Main Navigation */}
        <p className="px-3 mb-2 text-xs font-semibold text-muted-foreground uppercase tracking-wider">Main</p>
        {mainNavItems.map((item) => (
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

        {/* Admin Section */}
        <div className="pt-4">
          <p className="px-3 mb-2 text-xs font-semibold text-muted-foreground uppercase tracking-wider">Admin</p>
          {adminNavItems.map((item) => (
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
        </div>
      </nav>

      <div className="p-4 border-t mt-auto space-y-2">
        {user && (
          <div className="flex items-center gap-3 px-3 py-2">
            <div className="w-8 h-8 rounded-full bg-linear-to-br from-primary to-blue-600 flex items-center justify-center text-white text-xs font-bold">
              {user.firstName?.[0]}{user.lastName?.[0]}
            </div>
            <div className="min-w-0">
              <p className="text-sm font-medium truncate">{user.firstName} {user.lastName}</p>
              <p className="text-xs text-muted-foreground truncate">{user.role}</p>
            </div>
          </div>
        )}
        <Link
          to="/settings/users"
          className="flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium text-muted-foreground hover:bg-muted hover:text-foreground transition-all"
        >
          <Settings className="h-5 w-5" />
          Settings & Users
        </Link>
        <button
          onClick={logout}
          className="flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium text-red-500 hover:bg-red-500/10 hover:text-red-600 transition-all w-full text-left"
        >
          Logout
        </button>
      </div>
    </aside>
  );
}
