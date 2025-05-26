<h2 align="center">📚 Book Store API – Go</h2>

This is a simple Book Store API built in Go using GORM to interact with a MySQL database through Object-Relational Mapping (ORM) and maintains authentication using JWT.

#### 🛠️ Tech Stack
- github.com/gin-gonic/gin - Web framework
- gorm.io/gorm - ORM for Go
- gorm.io/driver/mysql - MySQL driver for GORM
- github.com/golang-jwt/jwt/v5 - JWT Authentication
- Standard Go libraries

> Note: Singnup feature for book manager is not yet implemented. It is planned for a future update.

#### 🚀 Project Setup
1. Clone the repository:
```bash
git clone https://github.com/utkarshkrsingh/bookStoreApi && cd bookStoreApi
```
2. Download dependencies
```bash
go mod tidy
```
3. Run the application
```bash
go run .
```

#### 🔧 Environment Configuration
The application requires DB credentials or configurations from `.env` file in project root directory. The `.env` file should contain this (example)
```env
PORT=8000
DSN="<username>:<password>@tcp(127.0.0.1:3306)/<dbname>?charset=UTF8mb4&parseTime=True&loc=Local"
```

#### 📌 To-Do / Future Improvements
- [ ] Add Swagger/OpenAPI docs
- [ ] Dockerize the application
- [ ] Write unit tests
