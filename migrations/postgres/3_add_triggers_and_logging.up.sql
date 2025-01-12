CREATE SEQUENCE IF NOT EXISTS public.log_changes__log_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.log_changes__log_id_seq
    OWNER TO admin;
	
CREATE TABLE IF NOT EXISTS public.log_changes
(
 	log_id bigint NOT NULL DEFAULT nextval('log_changes__log_id_seq'::regclass),
	operation text COLLATE pg_catalog."default" NOT NULL,
	tablename text COLLATE pg_catalog."default" NOT NULL,
	id_information uuid NOT NULL,
	status bool NOT NULL
)
TABLESPACE pg_default;
ALTER TABLE IF EXISTS public.log_changes
    OWNER to admin;
CREATE OR REPLACE FUNCTION public.before_operation_function()
RETURNS TRIGGER
AS $$
BEGIN
	IF TG_TABLE_NAME = 'participants_project' THEN
		IF TG_OP = 'DELETE' THEN
			INSERT INTO log_changes (tablename,operation,id_information,status) VALUES (TG_TABLE_NAME,TG_OP,uuid_generate_v5(uuid_nil(), OLD.participants_project_id::text),false);
		ELSE
			INSERT INTO log_changes (tablename,operation,id_information,status) VALUES (TG_TABLE_NAME,TG_OP,uuid_generate_v5(uuid_nil(), NEW.participants_project_id::text),false);
		END IF;
	ELSIF TG_TABLE_NAME = 'role' THEN
		IF TG_OP = 'DELETE' THEN
			INSERT INTO log_changes (tablename,operation,id_information,status) VALUES (TG_TABLE_NAME,TG_OP,uuid_generate_v5(uuid_nil(), OLD.role_id::text),false);
		ELSE
			INSERT INTO log_changes (tablename,operation,id_information,status) VALUES (TG_TABLE_NAME,TG_OP,uuid_generate_v5(uuid_nil(), NEW.role_id::text),false);
		END IF;
	ELSIF TG_TABLE_NAME = 'users' THEN
		IF TG_OP = 'DELETE' THEN
			INSERT INTO log_changes (tablename,operation,id_information,status) VALUES (TG_TABLE_NAME,TG_OP,uuid_generate_v5(uuid_nil(), OLD.user_id::text),false);
		ELSE
			INSERT INTO log_changes (tablename,operation,id_information,status) VALUES (TG_TABLE_NAME,TG_OP,uuid_generate_v5(uuid_nil(), NEW.user_id::text),false);
		END IF;
	ELSIF TG_TABLE_NAME = 'project' THEN
		IF TG_OP = 'DELETE' THEN
			INSERT INTO log_changes (tablename,operation,id_information,status) VALUES (TG_TABLE_NAME,TG_OP,uuid_generate_v5(uuid_nil(), OLD.project_id::text),false);
		ELSE
			INSERT INTO log_changes (tablename,operation,id_information,status) VALUES (TG_TABLE_NAME,TG_OP,uuid_generate_v5(uuid_nil(), NEW.project_id::text),false);
		END IF;
	ELSIF TG_TABLE_NAME = 'task' THEN
		IF TG_OP = 'DELETE' THEN
			INSERT INTO log_changes (tablename,operation,id_information,status) VALUES (TG_TABLE_NAME,TG_OP,uuid_generate_v5(uuid_nil(), OLD.task_id::text),false);
		ELSE
			INSERT INTO log_changes (tablename,operation,id_information,status) VALUES (TG_TABLE_NAME,TG_OP,uuid_generate_v5(uuid_nil(), NEW.task_id::text),false);
		END IF;
	END IF;
	IF TG_OP = 'DELETE' THEN
		RETURN OLD;
	ELSE
		RETURN NEW;
	END IF;
	EXCEPTION
		WHEN OTHERS THEN
			RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.after_operation_function()
RETURNS TRIGGER
AS $$
BEGIN
	UPDATE log_changes 
	SET status = true 
	WHERE log_id=currval('log_changes__log_id_seq');
	IF TG_OP = 'DELETE' THEN
		RETURN OLD;
	ELSE
		RETURN NEW;
	END IF;
	IF TG_TABLE_NAME = 'task' THEN
		IF TG_OP = "INSERT" THEN
			UPDATE task
			SET 
				projectname = (SELECT name FROM project WHERE project_id = NEW.project_id),
				firstname_author = (SELECT first_name FROM users WHERE user_id = NEW.author_id),
				lastname_author = (SELECT last_name FROM users WHERE user_id = NEW.author_id),
				firstname_executor = (SELECT first_name FROM users WHERE user_id = NEW.executor_id),
				lastname_executor = (SELECT last_name FROM users WHERE user_id = NEW.executor_id),
				firstname_tester = (SELECT first_name FROM users WHERE user_id = NEW.tester_id),
				lastname_tester = (SELECT last_name FROM users WHERE user_id = NEW.tester_id)
			WHERE id = NEW.id;
		END IF;
	END IF;
	EXCEPTION
		WHEN OTHERS THEN
			RETURN NULL;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE TRIGGER trigger_before_operation_participants_project BEFORE INSERT OR UPDATE OR DELETE ON public.participants_project 
FOR EACH ROW EXECUTE PROCEDURE public.before_operation_function();
CREATE OR REPLACE TRIGGER trigger_before_operation_role BEFORE INSERT OR UPDATE OR DELETE ON public.role
FOR EACH ROW EXECUTE PROCEDURE public.before_operation_function();
CREATE OR REPLACE TRIGGER trigger_before_operation_users BEFORE INSERT OR UPDATE OR DELETE ON public.users
FOR EACH ROW EXECUTE PROCEDURE public.before_operation_function();
CREATE OR REPLACE TRIGGER trigger_before_operation_project BEFORE INSERT OR UPDATE OR DELETE ON public.project
FOR EACH ROW EXECUTE PROCEDURE public.before_operation_function();
CREATE OR REPLACE TRIGGER trigger_before_operation_task BEFORE INSERT OR UPDATE OR DELETE ON public.task
FOR EACH ROW EXECUTE PROCEDURE public.before_operation_function();
CREATE OR REPLACE TRIGGER trigger_after_operation_participants_project AFTER INSERT OR UPDATE OR DELETE ON public.participants_project 
FOR EACH ROW EXECUTE PROCEDURE public.after_operation_function();
CREATE OR REPLACE TRIGGER trigger_after_operation_role AFTER INSERT OR UPDATE OR DELETE ON public.role
FOR EACH ROW EXECUTE PROCEDURE public.after_operation_function();
CREATE OR REPLACE TRIGGER trigger_after_operation_users AFTER INSERT OR UPDATE OR DELETE ON public.users
FOR EACH ROW EXECUTE PROCEDURE public.after_operation_function();
CREATE OR REPLACE TRIGGER trigger_after_operation_project AFTER INSERT OR UPDATE OR DELETE ON public.project
FOR EACH ROW EXECUTE PROCEDURE public.after_operation_function();
CREATE OR REPLACE TRIGGER trigger_after_operation_task AFTER INSERT OR UPDATE OR DELETE ON public.task
FOR EACH ROW EXECUTE PROCEDURE public.after_operation_function();



