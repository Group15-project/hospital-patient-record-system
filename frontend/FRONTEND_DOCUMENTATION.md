# Frontend Documentation

## Overview

This React-based frontend provides a complete Hospital Patient Record System with the following key features:

### ✨ Key Features Implemented

#### 1. **Authentication & Role-Based Access Control**
- User login and registration system
- Three user roles: Patient, Doctor, Admin
- Role-based UI rendering (show/hide elements based on user role)
- Persistent authentication with localStorage
- Protected routes with automatic redirection

#### 2. **Patient Management**
- **View**: List all patients with status indicators
- **Search**: Search patients by name, email, or phone
- **Filter**: Filter patients by status (Active, Inactive, Discharged)
- **Add**: Create new patient records with validation
- **Edit**: Update patient information
- **Delete**: Remove patient records
- **Responsive Grid Layout**: Works on mobile, tablet, and desktop

#### 3. **Appointment Scheduling**
- **Schedule**: Create new appointments with date/time validation
- **View**: Display appointments organized by upcoming and past
- **Edit**: Modify appointment details
- **Delete**: Cancel appointments
- **Status Tracking**: Track appointment status (Scheduled, Completed, Cancelled)
- **Type Selection**: General Checkup, Follow-up, Surgery, Consultation, Vaccination
- **Filtering**: Filter by appointment status

#### 4. **Medical Records Management**
- **Create Records**: Add medical records with details
- **Track Severity**: Mark records as Normal, High, or Critical
- **Record Types**: Consultation, Lab Test, Imaging, Surgery, Follow-up
- **Prescriptions**: Store prescription information
- **Doctor Association**: Link records to specific doctors
- **Organized by Patient**: Group records by patient for easy navigation

#### 5. **Interactive Dashboard**
- **Statistics Cards**: Display key metrics (Total Patients, Upcoming Appointments, Medical Records, Active Cases)
- **Charts & Visualizations**:
  - Patient Status Distribution (Pie Chart)
  - Appointments Trend (Line Chart)
  - Medical Records by Severity (Bar Chart)
  - Appointment Status Distribution (Pie Chart)
- **Recent Activity Panels**: Latest patients and upcoming appointments
- **Real-time Data**: Updates data from API

#### 6. **Forms & Validation**
- **React Hook Form Integration**: Powerful form handling
- **Real-time Validation**: Email, required fields, date ranges
- **Error Messages**: User-friendly validation feedback
- **Modal Forms**: Clean modal dialogs for add/edit operations
- **Controlled Components**: Proper React form patterns

#### 7. **Loading & Error States**
- **Spinner Component**: Loading indicators
- **Skeleton Loaders**: Content placeholders while loading
- **Error Messages**: Styled error notifications with dismiss option
- **Success Messages**: Confirmation messages with auto-dismiss
- **Empty States**: Helpful messages when no data available
- **Loading States**: Disabled buttons and feedback during async operations

#### 8. **Responsive Design**
- **Mobile-First Approach**: Works seamlessly on all devices
- **Tailwind CSS**: Utility-first CSS framework
- **Breakpoints**: sm (640px), md (768px), lg (1024px)
- **Mobile Menu**: Collapsible sidebar for mobile devices
- **Grid Layouts**: Responsive grid that adapts to screen size
- **Touch-Friendly**: Proper button sizes for touch input

## Component Structure

```
src/
├── components/
│   ├── Auth.jsx                 # Login and Register components
│   ├── Dashboard.jsx            # Dashboard with charts
│   ├── PatientManagement.jsx    # Patient CRUD operations
│   ├── AppointmentScheduling.jsx # Appointment management
│   ├── MedicalRecords.jsx       # Medical records management
│   ├── Layout.jsx               # Main layout with navigation
│   ├── LoadingAndErrors.jsx     # Reusable UI components
│   └── index.js                 # Components index
├── context/
│   ├── AuthContext.jsx          # Authentication state management
│   └── PatientContext.jsx       # Patient data management
├── services/
│   └── api.js                   # API service with axios
├── App.jsx                      # Main app component with routing
├── main.jsx                     # React entry point
├── index.css                    # Global styles
└── App.css                      # App-specific styles
```

