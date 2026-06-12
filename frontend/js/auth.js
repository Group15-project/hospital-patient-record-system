async function login(email, password) {
  const data = await apiRequest("/auth/login", "POST", {
    email,
    password,
  });

  localStorage.setItem("access_token", data.data.access_token);

  localStorage.setItem("refresh_token", data.data.refresh_token);

  localStorage.setItem("user", JSON.stringify(data.data.user));

  localStorage.setItem("user_role", data.data.user.role);

  return data;
}
//    document.addEventListener(
//     "DOMContentLoaded",
//     () => {
//         requireAuth();
//         applyRolePermissions();
//     }
// );

function logout() {
  localStorage.clear();
  window.location.href = "/";
}

function requireAuth() {
  const token = localStorage.getItem("access_token");

  if (!token) {
    window.location.href = "/";
  }
}

function getCurrentUser() {
  const user = localStorage.getItem("user");

  return user ? JSON.parse(user) : null;
}

function isAdmin() {
  const user = getCurrentUser();

  return ["ADMIN", "SUPER_ADMIN"].includes(user?.role);
}

function applyRolePermissions() {
  const user = getCurrentUser();

  if (!user || (user.role !== "ADMIN" && user.role !== "SUPER_ADMIN")) {
    document.querySelectorAll(".dashboard-menu").forEach((item) => {
      item.style.display = "none";
    });
  }

  if (!user || user.role !== "DOCTOR") {
    const doctorNotesMenu = document
      .querySelector('a[href="doctor-notes.html"]')
      ?.closest("li");

    if (doctorNotesMenu) {
      doctorNotesMenu.style.display = "none";
    }

    const doctorNotesButton = document.getElementById("doctorNotesLink");

    if (doctorNotesButton) {
      doctorNotesButton.style.display = "none";
    }
  }
  const staffMenu = document.getElementById("staffMenu");

  if (staffMenu) {
    if (user?.role === "SUPER_ADMIN") {
      staffMenu.style.display = "block";
    } else {
      staffMenu.style.display = "none";
    }
  }
}
