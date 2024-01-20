ALTER TABLE users
ALTER COLUMN created_at TYPE TIMESTAMPTZ USING created_at::timestamp with time zone;

ALTER TABLE users
ALTER COLUMN id TYPE UUID USING id::uuid;

ALTER TABLE user_securities
ALTER COLUMN user_id TYPE UUID USING user_id::uuid;