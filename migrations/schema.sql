CREATE TABLE todos (
    id CHAR(36) PRIMARY KEY,
    task TEXT NOT NULL,
    due_date DATETIME NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_due_date (due_date),
    INDEX idx_completed (completed)
);