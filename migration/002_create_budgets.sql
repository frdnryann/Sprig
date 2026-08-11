-- Tabel utama yang paling sering digunakan, dan yang paling membuat sy pusing
-- 12-08-2026

CREATE TABLE
    budgets (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        user_id BIGINT UNSIGNED NOT NULL,
        total_income BIGINT UNSIGNED NOT NULL,
        needs_percentage TINYINT UNSIGNED NOT NULL DEFAULT 50,
        wants_percentage TINYINT UNSIGNED NOT NULL DEFAULT 30,
        savings_percentage TINYINT UNSIGNED NOT NULL DEFAULT 20,
        year SMALLINT UNSIGNED NOT NULL,
        month TINYINT UNSIGNED NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

        -- RULE / Aturan untuk pengecekan dan melindungi database
        CONSTRAINT fk_budgets_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
        CONSTRAINT uq_user_budget_month UNIQUE (user_id, year, month), -- Karna Perbulan, Jadi tidak boleh ada lebih dari budget yang ditambahkan dalam bulan yang sama
        CONSTRAINT chk_budget_month CHECK (month BETWEEN 1 AND 12),
        CONSTRAINT chk_budget_percentage CHECK (
            needs_percentage + wants_percentage + savings_percentage = 100
        )
    );