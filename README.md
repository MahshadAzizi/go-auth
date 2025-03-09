# 🚀 Go Authentication API with JWT & PostgreSQL

A simple and secure authentication system using **Go (Golang)**, **JWT (JSON Web Token)**, and **PostgreSQL**. This project follows **clean architecture** and is structured with handlers, middleware, and utilities.

---

## 📌 Features
✅ **User Registration** – Secure password hashing with bcrypt  
✅ **User Login** – Generates JWT for authentication  
✅ **JWT Authentication** – Middleware protects secure routes  
✅ **Profile Retrieval** – Fetch user details from the database  
✅ **PostgreSQL Integration** – Stores user data securely  

---

## 🚀 Installation & Setup

### **1️⃣ Clone the Repository**
```bash
git clone https://github.com/MahshadAzizi/go-auth.git
cd go-auth
```

### 2️⃣ Environment Variables
```bash
DB_HOST=yourhost
DB_PORT=yourport
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdatabase
DB_SSLMODE=DB_SSLMODE
JWT_SECRET_KEY=your_secret_key
```
### 3️⃣ Install Dependencies
```bash
go mod tidy
```

### 4️⃣ Run the Application
```bash
go run .
```
