DO $$ 
DECLARE 
    post_id_1 INT;
    post_id_2 INT;
    post_id_3 INT;

    open_position_id_1 INT;
    open_position_id_2 INT;
    open_position_id_3 INT;

    required_skill_id_1 INT;
    required_skill_id_2 INT;

    benefit_id_1 INT;
    benefit_id_2 INT;
BEGIN
    INSERT INTO posts (uid, topic, description, period, how_to, created_at, updated_at) VALUES 
    (1, 'Internship at Agoda', '{ "root": { "children": ["child"] } }', '1 Aug - 31 Aug', '{ "root": { "children": ["child"] } }', 1695624208, 1695624208) RETURNING pid INTO post_id_1;
    INSERT INTO posts (uid, topic, description, period, how_to, created_at, updated_at) VALUES 
    (1, 'Internship at Google', '{ "root": { "children": ["child"] } }', '2 Aug - 30 Aug', '{ "root": { "children": ["child"] } }', 1695624208, 1695624208) RETURNING pid INTO post_id_2;
    INSERT INTO posts (uid, topic, description, period, how_to, created_at, updated_at) VALUES 
    (1, 'Internship at Dime', '{ "root": { "children": ["child"] } }', '3 Aug - 29 Aug', '{ "root": { "children": ["child"] } }', 1695624208, 1695624208) RETURNING pid INTO post_id_3;

    INSERT INTO open_positions (title, created_at, updated_at) VALUES
    ('Software Engineer', 1695624208, 1695624208) RETURNING oid INTO open_position_id_1;
    INSERT INTO open_positions (title, created_at, updated_at) VALUES
    ('Data Scientist', 1695624208, 1695624208) RETURNING oid INTO open_position_id_2;
    INSERT INTO open_positions (title, created_at, updated_at) VALUES
    ('Product Manager', 1695624208, 1695624208) RETURNING oid INTO open_position_id_3;


    INSERT INTO required_skills (title, created_at, updated_at) VALUES
    ('Python', 1695624208, 1695624208) RETURNING sid INTO required_skill_id_1;
    INSERT INTO required_skills (title, created_at, updated_at) VALUES
    ('Java', 1695624208, 1695624208) RETURNING sid INTO required_skill_id_2;

    INSERT INTO benefits (title, created_at, updated_at) VALUES
    ('Free Lunch', 1695624208, 1695624208) RETURNING bid INTO benefit_id_1;
    INSERT INTO benefits (title, created_at, updated_at) VALUES
    ('Free Snack', 1695624208, 1695624208) RETURNING bid INTO benefit_id_2;

    INSERT INTO posts_open_positions (pid, oid, created_at, updated_at) VALUES
    (post_id_1, open_position_id_1, 1695624208, 1695624208);
    INSERT INTO posts_open_positions (pid, oid, created_at, updated_at) VALUES
    (post_id_2, open_position_id_2, 1695624208, 1695624208);
    INSERT INTO posts_open_positions (pid, oid, created_at, updated_at) VALUES
    (post_id_3, open_position_id_3, 1695624208, 1695624208);

    INSERT INTO posts_required_skills (pid, sid, created_at, updated_at) VALUES
    (post_id_1, required_skill_id_1, 1695624208, 1695624208);
    INSERT INTO posts_required_skills (pid, sid, created_at, updated_at) VALUES
    (post_id_2, required_skill_id_2, 1695624208, 1695624208);
    INSERT INTO posts_required_skills (pid, sid, created_at, updated_at) VALUES
    (post_id_3, required_skill_id_2, 1695624208, 1695624208);

    INSERT INTO posts_benefits (pid, bid, created_at, updated_at) VALUES
    (post_id_1, benefit_id_1, 1695624208, 1695624208);
    INSERT INTO posts_benefits (pid, bid, created_at, updated_at) VALUES
    (post_id_2, benefit_id_2, 1695624208, 1695624208);
    INSERT INTO posts_benefits (pid, bid, created_at, updated_at) VALUES
    (post_id_3, benefit_id_1, 1695624208, 1695624208);
END $$;
