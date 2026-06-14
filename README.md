# Hospital Patient Record System

A full-stack healthcare patient record management system built with HTML/CSS/JavaScript frontend and Golang backend, deployed on Render.
Live Link: https://hospital-patient-record-system-1-omgx.onrender.com/index.html
## Features

### Frontend
- Patient management dashboard
- Appointment scheduling
- Medical records viewing
- Role-based access control
- Responsive design

### Backend
- RESTful API (Golang)
- Authentication & Authorization
- Patient management endpoints
- Appointment management
- Medical records storage
- Database persistence

## Project Structure
```
hospital-patient-record-system/
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   ├── pages/
│   │   ├── services/
│   │   └── styles/
│   ├── index.html
│   ├── main.js
│   ├── package.json
│   └── vite.config.js
├── backend/
│   ├── main.go
│   ├── handlers/
│   ├── models/
│   ├── routes/
│   ├── middleware/
│   ├── go.mod
│   └── go.sum
├── .github/
│   └── workflows/
│       └── deploy.yml
└── README.md
```

### Frontend (HTML/CSS/JavaScript)
- src/ - Source code directory
- components/ - Reusable UI components
- pages/ - Page components
- services/ - API service calls
- styles/ - CSS stylesheets
- index.html - Main HTML file
- main.js - JavaScript entry point

### Backend (Golang)
- main.go - Application entry point
- handlers/ - Request handlers
- models/ - Data models
- routes/ - API route definitions
- middleware/ - Custom middleware
- go.mod - Go module file
- go.sum - Go dependencies lock file

### DevOps
- .github/workflows/ - GitHub Actions workflows
- deploy.yml - CI/CD pipeline configuration
  
## Quick Start

### Prerequisites
- Golang 1.19+ installed
- Node.js 16+ installed 
- Git installed
- PostgreSQL 12+ (optional for local testing)

### Installation

#### 1. Clone the repository

```bash
git clone https://github.com/Group15-project/hospital-patient-record-system.git
cd hospital-patient-record-system
```

#### 2. Backend Setup (Golang)

```bash
cd backend

# Install dependencies
go mod download

# Create .env file
touch .env

# Configure environment variables (see below)

# Run the server
go run main.go

# Backend runs on http://localhost:8080
```

#### 3. Frontend Setup (HTML/CSS/JavaScript)

```bash
cd frontend

# If using a simple HTTP server
python -m http.server 8000
# OR
npx http-server

# Frontend runs on http://localhost:8000
```

### First Test

- Open http://localhost:8000 in browser
- You should see the Patient Dashboard
- It fetches data from http://localhost:8080/api/patients
- If you see patient data, everything works!

## Development Workflow

### Creating a Feature

#### 1. Update main branch

```bash
git pull origin main
```

#### 2. Create feature branch

```bash
git checkout -b feature/your-feature-name
```

#### 3. Make changes

- If frontend: Edit .html, .css, .js files
- If backend: Edit .go files

#### 4. Test locally

```bash
# Backend
cd backend && go run main.go

# Frontend (new terminal)
cd frontend && python -m http.server 8000
```

#### 5. Commit changes

```bash
git add .
git commit -m "feat: add patient search functionality"
```

#### 6. Push to GitHub

```bash
git push origin feature/your-feature-name
```

#### 7. Create Pull Request

- Go to GitHub
- Click "Compare & pull request"
- Describe your changes
- Submit PR

#### 8. Review & Merge

- Team reviews your code
- You make fixes if needed
- Click "Merge pull request"
- GitHub Actions runs tests automatically

#### 9. After Merge

- CI/CD pipeline runs tests
- If all pass, deployment happens
- Live URL updates with your changes

## Database Setup

### PostgreSQL

