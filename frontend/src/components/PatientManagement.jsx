import { useState, useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { usePatients } from '../context/PatientContext';
import { useAuth } from '../context/AuthContext';
import {
  LoadingState,
  EmptyState,
  ErrorMessage,
  SuccessMessage,
  Modal,
  Spinner,
  Button,
  Input,
  Select,
  Badge,
} from './LoadingAndErrors';

export const PatientManagement = () => {
  const { patients, loading, error, createPatient, updatePatient, deletePatient, fetchAllPatients, updateFilters, filters, getFilteredPatients } = usePatients();
  const { hasRole } = useAuth();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingPatient, setEditingPatient] = useState(null);
  const [successMsg, setSuccessMsg] = useState('');
  const [searchTerm, setSearchTerm] = useState('');
  const [statusFilter, setStatusFilter] = useState('');
  const [sortField, setSortField] = useState('name');
  const [sortOrder, setSortOrder] = useState('asc');

  useEffect(() => {
    fetchAllPatients();
  }, []);

  const filteredPatients = getFilteredPatients();

  // Sorting logic
  const getSortedPatients = () => {
    const sorted = [...filteredPatients].sort((a, b) => {
      let aVal = a[sortField];
      let bVal = b[sortField];

      // Handle null/undefined values
      if (aVal === null || aVal === undefined) aVal = '';
      if (bVal === null || bVal === undefined) bVal = '';

      // Convert to lowercase for string comparison
      if (typeof aVal === 'string') aVal = aVal.toLowerCase();
      if (typeof bVal === 'string') bVal = bVal.toLowerCase();

      if (aVal < bVal) return sortOrder === 'asc' ? -1 : 1;
      if (aVal > bVal) return sortOrder === 'asc' ? 1 : -1;
      return 0;
    });

    return sorted;
  };

  const toggleSort = (field) => {
    if (sortField === field) {
      setSortOrder(sortOrder === 'asc' ? 'desc' : 'asc');
    } else {
      setSortField(field);
      setSortOrder('asc');
    }
  };

  const getSortIcon = (field) => {
    if (sortField !== field) return '⇅';
    return sortOrder === 'asc' ? '↑' : '↓';
  };

  const handleSearchChange = (e) => {
    const value = e.target.value;
    setSearchTerm(value);
    updateFilters({ search: value });
  };

  const handleStatusChange = (e) => {
    const value = e.target.value;
    setStatusFilter(value);
    updateFilters({ status: value });
  };

  const openAddModal = () => {
    setEditingPatient(null);
    setIsModalOpen(true);
  };

  const openEditModal = (patient) => {
    setEditingPatient(patient);
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
    setEditingPatient(null);
  };

  const handleDelete = async (id) => {
    if (window.confirm('Are you sure you want to delete this patient?')) {
      try {
        await deletePatient(id);
        setSuccessMsg('Patient deleted successfully');
        setTimeout(() => setSuccessMsg(''), 3000);
      } catch (err) {
        console.error(err);
      }
    }
  };

  const sortedPatients = getSortedPatients();

  return (
    <div className="space-y-6">
      <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
        <div>
          <h2 className="text-3xl font-bold text-gray-900">Patient Management</h2>
          <p className="text-gray-600 text-sm mt-1">Total patients: {patients.length}</p>
        </div>
        {hasRole(['doctor', 'admin']) && (
          <button
            onClick={openAddModal}
            className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg font-semibold transition duration-200"
          >
            + Add Patient
          </button>
        )}
      </div>

      {error && <ErrorMessage message={error} />}
      {successMsg && <SuccessMessage message={successMsg} onDismiss={() => setSuccessMsg('')} />}

      {/* Search and Filter */}
      <div className="bg-white rounded-lg shadow p-4 space-y-4">
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <input
            type="text"
            placeholder="Search by name, email, or phone..."
            value={searchTerm}
            onChange={handleSearchChange}
            className="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
          />
          <select
            value={statusFilter}
            onChange={handleStatusChange}
            className="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
          >
            <option value="">All Status</option>
            <option value="active">Active</option>
            <option value="inactive">Inactive</option>
            <option value="discharged">Discharged</option>
          </select>
        </div>
        
        {/* Sort Options */}
        <div className="flex flex-wrap gap-2">
          <span className="text-sm font-medium text-gray-700 self-center">Sort by:</span>
          {['name', 'age', 'email', 'status'].map(field => (
            <button
              key={field}
              onClick={() => toggleSort(field)}
              className={`px-3 py-1 rounded text-sm font-medium transition ${
                sortField === field
                  ? 'bg-blue-600 text-white'
                  : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
              }`}
            >
              {field.charAt(0).toUpperCase() + field.slice(1)} {getSortIcon(field)}
            </button>
          ))}
        </div>
      </div>

      {/* Patient List */}
      {loading ? (
        <LoadingState message="Loading patients..." />
      ) : sortedPatients.length === 0 ? (
        <EmptyState
          title="No patients found"
          description="Start by adding a new patient to the system"
          action={
            hasRole(['doctor', 'admin']) && (
              <button
                onClick={openAddModal}
                className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-lg font-semibold"
              >
                Add Your First Patient
              </button>
            )
          }
        />
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {sortedPatients.map(patient => (
            <div key={patient.id} className="bg-white rounded-lg shadow hover:shadow-lg transition p-6">
              <div className="flex justify-between items-start mb-4">
                <h3 className="text-lg font-bold text-gray-900">{patient.name}</h3>
                <span className={`px-3 py-1 rounded-full text-sm font-semibold ${
                  patient.status === 'active' ? 'bg-green-100 text-green-800' :
                  patient.status === 'inactive' ? 'bg-yellow-100 text-yellow-800' :
                  'bg-gray-100 text-gray-800'
                }`}>
                  {patient.status || 'Active'}
                </span>
              </div>

              <div className="space-y-2 text-sm text-gray-600 mb-4">
                <p><span className="font-semibold">Age:</span> {patient.age || 'N/A'}</p>
                <p><span className="font-semibold">Email:</span> {patient.email}</p>
                <p><span className="font-semibold">Phone:</span> {patient.phone}</p>
                <p><span className="font-semibold">Condition:</span> {patient.condition || 'N/A'}</p>
              </div>

              {hasRole(['doctor', 'admin']) && (
                <div className="flex gap-2 pt-4 border-t">
                  <button
                    onClick={() => openEditModal(patient)}
                    className="flex-1 bg-blue-50 hover:bg-blue-100 text-blue-600 py-2 rounded font-semibold transition"
                  >
                    Edit
                  </button>
                  <button
                    onClick={() => handleDelete(patient.id)}
                    className="flex-1 bg-red-50 hover:bg-red-100 text-red-600 py-2 rounded font-semibold transition"
                  >
                    Delete
                  </button>
                </div>
              )}
            </div>
          ))}
        </div>
      )}

      {/* Add/Edit Patient Modal */}
      <PatientFormModal
        isOpen={isModalOpen}
        patient={editingPatient}
        onClose={closeModal}
        onSuccess={() => {
          setSuccessMsg(editingPatient ? 'Patient updated successfully' : 'Patient added successfully');
          setTimeout(() => setSuccessMsg(''), 3000);
          closeModal();
          fetchAllPatients();
        }}
      />
    </div>
  );
};

