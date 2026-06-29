import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { MainLayout } from "@/widgets/layout/MainLayout";
import { ProtectedRoute } from "./routes/ProtectedRoute";
import { LoginPage } from "@/pages/auth/LoginPage";
import { PipelinePage } from "@/pages/pipeline/PipelinePage";
import { AdminDashboard } from "@/pages/admin/AdminDashboard";
import { TenantsPage } from "@/pages/admin/TenantsPage";
import { BillingPage } from "@/pages/admin/BillingPage";
import { JobsPage } from "@/pages/jobs/JobsPage";
import { CandidatesPage } from "@/pages/candidates/CandidatesPage";
import { UsersPage } from "@/pages/users/UsersPage";
import { InterviewsPage } from "@/pages/interviews/InterviewsPage";
import { TasksPage } from "@/pages/tasks/TasksPage";

// Temporary placeholder pages
const Dashboard = () => <div><h1 className="text-3xl font-bold">Dashboard</h1><p className="text-muted-foreground mt-2">Welcome to TalentFlow ATS</p></div>;

const router = createBrowserRouter([
  {
    path: "/login",
    element: <LoginPage />,
  },
  {
    path: "/",
    element: <ProtectedRoute />,
    children: [
      {
        path: "/",
        element: (
          <MainLayout>
            <Dashboard />
          </MainLayout>
        ),
      },
      {
        path: "/jobs",
        element: (
          <MainLayout>
            <JobsPage />
          </MainLayout>
        ),
      },
      {
        path: "/candidates",
        element: (
          <MainLayout>
            <CandidatesPage />
          </MainLayout>
        ),
      },
      {
        path: "/pipeline",
        element: (
          <MainLayout>
            <PipelinePage />
          </MainLayout>
        ),
      },
      {
        path: "/interviews",
        element: (
          <MainLayout>
            <InterviewsPage />
          </MainLayout>
        ),
      },
      {
        path: "/tasks",
        element: (
          <MainLayout>
            <TasksPage />
          </MainLayout>
        ),
      },
      {
        path: "/admin",
        element: (
          <MainLayout>
            <AdminDashboard />
          </MainLayout>
        ),
      },
      {
        path: "/admin/tenants",
        element: (
          <MainLayout>
            <TenantsPage />
          </MainLayout>
        ),
      },
      {
        path: "/admin/billing",
        element: (
          <MainLayout>
            <BillingPage />
          </MainLayout>
        ),
      },
      {
        path: "/settings/users",
        element: (
          <MainLayout>
            <UsersPage />
          </MainLayout>
        ),
      },
    ],
  },
]);

export function AppRouter() {
  return <RouterProvider router={router} />;
}
