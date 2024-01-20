CREATE TABLE users (
  id VARCHAR(50) PRIMARY KEY,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  created_at VARCHAR(50) NOT NULL,
  email VARCHAR(100) NOT NULL
);

CREATE TABLE user_securities (
  is_confirmed BIT NOT NULL,
  password_hash VARCHAR(100) NOT NULL,
  user_id VARCHAR(50) NOT NULL
);