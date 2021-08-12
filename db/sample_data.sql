USE account;
INSERT INTO users (user_id, account_type, student_number, password)
VALUES (
        '1121141',
        'student',
        '1121141',
        'sample'
    ),
    (
        'tanaka',
        'teacher',
        NULL,
        'sample'
    ),
    (
        'suzuki',
        'general',
        NULL,
        'sample'
    );
INSERT INTO admins (name, password)
VALUES ('admin1', 'admin1'),
    ('admin2', 'admin2');
