# Go Backend Template

A simple Go backend application designed to be modular, maintainable, and scalable. This project follows best practices in structuring Go backend applications.

## Project Structure

```
go-backend-template/
├── cmd/
│   └── api/
│       └── main.go                # Entry point of the application
├── internal/
│   ├── config/
│   │   └── config.go              # Configuration loading and management
│   ├── handlers/
│   │   └── handler.go              # HTTP handlers for processing requests
│   ├── services/
│   │   └── service.go              # Business logic layer
│   ├── models/
│   │   └── model.go                # Data models
│   └── middleware/
│       └── middleware.go           # Custom middleware functions
├── pkg/
│   ├── responses/
│   │   └── response.go             # Standardized API response formats
│   └── utils/
│       └── utils.go                # Utility functions
├── api/
│   └── v1/
│       └── routes.go               # API route definitions
services
├── go.mod                          # Go module file
├── go.sum                          # Dependency checksums
└── README.md                       # Project documentation
```

### Directory Descriptions

- **cmd/**: Contains the main application entry point.

  - **api/**: This directory holds the `main.go` file, which starts the HTTP server.

- **internal/**: Contains application-specific logic that is not intended for external use.

  - **config/**: Manages configuration settings and application parameters. This will also contain files for db, caches, pub-sub, etc
  - **handlers/**: Contains HTTP handlers that define how requests are processed.
  - **services/**: Contains business logic and interacts with data models.
  - **models/**: Represents the data structures used in the application.
  - **middleware/**: Holds middleware functions for request processing.

- **pkg/**: Contains reusable code that can be imported by other applications.

  - **responses/**: Functions for creating standardized responses.
  - **utils/**: General utility functions that provide common functionality.

- **api/**: Contains versioning for your API to manage changes without breaking existing clients.

  - **v1/**: Holds routing setup for version 1 of the API.

- **go.mod** and **go.sum**: Standard Go module files for managing dependencies.

- **README.md**: Documentation providing an overview of the project and instructions for setup.

## Getting Started

### Prerequisites

- Go (version 1.24.4 or higher)
- Git

### Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/Aniketyadav44/go-backend-template.git
cd go-backend-template
```

### Renaming to your project

- Open the `go.mod` file and change the module path to your new desired path:

  ```go
  module github.com/yourusername/your-new-project-name
  ```

- Rename the Project Directory:

  ```bash
  mv go-project-strucutre your-new-project-name
  ```

- Update Repository URL (if using Git):
  ```bash
  git remote set-url origin https://github.com/yourusername/your-new-project-name.git
  ```

### Install Dependencies

Run the following command to install the necessary dependencies:

```bash
go mod tidy
```

### Run the Application

To start the application, run:

```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:<port from .env>`.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
