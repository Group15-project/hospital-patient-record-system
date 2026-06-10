document.addEventListener("DOMContentLoaded", async () => {
  await loadPatients();

  await loadDoctors();

  await loadAppointments();

  initializeForm();
});

async function loadPatients() {
  try {
    const response = await apiRequest("/patients");

    const patients = response.data || [];

    const select = document.getElementById("patientSelect");

    patients.forEach((patient) => {
      select.innerHTML += `
                    <option value="${patient.ID}">
                        ${patient.PatientNumber}
                        - 
                        ${patient.FirstName}
                        ${patient.LastName}
                    </option>
                `;
    });
  } catch (error) {
    console.error(error);
  }
}
async function loadDoctors() {
  try {
    const response = await apiRequest("/auth/doctors");

    const doctors = response.data || [];

    const select = document.getElementById("doctorSelect");

    doctors.forEach((doctor) => {
      select.innerHTML += `
                    <option value="${doctor.ID}">
                        Dr. ${doctor.FirstName}
                        ${doctor.LastName}
                    </option>
                `;
    });
  } catch (error) {
    console.error(error);
  }
}
async function loadAppointments() {
  try {

    const user = getCurrentUser();
    if (user?.role === "DOCTOR") {

    document.querySelector(".form-card").style.display = "none";

}

    let endpoint = "/appointments";

    if (
      user &&
      user.role === "DOCTOR"
    ) {
      endpoint =
        `/appointments/doctor/${user.id}`;
    }

    const response =
      await apiRequest(endpoint);

    renderAppointments(
      response.data || []
    );

  } catch (error) {

    console.error(
      "Failed to load appointments",
      error
    );
  }
}
function renderAppointments(appointments) {
  const tbody = document.getElementById("appointmentsTableBody");

  if (!tbody) return;

  tbody.innerHTML = "";

  appointments.forEach((appointment) => {
    const date = new Date(appointment.AppointmentDate);

    let priorityClass = "priority-normal";

    if (appointment.Priority === "URGENT") {
      priorityClass = "priority-urgent";
    }

    if (appointment.Priority === "EMERGENCY") {
      priorityClass = "priority-emergency";
    }

    tbody.innerHTML += `
        <tr>

            <td>
                ${appointment.Patient?.FirstName || ""}
                ${appointment.Patient?.LastName || ""}
            </td>

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
                <span class="priority ${priorityClass}">
                    ${appointment.Priority}
                </span>
            </td>

        </tr>
    `;
  });
}

function initializeForm() {
  const form = document.getElementById("appointmentForm");

  if (!form) return;

  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    try {
      const date = document.getElementById("appointmentDate").value;

      const time = document.getElementById("appointmentTime").value;

      const payload = {
        patient_id: Number(document.getElementById("patientSelect").value),

        doctor_id: Number(document.getElementById("doctorSelect").value),

        appointment_date: `${date}T${time}:00Z`,

        reason: document.getElementById("appointmentReason").value,
        priority: document.getElementById("priority").value,
      };

      await apiRequest("/appointments", "POST", payload);

      alert("Appointment scheduled successfully");

      form.reset();

      await loadAppointments();
    } catch (error) {
      console.error(error);

      alert("Failed to schedule appointment");
    }
  });
}
