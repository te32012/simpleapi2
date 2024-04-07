CREATE TABLE IF NOT EXISTS banners (
    id_banners SERIAL PRIMARY KEY,
    inpute TEXT NOT NULL,
);

CREATE TABLE IF NOT EXISTS bann (
    id_banners INTEGER references banners(id_banners),
    id_fitch INTEGER , 
    id_tag INTEGER,
);