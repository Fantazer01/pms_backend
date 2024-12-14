CREATE ROLE admin WITH LOGIN SUPERUSER PASSWORD 'strongPass123';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- SEQUENCE: public.role__role_id_seq

-- DROP SEQUENCE IF EXISTS public.role__role_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.role__role_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.role__role_id_seq
    OWNER TO admin;

-- Table: public.role

-- DROP TABLE IF EXISTS public.role;

CREATE TABLE IF NOT EXISTS public.role
(
    id integer NOT NULL DEFAULT nextval('role__role_id_seq'::regclass),
    name_role text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT role_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.role
    OWNER to admin;

INSERT INTO public.role (name_role)
    VALUES 
        ('Руководитель проекта'),
        ('Участник');

-- Table: public.user_app

-- DROP TABLE IF EXISTS public.user_app;

CREATE TABLE IF NOT EXISTS public.users
(
    id uuid PRIMARY KEY,
    first_name text COLLATE pg_catalog."default" NOT NULL,
    middle_name text COLLATE pg_catalog."default" NOT NULL,
    last_name text COLLATE pg_catalog."default" NOT NULL,
    "position" text COLLATE pg_catalog."default" NOT NULL,
    is_admin boolean NOT NULL,
    login text COLLATE pg_catalog."default" NOT NULL,
    password bytea NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to admin;

-- Table: public.project

-- DROP TABLE IF EXISTS public.project;

CREATE TABLE IF NOT EXISTS public.project
(
    id uuid PRIMARY KEY,
    name text COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default" NOT NULL,
    is_active boolean NOT NULL default true,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.project
    OWNER to admin;

-- Table: public.participants_project

-- DROP TABLE IF EXISTS public.participants_project;

CREATE TABLE IF NOT EXISTS public.participants_project
(
    user_id uuid NOT NULL,
    project_id uuid NOT NULL,
    role_id integer NOT NULL,
    is_admin_project boolean NOT NULL,
    
    CONSTRAINT user_in_project_pk PRIMARY KEY (user_id, project_id),
    CONSTRAINT user_in_project_user_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT user_in_project_project_fk FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.participants_project
    OWNER to admin;

CREATE TYPE status_task AS ENUM ('Открыта', 'Отложено', 'В работе', 'На тестировании', 'На ревью', 'Завершена');

-- Table: public.task

-- DROP TABLE IF EXISTS public.task;

CREATE TABLE IF NOT EXISTS public.task
(
    id uuid PRIMARY KEY,
    name text COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default" NOT NULL,
    status status_task NOT NULL,
    project_id uuid NOT NULL,
    created_at timestamptz NOT NULL,
    deadline timestamptz NOT NULL,
    author_id uuid NOT NULL,
    executor_id uuid NOT NULL,
    tester_id uuid NOT NULL,
    CONSTRAINT fkey_task_project FOREIGN KEY (project_id) REFERENCES project(id)ON DELETE CASCADE,
	projectname text COLLATE pg_catalog."default",
    CONSTRAINT fkey_task_author FOREIGN KEY (author_id) REFERENCES users(id),
	firstname_author text COLLATE pg_catalog."default",
    lastname_author text COLLATE pg_catalog."default",
    CONSTRAINT fkey_task_executor FOREIGN KEY (executor_id) REFERENCES users(id),
	firstname_executor text COLLATE pg_catalog."default",
    lastname_executor text COLLATE pg_catalog."default",
    CONSTRAINT fkey_task_tester FOREIGN KEY (tester_id) REFERENCES users(id),
	firstname_tester text COLLATE pg_catalog."default",
    lastname_tester text COLLATE pg_catalog."default"
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.task
    OWNER to admin;

