# Golang Scheduler Microservice

This is a backend microservice built using **Go**, **Gin**, and **PostgreSQL**. It allows users to create scheduled jobs through a REST API. Jobs are stored in a database and executed using cron logic in the background.

---

## 🔧 Features

- ✅ Create new jobs using API
- ✅ Store job data in PostgreSQL
- ✅ Schedule tasks using cron
- ✅ Swagger UI for API testing
- ✅ Clean and modular codebase

---

## 🛠 Tech Used

- Golang
- Gin Web Framework
- GORM (PostgreSQL ORM)
- robfig/cron (scheduler)
- swaggo/swag (Swagger documentation)

---
## 📦 Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/scheduler-microservice.git
   ```
2. Navigate to the project directory:
   ```bash
   cd scheduler-microservice
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Set up your PostgreSQL database and update the connection string in `config/config.go`.
5. Run the application:
   ```bash
   go run main.go
   ```                                          
## 📝 API Documentation

```bash
   swag init
   ```
This will generate Swagger documentation in the `docs` folder. You can access it at `http://localhost:8080/swagger/index.html` after running the application.
## 🚀 Running the Application
```bash
   go run main.go
   ```
## 📚 Usage
You can use tools like Postman or cURL to interact with the API. Here are some example endpoints:
- **Create Job**: `POST /jobs`
- **Get All Jobs**: `GET /jobs`
- **Get Job by ID**: `GET /jobs/:id`

## 🧪 Testing
To run tests, use the following command:
```bash
go test ./...
```
## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
## 👥 Contributing
Contributions are welcome! Please fork the repository and submit a pull request for any improvements or bug fixes.  
## 📫 Contact
You can reach me at [thayanithi606@gmail.com](mailto:thayanithi606@gmail.com)
## 📖 Documentation
For more detailed documentation, please refer to the [Wiki](https://github.com/yourusername/scheduler-microservice/wiki)
## 📜 Changelog
- **v1.0.0** - Initial release with basic job scheduling functionality.
- Added Swagger documentation and improved error handling.
- Refactored code for better modularity and maintainability.
- Added unit tests and improved API endpoints.
- Updated dependencies and fixed minor bugs. 
