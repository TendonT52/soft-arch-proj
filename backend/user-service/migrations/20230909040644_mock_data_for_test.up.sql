INSERT INTO "public"."users" ("id", "email", "password", "role", "verified", "created_at", "updated_at") VALUES
(105, 'not_verified@gmail.com', 'fake-password-a', 'student', 'false', '2023-08-31 14:35:05.4137', '2023-08-31 14:35:05.4137'),
(106, 'password_not_match@gmail.com', '$2a$10$yGkA68nAWsGgP8F0kVtPhekIGsQp.Q0Mgi1vU3ezQGKLHKqW0diBe', 'student', 'true', '2023-08-31 14:35:05.4137', '2023-08-31 14:35:05.4137');

INSERT INTO "public"."students" ("sid", "verification_code", "faculty", "major", "year", "created_at", "updated_at", "name", "description") VALUES
(105, 'Vmp5QlB0cnJPcWhab25iblNFMG0=', 'Engineering', 'Computer Engineering4', 3, '2023-09-01 04:26:33.710955', '2023-09-01 04:26:33.710955', 'Not Verified', 'Mock student'),
(106, 'Vmp5QlB0cnJPcWhab25iblNFMG0=', 'Engineering', 'Computer Engineering4', 3, '2023-09-01 04:26:33.710955', '2023-09-01 04:26:33.710955', 'Not Match', 'Mock student');