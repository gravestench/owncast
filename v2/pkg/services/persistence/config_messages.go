package persistence

import (
	"context"
	"database/sql"

	"github.com/owncast/owncast/db"
	"github.com/owncast/owncast/models"
)

// GetMessagesCount will return the number of messages in the database.
func (s *Service) GetMessagesCount() int64 {
	query := `SELECT COUNT(*) FROM messages`
	rows, err := s.db.GetDatabase().Query(query)
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

// BanIPAddress will persist a new IP address ban to the datastore.
func (s *Service) BanIPAddress(address, note string) error {
	return s.GetQueries().BanIPAddress(context.Background(), db.BanIPAddressParams{
		IpAddress: address,
		Notes:     sql.NullString{String: note, Valid: true},
	})
}

// IsIPAddressBanned will return if an IP address has been previously blocked.
func (s *Service) IsIPAddressBanned(address string) (bool, error) {
	blocked, error := s.GetQueries().IsIPAddressBlocked(context.Background(), address)
	return blocked > 0, error
}

// GetIPAddressBans will return all the banned IP addresses.
func (s *Service) GetIPAddressBans() ([]models.IPAddress, error) {
	result, err := s.GetQueries().GetIPAddressBans(context.Background())
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
	return s.GetQueries().RemoveIPAddressBan(context.Background(), address)
}
