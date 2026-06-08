import { createContext, useContext, useState, useCallback } from 'react';
import { patientAPI } from '../services/api';

const PatientContext = createContext();

export const usePatients = () => {
  const context = useContext(PatientContext);
  if (!context) {
    throw new Error('usePatients must be used within a PatientProvider');
  }
  return context;
};

export const PatientProvider = ({ children }) => {
  const [patients, setPatients] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [filters, setFilters] = useState({ search: '', status: '' });
  const [selectedPatient, setSelectedPatient] = useState(null);

  const fetchAllPatients = useCallback(async () => {
    try {
      setLoading(true);
      setError(null);
      const response = await patientAPI.getAll();
      setPatients(response.data.patients || []);
    } catch (err) {
      const errorMsg = err.response?.data?.message || 'Failed to fetch patients';
      setError(errorMsg);
      console.error(err);
    } finally {
      setLoading(false);
    }
  }, []);

  const fetchPatientById = useCallback(async (id) => {
    try {
      setLoading(true);
      setError(null);
      const response = await patientAPI.getById(id);
      setSelectedPatient(response.data.patient || response.data);
      return response.data.patient || response.data;
    } catch (err) {
      const errorMsg = err.response?.data?.message || 'Failed to fetch patient';
      setError(errorMsg);
      console.error(err);
    } finally {
      setLoading(false);
    }
  }, []);

  const createPatient = useCallback(async (data) => {
    try {
      setLoading(true);
      setError(null);
      const response = await patientAPI.create(data);
      const newPatient = response.data.patient || response.data;
      setPatients(prev => [...prev, newPatient]);
      return newPatient;
    } catch (err) {
      const errorMsg = err.response?.data?.message || 'Failed to create patient';
      setError(errorMsg);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  const updatePatient = useCallback(async (id, data) => {
    try {
      setLoading(true);
      setError(null);
      const response = await patientAPI.update(id, data);
      const updatedPatient = response.data.patient || response.data;
      setPatients(prev => prev.map(p => p.id === id ? updatedPatient : p));
      if (selectedPatient?.id === id) {
        setSelectedPatient(updatedPatient);
      }
      return updatedPatient;
    } catch (err) {
      const errorMsg = err.response?.data?.message || 'Failed to update patient';
      setError(errorMsg);
      throw err;
    } finally {
      setLoading(false);
    }
  }, [selectedPatient?.id]);

  const deletePatient = useCallback(async (id) => {
    try {
      setLoading(true);
      setError(null);
      await patientAPI.delete(id);
      setPatients(prev => prev.filter(p => p.id !== id));
      if (selectedPatient?.id === id) {
        setSelectedPatient(null);
      }
    } catch (err) {
      const errorMsg = err.response?.data?.message || 'Failed to delete patient';
      setError(errorMsg);
      throw err;
    } finally {
      setLoading(false);
    }
  }, [selectedPatient?.id]);

  // Filter and search patients
  const getFilteredPatients = useCallback(() => {
    return patients.filter(patient => {
      const matchesSearch = !filters.search || 
        patient.name?.toLowerCase().includes(filters.search.toLowerCase()) ||
        patient.email?.toLowerCase().includes(filters.search.toLowerCase()) ||
        patient.phone?.includes(filters.search);
      
      const matchesStatus = !filters.status || patient.status === filters.status;
      
      return matchesSearch && matchesStatus;
    });
  }, [patients, filters]);

  const updateFilters = useCallback((newFilters) => {
    setFilters(prev => ({ ...prev, ...newFilters }));
  }, []);

  const clearFilters = useCallback(() => {
    setFilters({ search: '', status: '' });
  }, []);

  const value = {
    patients,
    loading,
    error,
    filters,
    selectedPatient,
    fetchAllPatients,
    fetchPatientById,
    createPatient,
    updatePatient,
    deletePatient,
    getFilteredPatients,
    updateFilters,
    clearFilters,
    setSelectedPatient,
    setError,
  };

  return (
    <PatientContext.Provider value={value}>
      {children}
    </PatientContext.Provider>
  );
};

export default PatientContext;
