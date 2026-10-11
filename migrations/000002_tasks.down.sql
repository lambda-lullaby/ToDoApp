DROP INDEX todoapp.tasks_author_user_id_idx;

ALTER TABLE todoapp.tasks
    DROP CONSTRAINT tasks_completed_check,
    DROP CONSTRAINT tasks_title_length_check,
    DROP COLUMN description,
    DROP COLUMN version,
    ALTER COLUMN created_at DROP DEFAULT,
    ALTER COLUMN completed DROP DEFAULT,
    ALTER COLUMN id DROP DEFAULT,
    ADD CONSTRAINT tasks_check CHECK (completed = (completed_at IS NOT NULL));
