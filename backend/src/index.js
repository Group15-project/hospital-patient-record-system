import express from 'express';
import cors from 'cors';
import dotenv from 'dotenv';

dotenv.config();

const app = express();
const PORT = process.env.PORT || 5000;

// Middleware
app.use(cors({
  origin: process.env.FRONTEND_URL || 'http://localhost:5173',
  credentials: true,
}));
app.use(express.json());

// Routes
app.get('/api/health', (req, res) => {
  res.json({ status: 'Server is running', timestamp: new Date() });
});

// Mock patients endpoint (for now)
app.get('/api/patients', (req, res) => {
  res.json({
    patients: [
      { id: 1, name: 'John Doe', age: 35, phone: '555-0001', condition: 'Hypertension' },
      { id: 2, name: 'Jane Smith', age: 42, phone: '555-0002', condition: 'Diabetes' },
    ]
  });
});

app.post('/api/patients', (req, res) => {
  res.json({ message: 'Patient created', patient: req.body });
});

// Error handling
app.use((err, req, res, next) => {
  console.error(err.stack);
  res.status(500).json({ error: 'Something went wrong!' });
});

// Start server
app.listen(PORT, () => {
  console.log(`🚀 Server running on http://localhost:${PORT}`);
  console.log(`📝 API: http://localhost:${PORT}/api`);
});