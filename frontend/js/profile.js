document.addEventListener(
    "DOMContentLoaded",
    () => {

        const user = JSON.parse(
            localStorage.getItem("user")
        );

        if (!user) {
            logout();
            return;
        }

        document.getElementById(
            "fullName"
        ).textContent =
            `${user.first_name} ${user.last_name}`;

        document.getElementById(
            "userRole"
        ).textContent =
            user.role;

        document.getElementById(
            "email"
        ).textContent =
            user.email || "-";

        document.getElementById(
            "phone"
        ).textContent =
            user.phone || "-";

        document.getElementById(
            "role"
        ).textContent =
            user.role;

        document.getElementById(
            "employeeId"
        ).textContent =
            user.id;

        const initials =
            `${user.first_name?.[0] || ""}
             ${user.last_name?.[0] || ""}`;

        document.getElementById(
            "avatar"
        ).textContent =
            initials
                .replace(/\s/g, "")
                .toUpperCase();
    }
);