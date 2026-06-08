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
} from './LoadingAndErrors';
import api from '../services/api';

export const MedicalRecords = () => {
  const [records, setRecords] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [successMsg, setSuccessMsg] = useState('');
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [selectedRecord, setSelectedRecord] = useState(null);
  const [selectedPatientId, setSelectedPatientId] = useState(null);
  const { patients, fetchAllPatients } = usePatients();
  const { hasRole } = useAuth();

  useEffect(() => {
    fetchAllPatients();
    fetchRecords();
  }, []);

  const fetchRecords = async () => {
    try {
      setLoading(true);
      const response = await api.get('/medical-records');
      setRecords(response.data.records || []);
      setError(null);
    } catch (err) {
      setError('Failed to fetch medical records');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (id) => {
    if (window.confirm('Are you sure you want to delete this record?')) {
      try {
        await api.delete(`/medical-records/${id}`);
        setRecords(prev => prev.filter(r => r.id !== id));
        setSuccessMsg('Record deleted successfully');
        setTimeout(() => setSuccessMsg(''), 3000);
      } catch (err) {
        setError('Failed to delete record');
      }
    }
  };

  const openAddModal = (patientId = null) => {
    setSelectedRecord(null);
    setSelectedPatientId(patientId);
    setIsModalOpen(true);
  };

  const openEditModal = (record) => {
    setSelectedRecord(record);
    setSelectedPatientId(record.patientId);
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
    setSelectedRecord(null);
    setSelectedPatientId(null);
  };

  const recordsByPatient = records.reduce((acc, record) => {
    if (!acc[record.patientId]) {
      acc[record.patientId] = [];
    }
    acc[record.patientId].push(record);
    return acc;
  }, {});

  return (
    <div className="space-y-6">
      <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
        <h2 className="text-3xl font-bold text-gray-900">Medical Records</h2>
        {hasRole(['doctor', 'admin']) && (
          <button
            onClick={() => openAddModal()}
            className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg font-semibold transition duration-200"
          >
            + Add Medical Record
          </button>
        )}
      </div>

      {error && <ErrorMessage message={error} />}
      {successMsg && <SuccessMessage message={successMsg} onDismiss={() => setSuccessMsg('')} />}

      {loading ? (
        <LoadingState message="Loading medical records..." />
      ) : records.length === 0 ? (
        <EmptyState
          title="No medical records found"
          description="Add medical records to keep patient history updated"
          action={
            hasRole(['doctor', 'admin']) && (
              <button
                onClick={() => openAddModal()}
                className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-lg font-semibold"
              >
                Add Medical Record
              </button>
            )
          }
        />
      ) : (
        <div className="space-y-6">
          {patients.map(patient => {
            const patientRecords = recordsByPatient[patient.id] || [];
            if (patientRecords.length === 0) return null;

            return (
              <div key={patient.id} className="bg-white rounded-lg shadow">
                <div className="px-6 py-4 border-b bg-gray-50 flex justify-between items-center">
                  <h3 className="text-lg font-bold text-gray-900">{patient.name}</h3>
                  {hasRole(['doctor', 'admin']) && (
                    <button
                      onClick={() => openAddModal(patient.id)}
                      className="bg-blue-600 hover:bg-blue-700 text-white px-3 py-1 rounded text-sm font-semibold transition"
                    >
                      + Add Record
                    </button>
                  )}
                </div>

                <div className="divide-y">
                  {patientRecords.map(record => (
                    <MedicalRecordItem
                      key={record.id}
                      record={record}
                      onEdit={openEditModal}
                      onDelete={handleDelete}
                      canEdit={hasRole(['doctor', 'admin'])}
                    />
                  ))}
                </div>
              </div>
            );
          })}
        </div>
      )}

      <MedicalRecordFormModal
        isOpen={isModalOpen}
        record={selectedRecord}
        patientId={selectedPatientId}
        patients={patients}
        onClose={closeModal}
        onSuccess={() => {
          setSuccessMsg(selectedRecord ? 'Record updated' : 'Record added');
          setTimeout(() => setSuccessMsg(''), 3000);
          closeModal();
          fetchRecords();
        }}
      />
    </div>
  );
};

const MedicalRecordItem = ({ record, onEdit, onDelete, canEdit }) => {
  return (
    <div className="p-6 hover:bg-gray-50 transition">
      <div className="flex justify-between items-start mb-3">
        <div>
          <h4 className="text-lg font-semibold text-gray-900">{record.title}</h4>
          <p className="text-sm text-gray-600">{new Date(record.date).toLocaleDateString()}</p>
        </div>
        <span className={`px-3 py-1 rounded-full text-sm font-semibold ${
          record.severity === 'critical' ? 'bg-red-100 text-red-800' :
          record.severity === 'high' ? 'bg-orange-100 text-orange-800' :
          'bg-green-100 text-green-800'
        }`}>
          {record.severity || 'Normal'}
        </span>
      </div>

      <p className="text-gray-700 mb-3">{record.description}</p>

      <div className="grid grid-cols-2 gap-4 text-sm mb-4">
        <div>
          <span className="font-semibold text-gray-600">Type: </span>
          <span className="text-gray-900">{record.type || 'Consultation'}</span>
        </div>
        <div>
          <span className="font-semibold text-gray-600">Doctor: </span>
          <span className="text-gray-900">{record.doctorName || 'N/A'}</span>
        </div>
      </div>

      {record.prescription && (
        <div className="bg-blue-50 p-3 rounded mb-4">
          <p className="font-semibold text-blue-900 mb-1">Prescription:</p>
          <p className="text-blue-800 text-sm">{record.prescription}</p>
        </div>
      )}

      {canEdit && (
        <div className="flex gap-2 pt-4 border-t">
          <button
            onClick={() => onEdit(record)}
            className="flex-1 bg-blue-50 hover:bg-blue-100 text-blue-600 py-2 rounded font-semibold transition"
          >
            Edit
          </button>
          <button
            onClick={() => onDelete(record.id)}
            className="flex-1 bg-red-50 hover:bg-red-100 text-red-600 py-2 rounded font-semibold transition"
          >
            Delete
          </button>
        </div>
      )}
    </div>
  );
};

const MedicalRecordFormModal = ({ isOpen, record, patientId, patients, onClose, onSuccess }) => {
  const { register, handleSubmit, reset, formState: { errors } } = useForm({
    defaultValues: record || {
      severity: 'normal',
      type: 'Consultation'
    },
  });
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (record) {
      reset(record);
    } else {
      reset({ severity: 'normal', type: 'Consultation', date: new Date().toISOString().split('T')[0] });
    }
  }, [record, isOpen, reset]);

  const onSubmit = async (data) => {
    try {
      setLoading(true);
      if (record) {
        await api.put(`/medical-records/${record.id}`, data);
      } else {
        await api.post('/medical-records', data);
      }
      onSuccess();
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <Modal
      isOpen={isOpen}
      title={record ? 'Edit Medical Record' : 'Add Medical Record'}
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
            {record ? 'Update' : 'Add'}
          </button>
        </>
      }
    >
      <form className="space-y-4">
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Patient *</label>
          <select
            {...register('patientId', { required: 'Patient is required' })}
            defaultValue={patientId || ''}
            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
          >
            <option value="">Select a patient</option>
            {patients.map(patient => (
              <option key={patient.id} value={patient.id}>
                {patient.name}
              </option>
            ))}
          </select>
          {errors.patientId && <p className="text-red-600 text-sm mt-1">{errors.patientId.message}</p>}
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Title *</label>
          <input
            {...register('title', { required: 'Title is required' })}
            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
            placeholder="Record title (e.g., Annual Checkup)"
          />
          {errors.title && <p className="text-red-600 text-sm mt-1">{errors.title.message}</p>}
        </div>

        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Date *</label>
            <input
              type="date"
              {...register('date', { required: 'Date is required' })}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
            />
            {errors.date && <p className="text-red-600 text-sm mt-1">{errors.date.message}</p>}
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Type</label>
            <select
              {...register('type')}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
            >
              <option value="Consultation">Consultation</option>
              <option value="Lab Test">Lab Test</option>
              <option value="Imaging">Imaging</option>
              <option value="Surgery">Surgery</option>
              <option value="Follow-up">Follow-up</option>
            </select>
          </div>
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Description *</label>
          <textarea
            {...register('description', { required: 'Description is required' })}
            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
            placeholder="Detailed description of the medical record"
            rows={3}
          />
          {errors.description && <p className="text-red-600 text-sm mt-1">{errors.description.message}</p>}
        </div>

        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Severity</label>
            <select
              {...register('severity')}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
            >
              <option value="normal">Normal</option>
              <option value="high">High</option>
              <option value="critical">Critical</option>
            </select>
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Doctor Name</label>
            <input
              {...register('doctorName')}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
              placeholder="Doctor name"
            />
          </div>
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Prescription</label>
          <textarea
            {...register('prescription')}
            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
            placeholder="Prescription details (if applicable)"
            rows={2}
          />
        </div>
      </form>
    </Modal>
  );
};

export default MedicalRecords;
