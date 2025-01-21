-- Drop attendance histories first (due to foreign key constraint)
DELETE FROM attendance_histories WHERE attendance_id IN (SELECT id FROM attendances);

-- Drop attendances second (due to foreign key constraint)
DELETE FROM attendances WHERE employee_id IN (SELECT id FROM employees);

-- Drop employees third (due to foreign key constraint)
DELETE FROM employees WHERE departement_id IN (SELECT id FROM departements);

-- Drop departements last
DELETE FROM departements;
