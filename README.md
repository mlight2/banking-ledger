# 🏦 Golang Banking Ledger

A **scalable banking ledger system** built with **Golang**, **PostgreSQL**, **MongoDB**, **Kafka**, and **Docker**.

---

## 🚀 Features
- 🔹 Account management (Create, Deposit, Withdraw)
- 🔹 PostgreSQL for **account balances**
- 🔹 MongoDB for **transaction history**
- 🔹 Kafka for **event-driven transaction processing**
- 🔹 Dockerized for **easy deployment**
- 🔹 Follows **clean architecture** and **best practices**

---

## ⚙️ Installation & Setup

### **1️⃣ Create a `config.yml` File**
In the project root, create a file named **`config.yml`** and paste the following:

```yaml
SERVER_PORT: "8080"
POSTGRES_URL: "postgres://user:password@localhost:5432/banking?sslmode=disable"
MONGO_URI: "mongodb://localhost:27017"
KAFKA_BROKERS: "localhost:9092"


