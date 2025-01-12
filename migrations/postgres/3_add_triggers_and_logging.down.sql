DROP TRIGGER IF EXISTS trigger_before_operation_participants_project ON public.participants_project;
DROP TRIGGER IF EXISTS trigger_after_operation_participants_project ON public.participants_project;
DROP TRIGGER IF EXISTS trigger_before_operation_role ON public.role;
DROP TRIGGER IF EXISTS trigger_after_operation_role ON public.role;
DROP TRIGGER IF EXISTS trigger_before_operation_users ON public.users;
DROP TRIGGER IF EXISTS trigger_after_operation_users ON public.users;
DROP TRIGGER IF EXISTS trigger_before_operation_project ON public.project;
DROP TRIGGER IF EXISTS trigger_after_operation_project ON public.project;
DROP TRIGGER IF EXISTS trigger_before_operation_task ON public.task;
DROP TRIGGER IF EXISTS trigger_after_operation_task ON public.task;
DROP FUNCTION IF EXISTS public.before_operation_function;
DROP FUNCTION IF EXISTS public.after_operation_function;
DROP TABLE IF EXISTS public.log_changes;
DROP SEQUENCE IF EXISTS public.log_changes__log_id_seq;