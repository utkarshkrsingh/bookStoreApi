### 📚 Book Store API – Go

This is a simple Book Store API built in Go using GORM to interact with a MySQL database through Object-Relational Mapping (ORM).

#### 🛠️ Tech Stack
- ```gorm.io/gorm``` - ORM for Go
- ```gorm.io/driver/mysql``` - MySQL driver for GORM
- ```net/http``` - for building the http server
- Standard Go libraries

> Note: JWT Authorization is not yet implemented. It is planned for a future update.

#### 🚀 Project Setup
1. Clone the repository:
```bash
git clone https://github.com/utkarshkrsingh/bookStoreApi && cd bookStoreApi
```
2. Download dependencies
```bash
go mod download
```
3. Run the application
```bash
go run ./cmd
```

#### 🔧 Environment Configuration
The application requires DB credentials or configurations from `.env` file in project root directory. The `.env` file should contain this (example)
```env
DATABASE = myDatabaseName
DATABASE_URL = tcp(127.0.0.1:3306)
DB_USER = myUser
DB_PASSWORD = myPassword
```

#### 📌 To-Do / Future Improvements
- [ ] Add JWT Authentication
- [ ] Add Swagger/OpenAPI docs
- [ ] Dockerize the application
- [ ] Write unit tests
