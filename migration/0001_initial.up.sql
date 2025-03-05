-- Tabel "user" untuk menyimpan informasi pengguna
CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel appoinment untuk menyimpan informasi janji temu
CREATE TABLE appoinment (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    appoinment_date TIMESTAMP NOT NULL,
    description TEXT,
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE
);

-- Tabel notification untuk menyimpan informasi notifikasi
CREATE TABLE notification (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    appoinment_id INTEGER,
    message TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE,
    FOREIGN KEY (appoinment_id) REFERENCES appoinment(id) ON DELETE CASCADE
);