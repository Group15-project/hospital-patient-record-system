# Hospital Patient Record System

A full-stack healthcare patient record management system built with React and Node.js.

## 🌟 Features

### Frontend
- Patient management dashboard
- Appointment scheduling
- Medical records viewing
- Role-based access control
- Responsive design
- Dark mode support

### Backend
- RESTful API
- Authentication & Authorization
- Patient management endpoints
- Appointment management
- Medical records storage
- Database persistence

## 📁 Project Structure
hospital-patient-record-system/
├── frontend/              (React app)
│   ├── src/
│   │   ├── components/
│   │   ├── pages/
│   │   ├── services/
│   │   ├── context/
│   │   ├── App.jsx
│   │   └── main.jsx
│   ├── package.json
│   ├── vite.config.js
│   └── tailwind.config.js
│
├── backend/               (Node.js API)
│   ├── src/
│   │   ├── controllers/
│   │   ├── routes/
│   │   ├── models/
│   │   ├── middleware/
│   │   └── index.js
│   ├── package.json
│   └── .env
│
├── .github/
│   └── workflows/
│       └── deploy.yml
│
└── README.md

## 🛠️ Tech Stack

### Frontend
- **React 18+** - UI framework
- **Tailwind CSS** - Styling
- **React Router** - Routing
- **Axios** - HTTP client
- **Recharts** - Data visualization
- **Vite** - Build tool

### Backend
- **Node.js** - Runtime
- **Express.js** - Web framework
- **PostgreSQL** - Database
- **JWT** - Authentication
- **bcryptjs** - Password hashing
- **CORS** - Cross-origin support

## 🚀 Quick Start

### Prerequisites
- Node.js 16+ installed
- Git installed
- (Optional) PostgreSQL 12+ for database

### Installation

1. **Clone the repository**
```bash
git clone https://github.com/YOUR-USERNAME/hospital-patient-record-system.git
cd hospital-patient-record-system
```

2. **Backend Setup**
```bash
cd backend

# Install dependencies
npm install

# Create .env file and configure
# Update with your database credentials
npm run dev

# Backend runs on http://localhost:5000
```

3. **Frontend Setup** (in new terminal)
```bash
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev

# Frontend runs on http://localhost:5173
```

### First Test

- Open http://localhost:5173 in browser
- You should see the Patient Dashboard
- It fetches data from http://localhost:5000/api/patients
- If you see patient data, everything works! ✅

## 📝 Development Workflow

### Creating a Feature

1. **Update main branch**
```bash
git pull origin main
```

2. **Create feature branch**
```bash
git checkout -b feature/your-feature-name
```

3. **Make changes**
- If frontend: `cd frontend && npm run dev`
- If backend: `cd backend && npm run dev`

4. **Commit changes**
```bash
git add .
git commit -m "feat: add patient search functionality"
```

5. **Push to GitHub**
```bash
git push origin feature/your-feature-name
```

6. **Create Pull Request**
- Go to GitHub
- Click "Compare & pull request"
- Describe your changes
- Submit PR

7. **Review & Merge**
- Team reviews your code
- You make fixes if needed
- Click "Merge pull request"
- GitHub Actions runs tests automatically

### After Merge
- CI/CD pipeline runs tests
- If all pass, deployment happens
- Live URL updates with your changes

## 🗄️ Database Setup (Optional)

If using PostgreSQL:

```sql
CREATE DATABASE hospital_records;

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  role VARCHAR(50),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE patients (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  age INT,
  phone VARCHAR(20),
  condition VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE appointments (
  id SERIAL PRIMARY KEY,
  patient_id INT REFERENCES patients(id),
  doctor_id INT REFERENCES users(id),
  appointment_date TIMESTAMP,
  status VARCHAR(50),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 🔐 Environment Variables

### Backend (.env)
DB_HOST=localhost
DB_PORT=5432
DB_NAME=hospital_records
DB_USER=postgres
DB_PASSWORD=your_password
PORT=5000
NODE_ENV=development
JWT_SECRET=your_secret_key
FRONTEND_URL=http://localhost:5173

### Frontend (.env)
VITE_API_URL=http://localhost:5000/api
VITE_APP_NAME=Hospital Record System

## 📊 API Documentation

### Patients Endpoints

**Get all patients**
GET /api/patients

**Get patient by ID**
GET /api/patients/:id

**Create patient**
POST /api/patients
Body: { name, age, phone, condition }

**Update patient**
PUT /api/patients/:id
Body: { name, age, phone, condition }

**Delete patient**
DELETE /api/patients/:id

### Appointments Endpoints

**Get all appointments**
GET /api/appointments

**Create appointment**
POST /api/appointments
Body: { patient_id, doctor_id, appointment_date, status }

**Update appointment**
PUT /api/appointments/:id
Body: { appointment_date, status }

## 🧪 Testing

### Frontend Testing
```bash
cd frontend
npm test
```

### Backend Testing
```bash
cd backend
npm test
```

## 📦 Build for Production

### Frontend
```bash
cd frontend
npm run build
# Creates optimized build in dist/
```

### Backend
```bash
cd backend
npm run start
# Runs production server
```

## 🚀 Deployment

### Frontend Deployment
- Deployed to Render or Vercel
- Automatic on every push to main
- Live URL: [Will be set after deployment]

### Backend Deployment
- Deployed to Render or Railway
- Automatic on every push to main
- API URL: [Will be set after deployment]

## 👥 Team Members

| Name | Role | GitHub |
|------|------|--------|
| [Your Name] | Frontend Lead | @username |
| [Team Member 2] | Backend Dev | @username |
| [Team Member 3] | Feature Dev | @username |
| [Team Member 4] | DevOps Lead | @username |

## 🤝 Contributing

1. Create a feature branch
2. Make your changes
3. Write clear commit messages
4. Push to your branch
5. Create a Pull Request
6. Wait for review and approval

## 📝 Commit Message Convention

- `feat:` - New feature
- `fix:` - Bug fix
- `style:` - Styling changes
- `refactor:` - Code refactoring
- `docs:` - Documentation
- `test:` - Tests

Example: `git commit -m "feat: add patient search"`

## 🐛 Reporting Issues

Found a bug? Create an issue:
1. Go to Issues tab
2. Click "New issue"
3. Describe the problem
4. Add steps to reproduce

## 📚 Documentation

- [Architecture Diagram](./docs/architecture.md)
- [API Documentation](./docs/api.md)
- [Component Library](./docs/components.md)
- [Database Schema](./docs/database.md)

## 📄 License

MIT License - see LICENSE file for details

## 💬 Questions?

Ask in the team chat or create an issue on GitHub.

---

**Last updated:** 2024
**Project Status:** In Development