# TalentFlow ATS

TalentFlow ATS is a modern multi-tenant Applicant Tracking System (ATS) designed to streamline the recruitment lifecycle from candidate application to hiring.

The platform provides candidate management, recruitment pipeline tracking, interview scheduling, task management, email automation, document management, analytics, and role-based access control within a scalable SaaS architecture.

---

## Features

### Candidate Management
- Candidate profile management
- Resume and document storage
- Candidate notes and activity timeline
- Search, filter, and pagination

### Recruitment Pipeline
- Kanban-style recruitment workflow
- Drag & drop stage movement
- Candidate status tracking
- Recruitment history

### Interview Management
- Interview scheduling
- Interview feedback collection
- Interview history tracking

### Task Management
- Assign recruitment tasks
- Priority and due date tracking
- Recruiter collaboration

### Email Automation
- Gmail integration
- Email-to-Candidate workflow
- Automatic candidate creation from inbound email
- Attachment processing

### Multi-Tenant SaaS
- Tenant isolation
- Organization-based data separation
- Role-based access control (RBAC)

### Dashboard & Analytics
- Recruitment metrics
- Hiring funnel visualization
- Candidate statistics

### Audit & Security
- Audit logs
- JWT Authentication
- Permission-based authorization
- Activity tracking

---

## System Architecture

```text
Frontend (React)
       |
       v
Backend API (Golang + Gin)
       |
       +---- PostgreSQL
       |
       +---- MinIO
       |
       +---- Gmail API
```

---

## Technology Stack

### Frontend

- React
- TypeScript
- Vite
- Tailwind CSS
- shadcn/ui
- TanStack Query
- TanStack Table
- React Hook Form
- Zod
- DND Kit

### Backend

- Golang
- Gin
- GORM
- PostgreSQL
- JWT Authentication

### Infrastructure

- Docker
- Docker Compose
- Nginx

### Storage

- MinIO

### Automation

- Gmail API

### Testing

- Vitest
- React Testing Library
- Testify
- Playwright

---

## Project Structure

### Frontend

```text
frontend/
└── src/
    ├── app/
    ├── pages/
    ├── widgets/
    ├── features/
    ├── entities/
    └── shared/
```

### Backend

```text
backend/
├── cmd/
├── internal/
│   ├── domain/
│   ├── application/
│   ├── infrastructure/
│   └── interfaces/
├── migrations/
└── pkg/
```

---

## Core Modules

| Module | Status |
|----------|----------|
| Authentication | 🚧 |
| Multi Tenant | 🚧 |
| User Management | 🚧 |
| RBAC | 🚧 |
| Candidate Management | 🚧 |
| Recruitment Pipeline | 🚧 |
| Interview Management | 🚧 |
| Task Management | 🚧 |
| Gmail Integration | 🚧 |
| Analytics Dashboard | 🚧 |
| Notifications | 🚧 |

---

## Database

Main entities:

- Tenants
- Users
- Roles
- Permissions
- Positions
- Candidates
- Candidate Documents
- Candidate Emails
- Candidate Notes
- Interviews
- Interview Feedbacks
- Tasks
- Activities
- Notifications
- Audit Logs
- Subscriptions

Database engine:

```text
PostgreSQL
```

---

## Getting Started

### Prerequisites

- Go 1.24+
- Node.js 22+
- PostgreSQL 16+
- Docker
- Docker Compose

---

### Clone Repository

```bash
git clone https://github.com/your-username/talentflow-ats.git

cd talentflow-ats
```

---

### Backend Setup

```bash
cd backend

go mod tidy

cp .env.example .env

go run cmd/main.go
```

---

### Frontend Setup

```bash
cd frontend

npm install

npm run dev
```

---

### Docker Setup

```bash
docker compose up -d
```

---

## Environment Variables

Backend:

```env
APP_ENV=development

DB_HOST=localhost
DB_PORT=5432
DB_NAME=talentflow
DB_USER=postgres
DB_PASSWORD=postgres

JWT_SECRET=your-secret

MINIO_ENDPOINT=localhost:9000
MINIO_ACCESS_KEY=minioadmin
MINIO_SECRET_KEY=minioadmin

GMAIL_CLIENT_ID=
GMAIL_CLIENT_SECRET=
```

---

## API Documentation

Swagger:

```text
http://localhost:8080/swagger/index.html
```

---

## Testing

### Backend

```bash
go test ./...
```

### Frontend

```bash
npm run test
```

### E2E

```bash
npx playwright test
```

---

## Roadmap

### MVP

- [ ] Authentication
- [ ] Candidate Management
- [ ] Recruitment Pipeline
- [ ] Interview Management
- [ ] Task Management

### Phase 2

- [ ] Gmail Automation
- [ ] Notifications
- [ ] Dashboard Analytics

### Phase 3

- [ ] Subscription Billing
- [ ] AI Resume Parsing
- [ ] AI Candidate Scoring

---

## Author

Fajar Januar

Salesforce Consultant | Full Stack Developer

Skills:

- Salesforce
- Golang
- React
- PostgreSQL
- Laravel
- SQL
- System Analysis
- Software Architecture

---

## License

This project is created for educational, portfolio, and demonstration purposes.
