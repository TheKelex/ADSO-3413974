package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"

	"workshop/internal/domain"
)

const (
	duplicateEntryCode      = 1062
	foreignKeyViolationCode = 1452
	checkConstraintCode     = 3819
)

// Open connects to MySQL with a bounded pool and waits for the server to
// accept connections, so a database container that is still starting does not
// force a restart of the backend.
func Open(ctx context.Context, dsn string, timeout time.Duration) (*sql.DB, error) {
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("opening the database connection: %w", err)
	}
	database.SetMaxOpenConns(20)
	database.SetMaxIdleConns(10)
	database.SetConnMaxLifetime(30 * time.Minute)

	var lastErr error
	for attempt := 0; attempt < 30; attempt++ {
		pingCtx, cancel := context.WithTimeout(ctx, timeout)
		lastErr = database.PingContext(pingCtx)
		cancel()
		if lastErr == nil {
			return database, nil
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
	return nil, fmt.Errorf("the database did not accept connections: %w", lastErr)
}

// translate turns a driver error into the domain error the use cases expect,
// so no layer above persistence has to know a MySQL error number.
func translate(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return domain.ErrNotFound
	}
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case duplicateEntryCode:
			return fmt.Errorf("%w: the value is already registered", domain.ErrConflict)
		case foreignKeyViolationCode:
			return fmt.Errorf("%w: the referenced record does not exist", domain.ErrNotFound)
		case checkConstraintCode:
			return fmt.Errorf("%w: the value violates a storage rule", domain.ErrInvalidInput)
		}
	}
	return err
}
