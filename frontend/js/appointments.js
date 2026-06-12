let selectedAppointmentId = null;
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

    if (user && user.role === "DOCTOR") {
      endpoint = `/appointments/doctor/${user.id}`;
    }

    const response = await apiRequest(endpoint);

    renderAppointments(response.data || []);
  } catch (error) {
    console.error("Failed to load appointments", error);
  }
}
function renderAppointments(appointments) {
  const tbody = document.getElementById("appointmentsTableBody");

  if (!tbody) return;

  const user = getCurrentUser();

  tbody.innerHTML = "";

  appointments.forEach((appointment) => {
    const date = new Date(appointment.AppointmentDate);

    let statusClass = "status-scheduled";

    if (appointment.Status === "COMPLETED") {
      statusClass = "status-completed";
    }

    if (appointment.Status === "CANCELLED") {
      statusClass = "status-cancelled";
    }

    if (appointment.Status === "NO_SHOW") {
      statusClass = "status-no-show";
    }

    let priorityClass = "priority-normal";

    if (appointment.Priority === "URGENT") {
      priorityClass = "priority-urgent";
    }

    if (appointment.Priority === "EMERGENCY") {
      priorityClass = "priority-emergency";
    }

    const canManage =
      user?.role === "RECEPTIONIST" ||
      user?.role === "ADMIN" ||
      user?.role === "SUPER_ADMIN";

    const canModify = appointment.Status === "SCHEDULED";
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
        <td>

    <span class="status ${statusClass}">

        ${appointment.Status}

    </span>

</td>

<td>

    ${
      canManage && canModify
        ? `
        <div class="appointment-actions">

            <button
                class="reschedule-btn"
                onclick="openRescheduleModal(${appointment.ID})"
            >
                Reschedule
            </button>

            <button
                class="cancel-btn"
                onclick="openCancelModal(${appointment.ID})"
            >
                Cancel
            </button>

        </div>
        `
        : "-"
    }

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
async function saveReschedule() {

  const date =
    document.getElementById("rescheduleDate").value;

  const time =
    document.getElementById("rescheduleTime").value;

  if (!date || !time) {

    alert("Please select date and time");

    return;
  }

  try {

    await apiRequest(
      `/appointments/${selectedAppointmentId}`,
      "PUT",
      {
        appointment_date:
          `${date}T${time}:00Z`,
      }
    );

    alert(
      "Appointment rescheduled successfully"
    );

    closeRescheduleModal();

    await loadAppointments();

  } catch (error) {

    console.error(error);

    alert(
      error?.message ||
      "Failed to reschedule appointment"
    );
  }
}
async function confirmCancelAppointment() {

    try {

        await apiRequest(
            `/appointments/${selectedAppointmentId}/status`,
            "PATCH",
            {
                status: "CANCELLED",
            }
        );

        closeCancelModal();

        alert(
            "Appointment cancelled successfully"
        );

        await loadAppointments();

    } catch (error) {

        console.error(error);

        alert(
            error?.message ||
            "Failed to cancel appointment"
        );
    }
}

function openRescheduleModal(id) {
  selectedAppointmentId = id;

  document.getElementById("rescheduleModal").classList.add("show");
}

function closeRescheduleModal() {
  selectedAppointmentId = null;

  document.getElementById("rescheduleModal").classList.remove("show");

  document.getElementById("rescheduleDate").value = "";

  document.getElementById("rescheduleTime").value = "";
}

async function saveReschedule() {
  const date = document.getElementById("rescheduleDate").value;

  const time = document.getElementById("rescheduleTime").value;

  if (!date || !time) {
    alert("Please select date and time");
    return;
  }

  try {
    await apiRequest(`/appointments/${selectedAppointmentId}`, "PUT", {
      appointment_date: `${date}T${time}:00Z`,
    });

    closeRescheduleModal();

    alert("Appointment rescheduled successfully");

    await loadAppointments();
  } catch (error) {
    console.error(error);

    alert(error?.message || "Failed to reschedule appointment");
  }
}


function openCancelModal(id) {

    selectedAppointmentId = id;

    document
        .getElementById("cancelModal")
        .classList
        .add("show");
}

function closeCancelModal() {

    selectedAppointmentId = null;

    document
        .getElementById("cancelModal")
        .classList
        .remove("show");
}