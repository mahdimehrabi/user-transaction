CREATE TABLE transactions (
                              id SERIAL PRIMARY KEY,
                              user_id INTEGER NOT NULL REFERENCES users(id),
                              amount NUMERIC(10, 2) NOT NULL,
                              type VARCHAR(255) NOT NULL
);
