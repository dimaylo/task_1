CREATE TABLE Users (
                    id SERIAL PRIMARY KEY,
                    e_mail VARCHAR(255) NOT NULL,
                    password VARCHAR(255) NOT NULL,
                    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
                    deleted_at TIMESTAMP DEFAULT NULL
);
