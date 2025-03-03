CREATE table IF NOT EXISTS students (
    id UUID PRIMARY key default gen_random_uuid(),
    name VARCHAR(30) not NULL,
    last_name VARCHAR(30) not NULL,
    registration VARCHAR(12) not NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

create table if not exists users (
	id UUID primary key default gen_random_uuid(),
	email VARCHAR(200),
	password VARCHAR()
)

insert into students(name, last_name, registration) VALUES(
	'Luiz Carlos',
	'Marinho Junior',
	'202512530075'
);

select * from students;



ALTER TABLE students
ALTER COLUMN registration SET DEFAULT generate_registration();