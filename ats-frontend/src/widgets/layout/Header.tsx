import { Menu, User, Bell } from "lucide-react";

export function Header() {
  return (
    <header className="sticky top-0 z-50 flex h-16 w-full items-center justify-between border-b bg-background px-4 md:px-6 shadow-sm">
      <div className="flex items-center gap-4">
        <button className="md:hidden">
          <Menu className="h-6 w-6" />
        </button>
        <span className="text-xl font-bold text-primary tracking-tight md:hidden">TalentFlow</span>
      </div>
      
      <div className="flex items-center gap-4">
        <button className="relative rounded-full p-2 hover:bg-muted transition-colors">
          <Bell className="h-5 w-5 text-muted-foreground" />
          <span className="absolute top-1 right-1 h-2 w-2 rounded-full bg-destructive"></span>
        </button>
        <div className="flex items-center gap-2 rounded-full border p-1 pr-4 bg-muted/30">
          <div className="h-8 w-8 rounded-full bg-primary/20 flex items-center justify-center">
            <User className="h-4 w-4 text-primary" />
          </div>
          <span className="text-sm font-medium">Recruiter</span>
        </div>
      </div>
    </header>
  );
}
