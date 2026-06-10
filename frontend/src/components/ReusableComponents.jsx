// Reusable UI Components Library

// Button Component
export const Button = ({
  children,
  variant = 'primary',
  size = 'md',
  disabled = false,
  isLoading = false,
  icon = null,
  fullWidth = false,
  className = '',
  onClick,
  type = 'button',
  title = '',
}) => {
  const baseClasses = 'font-semibold transition duration-200 flex items-center justify-center gap-2 rounded-lg';
  
  const variants = {
    primary: 'bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white',
    secondary: 'bg-gray-100 hover:bg-gray-200 text-gray-900 border border-gray-300',
    danger: 'bg-red-600 hover:bg-red-700 disabled:bg-gray-400 text-white',
    success: 'bg-green-600 hover:bg-green-700 disabled:bg-gray-400 text-white',
    outline: 'border-2 border-blue-600 text-blue-600 hover:bg-blue-50',
    ghost: 'text-blue-600 hover:bg-blue-50',
  };

  const sizes = {
    sm: 'px-3 py-1 text-sm',
    md: 'px-4 py-2 text-base',
    lg: 'px-6 py-3 text-lg',
  };

  return (
    <button
      type={type}
      disabled={disabled || isLoading}
      onClick={onClick}
      title={title}
      className={`
        ${baseClasses}
        ${variants[variant]}
        ${sizes[size]}
        ${fullWidth ? 'w-full' : ''}
        ${disabled || isLoading ? 'cursor-not-allowed' : 'cursor-pointer'}
        ${className}
      `}
    >
      {isLoading && <span className="w-4 h-4 border-2 border-current border-t-transparent rounded-full animate-spin" />}
      {icon && !isLoading && <span>{icon}</span>}
      {children}
    </button>
  );
};

// Input Component
export const Input = ({
  type = 'text',
  placeholder = '',
  value = '',
  onChange = () => {},
  error = '',
  label = '',
  required = false,
  disabled = false,
  icon = null,
  className = '',
  helperText = '',
  ...props
}) => {
  return (
    <div className="w-full">
      {label && (
        <label className="block text-sm font-medium text-gray-700 mb-1">
          {label}
          {required && <span className="text-red-600">*</span>}
        </label>
      )}
      <div className="relative">
        {icon && (
          <span className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400">
            {icon}
          </span>
        )}
        <input
          type={type}
          placeholder={placeholder}
          value={value}
          onChange={onChange}
          disabled={disabled}
          className={`
            w-full px-4 py-2 border rounded-lg outline-none transition duration-200
            ${icon ? 'pl-10' : ''}
            ${error ? 'border-red-500 focus:ring-2 focus:ring-red-200' : 'border-gray-300 focus:ring-2 focus:ring-blue-500 focus:border-transparent'}
            ${disabled ? 'bg-gray-100 cursor-not-allowed' : ''}
            ${className}
          `}
          {...props}
        />
      </div>
      {error && <p className="text-red-600 text-sm mt-1">{error}</p>}
      {helperText && !error && <p className="text-gray-500 text-sm mt-1">{helperText}</p>}
    </div>
  );
};

// Select Component
export const Select = ({
  options = [],
  value = '',
  onChange = () => {},
  error = '',
  label = '',
  required = false,
  disabled = false,
  placeholder = 'Select an option',
  className = '',
  ...props
}) => {
  return (
    <div className="w-full">
      {label && (
        <label className="block text-sm font-medium text-gray-700 mb-1">
          {label}
          {required && <span className="text-red-600">*</span>}
        </label>
      )}
      <select
        value={value}
        onChange={onChange}
        disabled={disabled}
        className={`
          w-full px-4 py-2 border rounded-lg outline-none transition duration-200
          ${error ? 'border-red-500 focus:ring-2 focus:ring-red-200' : 'border-gray-300 focus:ring-2 focus:ring-blue-500 focus:border-transparent'}
          ${disabled ? 'bg-gray-100 cursor-not-allowed' : ''}
          ${className}
        `}
        {...props}
      >
        <option value="">{placeholder}</option>
        {options.map(opt => (
          <option key={opt.value || opt} value={opt.value || opt}>
            {opt.label || opt}
          </option>
        ))}
      </select>
      {error && <p className="text-red-600 text-sm mt-1">{error}</p>}
    </div>
  );
};

