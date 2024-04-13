
CREATE TABLE IF NOT EXISTS banner (
    id_banner SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL
);


-- Таблица, которая хранит информацию о версиях
CREATE TABLE IF NOT EXISTS version_banner (
    id_version SERIAL PRIMARY KEY,
    id_banner INTEGER REFERENCES banner(id_banner),
    content BYTEA,
    updated_at TIMESTAMP NOT NULL,
    is_active BOOLEAN NOT NULL
);

-- Тут храним айдишники банеров и время последнего обновления банера


-- Иногда полезно иметь такие таблицы - если появится какая-то фильтрация завязанная на содержание фич
-- Но мы это не используем из-за того, что работа с лишними таблицами - это время
/*
-- Список фич
CREATE TABLE IF NOT EXISTS features (
    id_future INTEGER NOT NULL
);

-- Список тегов
CREATE TABLE IF NOT EXISTS tags (
    id_tag INTEGER NOT NULL
);
*/

-- Сопоставляем фичи версиям
CREATE TABLE IF NOT EXISTS features_banner (
    id_version INTEGER REFERENCES version_banner(id_version),
    id_future INTEGER,
    id_banner INTEGER REFERENCES banner(id_banner)
);

-- Сопоставляем теги версиям
CREATE TABLE IF NOT EXISTS tags_banner (
    id_version INTEGER REFERENCES version_banner(id_version),
    id_tag INTEGER,
    id_banner INTEGER REFERENCES banner(id_banner)
);
