-- Untuk mengkategorikan tipe pengeluaran. Apakah user mengeluarkan uang ini untuk makanan, transportasi, dll?
-- 12-08-2026

CREATE TABLE
    categories (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        user_id BIGINT UNSIGNED NOT NULL,
        name VARCHAR(100) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        CONSTRAINT fk_categories_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
        CONSTRAINT uq_user_category_name UNIQUE (user_id, name)
    );