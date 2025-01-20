-- Seeder: Add sample departments
INSERT INTO departements (id, departement_name, max_clock_in_time, max_clock_out_time) 
VALUES 
  (UUID(), 'IT', NOW(), NOW() + INTERVAL 8 HOUR),
  (UUID(), 'HR', NOW(), NOW() + INTERVAL 8 HOUR);

-- Insert employees with UUID for id and departement_id
INSERT INTO employees (id, departement_id, name, address, created_at, updated_at)
VALUES
  (UUID(), (SELECT id FROM departements WHERE departement_name = 'IT'), 'John Doe', 'Jakarta', NOW(), NOW()),
  (UUID(), (SELECT id FROM departements WHERE departement_name = 'HR'), 'Jane Doe', 'Bandung', NOW(), NOW());

-- Insert attendances with UUID for id and employee_id
INSERT INTO attendances (id, employee_id, clock_in, clock_out, created_at, updated_at)
VALUES
  (UUID(), (SELECT id FROM employees WHERE name = 'John Doe'), '2025-01-20 08:00:00', '2025-01-20 17:00:00', NOW(), NOW()),
  (UUID(), (SELECT id FROM employees WHERE name = 'Jane Doe'), '2025-01-20 09:00:00', '2025-01-20 18:00:00', NOW(), NOW());

-- Insert attendance histories with UUID for id, employee_id, and attendance_id
INSERT INTO attendance_histories (id, employee_id, attendance_id, date_attendance, attendance_type, description, created_at, updated_at)
VALUES
  (UUID(), (SELECT id FROM employees WHERE name = 'John Doe'), (SELECT id FROM attendances WHERE employee_id = (SELECT id FROM employees WHERE name = 'John Doe')), '2025-01-20 08:00:00', 1, 'On Time', NOW(), NOW()),
  (UUID(), (SELECT id FROM employees WHERE name = 'Jane Doe'), (SELECT id FROM attendances WHERE employee_id = (SELECT id FROM employees WHERE name = 'Jane Doe')), '2025-01-20 09:00:00', 0, 'Late', NOW(), NOW());
