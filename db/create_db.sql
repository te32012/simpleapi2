
-- Таблица, которая хранит информацию о версиях
CREATE TABLE IF NOT EXISTS version_banner (
    id_version SERIAL PRIMARY KEY,
    id_banner INTEGER REFERENCES(id_banner),
    content BYTEA,
    created_at TIMESTAMP NOT NULL,
    is_active BOOLEAN NOT NULL
);

-- Тут храним айдишники банеров и время последнего обновления банера
CREATE TABLE IF NOT EXISTS banner (
    id_banner SERIAL PRIMARY KEY,
    updated_at TIMESTAMP NOT NULL
);

-- Список фич
CREATE TABLE IF NOT EXISTS features (
    id_future INTEGER NOT NULL
);

-- Список тегов
CREATE TABLE IF NOT EXISTS tags (
    id_tag INTEGER NOT NULL
);

-- Сопоставляем фичи тегам
CREATE TABLE IF NOT EXISTS features_banner (
    id_banner INTEGER REFERENCES(id_banner),
    id_future INTEGER REFERENCES(id_future)
);

-- Сопоставляем теги банерам
CREATE TABLE IF NOT EXISTS tags_banner (
    id_banner INTEGER REFERENCES(id_banner),
    id_tag INTEGER REFERENCES(id_tag)
);
