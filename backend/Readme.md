# Backend Deployment on Render

## 1. Create a Render Web Service

1. Log in to Render.
2. Click **New → Web Service**.
3. Connect the GitHub repository containing the backend.
4. If the backend is inside a subfolder, set the **Root Directory** to the backend folder.
5. Select **Go** as the runtime. Render will automatically install Go.

---

## 2. Configure Environment Variables

Add the following environment variables in Render:

```env
APP_NAME=Hospital Management API
APP_ENV=production
APP_PORT=8080

DB_HOST=<database-host>
DB_PORT=3306
DB_USER=<database-user>
DB_PASSWORD=<database-password>
DB_NAME=hospital_db

JWT_SECRET=<your-secret-key>
JWT_ACCESS_EXPIRY=15m
JWT_REFRESH_EXPIRY=168h

LOG_LEVEL=info
```

---

## 3. Build Configuration

### Build Command

```bash
go mod tidy && go build -o main .
```

### Start Command

```bash
./main
```

---

## 4. Deploy

Click **Create Web Service** and wait for deployment to complete.

After a successful deployment, Render will generate a URL similar to:

```text
https://hospital-backend.onrender.com
```

Copy this URL.

---

## Connect Frontend to Backend

Update the frontend API base URL.

### Before

```javascript
const API_BASE_URL = "http://localhost:8080/api/v1";
```

### After

```javascript
const API_BASE_URL =
  "https://your-render-url.onrender.com/api/v1";
```

Replace `your-render-url.onrender.com` with the actual Render backend URL.

---

## Redeploy Frontend

After updating the API URL:

1. Commit and push the frontend changes.
2. Redeploy the frontend.
3. Test the following features:

- Login
- Patient Registration
- Patient Search
- Medical History
- Appointments
- Doctor Notes

---

## Important

Render provides the application port through the `PORT` environment variable.

Make sure the application reads it:

```go
port := os.Getenv("PORT")

if port == "" {
    port = "8080"
}
```

Without this, the application may fail to start on Render.
