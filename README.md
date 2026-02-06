
---

# 📝 Todo API (Go + PostgreSQL + sqlc)

A simple yet production-style Todo REST API built with Go, PostgreSQL, sqlc, and Docker.

> 🚀 This project was developed and tested using GitHub Codespaces to ensure a fully cloud-based, reproducible development environment.


---

✨ Features

+ CRUD operations for Todos

+ PostgreSQL persistence

+ Type-safe SQL using sqlc

+ Dockerized database

+ Clean Go project structure

+ Ready for extension (auth, pagination, users)



---

🛠 Tech Stack

- Go 1.22

- PostgreSQL

- sqlc

- Docker & docker-compose

- GitHub Codespaces



---

📁 Project Structure
```
todo-app/
├── cmd/server/            # Application entry point
│   ├── main.go
│   └── mod.go
├── internal/
│   ├── handlers/          # HTTP handlers
│   └── db/sqlc/           # sqlc generated code
├── db/
│   └── queries.sql        # SQL queries
├── migrations/            # Database migrations
├── docker-compose.yml
├── sqlc.yaml
├── .env.example
├── go.mod
└── README.md

```
---

🚀 Running the Project (GitHub Codespaces)

1️⃣ Open in GitHub Codespaces

Click Code → Codespaces → Create codespace

Wait for the environment to finish initializing

Open the terminal


> 💡 Codespaces provides Docker and Go pre-installed.




---

2️⃣ Install PostgreSQL client (psql)
``
sudo apt update
sudo apt install postgresql-client -y
``
Verify:
```
psql --version
```

---

3️⃣ Start PostgreSQL (Docker)
```
docker-compose up -d
```
Confirm:
```
docker ps
```

---

4️⃣ Apply database migrations
```
psql postgres://postgres:postgres@localhost:5432/todoapp \
  -f migrations/0001_init.up.sql
```

  ---

  5️⃣ Generate sqlc code
```
  sqlc generate
```

  ---

  6️⃣ Install Go dependencies automatically
```
  go mod tidy
```

  ---

  7️⃣ Run the API server
```
  go run ./cmd/server
```
  Server starts on:
```
  http://localhost:8080
```

  ---

  🔌 API Endpoints

  Create a Todo
```
  POST /todos

  {
      "title": "Learn sqlc"
  }
```

  ---

  List Todos

  GET /todos


  ---

  Get Todo by ID
```
  GET /todos/{id}

```
  ---

  Update Todo
```
  PATCH /todos/{id}

  {
      "completed": true
  }
```

  ---

  Delete Todo
```
  DELETE /todos/{id}
```
---

  🧪 Quick Test (curl)
```
  curl -X POST http://localhost:8080/todos \
    -H "Content-Type: application/json" \
      -d '{"title":"First Todo"}'
```

 ---     

      🧠 Notes

      Environment variables are defined in .env.example

      PostgreSQL runs in Docker

      Dependencies are managed automatically via Go modules

      sqlc ensures compile-time SQL correctness




      ✅ Project Status

      ✔ Completed
      ✔ Fully functional
      ✔ GitHub Codespaces compatible
      ✔ Ready for extension and production hardening



      📌 Possible Improvements

      User authentication (JWT)

      Pagination & filtering

      Automated migrations

      Unit & integration tests

      CI/CD pipeline

      Deployment to Fly.io / Render




     


  

  

 
