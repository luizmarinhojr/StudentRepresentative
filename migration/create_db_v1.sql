CREATE TYPE user_role AS ENUM ('admin', 'student', 'representative');

create table if not exists users (
	id serial primary key,
	external_id UUID default gen_random_uuid() UNIQUE,
	email VARCHAR(200) not null UNIQUE,
	pass BYTEA not NULL, 
	role user_role default 'student',
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE table IF NOT EXISTS students (
	id serial primary key,
    external_id UUID default gen_random_uuid() UNIQUE,
    name VARCHAR(30) not NULL,
    last_name VARCHAR(30) not NULL,
    registration VARCHAR(12) not null UNIQUE,
    user_id INTEGER UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_students_users
      FOREIGN KEY(user_id)
        REFERENCES users(id)
);

create table if not exists classes (
	id serial primary key,
	external_id UUID default gen_random_uuid() UNIQUE,
	name VARCHAR(30) not null unique,
	start_year smallint not null CHECK (start_year >= EXTRACT(YEAR FROM CURRENT_DATE) - 5 AND start_year <= EXTRACT(YEAR FROM CURRENT_DATE) + 5),
	start_semester smallint not null CHECK (start_semester IN (1, 2)),
	end_year smallint not null CHECK (end_year >= start_year and end_year <= EXTRACT(YEAR FROM CURRENT_DATE) + 7),
	end_semester smallint not null CHECK (end_semester IN (1, 2)),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
	CHECK (
        (end_year > start_year) OR
        (end_year = start_year AND end_semester >= start_semester)
    )
);

create table if not exists students_classes (
	id_student INTEGER references students(id),
	id_class INTEGER references classes(id),
	enrollment_date DATE DEFAULT CURRENT_DATE,
	primary KEY(id_student, id_class)
);

