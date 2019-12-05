package mysql

import (
	"database/sql"

	"github.com/mono83/xray"
	"github.com/mono83/xray/args"
	"github.com/mono83/xray/text"
)

const loggerName = "mysql-log-writer"

// NewWriterToXRayLog constructs new writer (experimental), that stores data into MySQL table
// created in following format:
//
// CREATE TABLE `xrayLog` (
// `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
// `time` int(10) unsigned NOT NULL,
// `rayId` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
// `level` enum('error','alert','crititcal') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'error',
// `pattern` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
// `message` varchar(1000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
// PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
func NewWriterToXRayLog(db *sql.DB) func(...xray.Event) {
	if db == nil {
		panic("database connection for error logging not provided")
	}

	return func(events ...xray.Event) {
		xray.Waiter.Add(1)
		defer xray.Waiter.Done()

		// Filtering
		var filtered []xray.LogEvent
		for _, e := range events {
			if l, ok := e.(xray.LogEvent); ok && l.GetLogger() != loggerName && l.GetLevel() > xray.WARNING {
				filtered = append(filtered, l)
			}
		}

		if len(filtered) == 0 {
			// Do nothing
			return
		}

		// Starting SQL transaction
		tx, err := db.Begin()
		if err != nil {
			xray.ROOT.Fork().WithLogger(loggerName).Error(
				"Unable to start transaction to write log data - :err",
				args.Error{Err: err},
			)
			return
		}

		// Storing data
		for _, l := range filtered {
			ray := ""
			if a := l.Get("rayId"); a != nil {
				ray = a.Value()
			}

			_, err := tx.Exec(
				"INSERT INTO `xrayLog` (`time`, `rayId`, `level`, `pattern`, `message`) VALUES (?, ?, ?, ?, ?)",
				l.GetTime().Unix(),
				ray,
				stringLevel(l.GetLevel()),
				l.GetMessage(),
				text.InterpolatePlainText(l.GetMessage(), l, true),
			)
			if err != nil {
				xray.ROOT.Fork().WithLogger(loggerName).Error(
					"Unable to write error data - :err",
					args.Error{Err: err},
				)
				_ = tx.Rollback()
				return
			}
		}

		// Commit
		if err := tx.Commit(); err != nil {
			xray.ROOT.Fork().WithLogger(loggerName).Error(
				"Unable to commit log data - :err",
				args.Error{Err: err},
			)
		}
	}
}

func stringLevel(l xray.Level) string {
	switch l {
	case xray.ALERT:
		return "alert"
	case xray.CRITICAL:
		return "critical"
	default:
		return "error"
	}
}