## Context Providers

### AuthContext
Manages authentication state and role-based access control.

**Methods:**
- `login(email, password)`: Authenticate user
- `logout()`: Clear authentication
- `register(data)`: Register new user
- `hasRole(roles)`: Check if user has specific role
- `canView(requiredRoles)`: Verify access permissions

**State:**
- `user`: Current logged-in user object
- `isAuthenticated`: Boolean authentication status
- `loading`: Loading state during auth operations
- `error`: Error messages

### PatientContext
Manages patient data and operations.

**Methods:**
- `fetchAllPatients()`: Load all patients
- `fetchPatientById(id)`: Load specific patient
- `createPatient(data)`: Add new patient
- `updatePatient(id, data)`: Update patient info
- `deletePatient(id)`: Remove patient
- `getFilteredPatients()`: Get filtered/searched patients
- `updateFilters(newFilters)`: Update search/filter criteria
- `clearFilters()`: Reset all filters

**State:**
- `patients`: Array of all patients
- `loading`: Loading state
- `error`: Error messages
- `filters`: Current search/filter settings
- `selectedPatient`: Currently selected patient

## Reusable UI Components

### LoadingAndErrors.jsx
Exports multiple utility components:

1. **Spinner**: Loading animation
   ```jsx
   <Spinner size="md" />  // sm, md, lg
   ```

2. **SkeletonLoader**: Content placeholder
   ```jsx
   <SkeletonLoader count={3} />
   ```

3. **ErrorMessage**: Error notification
   ```jsx
   <ErrorMessage message="Error text" onDismiss={() => {}} />
   ```

4. **SuccessMessage**: Success notification
   ```jsx
   <SuccessMessage message="Success!" onDismiss={() => {}} />
   ```

5. **LoadingState**: Full page loading screen
   ```jsx
   <LoadingState message="Loading..." />
   ```

6. **EmptyState**: No data message with action
   ```jsx
   <EmptyState title="No data" description="Add data" action={<button>Add</button>} />
   ```

7. **Modal**: Dialog component
   ```jsx
   <Modal isOpen={true} title="Dialog" onClose={() => {}} actions={<button>Save</button>}>
     Content here
   </Modal>
   ```

## API Integration

All API calls go through `services/api.js` which provides:

- **Patient Endpoints**: `/patients`, `/patients/:id`
- **Appointment Endpoints**: `/appointments`, `/appointments/:id`
- **Medical Records Endpoints**: `/medical-records`, `/medical-records/:id`
- **Auth Endpoints**: `/auth/login`, `/auth/register`, `/auth/logout`
- **Token Management**: Automatic Bearer token injection in headers

## Styling

### Tailwind CSS Configuration
The project uses Tailwind CSS v4 with:
- Custom color palette
- Responsive grid system
- Pre-built utility classes
- Custom shadows and borders

### Key Color Classes
- `bg-blue-*`: Primary actions
- `bg-green-*`: Success states
- `bg-red-*`: Delete/cancel actions
- `bg-orange-*`: Warnings
- `bg-gray-*`: Neutral elements

## Environment Variables

Create `.env` file (copy from `.env.example`):

```env
VITE_API_URL=http://localhost:5000/api
```

## Dependencies

### Core Libraries
- **react**: ^19.2.6 - UI framework
- **react-dom**: ^19.2.6 - DOM rendering
- **react-router-dom**: ^7.17.0 - Client-side routing

### Forms & Validation
- **react-hook-form**: ^7.77.0 - Form management and validation

### Data & Charts
- **recharts**: ^3.8.1 - Charts and visualizations
- **date-fns**: ^4.4.0 - Date formatting and manipulation

