CREATE TABLE IF NOT EXISTS students (
    student_id INT NOT NULL,
    name VARCHAR(120),
    balance INT DEFAULT 0,
    PRIMARY KEY (student_id)
);
