let patients = [];

document.addEventListener("DOMContentLoaded", async () => {
  await loadPatients();

  document
    .getElementById("searchBtn")
    .addEventListener("click", searchPatients);

  document
    .getElementById("searchInput")
    .addEventListener("keyup", searchPatients);
});

async function loadPatients() {
  try {
    const response = await apiRequest("/patients");

    patients = response.data || [];
  } catch (error) {
    console.error(error);

    alert("Failed to load patients");
  }
}

function searchPatients() {
  const keyword = document

    .getElementById("searchInput")

    .value.trim()

    .toLowerCase();

  if (keyword.length < 2) {
    document.getElementById("searchResultsBody").innerHTML = `

            <tr>

                

            </tr>

        `;

    document.getElementById("resultCount").textContent = "0 Records Found";

    return;
  }

  const filtered = patients.filter((patient) => {
    return (
      patient.PatientNumber?.toLowerCase().includes(keyword) ||
      patient.FirstName?.toLowerCase().includes(keyword) ||
      patient.LastName?.toLowerCase().includes(keyword) ||
      patient.Phone?.toLowerCase().includes(keyword)
    );
  });

  renderPatients(filtered);
}

function renderPatients(data) {
  const tbody = document.getElementById("searchResultsBody");

  const count = document.getElementById("resultCount");

  tbody.innerHTML = "";

  count.textContent = `${data.length} Records Found`;

  if (data.length === 0) {
    tbody.innerHTML = `
            <tr>
                <td colspan="6">
                    No patients found
                </td>
            </tr>
        `;

    return;
  }

  data.forEach((patient) => {
    tbody.innerHTML += `
                <tr>

                    <td>
                        ${patient.PatientNumber}
                    </td>

                    <td>
                        ${patient.FirstName}
                        ${patient.LastName}
                    </td>

                    <td>
                        ${patient.Gender || "-"}
                    </td>

                    <td>
                        ${patient.Phone || "-"}
                    </td>

                    <td>
                        <span class="status active">
                            ${patient.IsActive ? "Active" : "Inactive"}
                        </span>
                    </td>

                    <td>
                        <a
                            href="patient-details.html?id=${patient.ID}"
                            class="view-btn"
                        >
                            View Record
                        </a>
                    </td>

                </tr>
            `;
  });
}
