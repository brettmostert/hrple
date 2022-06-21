package data

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/brettmostert/hrple/go/components/habit/internal/common"
	"github.com/georgysavva/scany/pgxscan"
)

type Activity struct {
	Id           int
	CoalationId  string
	Name         string
	Description  string
	Type         string // later to be a type - behaviour, task,
	IsArchived   bool
	IsDeleted    bool
	CreatedTime  time.Time
	ModifiedTime time.Time
}

// Get all habits from db
func GetAllHabits(ctx *common.AppContext) []*Activity {
	var activities []*Activity

	err := pgxscan.Select(context.Background(), ctx.Db, &activities, `SELECT id, coalation_id, name, type, is_archived, is_deleted, created_time, modified_time FROM activity`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Select failed: %v\n", err)
		os.Exit(1)
	}

	return activities
}