```sql
-- Create database
CREATE DATABASE hospital_records;

-- Users table
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  role VARCHAR(50),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Patients table
CREATE TABLE patients (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  age INT,
  phone VARCHAR(20),
  condition VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Appointments table
CREATE TABLE appointments (
  id SERIAL PRIMARY KEY,
  patient_id INT REFERENCES patients(id),
  doctor_id INT REFERENCES users(id),
  appointment_date TIMESTAMP,
  status VARCHAR(50),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Environment Variables

### Backend (.env)
DB_HOST=localhost
DB_PORT=5432
DB_NAME=hospital_records
DB_USER=postgres
DB_PASSWORD=your_password
PORT=8080
GIN_MODE=debug
JWT_SECRET=your_super_secret_key_change_this_in_production
FRONTEND_URL=http://localhost:8000

### Frontend (in index.html or config.js)

```javascript
const API_URL = 'http://localhost:8080/api';
const APP_NAME = 'Hospital Record System';
```

## API Documentation

### Base URL
http://localhost:8080/api

### Patients Endpoints

#### Get all patients
GET /api/patients
Response: { patients: [...] }

#### Get patient by ID
GET /api/patients/:id
Response: { patient: {...} }

#### Create patient
POST /api/patients
Body: { name, age, phone, condition }
Response: { patient: {...}, message: "Patient created" }

#### Update patient
PUT /api/patients/:id
Body: { name, age, phone, condition }
Response: { patient: {...}, message: "Patient updated" }

#### Delete patient
DELETE /api/patients/:id
Response: { message: "Patient deleted" }

### Appointments Endpoints

#### Get all appointments
GET /api/appointments
Response: { appointments: [...] }

#### Create appointment
POST /api/appointments
Body: { patient_id, doctor_id, appointment_date, status }
Response: { appointment: {...} }

#### Update appointment
PUT /api/appointments/:id
Body: { appointment_date, status }
Response: { appointment: {...} }

## Testing

### Backend Testing (Golang)

```bash
cd backend
go test ./...
```

### Frontend Testing (Manual)

```bash
# Test API connection
# Open browser console and run:
fetch('http://localhost:8080/api/patients')
  .then(res => res.json())
  .then(data => console.log(data))
```

## Build for Production

### Frontend

```bash
cd frontend
# If using Vite/bundler:
npm run build
# Creates optimized build in dist/

# If using plain HTML/CSS/JS:
# No build needed - files are ready to deploy
```

### Backend

```bash
cd backend
# Build Golang binary
go build -o hospital-backend main.go

