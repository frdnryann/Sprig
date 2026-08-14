-- Mencatat pengeluaran
-- 12-08-2026

CREATE TABLE
    expenses (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        user_id BIGINT UNSIGNED NOT NULL,
        category_id BIGINT UNSIGNED NOT NULL,
        amount BIGINT UNSIGNED NOT NULL,
        description VARCHAR(255),
        expense_date DATE NOT NULL,
        type ENUM('pengeluaran', 'pemasukan') NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        CONSTRAINT fk_expenses_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
        CONSTRAINT fk_expenses_category FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE RESTRICT,
        INDEX idx_expenses_user_date (user_id, expense_date),
        INDEX idx_expenses_category (category_id)
    );