package interview

import "database/sql"

// table taskhistory
type TaskHistoryRecord struct {
	id     int
	record string
}

func saveRecord(tx *sql.Tx, record TaskHistoryRecord) {
	// TODO implement
}
