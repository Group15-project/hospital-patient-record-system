import { useState, useEffect } from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import api from './services/api'
import './App.css'

// Components
function Header() {
  return (
    <header className="bg-primary-700 text-white p-4">
      <div className="container mx-auto">
        <h1 className="text-2xl font-bold">Hospital Patient Record System</h1>
      </div>
    </header>
  )
}

function Dashboard() {
  const [patients, setPatients] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => {
    const fetchPatients = async () => {
      try {
        setLoading(true)
        const response = await api.get('/patients')
        setPatients(response.data.patients || [])
        setError(null)
      } catch (err) {
        setError('Failed to fetch patients')
        console.error(err)
      } finally {
        setLoading(false)
      }
    }

    fetchPatients()
  }, [])

  return (
    <div className="container mx-auto p-4">
      <h2 className="text-2xl font-bold mb-4">Patient Dashboard</h2>
      
      {loading && <p className="text-gray-600">Loading...</p>}
      {error && <p className="text-red-600">{error}</p>}
      
      {patients.length > 0 && (
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          {patients.map(patient => (
            <div key={patient.id} className="bg-white p-4 rounded-lg shadow">
              <h3 className="text-lg font-bold">{patient.name}</h3>
              <p className="text-gray-600">Age: {patient.age}</p>
              <p className="text-gray-600">Phone: {patient.phone}</p>
              <p className="text-gray-600">Condition: {patient.condition}</p>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}

function App() {
  return (
    <Router>
      <Header />
      <Routes>
        <Route path="/" element={<Dashboard />} />
      </Routes>
    </Router>
  )
}

export default App