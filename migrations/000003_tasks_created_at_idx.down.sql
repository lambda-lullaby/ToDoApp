CREATE INDEX tasks_author_user_id_idx ON todoapp.tasks (author_user_id);
DROP INDEX todoapp.tasks_author_user_id_created_at_idx;
DROP INDEX todoapp.tasks_created_at_idx;