# Or let Render handle the build automatically
```

## Deployment

### Live Application
- Frontend URL: https://hospital-frontend.render.com
- Backend API URL: https://hospital-backend.render.com/api
- Status: Live & Operational

### Frontend Deployment
- Platform: Render
- Deployment method: GitHub Actions (automatic)
- Trigger: Every push to main branch
- Build: HTML/CSS/JavaScript (no build needed)
- Hosted as: Static site

### Backend Deployment
- Platform: Render
- Deployment method: GitHub Actions (automatic)
- Trigger: Every push to main branch
- Language: Golang
- Build: Automatic Go build via Render
- Runtime: Golang server

### CI/CD Pipeline
- Tool: GitHub Actions
- Workflow: .github/workflows/deploy.yml
- Tests: Run on every PR
- Deployment: Automatic on merge to main

## Team Members & Responsibilities

| Name | Role | Responsibility | Tech Focus | Status |
|------|------|-----------------|-----------|--------|
| Sanni Oluwadarasimi Peter | Team Lead & Devops Lead | Project coordination, DevOps | GitHub, Git Bash, GitHub Actions, VS Code | Active |
| Umeobi Chinwendu | Assistant Lead & Frontend Developer | UI/UX, JavaScript logic, responsive design | HTML/CSS/JavaScript | Active |
| Eminola Daniels | Frontend Developer | UI/UX, JavaScript logic, responsive design | HTML/CSS/JavaScript | Active |
| Ifeoma Chukwudum | Frontend Developer | UI/UX, JavaScript logic, responsive design | HTML/CSS/JavaScript | Active |
| Chidi Stanley | Backend Developer | API design, server logic, database management | Golang, PostgreSQL | Active |
| Adeniji Adetomiwa Precious | DevOps & Infrastructure | CI/CD setup, deployment, monitoring | GitHub Actions, Render | Active |
| Pelumi Olaosebikan | Presentation/Documentation | Detailed Project Documentation, Slides Designing | Google Docs, Google Slides | Active |
| Uche Miracle | Presentation/Documentation | Detailed Project Documentation, Slides Designing | Google Docs, Google Slides | Active |
| Adejuwon Oluwaferanmi Emmanuel | Documentation | Detailed Project  Documentation | Google Docs | Active |
| Ogbu Stanley | Architectural Diagram | Architectural Diagram Designer | Canva | Active |

### Frontend Development (HTML/CSS/JavaScript)

Developed by: Umeobi Chinwendu, Eminola Daniels, Ifeoma Chukwudum

Responsible for: User interface, user experience, client-side logic

Handles: Dashboard, patient profiles, appointment calendar

Technologies: Vanilla JavaScript, CSS3, HTML5

### Backend Development (Golang)

Developed by: Chidi Stanley

Responsible for: API endpoints, business logic, database

Handles: Patient CRUD, appointments, authentication

Technologies: Golang, Gin Framework, PostgreSQL

### DevOps & Deployment

Handled by: Sanni Oluwadarasimi Peter, Adeniji Adetomiwa Precious

Responsible for: CI/CD pipeline, deployment automation, hosting

Technologies: GitHub Actions, Render, 

## Contributing

1. Create a feature branch: git checkout -b feature/your-feature
2. Make your changes
3. Test locally before pushing
4. Write clear commit messages
5. Push to your branch: git push origin feature/your-feature
6. Create a Pull Request on GitHub
7. Wait for code review from team
8. Address feedback if needed
9. Merge when approved

## Commit Message Convention

Follow this format for clear commit history:

- feat: New feature (e.g., feat: add appointment scheduling)
- fix: Bug fix (e.g., fix: resolve patient search issue)
- style: CSS/styling changes (e.g., style: update dashboard colors)
- refactor: Code refactoring (e.g., refactor: optimize API response)
- docs: Documentation (e.g., docs: update README)
- test: Tests (e.g., test: add patient CRUD tests)
- chore: Build/tooling (e.g., chore: update dependencies)

### Examples

```bash
git commit -m "feat: add patient search functionality"
git commit -m "fix: resolve CORS issue with frontend"
git commit -m "docs: update API documentation"
```

## Reporting Issues

Found a bug? Follow these steps:

1. Go to Issues tab on GitHub
2. Click New issue
3. Provide:
   - Clear title
   - Detailed description
   - Steps to reproduce
   - Expected vs actual behavior
   - Screenshots if applicable

## Additional Documentation

- Project Objectives: ./docs/PROJECT_OBJECTIVES.md
- System Architecture: ./docs/ARCHITECTURE.md
- Team Members: ./docs/TEAM_MEMBERS.md
- API Documentation: ./docs/api.md
- Database Schema: ./docs/database.md
- Deployment Guide: ./docs/DEPLOYMENT.md

## Security Considerations

- JWT Authentication: Secure token-based auth
- Password Hashing: bcryptjs/golang-crypto
- CORS Configuration: Restricted to frontend domain
- Environment Variables: Sensitive data in .env
- Role-Based Access: Different permissions per user role
- HTTPS: Enforced on production (Render)

## Communication

- Daily Standup: [Time & Platform]
- Weekly Sync: [Time & Platform]
- Chat Channel: Slack/Discord/Teams
- Repository Issues: For bug reports
- GitHub Discussions: For questions & ideas

## License

MIT License - see LICENSE file for details

## Questions?

- Ask in the team chat
- Create an issue on GitHub
- Contact your team lead
- Check documentation in /docs folder

---

## Project Information

- Status: In Development
- Last Updated: 2026
- Repository: https://github.com/Group15-project/hospital-patient-record-system
- Frontend Live: https://hospital-frontend.render.com
- Backend API: https://hospital-backend.render.com/api
- Live Link: https://hospital-patient-record-system-1-omgx.onrender.com/index.html
- Team: Group 15
