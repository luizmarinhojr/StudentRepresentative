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
	email VARCHAR(200) not null,
	pass VARCHAR(254) not NULL, 
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

insert into students(name, last_name, registration) VALUES(
	'Luiz Carlos',
	'Marinho Junior',
	'202512530075'
);

select * from students;

ALTER TABLE students
ALTER COLUMN registration SET DEFAULT generate_registration();

SELECT EXISTS (SELECT 1 FROM students WHERE registration = '202512530075');

select name from students where registration = '202512530075';

ALTER TABLE students
ADD COLUMN user_id UUID,
ADD CONSTRAINT fk_students_users
FOREIGN KEY (user_id)
REFERENCES users(id);

SELECT s.id, s.name, s.last_name, s.registration, s.created_at, s.updated_at, u.id, u.email FROM students s full join users u on s.user_id = u.id;


insert into users (email, pass) VALUES('marweedofc@gmail.com', 'smadjsa1uh32u1ybhbdudybhdbaysudvasuyjhsd(*d78sAdusa');

update students set user_id = '5d806edb-cc34-46c6-97e1-9ef6a708a443' where id = '32181552-16df-46cf-8934-83dabe617c70';