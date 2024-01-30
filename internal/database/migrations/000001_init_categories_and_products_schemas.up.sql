CREATE TABLE IF NOT EXISTS categories (
    id SERIAL NOT NULL,
    name TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL NOT NULL,
    name TEXT UNIQUE NOT NULL,
	description TEXT,
	price DECIMAL(38, 18) NOT NULL,
	category_id INTEGER NOT NULL,
	image_url TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    PRIMARY KEY (id),
    FOREIGN KEY (category_id) REFERENCES categories (id)
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON categories
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON products
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
