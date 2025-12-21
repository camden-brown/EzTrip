-- Seed data for users table
INSERT INTO users (id, first_name, last_name, email, created_at) 
VALUES 
    (gen_random_uuid(), 'John', 'Doe', 'john@example.com', CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'Jane', 'Smith', 'jane@example.com', CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'Alice', 'Johnson', 'alice@example.com', CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'Bob', 'Williams', 'bob@example.com', CURRENT_TIMESTAMP)
ON CONFLICT (email) DO NOTHING;