### HTTP & Icons
- **axios**: ^1.17.0 - HTTP client
- **heroicons**: ^2.2.0 - Icon library

### Build & Dev
- **vite**: ^8.0.12 - Build tool
- **tailwindcss**: ^4.3.0 - CSS framework
- **postcss**: ^8.5.15 - CSS processing
- **autoprefixer**: ^10.5.0 - CSS prefixer

## Usage Examples

### Creating a New Patient

```jsx
import { usePatients } from './context/PatientContext';

function MyComponent() {
  const { createPatient } = usePatients();

  const handleAddPatient = async () => {
    try {
      const newPatient = await createPatient({
        name: 'John Doe',
        email: 'john@example.com',
        phone: '123-456-7890',
        age: 30,
        condition: 'Hypertension',
      });
      console.log('Patient added:', newPatient);
    } catch (error) {
      console.error('Failed to add patient:', error);
    }
  };

  return <button onClick={handleAddPatient}>Add Patient</button>;
}
```

### Checking User Role

```jsx
import { useAuth } from './context/AuthContext';

function AdminPanel() {
  const { hasRole } = useAuth();

  if (!hasRole(['admin', 'doctor'])) {
    return <p>Access denied</p>;
  }

  return <div>Admin content here</div>;
}
```

### Using Forms with Validation

```jsx
import { useForm } from 'react-hook-form';

function MyForm() {
  const { register, handleSubmit, formState: { errors } } = useForm();

  const onSubmit = (data) => {
    console.log(data);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <input {...register('name', { required: 'Name is required' })} />
      {errors.name && <p>{errors.name.message}</p>}
      <button type="submit">Submit</button>
    </form>
  );
}
```

## Running the Application

```bash
# Install dependencies
npm install

# Start development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview

# Run linter
npm run lint
```

## Testing the Application

### Demo Credentials
- **Email**: doctor@hospital.com
- **Password**: password123

### Test Scenarios
1. Login with demo credentials
2. View dashboard with statistics
3. Create a new patient
4. Schedule an appointment
5. Add medical records
6. Filter and search patients
7. Update/delete records
8. Logout

## Performance Optimization

- **Code Splitting**: Route-based lazy loading available
- **Memoization**: useCallback for expensive operations
- **API Caching**: Implement with context state management
- **Image Optimization**: SVG icons used throughout

## Accessibility Features

- Semantic HTML elements
- ARIA labels on buttons
- Keyboard navigation support
- Color contrast compliance
- Form validation feedback

## Browser Support

- Chrome/Edge: Latest
- Firefox: Latest
- Safari: Latest
- Mobile browsers: iOS Safari, Chrome Mobile

## Known Limitations

1. Medical records API endpoint may need creation if not in backend
2. Real-time updates require WebSocket integration
3. File uploads not yet implemented
4. Offline support not included

## Future Enhancements

1. Export reports to PDF
2. Email notifications
3. File uploads for medical records
4. Advanced search filters
5. Prescription generation
6. Patient communication chat
7. Video consultation integration
8. Mobile app (React Native)

## Troubleshooting

### Issue: "Cannot GET /api/patients"
- **Solution**: Ensure backend is running on `http://localhost:5000`
- Update `VITE_API_URL` in `.env` if backend is on different port

### Issue: "Login fails but request succeeds"
- **Solution**: Check that backend returns `{ token, user }` structure
- Verify token format and localStorage key names

### Issue: Components not showing based on role
- **Solution**: Check user.role value matches role names in components
- Verify AuthContext is properly providing user data

### Issue: Charts not displaying
- **Solution**: Ensure Recharts is installed: `npm install recharts`
- Check that chart data is being populated correctly

## Contributing

When adding new features:
1. Create components in `src/components/`
2. Add context if managing complex state
3. Use existing UI components for consistency
4. Follow react-hook-form patterns for forms
5. Add error handling and loading states
6. Test on multiple screen sizes
