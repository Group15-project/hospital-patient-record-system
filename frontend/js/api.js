const API_BASE_URL ="https://hospital-patient-record-system-1.onrender.com/api/v1" ;
///"http://localhost:8080/api/v1"
 async function apiRequest(
    endpoint,
    method = "GET",
    body = null
) {
    
    const token =
        localStorage.getItem("access_token");

    const headers = {
        "Content-Type": "application/json",
    };

    if (token) {
        headers["Authorization"] =
            `Bearer ${token}`;
    }

    const response = await fetch(
        `${API_BASE_URL}${endpoint}`,
        {
            method,
            headers,
            body: body
                ? JSON.stringify(body)
                : null,
        }
    );

const text = await response.text();

const data = text
    ? JSON.parse(text)
    : {};

if (!response.ok) {
    throw new Error(
        data.message || "Request failed"
    );
}

return data;
}