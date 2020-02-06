CREATE TABLE IF NOT EXISTS access_tokens (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);
