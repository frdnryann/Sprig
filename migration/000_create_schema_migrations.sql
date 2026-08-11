-- Untuk validasi, apakah ada tabel "yang sama persis" ketika melakukan migrate?, jika iya, maka bash akan melewatinya
-- 12-08-2026

CREATE TABLE
    IF NOT EXISTS schema_migrations (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        migration VARCHAR(255) NOT NULL UNIQUE,
        migrated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );