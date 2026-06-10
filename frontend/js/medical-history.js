let currentPatientId = null;

document.addEventListener("DOMContentLoaded", async () => {
  const patientId = new URLSearchParams(window.location.search).get(
    "patientId",
  );
  if (patientId) {
    document.querySelector(".table-header").style.display = "none";
  }
  await loadPatients();

  initializeForm();

  if (patientId) {
    currentPatientId = patientId;

    document.getElementById("patientSelect").value = patientId;

    await loadPatient(patientId);

    await loadMedicalHistory(patientId);
  }

  document
    .getElementById("patientSelect")
    .addEventListener("change", async function () {
      const id = this.value;

      if (!id) {
        clearPatientInfo();
        return;
      }

      currentPatientId = id;

      await loadPatient(id);

      await loadMedicalHistory(id);
    });
});

async function loadPatients() {
  try {
    const response = await apiRequest("/patients");

    console.log("Response:", response);

    const patients = response.data || [];

    console.log("Patients:", patients);

    const select = document.getElementById("patientSelect");

    console.log("Select:", select);

    patients.forEach((patient) => {
      console.log(patient);

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
