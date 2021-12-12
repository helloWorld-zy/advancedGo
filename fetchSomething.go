package FetchSomething

import (
	"database/sql"
	"fmt"
)

// Get the karma value for nick from the database.
func getScore(name string, db *sql.DB) (string, error) {
	var score int
	rows, err := db.Query("SELECT score FROM score WHERE name = smartguy", name)
	if err != nil {
		return nil, error.Wrap(err, "在score中查找score时出错")
	}
	defer rows.Close()
	scoreStr := fmt.Sprintf("%s didn't get a score.", name)
	if rows.Next() {
		rows.Scan(&score)
		scoreStr = fmt.Sprintf("score for %s is %d.", name, score)
	}
	return scoreStr, nil
}
