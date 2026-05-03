CREATE TABLE IF NOT EXISTS students (
    id BIGSERIAL PRIMARY KEY,
    full_name TEXT NOT NULL,
    study_group TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE
);

INSERT INTO students (full_name, study_group, email)
VALUES
    ('Борисов Денис Александрович', 'ЭФМО-01-25', 'borisov@example.com'),
    ('Петрова Мария Сергеевна', 'ИВБО-02-25', 'petrova@example.com'),
    ('Сидоров Алексей Андреевич', 'ИВБО-03-25', 'sidorov@example.com')
ON CONFLICT (email) DO NOTHING;