// Textarea Component
export const Textarea = ({
  placeholder = '',
  value = '',
  onChange = () => {},
  error = '',
  label = '',
  required = false,
  disabled = false,
  rows = 3,
  className = '',
  ...props
}) => {
  return (
    <div className="w-full">
      {label && (
        <label className="block text-sm font-medium text-gray-700 mb-1">
          {label}
          {required && <span className="text-red-600">*</span>}
        </label>
      )}
      <textarea
        placeholder={placeholder}
        value={value}
        onChange={onChange}
        disabled={disabled}
        rows={rows}
        className={`
          w-full px-4 py-2 border rounded-lg outline-none transition duration-200 resize-none
          ${error ? 'border-red-500 focus:ring-2 focus:ring-red-200' : 'border-gray-300 focus:ring-2 focus:ring-blue-500 focus:border-transparent'}
          ${disabled ? 'bg-gray-100 cursor-not-allowed' : ''}
          ${className}
        `}
        {...props}
      />
      {error && <p className="text-red-600 text-sm mt-1">{error}</p>}
    </div>
  );
};

// Checkbox Component
export const Checkbox = ({
  checked = false,
  onChange = () => {},
  label = '',
  disabled = false,
  className = '',
}) => {
  return (
    <div className="flex items-center gap-2">
      <input
        type="checkbox"
        checked={checked}
        onChange={onChange}
        disabled={disabled}
        className={`
          w-5 h-5 rounded border-gray-300 text-blue-600 focus:ring-2 focus:ring-blue-500 cursor-pointer
          ${disabled ? 'cursor-not-allowed opacity-50' : ''}
          ${className}
        `}
      />
      {label && (
        <label className={`text-sm text-gray-700 ${disabled ? 'opacity-50' : ''}`}>
          {label}
        </label>
      )}
    </div>
  );
};

// Badge Component
export const Badge = ({
  children,
  variant = 'default',
  size = 'md',
  icon = null,
  className = '',
}) => {
  const variants = {
    default: 'bg-gray-100 text-gray-800',
    primary: 'bg-blue-100 text-blue-800',
    success: 'bg-green-100 text-green-800',
    danger: 'bg-red-100 text-red-800',
    warning: 'bg-yellow-100 text-yellow-800',
    purple: 'bg-purple-100 text-purple-800',
    orange: 'bg-orange-100 text-orange-800',
  };

  const sizes = {
    sm: 'px-2 py-1 text-xs',
    md: 'px-3 py-1 text-sm',
    lg: 'px-4 py-2 text-base',
  };

  return (
    <span
      className={`
        inline-flex items-center gap-1 rounded-full font-semibold
        ${variants[variant]}
        ${sizes[size]}
        ${className}
      `}
    >
      {icon && <span>{icon}</span>}
      {children}
    </span>
  );
};

// Card Component
export const Card = ({
  children,
  title = '',
  subtitle = '',
  className = '',
  onClick = null,
  hoverable = false,
  footer = null,
}) => {
  return (
    <div
      onClick={onClick}
      className={`
        bg-white rounded-lg shadow transition duration-200
        ${hoverable ? 'hover:shadow-lg cursor-pointer' : ''}
        ${onClick ? 'cursor-pointer' : ''}
        ${className}
      `}
    >
      {title && (
        <div className="px-6 py-4 border-b">
          <h3 className="text-lg font-bold text-gray-900">{title}</h3>
          {subtitle && <p className="text-sm text-gray-600 mt-1">{subtitle}</p>}
        </div>
      )}
      <div className="px-6 py-4">
        {children}
      </div>
      {footer && (
        <div className="px-6 py-4 border-t bg-gray-50 rounded-b-lg">
          {footer}
        </div>
      )}
    </div>
  );
};

// FormField Component (Wrapper for form fields with label and error)
export const FormField = ({
  children,
  label = '',
  error = '',
  required = false,
  helperText = '',
}) => {
  return (
    <div className="w-full">
      {label && (
        <label className="block text-sm font-medium text-gray-700 mb-1">
          {label}
          {required && <span className="text-red-600">*</span>}
        </label>
      )}
      {children}
      {error && <p className="text-red-600 text-sm mt-1">{error}</p>}
      {helperText && !error && <p className="text-gray-500 text-sm mt-1">{helperText}</p>}
    </div>
  );
};

