# Reusable Components Usage Guide

This guide demonstrates how to use all the new reusable components available in `ReusableComponents.jsx`.

## Quick Reference

All components are exported from `src/components/index.js` and can be imported like:

```jsx
import { Button, Input, Select, Badge, Card, DataTable, Alert } from './components';
```

---

## Components Overview

### 1. **Button Component**

Versatile button component with multiple variants and sizes.

**Props:**
- `variant`: 'primary' | 'secondary' | 'danger' | 'success' | 'outline' | 'ghost' (default: 'primary')
- `size`: 'sm' | 'md' | 'lg' (default: 'md')
- `disabled`: boolean (default: false)
- `isLoading`: boolean - shows spinner (default: false)
- `icon`: React element or string
- `fullWidth`: boolean (default: false)
- `onClick`: function
- `type`: 'button' | 'submit' (default: 'button')

**Examples:**

```jsx
// Basic button
<Button onClick={() => console.log('clicked')}>Click Me</Button>

// Danger button with loading
<Button variant="danger" isLoading={isSubmitting}>
  Delete
</Button>

// Success button with icon
<Button variant="success" icon="✓">
  Confirm
</Button>

// Full width outline button
<Button variant="outline" fullWidth>
  Select Option
</Button>

// Ghost button for secondary actions
<Button variant="ghost">Learn More</Button>

// Submit button in forms
<Button type="submit" size="lg">
  Submit Form
</Button>
```

---

### 2. **Input Component**

Enhanced text input with validation support.

**Props:**
- `type`: 'text' | 'email' | 'password' | 'number' | 'tel' etc.
- `placeholder`: string
- `value`: string
- `onChange`: function
- `error`: string - shows error message if present
- `label`: string
- `required`: boolean
- `disabled`: boolean
- `icon`: React element or string (left side)
- `helperText`: string - helper text below input

**Examples:**

```jsx
// Basic input
const [name, setName] = useState('');
<Input
  type="text"
  placeholder="Enter your name"
  value={name}
  onChange={(e) => setName(e.target.value)}
/>

// Input with label and validation error
<Input
  type="email"
  label="Email Address"
  required
  error={emailError}
  placeholder="user@example.com"
  icon="📧"
/>

// Input with helper text
<Input
  type="password"
  label="Password"
  helperText="Must be at least 8 characters"
  icon="🔒"
/>

// Disabled input
<Input
  type="text"
  label="Reference Number"
  value="REF-12345"
  disabled
/>
```

---

### 3. **Select Component**

Dropdown/select input with validation.

**Props:**
- `options`: array of {label, value} or strings
- `value`: string
- `onChange`: function
- `error`: string
- `label`: string
- `required`: boolean
- `disabled`: boolean
- `placeholder`: string (default: 'Select an option')

**Examples:**

```jsx
// Basic select
<Select
  options={['Active', 'Inactive', 'Pending']}
  value={status}
  onChange={(e) => setStatus(e.target.value)}
/>

// Select with labels and values
<Select
  label="Patient Status"
  required
  options={[
    { label: 'Active', value: 'active' },
    { label: 'Inactive', value: 'inactive' },
    { label: 'Discharged', value: 'discharged' },
  ]}
  value={patientStatus}
  onChange={(e) => setPatientStatus(e.target.value)}
/>

// Select with error
<Select
  label="Appointment Type"
  error={typeError}
  options={['Checkup', 'Follow-up', 'Surgery']}
/>
```

---

### 4. **Textarea Component**

Multi-line text input with validation.

**Props:**
- `placeholder`: string
- `value`: string
- `onChange`: function
- `error`: string
- `label`: string
- `required`: boolean
- `disabled`: boolean
- `rows`: number (default: 3)

**Examples:**

```jsx
// Basic textarea
<Textarea
  placeholder="Enter your notes..."
  value={notes}
  onChange={(e) => setNotes(e.target.value)}
/>

// Textarea with label and validation
<Textarea
  label="Medical History"
  required
  error={historyError}
  rows={5}
  placeholder="Describe any relevant medical history"
/>
```

---

### 5. **Checkbox Component**

Styled checkbox input.

**Props:**
- `checked`: boolean
- `onChange`: function
- `label`: string
- `disabled`: boolean

**Examples:**

```jsx
// Basic checkbox
<Checkbox
  checked={agreeToTerms}
  onChange={(e) => setAgreeToTerms(e.target.checked)}
  label="I agree to the terms and conditions"
/>

// Disabled checkbox
<Checkbox
  checked={true}
  disabled
  label="This option is locked"
/>
```

---

### 6. **Badge Component**

Status badge/label component.

**Props:**
- `variant`: 'default' | 'primary' | 'success' | 'danger' | 'warning' | 'purple' | 'orange'
- `size`: 'sm' | 'md' | 'lg'
- `icon`: React element or string
- `children`: React node

**Examples:**

