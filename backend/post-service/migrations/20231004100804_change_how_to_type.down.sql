ALTER TABLE posts
ADD COLUMN new_how_to json;

UPDATE posts
SET new_how_to = json_build_object('key1', how_to);

ALTER TABLE posts
DROP COLUMN how_to;

ALTER TABLE posts
RENAME COLUMN new_how_to TO how_to;
