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
        ('Исполнитель'),
        ('Тестировщик');

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

INSERT INTO public.users (id, first_name, middle_name, last_name, position, is_admin, login, password, created_at, updated_at)
    VALUES ('445e2563-67d1-474e-9b8b-ba325dc17ac8', 'Главный администратор', '', '', 'Администратор системы', true, 'admin', '�iv��A��M�߱g��s�K��o*�H�', 'NOW()', 'NOW()');

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
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL,
    project_id uuid NOT NULL,
    role_id integer NOT NULL,
    is_admin_project boolean NOT NULL,
    CONSTRAINT fkey_participants_project_role FOREIGN KEY (role_id) REFERENCES role(id),
    CONSTRAINT fkey_participants_project_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fkey_participants_project_project FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE
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
    CONSTRAINT fkey_task_project FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE,
    CONSTRAINT fkey_task_author FOREIGN KEY (author_id) REFERENCES users(id),
    CONSTRAINT fkey_task_executor FOREIGN KEY (executor_id) REFERENCES users(id),
    CONSTRAINT fkey_task_tester FOREIGN KEY (tester_id) REFERENCES users(id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.task
    OWNER to admin;

