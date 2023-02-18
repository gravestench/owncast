package notifications

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/owncast/owncast/core/data"
	"github.com/owncast/owncast/db"
)

func (n *Notifier) createNotificationsTable(db *sql.DB) {
	log.Traceln("Creating federation followers table...")

	createTableSQL := `CREATE TABLE IF NOT EXISTS notifications (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"channel" TEXT NOT NULL,
		"destination" TEXT NOT NULL,
		"created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`

	data.MustExec(createTableSQL, db)
	data.MustExec(`CREATE INDEX IF NOT EXISTS idx_channel ON notifications (channel);`, db)
}

// AddNotification saves a new user notification destination.
func (n *Notifier) AddNotification(channel, destination string) error {
	return n.data.Store.GetQueries().AddNotification(context.Background(), db.AddNotificationParams{
		Channel:     channel,
		Destination: destination,
	})
}

// RemoveNotificationForChannel removes a notification destination.
func (n *Notifier) RemoveNotificationForChannel(channel, destination string) error {
	log.Debugln("Removing notification for channel", channel)
	return n.data.Store.GetQueries().RemoveNotificationDestinationForChannel(context.Background(), db.RemoveNotificationDestinationForChannelParams{
		Channel:     channel,
		Destination: destination,
	})
}

// GetNotificationDestinationsForChannel will return a collection of
// destinations to notify for a given channel.
func (n *Notifier) GetNotificationDestinationsForChannel(channel string) ([]string, error) {
	result, err := n.data.Store.GetQueries().GetNotificationDestinationsForChannel(context.Background(), channel)
	if err != nil {
		return nil, errors.Wrap(err, "unable to query notification destinations for channel "+channel)
	}

	return result, nil
}
