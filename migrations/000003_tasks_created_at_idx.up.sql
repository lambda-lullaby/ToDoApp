CREATE INDEX tasks_created_at_idx ON todoapp.tasks (created_at);
CREATE INDEX tasks_author_user_id_created_at_idx ON todoapp.tasks (author_user_id, created_at);
DROP INDEX todoapp.tasks_author_user_id_idx;
