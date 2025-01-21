CREATE TABLE departements (
  id CHAR(36) PRIMARY KEY,
  departement_name VARCHAR(255) NOT NULL,
  max_clock_in_time DATETIME(3) DEFAULT NULL,
  max_clock_out_time DATETIME(3) DEFAULT NULL
);

CREATE TABLE employees (
  id CHAR(36) NOT NULL PRIMARY KEY,
  departement_id CHAR(36),
  name VARCHAR(255) NOT NULL,
  address TEXT,
  created_at DATETIME(3) NOT NULL,
  updated_at DATETIME(3) NOT NULL,
  FOREIGN KEY (departement_id) REFERENCES departements(id)
);

CREATE TABLE attendances (
  id CHAR(36) NOT NULL PRIMARY KEY,
  employee_id CHAR(36) NOT NULL,
  clock_in TIMESTAMP NULL,
  clock_out TIMESTAMP NULL,
  created_at DATETIME(3),
  updated_at DATETIME(3),
  FOREIGN KEY (employee_id) REFERENCES employees(id)
);

CREATE TABLE attendance_histories (
  id CHAR(36) NOT NULL PRIMARY KEY,
  employee_id CHAR(36) NOT NULL,
  attendance_id CHAR(36) NOT NULL,
  date_attendance TIMESTAMP NULL,
  attendance_type TINYINT(1),
  description TEXT,
  created_at DATETIME(3),
  updated_at DATETIME(3),
  FOREIGN KEY (attendance_id) REFERENCES attendances(id),
  FOREIGN KEY (employee_id) REFERENCES employees(id)
);
