let allPatients = [];
let selectedPatientId = null;
let allergies = [];


document.addEventListener("DOMContentLoaded", () => {
  loadPatients();
  initializePatientForm();
    initializeSearch();
});

document.querySelector(".add-allergy-btn")
  ?.addEventListener("click", addAllergy);

async function loadPatients() {
  try {
    const response = await apiRequest("/patients");

    console.log(response);

    console.log(response.data);

  allPatients = response.data || [];

renderPatients(allPatients);
   
  } catch (error) {
    console.error("Failed to load patients:", error);

    if (error.message?.toLowerCase().includes("unauthorized")) {
      logout();
    }
  }
}

function renderPatients(patients) {
  const user = getCurrentUser();
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

    ${
      
      user?.role === "SUPER_ADMIN"
      
        ? `
        <button
          class="danger-btn"
          onclick="openDeletePatientModal(${patient.ID})"
        >
          Delete
        </button>
      `
        : ""
    }
</td>
            </tr>
        `;
  });
console.log("Current User:", user);
console.log("Role:", user?.role);
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
        allergies: allergies
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

async function confirmDeletePatient() {

  if (!selectedPatientId) {
    return;
  }

  try {

    await apiRequest(
      `/patients/${selectedPatientId}`,
      "DELETE"
    );

    closeDeletePatientModal();

    alert(
      "Patient deleted successfully"
    );

    await loadPatients();

  } catch (error) {

    console.error(error);

    alert(
      error?.message ||
      "Failed to delete patient"
    );
  }
}


function initializeSearch() {
  const searchInput =
    document.querySelector(".search-input");

  if (!searchInput) return;

  searchInput.addEventListener("input", () => {

    const keyword =
      searchInput.value
      .trim()
      .toLowerCase();

    if (!keyword) {
      renderPatients(allPatients);
      return;
    }

    const filtered =
      allPatients.filter((patient) => {

        return (
          patient.PatientNumber
            ?.toLowerCase()
            .includes(keyword) ||

          patient.FirstName
            ?.toLowerCase()
            .includes(keyword) ||

          patient.LastName
            ?.toLowerCase()
            .includes(keyword) ||

          `${patient.FirstName} ${patient.LastName}`
            .toLowerCase()
            .includes(keyword) ||

          patient.Phone
            ?.toLowerCase()
            .includes(keyword)
        );
      });

    renderPatients(filtered);
  });
}


function openDeletePatientModal(id) {

  selectedPatientId = id;

  document
    .getElementById("deletePatientModal")
    .classList
    .add("show");
}

function closeDeletePatientModal() {

  selectedPatientId = null;

  document
    .getElementById("deletePatientModal")
    .classList
    .remove("show");
}

function addAllergy() {
  const input = document.getElementById("allergyInput");

  const value = input.value.trim();

  if (!value) return;

  allergies.push(value);

  renderAllergies();

  input.value = "";
}

function renderAllergies() {
  const container =
    document.getElementById("allergyList");

  container.innerHTML = allergies
    .map(
      (allergy, index) => `
        <span class="medical-tag">
          ${allergy}
          <button
            type="button"
            onclick="removeAllergy(${index})"
          >
            ×
          </button>
        </span>
      `
    )
    .join("");
}

function removeAllergy(index) {
  allergies.splice(index, 1);
  renderAllergies();
}

window.removeAllergy = removeAllergy;