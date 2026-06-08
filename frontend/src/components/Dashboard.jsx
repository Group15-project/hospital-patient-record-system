import { useState, useEffect } from 'react';
import {
  LineChart,
  Line,
  BarChart,
  Bar,
  PieChart,
  Pie,
  Cell,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from 'recharts';
import { usePatients } from '../context/PatientContext';
import { useAuth } from '../context/AuthContext';
import api from '../services/api';
import { LoadingState, ErrorMessage } from './LoadingAndErrors';

export const Dashboard = () => {
  const { patients, loading: patientsLoading, fetchAllPatients } = usePatients();
  const { user } = useAuth();
  const [appointments, setAppointments] = useState([]);
  const [records, setRecords] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        await fetchAllPatients();
        const [appointmentsRes, recordsRes] = await Promise.all([
          api.get('/appointments'),
          api.get('/medical-records'),
        ]);
        setAppointments(appointmentsRes.data.appointments || []);
        setRecords(recordsRes.data.records || []);
        setError(null);
      } catch (err) {
        setError('Failed to load dashboard data');
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading || patientsLoading) {
    return <LoadingState message="Loading dashboard..." />;
  }

  return (
    <div className="space-y-6">
      <div>
        <h2 className="text-3xl font-bold text-gray-900">Dashboard</h2>
        {user && <p className="text-gray-600">Welcome, {user.name}!</p>}
      </div>

      {error && <ErrorMessage message={error} />}

      {/* Stats Cards */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
        <StatCard
          title="Total Patients"
          value={patients.length}
          icon="👥"
          color="bg-blue-500"
        />
        <StatCard
          title="Upcoming Appointments"
          value={appointments.filter(a => {
            const date = new Date(a.date);
            return date >= new Date() && a.status === 'scheduled';
          }).length}
          icon="📅"
          color="bg-green-500"
        />
        <StatCard
          title="Medical Records"
          value={records.length}
          icon="📋"
          color="bg-purple-500"
        />
        <StatCard
          title="Active Cases"
          value={patients.filter(p => p.status === 'active').length}
          icon="⚕️"
          color="bg-orange-500"
        />
      </div>

      {/* Charts Section */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {/* Patients by Status - Pie Chart */}
        <div className="bg-white rounded-lg shadow p-6">
          <h3 className="text-lg font-bold text-gray-900 mb-4">Patients by Status</h3>
          <PatientStatusChart patients={patients} />
        </div>

        {/* Appointments Trend - Line Chart */}
        <div className="bg-white rounded-lg shadow p-6">
          <h3 className="text-lg font-bold text-gray-900 mb-4">Appointments Trend</h3>
          <AppointmentsTrendChart appointments={appointments} />
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {/* Medical Records by Severity - Bar Chart */}
        <div className="bg-white rounded-lg shadow p-6">
          <h3 className="text-lg font-bold text-gray-900 mb-4">Medical Records by Severity</h3>
          <RecordsSeverityChart records={records} />
        </div>

        {/* Appointment Status Distribution */}
        <div className="bg-white rounded-lg shadow p-6">
          <h3 className="text-lg font-bold text-gray-900 mb-4">Appointment Status</h3>
          <AppointmentStatusChart appointments={appointments} />
        </div>
      </div>

      {/* Recent Activity */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <RecentPatients patients={patients} />
        <RecentAppointments appointments={appointments} />
      </div>
    </div>
  );
};

const StatCard = ({ title, value, icon, color }) => {
  return (
    <div className="bg-white rounded-lg shadow p-6 hover:shadow-lg transition">
      <div className="flex justify-between items-start">
        <div>
          <p className="text-gray-600 text-sm font-medium mb-1">{title}</p>
          <p className="text-3xl font-bold text-gray-900">{value}</p>
        </div>
        <span className={`${color} w-12 h-12 rounded-lg flex items-center justify-center text-2xl`}>
          {icon}
        </span>
      </div>
    </div>
  );
};

const PatientStatusChart = ({ patients }) => {
  const statusData = [
    { name: 'Active', value: patients.filter(p => p.status === 'active').length },
    { name: 'Inactive', value: patients.filter(p => p.status === 'inactive').length },
    { name: 'Discharged', value: patients.filter(p => p.status === 'discharged').length },
  ].filter(item => item.value > 0);

  const COLORS = ['#3b82f6', '#f59e0b', '#6b7280'];

  return (
    <ResponsiveContainer width="100%" height={300}>
      <PieChart>
        <Pie
          data={statusData}
          cx="50%"
          cy="50%"
          labelLine={false}
          label={({ name, value }) => `${name}: ${value}`}
          outerRadius={80}
          fill="#8884d8"
          dataKey="value"
        >
          {statusData.map((entry, index) => (
            <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
          ))}
        </Pie>
        <Tooltip />
      </PieChart>
    </ResponsiveContainer>
  );
};

const AppointmentsTrendChart = ({ appointments }) => {
  const last7Days = Array.from({ length: 7 }).map((_, i) => {
    const date = new Date();
    date.setDate(date.getDate() - (6 - i));
    return date.toISOString().split('T')[0];
  });

  const chartData = last7Days.map(date => ({
    date: new Date(date).toLocaleDateString('en-US', { month: 'short', day: 'numeric' }),
    appointments: appointments.filter(a => a.date.startsWith(date)).length,
  }));

  return (
    <ResponsiveContainer width="100%" height={300}>
      <LineChart data={chartData}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="date" />
        <YAxis />
        <Tooltip />
        <Legend />
        <Line
          type="monotone"
          dataKey="appointments"
          stroke="#3b82f6"
          strokeWidth={2}
          dot={{ fill: '#3b82f6', r: 4 }}
        />
      </LineChart>
    </ResponsiveContainer>
  );
};

const RecordsSeverityChart = ({ records }) => {
  const severityData = [
    { name: 'Normal', value: records.filter(r => r.severity === 'normal').length },
    { name: 'High', value: records.filter(r => r.severity === 'high').length },
    { name: 'Critical', value: records.filter(r => r.severity === 'critical').length },
  ].filter(item => item.value > 0);

  return (
    <ResponsiveContainer width="100%" height={300}>
      <BarChart data={severityData}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="name" />
        <YAxis />
        <Tooltip />
        <Bar dataKey="value" fill="#3b82f6" />
      </BarChart>
    </ResponsiveContainer>
  );
};

const AppointmentStatusChart = ({ appointments }) => {
  const statusData = [
    { name: 'Scheduled', value: appointments.filter(a => a.status === 'scheduled').length },
    { name: 'Completed', value: appointments.filter(a => a.status === 'completed').length },
    { name: 'Cancelled', value: appointments.filter(a => a.status === 'cancelled').length },
  ].filter(item => item.value > 0);

  const COLORS = ['#3b82f6', '#10b981', '#ef4444'];

  return (
    <ResponsiveContainer width="100%" height={300}>
      <PieChart>
        <Pie
          data={statusData}
          cx="50%"
          cy="50%"
          labelLine={false}
          label={({ name, value }) => `${name}: ${value}`}
          outerRadius={80}
          fill="#8884d8"
          dataKey="value"
        >
          {statusData.map((entry, index) => (
            <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
          ))}
        </Pie>
        <Tooltip />
      </PieChart>
    </ResponsiveContainer>
  );
};

const RecentPatients = ({ patients }) => {
  const recentPatients = patients.slice(0, 5);

  return (
    <div className="bg-white rounded-lg shadow p-6">
      <h3 className="text-lg font-bold text-gray-900 mb-4">Recent Patients</h3>
      {recentPatients.length === 0 ? (
        <p className="text-gray-600">No patients yet</p>
      ) : (
        <div className="space-y-3">
          {recentPatients.map(patient => (
            <div key={patient.id} className="flex justify-between items-center p-3 bg-gray-50 rounded">
              <div>
                <p className="font-semibold text-gray-900">{patient.name}</p>
                <p className="text-sm text-gray-600">{patient.email}</p>
              </div>
              <span className={`px-3 py-1 rounded-full text-xs font-semibold ${
                patient.status === 'active' ? 'bg-green-100 text-green-800' :
                'bg-gray-100 text-gray-800'
              }`}>
                {patient.status || 'Active'}
              </span>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

const RecentAppointments = ({ appointments }) => {
  const upcomingAppointments = appointments
    .filter(a => new Date(a.date) >= new Date())
    .sort((a, b) => new Date(a.date) - new Date(b.date))
    .slice(0, 5);

  return (
    <div className="bg-white rounded-lg shadow p-6">
      <h3 className="text-lg font-bold text-gray-900 mb-4">Upcoming Appointments</h3>
      {upcomingAppointments.length === 0 ? (
        <p className="text-gray-600">No upcoming appointments</p>
      ) : (
        <div className="space-y-3">
          {upcomingAppointments.map(appointment => (
            <div key={appointment.id} className="flex justify-between items-center p-3 bg-gray-50 rounded">
              <div>
                <p className="font-semibold text-gray-900">{appointment.patientName}</p>
                <p className="text-sm text-gray-600">
                  {new Date(appointment.date).toLocaleDateString()} at{' '}
                  {appointment.time}
                </p>
              </div>
              <span className="px-3 py-1 rounded-full text-xs font-semibold bg-blue-100 text-blue-800">
                {appointment.type || 'Checkup'}
              </span>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default Dashboard;
