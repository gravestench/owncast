package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/owncast/owncast/models"
)

func (s *Service) createAccessTokenTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS user_access_tokens (
    "token" TEXT NOT NULL PRIMARY KEY,
    "user_id" TEXT NOT NULL,
    "timestamp" DATE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
  );`

	stmt, err := s.db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Warnln(err)
	}
}

func (s *Service) createWebhooksTable() {
	log.Traceln("Creating webhooks table...")

	createTableSQL := `CREATE TABLE IF NOT EXISTS webhooks (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"url" string NOT NULL,
		"events" TEXT NOT NULL,
		"timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP,
		"last_used" DATETIME
	);`

	stmt, err := s.db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	if _, err = stmt.Exec(); err != nil {
		log.Warnln(err)
	}
}

// InsertWebhook will add a new webhook to the database.
func (s *Service) InsertWebhook(url string, events []models.EventType) (int, error) {
	log.Traceln("Adding new webhook")

	eventsString := strings.Join(events, ",")

	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}
	stmt, err := tx.Prepare("INSERT INTO webhooks(url, events) values(?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(url, eventsString)
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	newID, err := insertResult.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(newID), err
}

// DeleteWebhook will delete a webhook from the database.
func (s *Service) DeleteWebhook(id int) error {
	log.Traceln("Deleting webhook")

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("DELETE FROM webhooks WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	if rowsDeleted, _ := result.RowsAffected(); rowsDeleted == 0 {
		_ = tx.Rollback()
		return errors.New(fmt.Sprint(id) + " not found")
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

// GetWebhooksForEvent will return all of the webhooks that want to be notified about an event type.
func (s *Service) GetWebhooksForEvent(event models.EventType) []models.Webhook {
	webhooks := make([]models.Webhook, 0)

	query := `SELECT * FROM (
		WITH RECURSIVE split(id, url, event, rest) AS (
		  SELECT id, url, '', events || ',' FROM webhooks
		   UNION ALL
		  SELECT id, url,
				 substr(rest, 0, instr(rest, ',')),
				 substr(rest, instr(rest, ',')+1)
			FROM split
		   WHERE rest <> '')
		SELECT id, url, event
		  FROM split
		 WHERE event <> ''
	  ) AS webhook WHERE event IS "` + event + `"`

	rows, err := s.db.Query(query)
	if err != nil || rows.Err() != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var url string

		if err := rows.Scan(&id, &url, &event); err != nil {
			log.Debugln(err)
			log.Error("There is a problem with the database.")
			break
		}

		singleWebhook := models.Webhook{
			ID:  id,
			URL: url,
		}

		webhooks = append(webhooks, singleWebhook)
	}
	return webhooks
}

// GetWebhooks will return all the webhooks.
func (s *Service) GetWebhooks() ([]models.Webhook, error) { //nolint
	webhooks := make([]models.Webhook, 0)

	query := "SELECT * FROM webhooks"

	rows, err := s.db.Query(query)
	if err != nil {
		return webhooks, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var url string
		var events string
		var timestampString string
		var lastUsedString *string

		if err := rows.Scan(&id, &url, &events, &timestampString, &lastUsedString); err != nil {
			log.Error("There is a problem reading the database.", err)
			return webhooks, err
		}

		timestamp, err := time.Parse(time.RFC3339, timestampString)
		if err != nil {
			return webhooks, err
		}

		var lastUsed *time.Time
		if lastUsedString != nil {
			lastUsedTime, _ := time.Parse(time.RFC3339, *lastUsedString)
			lastUsed = &lastUsedTime
		}

		singleWebhook := models.Webhook{
			ID:        id,
			URL:       url,
			Events:    strings.Split(events, ","),
			Timestamp: timestamp,
			LastUsed:  lastUsed,
		}

		webhooks = append(webhooks, singleWebhook)
	}

	if err := rows.Err(); err != nil {
		return webhooks, err
	}

	return webhooks, nil
}

// SetWebhookAsUsed will update the last used time for a webhook.
func (s *Service) SetWebhookAsUsed(webhook models.Webhook) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("UPDATE webhooks SET last_used = CURRENT_TIMESTAMP WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(webhook.ID); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *Service) createUsersTable() {
	log.Traceln("Creating users table...")

	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"id" TEXT,
		"display_name" TEXT NOT NULL,
		"display_color" NUMBER NOT NULL,
		"created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		"disabled_at" TIMESTAMP,
		"previous_names" TEXT DEFAULT '',
		"namechanged_at" TIMESTAMP,
    "authenticated_at" TIMESTAMP,
		"scopes" TEXT,
		"type" TEXT DEFAULT 'STANDARD',
		"last_used" DATETIME DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	);`

	s.MustExec(createTableSQL)
	s.MustExec(`CREATE INDEX IF NOT EXISTS idx_user_id ON users (id);`)
	s.MustExec(`CREATE INDEX IF NOT EXISTS idx_user_id_disabled ON users (id, disabled_at);`)
	s.MustExec(`CREATE INDEX IF NOT EXISTS idx_user_disabled_at ON users (disabled_at);`)
}

// CreateMessagesTable will create the chat messages table if needed.
func (s *Service) CreateMessagesTable() {
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
	s.MustExec(createTableSQL)

	// Create indexes
	s.MustExec(`CREATE INDEX IF NOT EXISTS user_id_hidden_at_timestamp ON messages (id, user_id, hidden_at, timestamp);`)
	s.MustExec(`CREATE INDEX IF NOT EXISTS idx_id ON messages (id);`)
	s.MustExec(`CREATE INDEX IF NOT EXISTS idx_user_id ON messages (user_id);`)
	s.MustExec(`CREATE INDEX IF NOT EXISTS idx_hidden_at ON messages (hidden_at);`)
	s.MustExec(`CREATE INDEX IF NOT EXISTS idx_timestamp ON messages (timestamp);`)
	s.MustExec(`CREATE INDEX IF NOT EXISTS idx_messages_hidden_at_timestamp on messages(hidden_at, timestamp);`)
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
		s.Logger().Fatal("error creating ip ban table", err)
	}
	defer stmt.Close()
	if _, err := stmt.Exec(); err != nil {
		s.Logger().Fatal("error creating ip ban table", err)
	}
}
