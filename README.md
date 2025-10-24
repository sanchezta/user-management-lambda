# 🧠 User Management Lambda

This project is a **Go-based AWS Lambda function** designed to handle **Cognito Post Confirmation events**.  
It initializes the AWS SDK configuration and processes user data after a successful registration in Cognito.

---

## 🚀 Features

- Written in **Go** using **AWS SDK v2**.
- Uses **AWS Lambda Go Runtime** (`aws-lambda-go`).
- Automatically loads AWS configuration from environment or local `.aws/config`.
- Supports **Cognito Post Confirmation trigger** events.

---

## 🧩 Project Structure

user-management-lambda/
├── aws/
│ └── init.go # AWS SDK initialization
├── models/
│ └── secret.go # Models for RDS secrets and signup
├── main.go # Lambda entry point
├── go.mod # Module definition
└── go.sum # Dependency checksums

---

## ⚙️ Setup

### 1. Clone the repository

```bash
git clone hhttps://github.com/sanchezta/user-management-lambda.git
cd user-management-lambda
```
