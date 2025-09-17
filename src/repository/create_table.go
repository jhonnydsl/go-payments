package repository

type TableRepository struct{}

func (r *TableRepository) CreateTablePayments() error {
	query := `
	CREATE TABLE IF NOT EXISTS payments (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL,
	amount NUMERIC NOT NULL,
	currency VARCHAR(3) NOT NULL,
	payment_method_id INT REFERENCES payment_methods(id),
	status VARCHAR(20) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(query)
	return err
}

func (r *TableRepository) CreateTableUsers() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	email VARCHAR(100) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(query)
	return err
}

func (r *TableRepository) CreateTablePaymentEvents() error {
	query := `
	CREATE TABLE IF NOT EXISTS payment_events (
	id SERIAL PRIMARY KEY,
	payment_id INT NOT NULL REFERENCES payments(id),
	event_type VARCHAR(50) NOT NULL,
	payload JSONB,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(query)
	return err
}

func (r *TableRepository) CreateTablePaymentMethods() error {
	query := `
	CREATE TABLE IF NOT EXISTS payment_methods (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	description TEXT
	);`

	_, err := DB.Exec(query)
	return err
}