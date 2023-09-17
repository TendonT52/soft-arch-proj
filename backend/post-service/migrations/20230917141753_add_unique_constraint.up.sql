ALTER TABLE open_positions ADD CONSTRAINT unique_open_positions_title UNIQUE (title);
ALTER TABLE required_skills ADD CONSTRAINT unique_required_skills_title UNIQUE (title);
ALTER TABLE benefits ADD CONSTRAINT unique_benefits_title UNIQUE (title);