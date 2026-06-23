import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { MainLayout } from "@/widgets/layout/MainLayout";

// Temporary placeholder pages
const Dashboard = () => <div><h1 className="text-3xl font-bold">Dashboard</h1><p className="text-muted-foreground mt-2">Welcome to TalentFlow ATS</p></div>;
const Jobs = () => <div><h1 className="text-3xl font-bold">Jobs</h1></div>;
const Candidates = () => <div><h1 className="text-3xl font-bold">Candidates</h1></div>;
const Interviews = () => <div><h1 className="text-3xl font-bold">Interviews</h1></div>;
const Tasks = () => <div><h1 className="text-3xl font-bold">Tasks</h1></div>;

const router = createBrowserRouter([
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
        <Jobs />
      </MainLayout>
    ),
  },
  {
    path: "/candidates",
    element: (
      <MainLayout>
        <Candidates />
      </MainLayout>
    ),
  },
  {
    path: "/interviews",
    element: (
      <MainLayout>
        <Interviews />
      </MainLayout>
    ),
  },
  {
    path: "/tasks",
    element: (
      <MainLayout>
        <Tasks />
      </MainLayout>
    ),
  },
]);

export function AppRouter() {
  return <RouterProvider router={router} />;
}
