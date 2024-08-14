CREATE TYPE roles AS ENUM ('admin','patient', 'doctor');
CREATE TYPE genders AS ENUM ('male', 'female');

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    date_of_birth DATE NOT NULL,
    gender genders NOT NULL,
    role roles DEFAULT 'patient',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at Bigint DEFAULT 0
);

CREATE TABLE IF NOT EXISTS notifications(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    messages TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE
);