```jsx
// Basic badges
<Badge variant="success">Active</Badge>
<Badge variant="danger">Critical</Badge>
<Badge variant="warning">Pending</Badge>

// Badge with icon
<Badge variant="primary" icon="✓">
  Verified
</Badge>

// Different sizes
<Badge size="sm" variant="default">Small</Badge>
<Badge size="md" variant="primary">Medium</Badge>
<Badge size="lg" variant="success">Large</Badge>
```

---

### 7. **Card Component**

Container component with header, body, and footer.

**Props:**
- `title`: string - card header title
- `subtitle`: string - subtitle below title
- `hoverable`: boolean - adds hover effect
- `onClick`: function - makes card clickable
- `footer`: React node - footer content
- `children`: React node - main content

**Examples:**

```jsx
// Basic card
<Card title="Patient Information">
  <p>John Doe, 30 years old</p>
  <p>Condition: Hypertension</p>
</Card>

// Hoverable clickable card
<Card
  title="Dr. Smith"
  subtitle="Cardiologist"
  hoverable
  onClick={() => navigate(`/doctor/${drId}`)}
>
  <p>Available: Mon-Fri, 9AM-5PM</p>
</Card>

// Card with footer
<Card
  title="Appointment"
  footer={
    <div className="flex gap-2">
      <Button variant="secondary" size="sm">Cancel</Button>
      <Button size="sm">Confirm</Button>
    </div>
  }
>
  <p>Date: June 10, 2026</p>
  <p>Time: 2:00 PM</p>
</Card>
```

---

### 8. **FormField Component**

Wrapper component for consistent form field styling.

**Props:**
- `label`: string
- `error`: string
- `required`: boolean
- `helperText`: string
- `children`: React node

**Examples:**

```jsx
// Using FormField with any input
<FormField label="Email" required error={emailError}>
  <input
    type="email"
    placeholder="user@example.com"
    value={email}
    onChange={(e) => setEmail(e.target.value)}
  />
</FormField>

// Wrapping custom components
<FormField label="Custom Input" helperText="This is a custom field">
  <input type="text" />
</FormField>
```

---

### 9. **Pagination Component**

Navigation for paginated data.

**Props:**
- `currentPage`: number (default: 1)
- `totalPages`: number (default: 1)
- `onPageChange`: function(pageNumber)

**Examples:**

```jsx
// Basic pagination
const [page, setPage] = useState(1);
const itemsPerPage = 10;
const totalPages = Math.ceil(items.length / itemsPerPage);

<Pagination
  currentPage={page}
  totalPages={totalPages}
  onPageChange={setPage}
/>

// In a list component
const startIdx = (page - 1) * itemsPerPage;
const paginatedItems = items.slice(startIdx, startIdx + itemsPerPage);
```

---

### 10. **StatusIndicator Component**

Visual status indicator with dot and label.

**Props:**
- `status`: 'active' | 'inactive' | 'discharged' | 'scheduled' | 'completed' | 'cancelled' | 'normal' | 'high' | 'critical'
- `showLabel`: boolean (default: true)
- `size`: 'sm' | 'md' | 'lg'

**Examples:**

```jsx
// Basic status indicator
<StatusIndicator status="active" />

// Without label
<StatusIndicator status="critical" showLabel={false} />

// Different sizes
<StatusIndicator status="scheduled" size="sm" />
<StatusIndicator status="completed" size="lg" />
```

---

### 11. **DataTable Component**

Sortable, filterable data table.

**Props:**
- `columns`: array of {key, label, width, render}
- `data`: array of objects
- `onRowClick`: function(row) - optional
- `hoverable`: boolean (default: true)
- `striped`: boolean (default: true)
- `isLoading`: boolean
- `emptyMessage`: string

**Examples:**

```jsx
// Basic table
const columns = [
  { key: 'name', label: 'Name' },
  { key: 'email', label: 'Email' },
  { key: 'status', label: 'Status' },
];

<DataTable columns={columns} data={patients} />

// Table with custom rendering and click handler
<DataTable
  columns={[
    { key: 'name', label: 'Name' },
    {
      key: 'age',
      label: 'Age',
      render: (age) => `${age} years old`,
    },
    {
      key: 'status',
      label: 'Status',
      render: (status) => <Badge variant="success">{status}</Badge>,
    },
  ]}
  data={patients}
  onRowClick={(patient) => navigate(`/patient/${patient.id}`)}
/>

// Loading state
<DataTable
  columns={columns}
  data={patients}
  isLoading={loading}
/>

// Empty state with custom message
<DataTable
  columns={columns}
  data={[]}
  emptyMessage="No patients found matching your criteria"
/>
```

---

### 12. **Alert Component**

Notification/alert message.

**Props:**
- `type`: 'info' | 'success' | 'warning' | 'error'
- `title`: string
- `message`: string
- `onClose`: function
- `icon`: React element or string

**Examples:**

