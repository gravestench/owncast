package data

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/owncast/owncast/db"
	"github.com/owncast/owncast/models"
)

// CreateMessagesTable will create the chat messages table if needed.
func (s *Service) CreateMessagesTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS messages (
		"id" string NOT NULL,
		"user_id" TEXT,
		"body" TEXT,
		"eventType" TEXT,
		"hidden_at" DATETIME,
		"timestamp" DATETIME,
		"title" TEXT,
		"subtitle" TEXT,
		"image" TEXT,
		"link" TEXT,
		PRIMARY KEY (id)
	);`
	MustExec(createTableSQL, db)

	// Create indexes
	MustExec(`CREATE INDEX IF NOT EXISTS user_id_hidden_at_timestamp ON messages (id, user_id, hidden_at, timestamp);`, db)
	MustExec(`CREATE INDEX IF NOT EXISTS idx_id ON messages (id);`, db)
	MustExec(`CREATE INDEX IF NOT EXISTS idx_user_id ON messages (user_id);`, db)
	MustExec(`CREATE INDEX IF NOT EXISTS idx_hidden_at ON messages (hidden_at);`, db)
	MustExec(`CREATE INDEX IF NOT EXISTS idx_timestamp ON messages (timestamp);`, db)
	MustExec(`CREATE INDEX IF NOT EXISTS idx_messages_hidden_at_timestamp on messages(hidden_at, timestamp);`, db)
}

// GetMessagesCount will return the number of messages in the database.
func (s *Service) GetMessagesCount() int64 {
	query := `SELECT COUNT(*) FROM messages`
	rows, err := s.Store.DB.Query(query)
	if err != nil || rows.Err() != nil {
		return 0
	}
	defer rows.Close()
	var count int64
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0
		}
	}
	return count
}

// CreateBanIPTable will create the IP ban table if needed.
func (s *Service) CreateBanIPTable(db *sql.DB) {
	createTableSQL := `  CREATE TABLE IF NOT EXISTS ip_bans (
    "ip_address" TEXT NOT NULL PRIMARY KEY,
    "notes" TEXT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );`

	stmt, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal("error creating ip ban table", err)
	}
	defer stmt.Close()
	if _, err := stmt.Exec(); err != nil {
		log.Fatal("error creating ip ban table", err)
	}
}

// BanIPAddress will persist a new IP address ban to the Datastore.
func (s *Service) BanIPAddress(address, note string) error {
	return s.Store.GetQueries().BanIPAddress(context.Background(), db.BanIPAddressParams{
		IpAddress: address,
		Notes:     sql.NullString{String: note, Valid: true},
	})
}

// IsIPAddressBanned will return if an IP address has been previously blocked.
func (s *Service) IsIPAddressBanned(address string) (bool, error) {
	blocked, error := s.Store.GetQueries().IsIPAddressBlocked(context.Background(), address)
	return blocked > 0, error
}

// GetIPAddressBans will return all the banned IP addresses.
func (s *Service) GetIPAddressBans() ([]models.IPAddress, error) {
	result, err := s.Store.GetQueries().GetIPAddressBans(context.Background())
	if err != nil {
		return nil, err
	}

	response := []models.IPAddress{}
	for _, ip := range result {
		response = append(response, models.IPAddress{
			IPAddress: ip.IpAddress,
			Notes:     ip.Notes.String,
			CreatedAt: ip.CreatedAt.Time,
		})
	}
	return response, err
}

// RemoveIPAddressBan will remove a previously banned IP address.
func (s *Service) RemoveIPAddressBan(address string) error {
	return s.Store.GetQueries().RemoveIPAddressBan(context.Background(), address)
}
