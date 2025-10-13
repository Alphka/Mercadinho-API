CREATE TABLE IF NOT EXISTS customers (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	email TEXT,
	phone VARCHAR(13),
	address TEXT,
	complement TEXT,
	number TEXT,
	neighborhood TEXT,
	city TEXT,
	state VARCHAR(2),
	zip_code VARCHAR(8),
	payment_day INT,
	created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
	updated_at TIMESTAMPTZ DEFAULT NULL
);
