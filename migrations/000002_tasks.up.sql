DROP TABLE todoapp.tasks;

CREATE TABLE todoapp.tasks (
    id             UUID PRIMARY KEY,
    version        BIGINT NOT NULL DEFAULT 1,
    title          VARCHAR(100) NOT NULL CHECK (char_length(title) BETWEEN 1 AND 100),
    description    VARCHAR(1000) CHECK (char_length(description) BETWEEN 1 AND 1000),
    completed      BOOLEAN NOT NULL,
    created_at     TIMESTAMPTZ NOT NULL,
    completed_at   TIMESTAMPTZ,
    author_user_id UUID NOT NULL REFERENCES todoapp.users(id),

    CHECK (
        (completed=FALSE AND completed_at IS NULL)
        OR
        (completed=TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
    )
);

CREATE INDEX tasks_author_user_id_idx ON todoapp.tasks (author_user_id);
