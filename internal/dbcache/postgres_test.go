package dbcache_test

import (
	"context"
	"errors"
	"strconv"
	"testing"
	"time"
	"translateapp/internal/dbcache"
	"translateapp/internal/logger"
	_ "translateapp/internal/logger"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	const key = "key"
	const expected = "true"

	tt := time.Now().Add(30 * time.Second)

	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())

	rows := mock.NewRows([]string{"value", "expiration"}).AddRow(expected, tt)
	mock.ExpectQuery("SELECT value, expiration from cache WHERE").WithArgs(key).WillReturnRows(rows)

	repo := dbcache.NewRepo(mock)
	loger := logger.DefaultLogger()
	var cache = dbcache.NewDBCache(repo, loger)
	val, _, err := cache.Get(key)
	require.NoError(t, err)
	str := strconv.FormatBool(val)
	require.Equal(t, expected, str)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSetError(t *testing.T) {
	const key = "key"

	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())

	mock.ExpectQuery("SELECT value, timeout from cache WHERE").WithArgs(key).WillReturnError(errors.New("error"))
	repo := dbcache.NewRepo(mock)

	cache := dbcache.NewDBCache(repo, logger.DefaultLogger())
	val, _, err := cache.Get(key)
	str := strconv.FormatBool(val)

	require.Error(t, err)
	require.Equal(t, "false", str)

	require.Error(t, mock.ExpectationsWereMet())
}
