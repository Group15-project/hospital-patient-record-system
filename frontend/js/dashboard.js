
async function loadDashboard() {
  try {
    const response = await apiRequest("/dashboard/summary");

    const data = response.data;

    document.getElementById("totalPatients").textContent =
      data.total_patients ?? 0;

    document.getElementById("appointmentsToday").textContent =
      data.total_appointments ?? 0;

    document.getElementById("doctorsAvailable").textContent =
      data.total_doctors ?? 0;

   document.getElementById("emergencyCases").textContent =
    data.emergency_cases ?? 0;
    renderPatients(data.recent_patients);
  } catch (error) {
    console.error(error);

    if (error.message.toLowerCase().includes("unauthorized")) {
      logout();
    }
  }
}
function renderPatients(patients = []) {
  const tbody = document.getElementById("recentPatientsBody");

  tbody.innerHTML = "";

  patients.forEach((patient) => {
    tbody.innerHTML += `
            <tr>
                <td>${patient.patient_number}</td>
                <td>${patient.first_name} ${patient.last_name}</td>
                <td>${patient.gender}</td>
                <td>${patient.phone}</td>
                <td>
                    <span class="status active">
                        ${patient.status}
                    </span>
                </td>
            </tr>
        `;
  });
}
const user = JSON.parse(
    localStorage.getItem("user")
);

if (user) {
    document.getElementById(
        "loggedInUser"
    ).textContent =
        `${user.first_name} ${user.last_name}`;
}
if (user) {
    document.getElementById(
        "welcomeText"
    ).textContent =
        `Welcome back, ${user.first_name}`;
}

loadDashboard();
