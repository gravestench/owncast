package database

// GetUsersCount will return the number of users in the database.
func (s *Service) GetUsersCount() int64 {
	const query = `SELECT COUNT(*) FROM users`

	rows, err := s.db.Query(query)
	if err != nil || rows.Err() != nil {
		return 0
	}
	defer rows.Close()

	var count int64

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0
		}
	}

	return count
}
