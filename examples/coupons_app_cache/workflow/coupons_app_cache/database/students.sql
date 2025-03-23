CREATE TABLE IF NOT EXISTS students (
    student_id SERIAL PRIMARY KEY,
    name VARCHAR(120),
    balance INT DEFAULT 0 CHECK (balance > 0),
    claimed_coupons INT DEFAULT 0 CHECK (claimed_coupons < 10)
);
