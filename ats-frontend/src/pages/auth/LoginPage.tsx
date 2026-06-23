import { LoginForm } from '@/features/auth/ui/LoginForm';

export function LoginPage() {
  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-zinc-50 dark:bg-zinc-950 p-4">
      {/* Background decoration */}
      <div className="absolute inset-0 z-0 overflow-hidden pointer-events-none">
        <div className="absolute -top-1/4 -right-1/4 w-1/2 h-1/2 bg-primary/20 blur-[120px] rounded-full mix-blend-multiply opacity-70"></div>
        <div className="absolute -bottom-1/4 -left-1/4 w-1/2 h-1/2 bg-blue-600/20 blur-[120px] rounded-full mix-blend-multiply opacity-70"></div>
      </div>
      
      <div className="relative z-10 w-full flex justify-center">
        <LoginForm />
      </div>
    </div>
  );
}
