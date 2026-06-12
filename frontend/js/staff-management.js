document.addEventListener(
  "DOMContentLoaded",
  async () => {

    const user =
      getCurrentUser();

    if (
      user?.role !==
      "SUPER_ADMIN"
    ) {

      window.location.href =
        "dashboard.html";

      return;
    }

    await loadRoles();

    initializeForm();
  }
);
document.addEventListener(
  "DOMContentLoaded",
  async () => {

    const user = getCurrentUser();

    if (user?.role !== "SUPER_ADMIN") {
      window.location.href = "dashboard.html";
      return;
    }

    await loadRoles();
    await loadStaff();

    initializeForm();
  }
);

async function loadRoles() {

  try {

    const response =
      await apiRequest(
        "/auth/roles"
      );

    const roles =
      response.data || [];

    const select =
      document.getElementById(
        "roleSelect"
      );

    roles.forEach(
      (role) => {

        if (
          role.Name ===
          "SUPER_ADMIN"
        ) {
          return;
        }

        select.innerHTML += `
          <option value="${role.ID}">
            ${role.Name}
          </option>
        `;
      }
    );

  } catch (error) {

    console.error(error);
  }
}

function initializeForm() {

  const form =
    document.getElementById(
      "staffForm"
    );

  form.addEventListener(
    "submit",
    async (e) => {

      e.preventDefault();

      try {

        const payload = {

          first_name:
            document.getElementById(
              "firstName"
            ).value,

          last_name:
            document.getElementById(
              "lastName"
            ).value,

          email:
            document.getElementById(
              "email"
            ).value,

          phone:
            document.getElementById(
              "phone"
            ).value,

          password:
            document.getElementById(
              "password"
            ).value,

          role_id: Number(
            document.getElementById(
              "roleSelect"
            ).value
          ),
        };

        await apiRequest(
          "/auth/users",
          "POST",
          payload
        );

        alert(
          "User created successfully"
        );

        form.reset();
        await loadStaff();

      } catch (error) {

        console.error(
          error
        );

        alert(
          error?.message ||
          "Failed to create user"
        );
      }
    }
  );
}

let allStaff = [];

async function loadStaff() {

  try {

    const response =
      await apiRequest(
        "/auth/users"
      );

    allStaff =
      response.data || [];

    renderStaff(allStaff);

  } catch (error) {

    console.error(
      "Failed to load staff:",
      error
    );
  }
}

function renderStaff(staff) {

  const tbody =
    document.getElementById(
      "staffTableBody"
    );

  if (!tbody) return;

  tbody.innerHTML = "";

  staff.forEach((user) => {

    tbody.innerHTML += `
      <tr>
        <td>
          ${user.FirstName}
          ${user.LastName}
        </td>

        <td>
          ${user.Email}
        </td>

        <td>
          ${user.Phone || "-"}
        </td>

        <td>
          ${user.Role?.Name || "-"}
        </td>

        <td>
          ${
            user.IsActive
              ? "Active"
              : "Inactive"
          }
        </td>

        <td>
          <button
            class="danger-btn"
            onclick="openDeleteStaffModal(${user.ID})"
          >
            Delete
          </button>
        </td>
      </tr>
    `; 
  });
}

let selectedStaffId = null;

function openDeleteStaffModal(id) {

  selectedStaffId = id;

  document
    .getElementById("deleteStaffModal")
    .classList
    .add("show");
}

function closeDeleteStaffModal() {

  selectedStaffId = null;

  document
    .getElementById("deleteStaffModal")
    .classList
    .remove("show");
}

async function confirmDeleteStaff() {

  if (!selectedStaffId) {
    return;
  }

  try {

    await apiRequest(
      `/auth/users/${selectedStaffId}`,
      "DELETE"
    );

    closeDeleteStaffModal();

    await loadStaff();

    alert("Staff account deleted successfully");

  } catch (error) {

    console.error(error);

    alert(
      error.message ||
      "Failed to delete staff account"
    );
  }
}