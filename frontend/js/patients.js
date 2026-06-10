document.addEventListener("DOMContentLoaded", () => {
  loadPatients();
  initializePatientForm();
});

async function loadPatients() {
  try {
    const response = await apiRequest("/patients");

    console.log(response);

    console.log(response.data);

    const patients = response.data || [];

    renderPatients(patients);
  } catch (error) {
    console.error("Failed to load patients:", error);

    if (error.message?.toLowerCase().includes("unauthorized")) {
      logout();
    }
  }
}

function renderPatients(patients) {
  const tbody = document.getElementById("patientsTableBody");

  if (!tbody) return;

  tbody.innerHTML = "";

  patients.forEach((patient) => {
    tbody.innerHTML += `
            <tr>
                <td>${patient.PatientNumber}</td>
                <td>${patient.FirstName} ${patient.LastName}</td>
                <td>${patient.Gender}</td>
                <td>${patient.Phone || "-"}</td>
                <td>
                    <a
                        href="patient-details.html?id=${patient.ID}"
                        class="view-btn"
                    >
                        View
                    </a>
                </td>
            </tr>
        `;
  });
}

function initializePatientForm() {
  const form = document.getElementById("patientForm");

  if (!form) return;

  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    try {
      const payload = {
        first_name: document.getElementById("firstName").value,

        last_name: document.getElementById("lastName").value,

        gender: document.getElementById("gender").value,

        date_of_birth: document.getElementById("dateOfBirth").value,

        phone: document.getElementById("phone").value,

        email: document.getElementById("email").value,

        address: document.getElementById("address").value,

        emergency_contact_name: document.getElementById("emergencyName").value,

        emergency_contact_phone:
          document.getElementById("emergencyPhone").value,

        blood_group: document.getElementById("bloodGroup").value,
      };
      console.log(payload);
      await apiRequest(
        "/patients",

        "POST",

        payload,
      );
    

      alert("Patient created successfully");

      form.reset();

      await loadPatients();
    } catch (error) {
      console.error(error);

      alert("Failed to create patient");
    }
  });
}
