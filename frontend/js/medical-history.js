let currentPatientId = null;
let allPatients = [];

document.addEventListener("DOMContentLoaded", async () => {
  const user = getCurrentUser();

  if (user?.role !== "DOCTOR") {
    const formCard = document.querySelector(".form-card");

    if (formCard) {
      formCard.style.display = "none";
    }
  }

  await loadPatients();

  initializeForm();

  initializeSearch();

  const patientId = new URLSearchParams(window.location.search).get(
    "patientId",
  );

  if (patientId) {
    currentPatientId = Number(patientId);

    const searchCard = document.getElementById("searchCard");

    if (searchCard) {
      searchCard.style.display = "none";
    }

    await loadPatient(currentPatientId);

    await loadMedicalHistory(currentPatientId);
  }
});

async function loadPatients() {
  try {
    const response = await apiRequest("/patients");

    allPatients = response.data || [];

    console.log("Patients:", allPatients);
  } catch (error) {
    console.error(error);
  }
}

async function loadPatient(patientId) {
  try {
    const response = await apiRequest(`/patients/${patientId}`);

    const patient = response.data;

    document.getElementById("patientNumber").textContent =
      patient.PatientNumber;

    document.getElementById("patientName").textContent =
      `${patient.FirstName} ${patient.LastName}`;

    document.getElementById("patientGender").textContent =
      patient.Gender || "-";

    document.getElementById("patientAge").textContent = calculateAge(
      patient.DateOfBirth,
    );
  } catch (error) {
    console.error(error);
  }
}

function initializeSearch() {
  const searchBtn = document.getElementById("searchBtn");

  const searchInput = document.getElementById("searchInput");

  const suggestions = document.getElementById("patientSuggestions");

  if (!searchBtn || !searchInput || !suggestions) {
    return;
  }

  searchInput.addEventListener("input", () => {
    const keyword = searchInput.value.trim().toLowerCase();

    if (!keyword) {
      suggestions.innerHTML = "";

      return;
    }

    const matches = allPatients.filter((patient) => {
      return (
        patient.PatientNumber?.toLowerCase().includes(keyword) ||
        patient.FirstName?.toLowerCase().includes(keyword) ||
        patient.LastName?.toLowerCase().includes(keyword) ||
        `${patient.FirstName} ${patient.LastName}`
          .toLowerCase()
          .includes(keyword) ||
        patient.Phone?.toLowerCase().includes(keyword)
      );
    });

    suggestions.innerHTML = matches
      .slice(0, 5)
      .map(
        (patient) => `
              <div
                class="suggestion-item"
                onclick="selectPatient(${patient.ID})"
              >
                ${patient.PatientNumber}
                -
                ${patient.FirstName}
                ${patient.LastName}
              </div>
            `,
      )
      .join("");
  });

  searchBtn.addEventListener("click", searchPatient);

  searchInput.addEventListener("keydown", (e) => {
    if (e.key === "Enter") {
      e.preventDefault();

      searchPatient();
    }
  });
}

async function selectPatient(id) {
  const patient = allPatients.find((p) => p.ID === id);

  if (!patient) {
    return;
  }

  currentPatientId = id;

  document.getElementById("searchInput").value =
    `${patient.FirstName} ${patient.LastName}`;

  document.getElementById("patientSuggestions").innerHTML = "";

  await loadPatient(id);

  await loadMedicalHistory(id);
}

window.selectPatient = selectPatient;

async function searchPatient() {
  const keyword = document
    .getElementById("searchInput")
    .value.trim()
    .toLowerCase();

  if (!keyword) {
    clearPatientInfo();

    return;
  }

  const patient = allPatients.find((p) => {
    return (
      p.PatientNumber?.toLowerCase().includes(keyword) ||
      p.FirstName?.toLowerCase().includes(keyword) ||
      p.LastName?.toLowerCase().includes(keyword) ||
      `${p.FirstName} ${p.LastName}`.toLowerCase().includes(keyword) ||
      p.Phone?.toLowerCase().includes(keyword)
    );
  });

  if (!patient) {
    alert("Patient not found");

    clearPatientInfo();

    return;
  }

  currentPatientId = patient.ID;

  await loadPatient(patient.ID);

  await loadMedicalHistory(patient.ID);
}

async function loadMedicalHistory(patientId) {
  try {
    const response = await apiRequest(`/medical-records/patient/${patientId}`);

    const histories = response.data || [];

    renderHistory(histories);
  } catch (error) {
    console.error(error);

    renderHistory([]);
  }
}

function initializeForm() {
  const form = document.getElementById("medicalHistoryForm");

  if (!form) {
    return;
  }

  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    if (!currentPatientId) {
      alert("Please select a patient");

      return;
    }

    try {
      const payload = {
        patientId: Number(currentPatientId),

        title: document.getElementById("previousIllnesses").value,

        description: `
        Surgeries: ${document.getElementById("surgeries").value}
        Allergies: ${document.getElementById("allergies").value}
        Chronic Conditions: ${document.getElementById("chronicConditions").value}
    `,

        type: "Medical History",

        severity: "Normal",

        doctorName: "System",

        prescription: "",

        date: new Date().toISOString().split("T")[0],
      };
      await apiRequest("/medical-records", "POST", payload);

      alert("Medical history saved successfully");

      form.reset();

      await loadMedicalHistory(currentPatientId);
    } catch (error) {
      console.error(error);

      alert("Failed to save medical history");
    }
  });
}

function renderHistory(histories) {
  const timeline = document.getElementById("timeline");

  const allergyTags = document.getElementById("allergyTags");

  timeline.innerHTML = "";
  allergyTags.innerHTML = "";

  if (histories.length === 0) {
    timeline.innerHTML = `
            <div class="timeline-item">
                <div class="timeline-content">
                    <h4>No medical history found</h4>
                </div>
            </div>
        `;

    return;
  }

  histories.forEach((history) => {
    const recordDate = history.date || history.createdAt;

    timeline.innerHTML += `
            <div class="timeline-item">

                <div class="timeline-date">
                    ${
                      recordDate
                        ? new Date(recordDate).toLocaleDateString()
                        : "-"
                    }
                </div>

                <div class="timeline-content">

                    <h4>
                        ${history.title || "-"}
                    </h4>

                    <p>
                        ${history.description || "-"}
                    </p>

                    <p>
                        <strong>Type:</strong>
                        ${history.type || "-"}
                    </p>

                    <p>
                        <strong>Doctor:</strong>
                        ${history.doctorName || "-"}
                    </p>

                    <p>
                        <strong>Severity:</strong>
                        ${history.severity || "-"}
                    </p>

                </div>

            </div>
        `;

    const match = history.description?.match(/Allergies:\s*(.*)/i);

    if (match) {
      const allergies = match[1].split(",");

      allergies.forEach((allergy) => {
        allergyTags.innerHTML += `
                        <span class="medical-tag">
                            ${allergy.trim()}
                        </span>
                    `;
      });
    }
  });
}

function clearPatientInfo() {
  document.getElementById("patientNumber").textContent = "--";

  document.getElementById("patientName").textContent = "--";

  document.getElementById("patientGender").textContent = "--";

  document.getElementById("patientAge").textContent = "--";

  document.getElementById("timeline").innerHTML = "";

  document.getElementById("allergyTags").innerHTML = "";
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
searchInput.addEventListener("keyup", (e) => {
  if (e.key === "Enter") {
    searchPatient();
  }
});
