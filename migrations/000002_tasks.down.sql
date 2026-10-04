DROP TABLE todoapp.tasks;

CREATE TABLE todoapp.tasks (
    id             UUID PRIMARY KEY,
    title          VARCHAR(100) NOT NULL,
    completed      BOOLEAN NOT NULL,
    created_at     TIMESTAMPTZ NOT NULL,
    completed_at   TIMESTAMPTZ,
    author_user_id UUID NOT NULL REFERENCES todoapp.users(id),
    CHECK (completed = (completed_at IS NOT NULL))
);
