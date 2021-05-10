USE account;
INSERT INTO user (name, email, type, student_number, password)
VALUES (
        '山田太郎',
        's1121141@exsample.com',
        'student',
        '1121141',
        'sample'
    ),
    (
        '田中太郎',
        'tanaka@exsample.com',
        'teacher',
        NULL,
        'sample'
    ),
    (
        '鈴木太郎',
        'suzuki@exsample.com',
        'general',
        NULL,
        'sample'
    );
INSERT INTO admin (name, password)
VALUES ('admin1', 'admin1'),
    ('admin2', 'admin2');
