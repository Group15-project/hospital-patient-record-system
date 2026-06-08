import { useState, useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { format, addDays, startOfToday } from 'date-fns';
import { appointmentAPI } from '../services/api';
import { useAuth } from '../context/AuthContext';
import {
  LoadingState,
  EmptyState,
  ErrorMessage,
  SuccessMessage,
  Modal,
  Spinner,
} from './LoadingAndErrors';

export const AppointmentScheduling = () => {
  const [appointments, setAppointments] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [successMsg, setSuccessMsg] = useState('');
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingAppointment, setEditingAppointment] = useState(null);
  const [statusFilter, setStatusFilter] = useState('');
  const { hasRole } = useAuth();

  useEffect(() => {
    fetchAppointments();
  }, []);

  const fetchAppointments = async () => {
    try {
      setLoading(true);
      const response = await appointmentAPI.getAll();
      setAppointments(response.data.appointments || []);
      setError(null);
    } catch (err) {
      setError('Failed to fetch appointments');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (id) => {
    if (window.confirm('Are you sure you want to delete this appointment?')) {
      try {
        await appointmentAPI.delete(id);
        setAppointments(prev => prev.filter(a => a.id !== id));
        setSuccessMsg('Appointment deleted successfully');
        setTimeout(() => setSuccessMsg(''), 3000);
      } catch (err) {
        setError('Failed to delete appointment');
      }
    }
  };

  const openAddModal = () => {
    setEditingAppointment(null);
    setIsModalOpen(true);
  };

  const openEditModal = (appointment) => {
    setEditingAppointment(appointment);
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
    setEditingAppointment(null);
  };

  const filteredAppointments = statusFilter
    ? appointments.filter(a => a.status === statusFilter)
    : appointments;

  const upcomingAppointments = filteredAppointments.filter(a => {
    const appDate = new Date(a.date);
    return appDate >= startOfToday();
  });

  const pastAppointments = filteredAppointments.filter(a => {
    const appDate = new Date(a.date);
    return appDate < startOfToday();
  });

  return (
    <div className="space-y-6">
      <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
        <h2 className="text-3xl font-bold text-gray-900">Appointments</h2>
        {hasRole(['doctor', 'admin']) && (
          <button
            onClick={openAddModal}
            className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg font-semibold transition duration-200"
          >
            + Schedule Appointment
          </button>
        )}
      </div>

      {error && <ErrorMessage message={error} />}
      {successMsg && <SuccessMessage message={successMsg} onDismiss={() => setSuccessMsg('')} />}

      {/* Filter */}
      <div className="bg-white rounded-lg shadow p-4">
        <select
          value={statusFilter}
          onChange={(e) => setStatusFilter(e.target.value)}
          className="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
        >
          <option value="">All Appointments</option>
          <option value="scheduled">Scheduled</option>
          <option value="completed">Completed</option>
          <option value="cancelled">Cancelled</option>
        </select>
      </div>

      {loading ? (
        <LoadingState message="Loading appointments..." />
      ) : appointments.length === 0 ? (
        <EmptyState
          title="No appointments yet"
          description="Schedule your first appointment"
          action={
            hasRole(['doctor', 'admin']) && (
              <button
                onClick={openAddModal}
                className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-lg font-semibold"
              >
                Schedule Appointment
              </button>
            )
          }
        />
      ) : (
        <div className="space-y-6">
          {/* Upcoming Appointments */}
          {upcomingAppointments.length > 0 && (
            <div>
              <h3 className="text-xl font-bold text-gray-900 mb-4">Upcoming Appointments</h3>
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                {upcomingAppointments.map(appointment => (
                  <AppointmentCard
                    key={appointment.id}
                    appointment={appointment}
                    onEdit={openEditModal}
                    onDelete={handleDelete}
                    canEdit={hasRole(['doctor', 'admin'])}
                  />
                ))}
              </div>
            </div>
          )}

          {/* Past Appointments */}
          {pastAppointments.length > 0 && (
            <div>
              <h3 className="text-xl font-bold text-gray-900 mb-4">Past Appointments</h3>
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                {pastAppointments.map(appointment => (
                  <AppointmentCard
                    key={appointment.id}
                    appointment={appointment}
                    onEdit={openEditModal}
                    onDelete={handleDelete}
                    canEdit={hasRole(['doctor', 'admin'])}
                    isPast={true}
                  />
                ))}
              </div>
            </div>
          )}
        </div>
      )}

      <AppointmentFormModal
        isOpen={isModalOpen}
        appointment={editingAppointment}
        onClose={closeModal}
        onSuccess={() => {
          setSuccessMsg(editingAppointment ? 'Appointment updated' : 'Appointment scheduled');
          setTimeout(() => setSuccessMsg(''), 3000);
          closeModal();
          fetchAppointments();
        }}
      />
    </div>
  );
};

const AppointmentCard = ({ appointment, onEdit, onDelete, canEdit, isPast }) => {
  return (
    <div className={`bg-white rounded-lg shadow hover:shadow-lg transition p-6 ${isPast ? 'opacity-75' : ''}`}>
      <div className="flex justify-between items-start mb-4">
        <div>
          <h3 className="text-lg font-bold text-gray-900">{appointment.patientName}</h3>
          <p className="text-sm text-gray-600">{appointment.doctorName}</p>
        </div>
        <span className={`px-3 py-1 rounded-full text-sm font-semibold ${
          appointment.status === 'scheduled' ? 'bg-blue-100 text-blue-800' :
          appointment.status === 'completed' ? 'bg-green-100 text-green-800' :
          'bg-red-100 text-red-800'
        }`}>
          {appointment.status || 'Scheduled'}
        </span>
      </div>

      <div className="space-y-2 text-sm text-gray-600 mb-4">
        <p>
          <span className="font-semibold">Date & Time:</span>{' '}
          {format(new Date(appointment.date), 'MMM dd, yyyy')} at{' '}
          {format(new Date(`2000-01-01T${appointment.time}`), 'hh:mm a')}
        </p>
        <p>
          <span className="font-semibold">Type:</span> {appointment.type || 'General Checkup'}
        </p>
        {appointment.notes && (
          <p>
            <span className="font-semibold">Notes:</span> {appointment.notes}
          </p>
        )}
      </div>

      {canEdit && (
        <div className="flex gap-2 pt-4 border-t">
          <button
            onClick={() => onEdit(appointment)}
            className="flex-1 bg-blue-50 hover:bg-blue-100 text-blue-600 py-2 rounded font-semibold transition"
          >
            Edit
          </button>
          <button
            onClick={() => onDelete(appointment.id)}
            className="flex-1 bg-red-50 hover:bg-red-100 text-red-600 py-2 rounded font-semibold transition"
          >
            Delete
          </button>
        </div>
      )}
    </div>
  );
};

const AppointmentFormModal = ({ isOpen, appointment, onClose, onSuccess }) => {
  const { register, handleSubmit, reset, formState: { errors } } = useForm({
    defaultValues: appointment || {
      type: 'General Checkup',
      status: 'scheduled'
    },
  });
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (appointment) {
      reset(appointment);
    } else {
      reset({ type: 'General Checkup', status: 'scheduled' });
    }
  }, [appointment, isOpen, reset]);

  const onSubmit = async (data) => {
    try {
      setLoading(true);
      if (appointment) {
        await appointmentAPI.update(appointment.id, data);
      } else {
        await appointmentAPI.create(data);
      }
      onSuccess();
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const tomorrow = format(addDays(startOfToday(), 1), 'yyyy-MM-dd');

  return (
    <Modal
      isOpen={isOpen}
      title={appointment ? 'Edit Appointment' : 'Schedule Appointment'}
      onClose={onClose}
      actions={
        <>
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
            {appointment ? 'Update' : 'Schedule'}
          </button>
        </>
      }
    >
      <form className="space-y-4">
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Patient Name *</label>
          <input
            {...register('patientName', { required: 'Patient name is required' })}
            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
            placeholder="Patient name"
          />
          {errors.patientName && <p className="text-red-600 text-sm mt-1">{errors.patientName.message}</p>}
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Doctor Name *</label>
          <input
            {...register('doctorName', { required: 'Doctor name is required' })}
            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
            placeholder="Doctor name"
          />
          {errors.doctorName && <p className="text-red-600 text-sm mt-1">{errors.doctorName.message}</p>}
        </div>

        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Date *</label>
            <input
              type="date"
              {...register('date', { required: 'Date is required' })}
              min={tomorrow}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
            />
            {errors.date && <p className="text-red-600 text-sm mt-1">{errors.date.message}</p>}
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Time *</label>
            <input
              type="time"
              {...register('time', { required: 'Time is required' })}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
            />
            {errors.time && <p className="text-red-600 text-sm mt-1">{errors.time.message}</p>}
          </div>
        </div>

        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Type</label>
            <select
              {...register('type')}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
            >
              <option value="General Checkup">General Checkup</option>
              <option value="Follow-up">Follow-up</option>
              <option value="Surgery">Surgery</option>
              <option value="Consultation">Consultation</option>
              <option value="Vaccination">Vaccination</option>
            </select>
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Status</label>
            <select
              {...register('status')}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
            >
              <option value="scheduled">Scheduled</option>
              <option value="completed">Completed</option>
              <option value="cancelled">Cancelled</option>
            </select>
          </div>
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Notes</label>
          <textarea
            {...register('notes')}
            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
            placeholder="Additional notes about the appointment"
            rows={3}
          />
        </div>
      </form>
    </Modal>
  );
};

export default AppointmentScheduling;
