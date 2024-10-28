DROP TABLE IF EXISTS  project;
DROP TABLE IF EXISTS  participants_project;
DROP TABLE IF EXISTS  task;
DROP TABLE IF EXISTS  user_app;
DROP TABLE IF EXISTS  role;
DROP EXTENSION IF EXISTS  "uuid-ossp";
DROP SEQUENCE IF EXISTS  public.role__role_id_seq;
DROP SEQUENCE IF EXISTS  public.project__project_id_seq;
DROP SEQUENCE IF EXISTS  public.participants_project__participants_project_id_seq;
DROP SEQUENCE IF EXISTS  public.task__task_id_seq;