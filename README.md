# Attendance API

This application is an attendance and employee management system built using Go, Gin, and MySQL. The API allows users to manage employee data, departments, and attendance history.

## Prerequisites
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Running the Application

To run the application using Docker and Docker Compose, follow these steps:

1. **Clone the Repository**

   First, clone the repository to your local directory:

   ```bash
   git clone https://github.com/auliardana/fleetify-test.git
   cd fleetify-test/deployments
   ```

2. Run docker compose

   then, run docker compose:

   ```bash
   docker-compose up -d
   ```
3. Access the API on http://localhost:9999

    import **Fleetify-test.postman_collection.json** in api folder to postman

    and here some available endpoints:

    - POST /api/v1/employee: Add a new employee.
    - GET /api/v1/employee: Get a list of all employees.
    - PATCH /api/v1/employee/:id: Update employee data by ID.
    - DELETE /api/v1/employee/:id: Delete an employee by ID.
    - POST /api/v1/departement: Add a new department.
    - GET /api/v1/departement: Get a list of all departments.
    - PATCH /api/v1/departement/:id: Update department data by ID.
    - DELETE /api/v1/departement/:id: Delete a department by ID.
    - POST /api/v1/attendance: Clock in for an employee.
    - PUT /api/v1/attendance/:id: Clock out for an employee.
    - GET /api/v1/attendance: Get attendance history of employees


4. if you want to customized environment variable, the file is on configs folder.

    ```bash
    database:
        uri: root:password@tcp(mysql:3306)/absensi?charset=utf8mb4&parseTime=True&loc=Local
        pool:
            idle: 10
            maxconnection: 50
            lifetime: 300
        test: Hello World
    app:
        port: 9999
        log:
        level: 2
        format: json
        output: stdout
  ```
