document.addEventListener("DOMContentLoaded", () => {
  loadPatient();
});

async function loadPatient() {
  try {
    const params = new URLSearchParams(window.location.search);

    const patientId = params.get("id");

    if (!patientId) {
      alert("Patient ID missing");

      return;
    }

    const response = await apiRequest(`/patients/${patientId}`);

    const patient = response.data;

    renderPatient(patient);

    await loadAppointments(patientId);
  } catch (error) {
    console.error(error);

    alert("Failed to load patient");
  }
}

function renderPatient(patient) {
  document.getElementById("patientName").textContent =
    `${patient.FirstName} ${patient.LastName}`;

  document.getElementById("patientNumber").textContent =
    `Patient ID: ${patient.PatientNumber}`;

  document.getElementById("patientGender").textContent = patient.Gender || "-";

  document.getElementById("patientBloodGroup").textContent =
    patient.BloodGroup || "-";

  document.getElementById("patientPhone").textContent = patient.Phone || "-";

  document.getElementById("patientEmail").textContent = patient.Email || "-";

  document.getElementById("patientAddress").textContent =
    patient.Address || "-";

  document.getElementById("patientEmergencyPhone").textContent =
    patient.EmergencyContactPhone || "-";

  document.getElementById("patientStatus").textContent = patient.IsActive
    ? "Active"
    : "Inactive";

  document.getElementById("medicalHistoryLink").href =
    `medical-history.html?patientId=${patient.ID}`;

  const initials = `${patient.FirstName?.[0] || ""}${patient.LastName?.[0] || ""}`;

  document.getElementById("avatar").textContent = initials.toUpperCase();

  document.getElementById("patientAge").textContent = calculateAge(
    patient.DateOfBirth,
  );
  document.getElementById("medicalHistoryLink").href =
    `medical-history.html?patientId=${patient.ID}`;
  document.getElementById("doctorNotesLink").href =
    `doctor-notes.html?patientId=${patient.ID}`;
}

function calculateAge(date) {
  if (!date) {
    return "-";
  }

  const dob = new Date(date);

  const today = new Date();

  let age = today.getFullYear() - dob.getFullYear();

  const monthDiff = today.getMonth() - dob.getMonth();

  if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < dob.getDate())) {
    age--;
  }

  return age;
}

async function loadAppointments(patientId) {
  try {
    const response = await apiRequest(`/appointments/patient/${patientId}`);

    renderAppointments(response.data || []);
  } catch (error) {
    console.error(error);
  }
}

function renderAppointments(appointments) {
  const tbody = document.getElementById("appointmentsTableBody");

  if (!tbody) return;

  tbody.innerHTML = "";

  if (appointments.length === 0) {
    tbody.innerHTML = `
            <tr>
                <td colspan="4">
                    No appointments found
                </td>
            </tr>
        `;

    return;
  }

  appointments.forEach((appointment) => {
    const date = new Date(appointment.AppointmentDate);

    tbody.innerHTML += `
                <tr>

                    <td>
                        Dr.
                        ${appointment.Doctor?.FirstName || ""}
                        ${appointment.Doctor?.LastName || ""}
                    </td>

                    <td>
                        ${date.toLocaleDateString()}
                    </td>

                    <td>
                        ${date.toLocaleTimeString([], {
                          hour: "2-digit",
                          minute: "2-digit",
                        })}
                    </td>

                    <td>
                        ${appointment.Status}
                    </td>

                </tr>
            `;
  });
}
