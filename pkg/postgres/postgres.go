package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewDB(host, port, user, password, dbname string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitializeTable(db *sql.DB) error {
	q1 := `
		CREATE TABLE IF NOT EXISTS users_octo (
  		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
  		name VARCHAR(255) NOT NULL,
  		phone_number VARCHAR(255) UNIQUE NOT NULL,
  		created_at TIMESTAMP DEFAULT NOW(),
  		updated_at TIMESTAMP DEFAULT NOW()
	);`

	q2 := `
		CREATE TABLE IF NOT EXISTS accounts (
  		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
  		name VARCHAR(255) NOT NULL,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		octopay_id VARCHAR(255),
  		created_at TIMESTAMP DEFAULT NOW(),
  		updated_at TIMESTAMP DEFAULT NOW()
	);`

	q3 := `
		CREATE TABLE IF NOT EXISTS users (
  		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
  		name VARCHAR(255) NOT NULL,
  		phone_number VARCHAR(255) UNIQUE NOT NULL,
		has_octopay BOOLEAN NOT NULL,
  		created_at TIMESTAMP DEFAULT NOW(),
  		updated_at TIMESTAMP DEFAULT NOW()
	);`

	q4 := `
		CREATE TABLE IF NOT EXISTS payments (
  		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
  		payment_method VARCHAR(255) NOT NULL CHECK (payment_method IN ('octopay', 'qris')) DEFAULT 'qris',
		payment_value TEXT NOT NULL,
		payment_status VARCHAR(255) NOT NULL CHECK (payment_status IN ('paid', 'unpaid')) DEFAULT 'unpaid',
		expire_time TIMESTAMP DEFAULT (NOW() + interval '1 hours'),
  		created_at TIMESTAMP DEFAULT NOW(),
  		updated_at TIMESTAMP DEFAULT NOW(),
		receipt_id UUID REFERENCES receipts(id) ON DELETE CASCADE,
		user_id UUID REFERENCES users(id) ON DELETE CASCADE
	);`

	q5 := `
		CREATE TABLE IF NOT EXISTS receipts (
  		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
		name VARCHAR(255) NOT NULL,
  		sum FLOAT NOT NULL,
		discount FLOAT NOT NULL,
		total FLOAT NOT NULL,
		category VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
  		updated_at TIMESTAMP DEFAULT NOW(),
		creator_id UUID REFERENCES users(id) ON DELETE CASCADE
	);`

	q6 := `
		CREATE TABLE IF NOT EXISTS items (
  		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
  		product VARCHAR(255) NOT NULL,
		price FLOAT NOT NULL,
		quantity INTEGER NOT NULL,
		discount FLOAT NOT NULL,
		tax FLOAT NOT NULL,
		service FLOAT NOT NULL,
		total FLOAT NOT NULL,
  		created_at TIMESTAMP DEFAULT NOW(),
  		updated_at TIMESTAMP DEFAULT NOW(),
		receipt_id UUID REFERENCES receipts(id) ON DELETE CASCADE,
		user_id UUID REFERENCES users(id) ON DELETE CASCADE
	);`

	// q7 := `
	// 	INSERT INTO users (id, name, phone_number, has_octopay) VALUES 
	// 	('f5063dca-556c-4723-931b-cbade7ca139a', 'Irfan', '081111111111', true),
	// 	('3d91b359-b840-470b-a069-a4627900ddcf', 'Adit', '081290882428', false);
	// `

	// q8 := `
	// 	INSERT INTO receipts (id, name, sum, discount, total, category, creator_id) VALUES 
	// 	('ad9197e5-cb9a-419c-b055-596962d5501e', 'Test Toko', 38850, 0, 38850, 'Makanan', 'f5063dca-556c-4723-931b-cbade7ca139a');
	// `

	// q9 := `
	// 	INSERT INTO payments (payment_method, payment_value, receipt_id, user_id) VALUES 
	// 	('octopay', '081111111111', 'ad9197e5-cb9a-419c-b055-596962d5501e', 'f5063dca-556c-4723-931b-cbade7ca139a'),
	// 	('qris', '00020101021126680016ID.CO.TELKOM.WWW011893600898026724611002150001952559966270303UMI51440014ID.CO.QRIS.WWW02150000002672461100303UMI5204573253033605502015802ID5916InterActive Corp6015Permata Safira 610560136630403A2', 'ad9197e5-cb9a-419c-b055-596962d5501e', '3d91b359-b840-470b-a069-a4627900ddcf');
	// `

	// q10 := `
	// 	INSERT INTO items (product, price, quantity, discount, tax, service, total, receipt_id, user_id) VALUES 
	// 	('Nasi Goreng', '20000', '1', '0', '11', '0', '22200', 'ad9197e5-cb9a-419c-b055-596962d5501e', 'f5063dca-556c-4723-931b-cbade7ca139a'),
	// 	('Bakso', '15000', '1', '0', '11', '0', '16650', 'ad9197e5-cb9a-419c-b055-596962d5501e', '3d91b359-b840-470b-a069-a4627900ddcf');
	// `
	
	if _, err := db.Exec(q1); err != nil {
		log.Printf("initialize table users_octo: %s\n", err.Error())
		return err
	}

	if _, err := db.Exec(q2); err != nil {
		log.Printf("initialize table accounts: %s\n", err.Error())
		return err
	}

	if _, err := db.Exec(q3); err != nil {
		log.Printf("initialize table users: %s\n", err.Error())
		return err
	}

	if _, err := db.Exec(q5); err != nil {
		log.Printf("initialize table receipts: %s\n", err.Error())
		return err
	}

	if _, err := db.Exec(q4); err != nil {
		log.Printf("initialize table payments: %s\n", err.Error())
		return err
	}

	if _, err := db.Exec(q6); err != nil {
		log.Printf("initialize table items: %s\n", err.Error())
		return err
	}

	// if _, err := db.Exec(q7); err != nil {
	// 	log.Printf("initialize table items: %s\n", err.Error())
	// 	return err
	// }

	// if _, err := db.Exec(q8); err != nil {
	// 	log.Printf("initialize table items: %s\n", err.Error())
	// 	return err
	// }

	// if _, err := db.Exec(q9); err != nil {
	// 	log.Printf("initialize table items: %s\n", err.Error())
	// 	return err
	// }

	// if _, err := db.Exec(q10); err != nil {
	// 	log.Printf("initialize table items: %s\n", err.Error())
	// 	return err
	// }

	return nil
}
