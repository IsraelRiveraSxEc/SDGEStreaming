-- Crea tablas relacionadas con suscripciones y pagos
-- Depende de: users, plans

CREATE TABLE IF NOT EXISTS subscriptions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    plan_id INTEGER NOT NULL,
    start_date DATETIME NOT NULL,
    end_date DATETIME,
    is_active BOOLEAN NOT NULL DEFAULT 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (plan_id) REFERENCES plans(id)
);

CREATE INDEX IF NOT EXISTS idx_subscriptions_user
ON subscriptions(user_id);

CREATE INDEX IF NOT EXISTS idx_subscriptions_active
ON subscriptions(is_active);

-- Tabla de métodos de pago (ya existe en 001, aquí solo reforzamos índices)
CREATE INDEX IF NOT EXISTS idx_payment_methods_user
ON payment_methods(user_id);

CREATE INDEX IF NOT EXISTS idx_payment_methods_default
ON payment_methods(user_id, is_default);