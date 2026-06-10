async function loadProfile() {
  try {
    const response = await apiRequest("/profile");

    const user = response.data;

    document.getElementById("fullName").value =
      `${user.first_name} ${user.last_name}`;

    document.getElementById("email").value =
      user.email || "";

    document.getElementById("phone").value =
      user.phone || "";

    document.getElementById("role").value =
      user.role || "";
  } catch (error) {
    console.error(error);
  }
}

document.addEventListener(
  "DOMContentLoaded",
  loadProfile
);