const PatientFormModal = ({ isOpen, patient, onClose, onSuccess }) => {
  const { createPatient, updatePatient, loading } = usePatients();
  const { register, handleSubmit, reset, formState: { errors } } = useForm({
    defaultValues: patient || {},
  });

  useEffect(() => {
    if (patient) {
      reset(patient);
    } else {
      reset({});
    }
  }, [patient, isOpen, reset]);

  const onSubmit = async (data) => {
    try {
      if (patient) {
        await updatePatient(patient.id, data);
      } else {
        await createPatient(data);
      }
      onSuccess();
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div className={`fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 ${isOpen ? '' : 'hidden'}`}>
      <div className="bg-white rounded-lg shadow-xl max-w-md w-full mx-4">
        <div className="flex justify-between items-center p-6 border-b">
          <h2 className="text-xl font-bold text-gray-900">{patient ? 'Edit Patient' : 'Add New Patient'}</h2>
          <button
            onClick={onClose}
            className="text-gray-400 hover:text-gray-600"
          >
            <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        <div className="p-6">
          <form className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Name *</label>
              <input
                {...register('name', { required: 'Name is required' })}
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
                placeholder="Patient name"
              />
              {errors.name && <p className="text-red-600 text-sm mt-1">{errors.name.message}</p>}
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Email *</label>
              <input
                type="email"
                {...register('email', {
                  required: 'Email is required',
                  pattern: { value: /^[^\s@]+@[^\s@]+\.[^\s@]+$/, message: 'Invalid email' }
                })}
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
                placeholder="patient@example.com"
              />
              {errors.email && <p className="text-red-600 text-sm mt-1">{errors.email.message}</p>}
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Phone</label>
              <input
                {...register('phone')}
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
                placeholder="123-456-7890"
              />
            </div>

            <div className="grid grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Age</label>
                <input
                  type="number"
                  {...register('age')}
                  className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
                  placeholder="30"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Status</label>
                <select
                  {...register('status')}
                  className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
                >
                  <option value="active">Active</option>
                  <option value="inactive">Inactive</option>
                  <option value="discharged">Discharged</option>
                </select>
              </div>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Medical Condition</label>
              <textarea
                {...register('condition')}
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
                placeholder="Current medical condition or diagnosis"
                rows={3}
              />
            </div>
          </form>
        </div>
        <div className="flex gap-3 p-6 border-t bg-gray-50 rounded-b-lg">
          <button
            onClick={onClose}
            className="flex-1 px-4 py-2 border border-gray-300 rounded-lg text-gray-700 font-semibold hover:bg-gray-50"
          >
            Cancel
          </button>
          <button
            onClick={handleSubmit(onSubmit)}
            disabled={loading}
            className="flex-1 bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white px-4 py-2 rounded-lg font-semibold flex items-center justify-center gap-2"
          >
            {loading && <Spinner size="sm" />}
            {patient ? 'Update' : 'Add'}
          </button>
        </div>
      </div>
    </div>
  );
};

export default PatientManagement;
