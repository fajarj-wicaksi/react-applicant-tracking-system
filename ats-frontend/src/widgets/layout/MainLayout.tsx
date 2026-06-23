import { ReactNode } from "react";
import { Sidebar } from "./Sidebar";
import { Header } from "./Header";

export function MainLayout({ children }: { children: ReactNode }) {
  return (
    <div className="flex min-h-screen w-full bg-background/95">
      <Sidebar />
      <div className="flex flex-1 flex-col">
        <Header />
        <main className="flex-1 p-6 md:p-8 animate-in fade-in duration-500">
          {children}
        </main>
      </div>
    </div>
  );
}
