-- Seed Students --
DO $$ 
DECLARE 
    user_id_1 INT;
    user_id_2 INT;
    user_id_3 INT;
    user_id_4 INT;
BEGIN
    INSERT INTO users (email, password, role, verified, created_at, updated_at) VALUES 
    ('6330203521@student.chula.ac.th', 'password', 'student', true, 1609459200, 1609459200) RETURNING id INTO user_id_1;
    INSERT INTO students (sid, faculty, major, year, name, description, created_at, updated_at) VALUES 
    (user_id_1, 'Engineering', 'Computer Engineering', 4, 'Tikhamporn Tepsut', 'I am a student at Chulalongkorn University', 1609459200, 1609459200);

    INSERT INTO users (email, password, role, verified, created_at, updated_at) VALUES 
    ('6330203621@student.chula.ac.th', 'password', 'student', true, 1609459200, 1609459200) RETURNING id INTO user_id_2;
    INSERT INTO students (sid, faculty, major, year, name, description, created_at, updated_at) VALUES 
    (user_id_2, 'Engineering', 'Computer Engineering', 4, 'Sarah Teppa', 'I am a ENGINEER student', 1609459200, 1609459200);

    INSERT INTO users (email, password, role, verified, created_at, updated_at) VALUES 
    ('6330203721@student.chula.ac.th', 'password', 'student', false, 1609459200, 1609459200) RETURNING id INTO user_id_3;
    INSERT INTO students (sid, faculty, major, year, name, description, created_at, updated_at) VALUES 
    (user_id_3, 'Engineering', 'Computer Engineering', 4, 'Unverified One', 'Not Verified Student 1', 1609459200, 1609459200);

    INSERT INTO users (email, password, role, verified, created_at, updated_at) VALUES 
    ('6330203821@student.chula.ac.th', 'password', 'student', false, 1609459200, 1609459200) RETURNING id INTO user_id_4;
    INSERT INTO students (sid, faculty, major, year, name, description, created_at, updated_at) VALUES 
    (user_id_4, 'Engineering', 'Computer Engineering', 4, 'Unverified Two', 'Not Verified Student 2', 1609459200, 1609459200);
END $$;


-- Seed Companies --
DO $$ 
DECLARE 
    user_id_5 INT;
    user_id_6 INT;
    user_id_7 INT;
    user_id_8 INT;
    user_id_9 INT;
    user_id_10 INT;
BEGIN
    INSERT INTO users (email, password, role, verified, created_at, updated_at) VALUES 
    ('company1@gmail.com', 'password', 'company', false, 1609459200, 1609459200) RETURNING id INTO user_id_5;
    INSERT INTO companies (cid, name, description, location, phone, category, status, created_at, updated_at) VALUES 
    (user_id_5, 'Google Bank', 'I am a company', 'Bangkok', '0851231122', 'Global Tech', 'Pending', 1609459200, 1609459200);

    INSERT INTO users (email, password, role, verified, created_at, updated_at) VALUES 
    ('company2@gmail.com', 'password', 'company', true, 1609459200, 1609459200) RETURNING id INTO user_id_6;
    INSERT INTO companies (cid, name, description, location, phone, category, status, created_at, updated_at) VALUES 
    (user_id_6, 'Facebook Technology', 'Technical company', 'Bangkok', '0851231122', 'Social Media', 'Approve', 1609459200, 1609459200);

    INSERT INTO users (email, password, role, verified, created_at, updated_at) VALUES 
    ('company3@gmail.com', 'password', 'company', false, 1609459200, 1609459200) RETURNING id INTO user_id_7;
    INSERT INTO companies (cid, name, description, location, phone, category, status, created_at, updated_at) VALUES 
    (user_id_7, 'Apple', 'I am a company', 'Bangkok', '0851231122', 'IoT and Mobile', 'Reject', 1609459200, 1609459200);

    INSERT INTO users (email, password, role, verified, created_at, updated_at) VALUES 
    ('company4@gmail.com', 'password', 'company', false, 1609459200, 1609459200) RETURNING id INTO user_id_8;
    INSERT INTO companies (cid, name, description, location, phone, category, status, created_at, updated_at) VALUES 
    (user_id_8, 'Dime! Trending', 'Financial', 'Bangkok', '0851231122', 'Banking and Trending', 'Pending', 1609459200, 1609459200);

    INSERT INTO users (email, password, role, verified, created_at, updated_at) VALUES 
    ('company5@gmail.com', 'password', 'company', true, 1609459200, 1609459200) RETURNING id INTO user_id_9;
    INSERT INTO companies (cid, name, description, location, phone, category, status, created_at, updated_at) VALUES 
    (user_id_9, 'Agoda', 'Technical company', 'Bangkok', '0851231122', 'Travelling', 'Approve', 1609459200, 1609459200);

    INSERT INTO users (email, password, role, verified, created_at, updated_at) VALUES 
    ('company6@gmail.com', 'password', 'company', false, 1609459200, 1609459200) RETURNING id INTO user_id_10;
    INSERT INTO companies (cid, name, description, location, phone, category, status, created_at, updated_at) VALUES 
    (user_id_10, 'Rotten Tomatoes', 'I am a company', 'Bangkok', '0851231122', 'Entertainment Tech', 'Reject', 1609459200, 1609459200);
END $$;

-- Seed Admin --
INSERT INTO users (email, password, role, verified, created_at, updated_at) VALUES 
('admin1@admin.com', 'password', 'admin', true, 1609459200, 1609459200);

INSERT INTO users (email, password, role, verified, created_at, updated_at) VALUES 
('admin2@admin.com', 'password', 'admin', true, 1609459200, 1609459200);