-- Índices para mejorar búsquedas frecuentes

CREATE INDEX IF NOT EXISTS idx_users_email
ON users(email);

CREATE INDEX IF NOT EXISTS idx_audiovisual_age_rating
ON audiovisual_content(age_rating);

CREATE INDEX IF NOT EXISTS idx_audio_age_rating
ON audio_content(age_rating);

CREATE INDEX IF NOT EXISTS idx_user_ratings_content
ON user_ratings(content_id, content_type);

CREATE INDEX IF NOT EXISTS idx_playback_user
ON playback_history(user_id);

CREATE INDEX IF NOT EXISTS idx_favorites_user
ON favorites(user_id);
