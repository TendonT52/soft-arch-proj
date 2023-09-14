-- Seed Students --
INSERT INTO users (id, email, password, role, verified, created_at, updated_at) VALUES 
(1, '6330203521@student.chula.ac.th', 'password', 'student', true, 1609459200, 1609459200);
INSERT INTO students (sid, faculty, major, year, name, description, created_at, updated_at) VALUES 
(1, 'Engineering', 'Computer Engineering', 4, 'Tikhamporn Tepsut', 'I am a student at Chulalongkorn University', 1609459200, 1609459200);
INSERT INTO users (id, email, password, role, verified, created_at, updated_at) VALUES 

(2, '6330203621@student.chula.ac.th', 'password', 'student', true, 1609459200, 1609459200);
INSERT INTO students (sid, faculty, major, year, name, description, created_at, updated_at) VALUES 
(2, 'Engineering', 'Computer Engineering', 4, 'Sarah Teppa', 'I am a ENGINEER student', 1609459200, 1609459200);

INSERT INTO users (id, email, password, role, verified, created_at, updated_at) VALUES 
(3, '6330203721@student.chula.ac.th', 'password', 'student', false, 1609459200, 1609459200);
INSERT INTO students (sid, faculty, major, year, name, description, created_at, updated_at) VALUES 
(3, 'Engineering', 'Computer Engineering', 4, 'Unverified One', 'Not Verified Student 1', 1609459200, 1609459200);

INSERT INTO users (id, email, password, role, verified, created_at, updated_at) VALUES 
(4, '6330203821@student.chula.ac.th', 'password', 'student', false, 1609459200, 1609459200);
INSERT INTO students (sid, faculty, major, year, name, description, created_at, updated_at) VALUES 
(4, 'Engineering', 'Computer Engineering', 4, 'Unverified Two', 'Not Verified Student 2', 1609459200, 1609459200);


-- Seed Companies --
INSERT INTO users (id, email, password, role, verified, created_at, updated_at) VALUES 
(5, 'company1@gmail.com', 'password', 'company', false, 1609459200, 1609459200);
INSERT INTO companies (cid, name, description, location, phone, category, status, created_at, updated_at) VALUES 
(5, 'Google Bank', 'I am a company', 'Bangkok', '0851231122', 'Global Tech', 'Pending', 1609459200, 1609459200);

INSERT INTO users (id, email, password, role, verified, created_at, updated_at) VALUES 
(6, 'company2@gmail.com', 'password', 'company', true, 1609459200, 1609459200);
INSERT INTO companies (cid, name, description, location, phone, category, status, created_at, updated_at) VALUES 
(6, 'Facebook Technology', 'Technical company', 'Bangkok', '0851231122', 'Social Media', 'Approve', 1609459200, 1609459200);

INSERT INTO users (id, email, password, role, verified, created_at, updated_at) VALUES 
(7, 'company3@gmail.com', 'password', 'company', false, 1609459200, 1609459200);
INSERT INTO companies (cid, name, description, location, phone, category, status, created_at, updated_at) VALUES 
(7, 'Apple', 'I am a company', 'Bangkok', '0851231122', 'IoT and Mobile', 'Reject', 1609459200, 1609459200);

INSERT INTO users (id, email, password, role, verified, created_at, updated_at) VALUES 
(8, 'company4@gmail.com', 'password', 'company', false, 1609459200, 1609459200);
INSERT INTO companies (cid, name, description, location, phone, category, status, created_at, updated_at) VALUES 
(8, 'Dime! Trending', 'Financial', 'Bangkok', '0851231122', 'Banking and Trending', 'Pending', 1609459200, 1609459200);

INSERT INTO users (id, email, password, role, verified, created_at, updated_at) VALUES 
(9, 'company5@gmail.com', 'password', 'company', true, 1609459200, 1609459200);
INSERT INTO companies (cid, name, description, location, phone, category, status, created_at, updated_at) VALUES 
(9, 'Agoda', 'Technical company', 'Bangkok', '0851231122', 'Travelling', 'Approve', 1609459200, 1609459200);

INSERT INTO users (id, email, password, role, verified, created_at, updated_at) VALUES 
(10, 'company6@gmail.com', 'password', 'company', false, 1609459200, 1609459200);
INSERT INTO companies (cid, name, description, location, phone, category, status, created_at, updated_at) VALUES 
(10, 'Rotten Tomatoes', 'I am a company', 'Bangkok', '0851231122', 'Entertainment Tech', 'Reject', 1609459200, 1609459200);

-- Seed Admin --
INSERT INTO users (id, email, password, role, verified, created_at, updated_at) VALUES 
(11, 'admin1@admin.com', 'password', 'admin', true, 1609459200, 1609459200);

INSERT INTO users (id, email, password, role, verified, created_at, updated_at) VALUES 
(12, 'admin2@admin.com', 'password', 'admin', true, 1609459200, 1609459200);