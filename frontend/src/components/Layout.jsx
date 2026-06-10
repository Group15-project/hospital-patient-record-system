import { useState } from 'react';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

export const Layout = ({ children }) => {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);
  const { user, logout, isAuthenticated } = useAuth();
  const location = useLocation();
  const navigate = useNavigate();

  const handleLogout = async () => {
    await logout();
    navigate('/login');
  };

  const isActive = (path) => location.pathname === path;

  const navItems = [
    { path: '/dashboard', label: 'Dashboard', icon: '📊', roles: ['patient', 'doctor', 'admin'] },
    { path: '/patients', label: 'Patients', icon: '👥', roles: ['doctor', 'admin'] },
    { path: '/appointments', label: 'Appointments', icon: '📅', roles: ['patient', 'doctor', 'admin'] },
    { path: '/medical-records', label: 'Medical Records', icon: '📋', roles: ['patient', 'doctor', 'admin'] },
  ];

  const filteredNavItems = navItems.filter(item =>
    user && item.roles.includes(user.role)
  );

  return (
    <div className="min-h-screen bg-gray-100">
      {/* Mobile Menu Button */}
      <div className="lg:hidden flex items-center justify-between bg-white shadow p-4 sticky top-0 z-40">
        <h1 className="text-xl font-bold text-gray-900">Hospital System</h1>
        <button
          onClick={() => setIsSidebarOpen(!isSidebarOpen)}
          className="text-gray-600 hover:text-gray-900"
        >
          <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16M4 18h16" />
          </svg>
        </button>
      </div>

      <div className="flex">
        {/* Sidebar */}
        <aside
          className={`bg-gray-900 text-white w-64 min-h-screen transition-all duration-300 ${
            isSidebarOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'
          } fixed lg:static z-30`}
        >
          <div className="p-6">
            <h1 className="text-2xl font-bold">Hospital</h1>
            <p className="text-gray-400 text-sm">Patient Record System</p>
          </div>

          <nav className="mt-8 space-y-2 px-4">
            {filteredNavItems.map(item => (
              <Link
                key={item.path}
                to={item.path}
                onClick={() => setIsSidebarOpen(false)}
                className={`block px-4 py-3 rounded-lg transition duration-200 font-medium ${
                  isActive(item.path)
                    ? 'bg-blue-600 text-white'
                    : 'text-gray-300 hover:bg-gray-800'
                }`}
              >
                <span className="mr-3">{item.icon}</span>
                {item.label}
              </Link>
            ))}
          </nav>

          {/* User Profile */}
          {user && (
            <div className="absolute bottom-0 left-0 right-0 p-4 border-t border-gray-700">
              <div className="mb-4">
                <p className="text-sm text-gray-400">Logged in as:</p>
                <p className="font-semibold text-white">{user.name}</p>
                <p className="text-xs text-gray-400 capitalize">{user.role}</p>
              </div>
              <button
                onClick={handleLogout}
                className="w-full bg-red-600 hover:bg-red-700 text-white py-2 rounded-lg font-semibold transition duration-200"
              >
                Logout
              </button>
            </div>
          )}
        </aside>

        {/* Main Content */}
        <main className="flex-1 p-4 lg:p-8 w-full lg:w-auto overflow-auto">
          {children}
        </main>
      </div>

      {/* Overlay for mobile */}
      {isSidebarOpen && (
        <div
          className="fixed inset-0 bg-black bg-opacity-50 lg:hidden z-20"
          onClick={() => setIsSidebarOpen(false)}
        />
      )}
    </div>
  );
};

export const ProtectedRoute = ({ children }) => {
  const { isAuthenticated, loading } = useAuth();
  const navigate = useNavigate();

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="text-center">
          <div className="w-12 h-12 border-4 border-gray-200 border-t-blue-600 rounded-full animate-spin mx-auto mb-4" />
          <p className="text-gray-600">Loading...</p>
        </div>
      </div>
    );
  }

  if (!isAuthenticated) {
    navigate('/login');
    return null;
  }

  return children;
};

export default Layout;
