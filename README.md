# Neo-Cloud-Pro Message Server (SMS & EMAIL SERVER)

[![CI](https://github.com/kiddo9/SMS-MAIL-SERVER/actions/workflows/CI.yml/badge.svg)](https://github.com/kiddo9/SMS-MAIL-SERVER/actions/workflows/CI.yml)

A high-performance, gRPC-powered notification service designed for automated batch processing of SMS and Email alerts. This platform provides a streamlined workflow for institutions and businesses to manage templates and dispatch high-volume notifications via file uploads.

## 🚀 Overview

The Neo-Cloud-Pro Message Server is a full-stack solution built with a focus on speed, security, and scalability. It allows administrators to upload recipient data in bulk (Excel/CSV) and trigger personalized notifications across multiple channels using a modern web dashboard.

## ✨ Key Features

-   **Dual-Channel Notifications:** Native support for both Email and SMS (via EBulksms integration).
-   **Automated Batch Processing:** Upload `.xlsx` or `.csv` files to process thousands of notifications in seconds.
-   **Dynamic Template Management:** Create, edit, and manage reusable templates for different notification scenarios.
-   **Secure Authentication:** JWT-based session management with OTP verification for admin users.
-   **Abuse Prevention:** Integrated Google reCAPTCHA v3 and gRPC interceptors for robust security.
-   **Modern Web Dashboard:** A responsive React-based frontend for seamless template management and data uploads.
-   **gRPC-Web Architecture:** Optimized communication between the frontend and backend using Envoy Proxy.

## 🛠️ Technology Stack

### Backend
-   **Language:** Go (Golang)
-   **API:** gRPC & Protocol Buffers
-   **Proxy:** Envoy Proxy (gRPC-Web)
-   **Authentication:** JWT (JSON Web Tokens) & OTP

### Frontend
-   **Framework:** React (TypeScript)
-   **Build Tool:** Vite
-   **Styling:** Tailwind CSS
-   **Icons:** Lucide-React

### Infrastructure
-   **Containerization:** Docker & Docker Compose
-   **Web Server:** Nginx (for client serving)

## 🏗️ Technical Architecture

The Neo-Cloud-Pro Message Server employs a modern, high-performance architecture designed to handle concurrent workloads with minimal latency.

### 🔌 gRPC & Protocol Buffers
At the core of the system is **gRPC**, a high-performance RPC framework. 
-   **Type Safety:** All service interfaces and message structures are defined using **Protocol Buffers (proto3)**, ensuring strict type safety across the Go backend and TypeScript frontend.
-   **Efficiency:** Uses HTTP/2 for transport, enabling features like header compression and multiplexing.
-   **Streaming:** Supports server-side streaming (e.g., for fetching large lists of templates) to improve perceived performance.

### 🛡️ Envoy Proxy (gRPC-Web Bridge)
Since web browsers cannot natively initiate standard gRPC calls, the system utilizes **Envoy Proxy** as a specialized gateway.
-   **gRPC-Web Translation:** Envoy listens for gRPC-Web requests from the React client and translates them into pure gRPC for the backend.
-   **CORS Management:** Handles complex Cross-Origin Resource Sharing (CORS) preflight requests and headers.
-   **Routing:** Acts as a reverse proxy, directing traffic to the appropriate backend services while abstracting the internal network topology.

### 🐹 Go Microservice Design
The backend is a lightweight, concurrent microservice built in Go.
-   **Concurrent Processing:** The `BatchUploadHandler` utilizes a **Worker Pool pattern**. When a file is uploaded, it spawns multiple goroutines to process and dispatch notifications in parallel, significantly reducing processing time.
-   **Interceptor Pipeline:** Employs a custom interceptor chain for modular security. Every request passes through a validation pipeline:
    1.  **reCAPTCHA Interceptor:** Validates the frontend-provided token against Google's API.
    2.  **Auth Interceptor:** Verifies JWT tokens and handles automatic session renewal.
-   **Modular Handlers:** Business logic is decoupled into specific handlers (Admin, Template, Wallet/SMS, FileUpload) for better maintainability.

## 👥 Who Is This For?

-   **Educational Institutions:** Perfect for automating fee balance reminders, attendance alerts, and parent-teacher communications.
-   **Businesses & SMEs:** Ideal for sending bulk marketing updates, transactional alerts, or organizational announcements.
-   **Developers:** A comprehensive boilerplate for building gRPC-web applications with a Go backend and React frontend.

## 📋 File Format Requirements

For batch uploads, the system expects an Excel or CSV file with the following columns:
-   `Name`: Recipient's full name.
-   `Phone`: Mobile number (international format recommended).
-   `Email`: Recipient's email address.
-   `Course`: Associated course or category.
-   `PendingPrice`: The balance or amount to be notified about.

## 🚀 Getting Started

### Prerequisites
-   Docker and Docker Compose
-   Go 1.21+ (for local development)
-   Node.js & NPM (for frontend development)

### Deployment
1.  **Clone the repository:**
    ```bash
    git clone https://github.com/kiddo9/SMS-MAIL-SERVER.git
    ```
2.  **Configure Environment Variables:**
    Create a `.env` file in the `server` directory and the `client` directory based on the provided examples.
3.  **Run with Docker Compose:**
    ```bash
    docker-compose up --build
    ```

## 📜 License

Distributed under the MIT License. See `LICENSE` for more information.
