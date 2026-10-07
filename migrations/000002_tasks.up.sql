ALTER TABLE todoapp.tasks
    ALTER COLUMN id SET DEFAULT gen_random_uuid(),
    ALTER COLUMN completed SET DEFAULT FALSE,
    ALTER COLUMN created_at SET DEFAULT now(),
    ADD COLUMN version BIGINT NOT NULL DEFAULT 1,
    ADD COLUMN description VARCHAR(1000),
    DROP CONSTRAINT tasks_check,
    ADD CONSTRAINT tasks_title_length_check CHECK (char_length(title) BETWEEN 1 AND 100),
    ADD CONSTRAINT tasks_description_length_check CHECK (char_length(description) BETWEEN 1 AND 1000),
    ADD CONSTRAINT tasks_completed_check CHECK (
        (completed=FALSE AND completed_at IS NULL)
        OR
        (completed=TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
    );

CREATE INDEX tasks_author_user_id_idx ON todoapp.tasks (author_user_id);
