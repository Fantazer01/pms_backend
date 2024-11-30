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

-- SEQUENCE: public.project__project_id_seq

-- DROP SEQUENCE IF EXISTS public.project__project_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.project__project_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.project__project_id_seq
    OWNER TO admin;

-- SEQUENCE: public.participants_project__project_id_seq

-- DROP SEQUENCE IF EXISTS public.participants_project__project_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.participants_project__participants_project_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.participants_project__participants_project_id_seq
    OWNER TO admin;

-- SEQUENCE: public.task__task_id_seq

-- DROP SEQUENCE IF EXISTS public.task__task_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.task__task_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.task__task_id_seq
    OWNER TO admin;

-- Table: public.role

-- DROP TABLE IF EXISTS public.role;

CREATE TABLE IF NOT EXISTS public.role
(
    role_id integer NOT NULL DEFAULT nextval('role__role_id_seq'::regclass),
    name_role text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT role_pkey PRIMARY KEY (role_id)
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

CREATE TABLE IF NOT EXISTS public.user_app
(
    user_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    first_name text COLLATE pg_catalog."default" NOT NULL,
    last_name text COLLATE pg_catalog."default" NOT NULL,
    "position" text COLLATE pg_catalog."default" NOT NULL,
    is_admin boolean NOT NULL,
    login text COLLATE pg_catalog."default" NOT NULL,
    password text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT user_pkey PRIMARY KEY (user_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.user_app
    OWNER to admin;

INSERT INTO public.user_app (first_name, last_name, position, is_admin, login, password)
    VALUES ('Главный администратор', '', 'Администратор системы', true, 'admin', 'admin');

-- Table: public.project

-- DROP TABLE IF EXISTS public.project;

CREATE TABLE IF NOT EXISTS public.project
(
    project_id bigint NOT NULL DEFAULT nextval('project__project_id_seq'::regclass),
    name_project text COLLATE pg_catalog."default" NOT NULL,
    data_creation date NOT NULL,
    description text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT project_pkey PRIMARY KEY (project_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.project
    OWNER to admin;

-- Table: public.participants_project

-- DROP TABLE IF EXISTS public.participants_project;

CREATE TABLE IF NOT EXISTS public.participants_project
(
    participants_project_id bigint NOT NULL DEFAULT nextval('participants_project__participants_project_id_seq'::regclass),
    user_id uuid NOT NULL,
    project_id bigint NOT NULL,
    role_id integer NOT NULL,
    is_admin_project boolean NOT NULL,
    CONSTRAINT participants_project_pkey PRIMARY KEY (participants_project_id),
    CONSTRAINT fkey_role FOREIGN KEY (role_id)
        REFERENCES public.role (role_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fkey_user FOREIGN KEY (user_id)
        REFERENCES public.user_app (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.participants_project
    OWNER to admin;

-- Table: public.task

-- DROP TABLE IF EXISTS public.task;

CREATE TABLE IF NOT EXISTS public.task
(
    task_id bigint NOT NULL DEFAULT nextval('task__task_id_seq'::regclass),
    name text COLLATE pg_catalog."default" NOT NULL,
    project_id bigint NOT NULL,
    status text COLLATE pg_catalog."default" NOT NULL,
    data_creation timestamp without time zone NOT NULL,
    date_deadline time without time zone NOT NULL,
    description text COLLATE pg_catalog."default" NOT NULL,
    user_author_id uuid NOT NULL,
    user_executor_id uuid NOT NULL,
    user_tester_id uuid NOT NULL,
    CONSTRAINT task_pkey PRIMARY KEY (task_id),
    CONSTRAINT fkey_project FOREIGN KEY (project_id)
        REFERENCES public.role (role_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fkey_user_author FOREIGN KEY (user_author_id)
        REFERENCES public.user_app (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fkey_user_executor FOREIGN KEY (user_executor_id)
        REFERENCES public.user_app (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fkey_user_tester FOREIGN KEY (user_tester_id)
        REFERENCES public.user_app (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.task
    OWNER to admin;