```jsx
// Info alert
<Alert
  type="info"
  title="New Feature"
  message="You can now schedule appointments online"
/>

// Success alert with close button
<Alert
  type="success"
  title="Success"
  message="Patient record updated successfully"
  onClose={() => setShowAlert(false)}
/>

// Error alert with custom icon
<Alert
  type="error"
  title="Error"
  message="Failed to save changes. Please try again."
  icon="❌"
/>

// Warning alert
<Alert
  type="warning"
  message="This action cannot be undone"
/>
```

---

## Real-World Usage Examples

### Example 1: Patient Form

```jsx
import { Button, Input, Select, Textarea, Alert } from './components';
import { useForm } from 'react-hook-form';

export function PatientForm({ onSubmit }) {
  const { register, handleSubmit, formState: { errors, isSubmitting } } = useForm();
  const [alert, setAlert] = useState(null);

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <Input
        label="Full Name"
        required
        error={errors.name?.message}
        {...register('name', { required: 'Name is required' })}
      />

      <Input
        type="email"
        label="Email"
        required
        error={errors.email?.message}
        {...register('email', { 
          required: 'Email is required',
          pattern: { value: /^[^\s@]+@[^\s@]+\.[^\s@]+$/, message: 'Invalid email' }
        })}
      />

      <Select
        label="Status"
        options={['active', 'inactive', 'discharged']}
        {...register('status')}
      />

      <Textarea
        label="Medical Condition"
        rows={4}
        {...register('condition')}
      />

      {alert && <Alert {...alert} />}

      <Button
        type="submit"
        fullWidth
        isLoading={isSubmitting}
      >
        Save Patient
      </Button>
    </form>
  );
}
```

### Example 2: Data Table with Pagination

```jsx
import { DataTable, Pagination, Button, Badge } from './components';
import { useState } from 'react';

export function PatientList() {
  const [page, setPage] = useState(1);
  const itemsPerPage = 10;
  const patients = [...]; // your data

  const totalPages = Math.ceil(patients.length / itemsPerPage);
  const startIdx = (page - 1) * itemsPerPage;
  const paginatedData = patients.slice(startIdx, startIdx + itemsPerPage);

  const columns = [
    { key: 'name', label: 'Name' },
    { key: 'email', label: 'Email' },
    {
      key: 'status',
      label: 'Status',
      render: (status) => <Badge variant="success">{status}</Badge>,
    },
    {
      key: 'actions',
      label: 'Actions',
      render: (_, patient) => (
        <div className="flex gap-2">
          <Button size="sm" onClick={() => editPatient(patient)}>Edit</Button>
          <Button size="sm" variant="danger" onClick={() => deletePatient(patient.id)}>Delete</Button>
        </div>
      ),
    },
  ];

  return (
    <>
      <DataTable columns={columns} data={paginatedData} />
      <Pagination
        currentPage={page}
        totalPages={totalPages}
        onPageChange={setPage}
      />
    </>
  );
}
```

### Example 3: Status Dashboard

```jsx
import { Card, Badge, StatusIndicator, Alert } from './components';

export function Dashboard() {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <Card title="Patient Overview">
        <div className="space-y-3">
          <div className="flex justify-between items-center">
            <span>Total Patients</span>
            <Badge variant="primary">152</Badge>
          </div>
          <div className="flex justify-between items-center">
            <span>Active Cases</span>
            <StatusIndicator status="active" />
          </div>
          <div className="flex justify-between items-center">
            <span>Critical Cases</span>
            <Badge variant="danger">3</Badge>
          </div>
        </div>
      </Card>

      <Card title="System Status">
        <div className="space-y-2">
          <StatusIndicator status="active" />
          <p className="text-sm text-gray-600">All systems operational</p>
        </div>
      </Card>

      <Card title="Alerts">
        <Alert
          type="warning"
          message="2 pending approvals"
        />
      </Card>
    </div>
  );
}
```

---

## Styling Guide

All components use **Tailwind CSS** and follow these color schemes:

- **Primary**: Blue (#3b82f6)
- **Success**: Green (#10b981)
- **Danger**: Red (#ef4444)
- **Warning**: Yellow (#f59e0b)
- **Info**: Blue (#3b82f6)

### Custom Styling

All components accept a `className` prop for additional custom styles:

```jsx
<Button className="custom-class">Custom Button</Button>
<Input className="border-2" />
<Badge className="text-lg">Large Badge</Badge>
```

---

## Best Practices

1. **Use semantic HTML**: Components follow accessibility best practices
2. **Error handling**: Always provide error messages for validation
3. **Loading states**: Use `isLoading` prop on buttons during async operations
4. **Consistent spacing**: Use Tailwind's gap/margin utilities for consistent spacing
5. **Dark mode ready**: Components are designed to support dark mode (future enhancement)
6. **Responsive design**: Components are mobile-first and responsive
7. **Type safety**: Consider adding TypeScript for better developer experience

---

## Migration Guide

If updating existing components to use these reusable components:

### Before:
```jsx
<button className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg">
  Click
</button>
```

### After:
```jsx
<Button>Click</Button>
```

All styling is handled by the component, reducing code duplication!
