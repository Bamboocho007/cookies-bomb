ALTER TABLE users
ADD CONSTRAINT UniqueEmail UNIQUE (email);

ALTER TABLE user_securities
ADD CONSTRAINT FK_PersonSequrities FOREIGN KEY (user_id) REFERENCES users(id);