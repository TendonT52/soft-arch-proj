CREATE INDEX open_positions_search_idx ON open_positions USING GIN (to_tsvector('english', open_positions.title));
CREATE INDEX required_skills_search_idx ON required_skills USING GIN (to_tsvector('english', required_skills.title));
CREATE INDEX benefits_search_idx ON benefits USING GIN (to_tsvector('english', benefits.title));