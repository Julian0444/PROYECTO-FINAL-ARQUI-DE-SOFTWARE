
CREATE TABLE categories (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    category_name LONGTEXT NULL
);

CREATE INDEX idx_categories_deleted_at ON categories (deleted_at);

CREATE TABLE courses (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    course_name LONGTEXT NULL,
    course_description LONGTEXT NULL,
    course_price DOUBLE NULL,
    course_duration BIGINT NULL,
    course_init_date LONGTEXT NULL,
    course_state TINYINT(1) NULL,
    course_capacity BIGINT NULL,
    course_image LONGTEXT NULL,
    category_id BIGINT UNSIGNED NULL,
    CONSTRAINT fk_courses_category FOREIGN KEY (category_id) REFERENCES categories (id)
);

CREATE INDEX idx_courses_deleted_at ON courses (deleted_at);

CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    password LONGTEXT NULL,
    email LONGTEXT NULL,
    role BIGINT NULL,
    name LONGTEXT NULL,
    avatar LONGTEXT NULL
);

CREATE INDEX idx_users_deleted_at ON users (deleted_at);

CREATE TABLE inscriptos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    course_id BIGINT UNSIGNED NULL,
    user_id BIGINT UNSIGNED NULL,
    CONSTRAINT fk_inscriptos_course FOREIGN KEY (course_id) REFERENCES courses (id),
    CONSTRAINT fk_inscriptos_user FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE INDEX idx_inscriptos_deleted_at ON inscriptos (deleted_at);



INSERT INTO categories (created_at, updated_at, deleted_at, category_name)
VALUES
    (NOW(), NOW(), NULL, 'Programming'),
    (NOW(), NOW(), NULL, 'Mathematics'),
    (NOW(), NOW(), NULL, 'Science'),
    (NOW(), NOW(), NULL, 'Languages'),
    (NOW(), NOW(), NULL, 'Arts'),
    (NOW(), NOW(), NULL, 'Business'),
    (NOW(), NOW(), NULL, 'Design'),
    (NOW(), NOW(), NULL, 'Engineering'),
    (NOW(), NOW(), NULL, 'Healthcare'),
    (NOW(), NOW(), NULL, 'Marketing');

INSERT INTO courses (created_at, updated_at, deleted_at, course_name, course_description, course_price, course_duration, course_init_date, course_state, course_capacity, course_image, category_id)
VALUES
    (NOW(), NOW(), NULL, 'Intro to Python', 'Basic Python programming', 100.0, 30, '2024-11-01', 1, 50, 'python.jpg', 1),
    (NOW(), NOW(), NULL, 'Advanced Mathematics', 'Advanced topics in algebra and calculus', 150.0, 45, '2024-12-01', 1, 30, 'math.jpg', 2),
    (NOW(), NOW(), NULL, 'Physics for Engineers', 'Physics principles applied to engineering', 200.0, 40, '2024-11-15', 1, 40, 'physics.jpg', 3),
    (NOW(), NOW(), NULL, 'Spanish Language', 'Beginner course in Spanish', 80.0, 20, '2024-11-20', 1, 20, 'spanish.jpg', 4),
    (NOW(), NOW(), NULL, 'Creative Painting', 'Introduction to painting techniques', 120.0, 25, '2024-10-30', 1, 15, 'painting.jpg', 5),
    (NOW(), NOW(), NULL, 'Business Management', 'Principles of managing a business', 180.0, 35, '2024-12-01', 1, 25, 'business.jpg', 6),
    (NOW(), NOW(), NULL, 'Graphic Design Basics', 'Learn the basics of graphic design', 110.0, 28, '2024-11-10', 1, 30, 'design.jpg', 7),
    (NOW(), NOW(), NULL, 'Mechanical Engineering', 'Fundamentals of mechanical engineering', 220.0, 50, '2024-12-05', 1, 40, 'mech_eng.jpg', 8),
    (NOW(), NOW(), NULL, 'Introduction to Healthcare', 'Basics of healthcare and medicine', 130.0, 35, '2024-10-25', 1, 20, 'health.jpg', 9),
    (NOW(), NOW(), NULL, 'Digital Marketing', 'Strategies for digital marketing', 140.0, 30, '2024-12-15', 1, 35, 'marketing.jpg', 10);


INSERT INTO users (created_at, updated_at, deleted_at, password, email, role, name, avatar)
VALUES
    (NOW(), NOW(), NULL, 'pass123', 'john.doe@example.com', 1, 'John Doe', 'avatar1.jpg'),
    (NOW(), NOW(), NULL, 'pass234', 'jane.smith@example.com', 2, 'Jane Smith', 'avatar2.jpg'),
    (NOW(), NOW(), NULL, 'pass345', 'mike.jones@example.com', 1, 'Mike Jones', 'avatar3.jpg'),
    (NOW(), NOW(), NULL, 'pass456', 'emily.white@example.com', 2, 'Emily White', 'avatar4.jpg'),
    (NOW(), NOW(), NULL, 'pass567', 'linda.brown@example.com', 1, 'Linda Brown', 'avatar5.jpg'),
    (NOW(), NOW(), NULL, 'pass678', 'paul.johnson@example.com', 2, 'Paul Johnson', 'avatar6.jpg'),
    (NOW(), NOW(), NULL, 'pass789', 'karen.davis@example.com', 1, 'Karen Davis', 'avatar7.jpg'),
    (NOW(), NOW(), NULL, 'pass890', 'tom.williams@example.com', 2, 'Tom Williams', 'avatar8.jpg'),
    (NOW(), NOW(), NULL, 'pass901', 'nancy.taylor@example.com', 1, 'Nancy Taylor', 'avatar9.jpg'),
    (NOW(), NOW(), NULL, 'pass012', 'david.martin@example.com', 2, 'David Martin', 'avatar10.jpg');


INSERT INTO inscriptos (created_at, updated_at, deleted_at, course_id, user_id)
VALUES
    (NOW(), NOW(), NULL, 1, 1),
    (NOW(), NOW(), NULL, 2, 2),
    (NOW(), NOW(), NULL, 3, 3),
    (NOW(), NOW(), NULL, 4, 4),
    (NOW(), NOW(), NULL, 5, 5),
    (NOW(), NOW(), NULL, 6, 6),
    (NOW(), NOW(), NULL, 7, 7),
    (NOW(), NOW(), NULL, 8, 8),
    (NOW(), NOW(), NULL, 9, 9),
    (NOW(), NOW(), NULL, 10, 10);