// Pagination Component
export const Pagination = ({
  currentPage = 1,
  totalPages = 1,
  onPageChange = () => {},
  className = '',
}) => {
  return (
    <div className={`flex items-center gap-2 justify-center ${className}`}>
      <Button
        variant="outline"
        size="sm"
        disabled={currentPage === 1}
        onClick={() => onPageChange(currentPage - 1)}
      >
        ← Previous
      </Button>

      <div className="flex items-center gap-1">
        {Array.from({ length: totalPages }).map((_, i) => {
          const page = i + 1;
          const show = totalPages <= 7 || 
            page === 1 || 
            page === totalPages || 
            (page >= currentPage - 1 && page <= currentPage + 1);

          if (!show && page === 2) {
            return <span key="dots1">...</span>;
          }
          if (!show && page === totalPages - 1) {
            return <span key="dots2">...</span>;
          }
          if (!show) return null;

          return (
            <Button
              key={page}
              variant={page === currentPage ? 'primary' : 'outline'}
              size="sm"
              onClick={() => onPageChange(page)}
            >
              {page}
            </Button>
          );
        })}
      </div>

      <Button
        variant="outline"
        size="sm"
        disabled={currentPage === totalPages}
        onClick={() => onPageChange(currentPage + 1)}
      >
        Next →
      </Button>
    </div>
  );
};

// Status Indicator Component
export const StatusIndicator = ({
  status = 'active',
  showLabel = true,
  size = 'md',
}) => {
  const statusMap = {
    active: { color: 'bg-green-500', label: 'Active' },
    inactive: { color: 'bg-yellow-500', label: 'Inactive' },
    discharged: { color: 'bg-gray-500', label: 'Discharged' },
    scheduled: { color: 'bg-blue-500', label: 'Scheduled' },
    completed: { color: 'bg-green-500', label: 'Completed' },
    cancelled: { color: 'bg-red-500', label: 'Cancelled' },
    normal: { color: 'bg-green-500', label: 'Normal' },
    high: { color: 'bg-orange-500', label: 'High' },
    critical: { color: 'bg-red-500', label: 'Critical' },
  };

  const config = statusMap[status] || statusMap.active;
  const sizes = { sm: 'w-2 h-2', md: 'w-3 h-3', lg: 'w-4 h-4' };

  return (
    <div className="flex items-center gap-2">
      <div className={`rounded-full ${config.color} ${sizes[size]}`} />
      {showLabel && <span className="text-sm text-gray-700">{config.label}</span>}
    </div>
  );
};

// Data Table Component
export const DataTable = ({
  columns = [],
  data = [],
  onRowClick = null,
  hoverable = true,
  striped = true,
  isLoading = false,
  emptyMessage = 'No data available',
}) => {
  if (isLoading) {
    return (
      <div className="flex justify-center py-8">
        <div className="w-8 h-8 border-4 border-gray-200 border-t-blue-600 rounded-full animate-spin" />
      </div>
    );
  }

  if (data.length === 0) {
    return (
      <div className="text-center py-8 text-gray-500">
        {emptyMessage}
      </div>
    );
  }

  return (
    <div className="overflow-x-auto rounded-lg border border-gray-200">
      <table className="w-full">
        <thead>
          <tr className="bg-gray-50 border-b">
            {columns.map(col => (
              <th
                key={col.key}
                className="px-6 py-3 text-left text-sm font-semibold text-gray-700"
                style={{ width: col.width }}
              >
                {col.label}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {data.map((row, idx) => (
            <tr
              key={idx}
              onClick={() => onRowClick && onRowClick(row)}
              className={`
                border-b transition duration-150
                ${striped && idx % 2 === 0 ? 'bg-gray-50' : 'bg-white'}
                ${hoverable && onRowClick ? 'hover:bg-blue-50 cursor-pointer' : ''}
              `}
            >
              {columns.map(col => (
                <td key={col.key} className="px-6 py-4 text-sm text-gray-900">
                  {col.render ? col.render(row[col.key], row) : row[col.key]}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

// Alert Component
export const Alert = ({
  type = 'info',
  title = '',
  message = '',
  onClose = null,
  icon = null,
}) => {
  const types = {
    info: { bg: 'bg-blue-50', border: 'border-blue-200', text: 'text-blue-800', icon: 'ℹ️' },
    success: { bg: 'bg-green-50', border: 'border-green-200', text: 'text-green-800', icon: '✓' },
    warning: { bg: 'bg-yellow-50', border: 'border-yellow-200', text: 'text-yellow-800', icon: '⚠️' },
    error: { bg: 'bg-red-50', border: 'border-red-200', text: 'text-red-800', icon: '✕' },
  };

  const config = types[type] || types.info;

  return (
    <div className={`${config.bg} border ${config.border} rounded-lg p-4 mb-4 flex justify-between items-start`}>
      <div className="flex items-start gap-3">
        <span className="text-xl">{icon || config.icon}</span>
        <div>
          {title && <h4 className={`font-semibold ${config.text}`}>{title}</h4>}
          {message && <p className={`text-sm ${config.text}`}>{message}</p>}
        </div>
      </div>
      {onClose && (
        <button onClick={onClose} className={`text-lg leading-none ${config.text} hover:opacity-70`}>
          ✕
        </button>
      )}
    </div>
  );
};
