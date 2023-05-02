package walletBusiness_test

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"game-inventory-management/internal/application/walletBusiness"
	"game-inventory-management/internal/domain/wallet"
)

func TestGetWallet_Success(t *testing.T) {
	id := uuid.New()
	db, mock, _ := sqlmock.New()
	defer db.Close()

	sql.Register("postgres", db.Driver())

	log := zap.NewNop().Sugar()

	expectedWallet := wallet.Wallet{
		Id:    id,
		Value: 0,
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, value FROM wallet WHERE id = $1")).
		WithArgs(id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "value"}).
				AddRow(id.String(), 0),
		)

	result, err := walletBusiness.GetWallet(id, db, log)

	assert.NoError(t, err)
	assert.Equal(t, expectedWallet, result)
}

func TestGetWallet_Error(t *testing.T) {
	id := uuid.New()
	db, mock, _ := sqlmock.New()
	defer db.Close()
	log := zap.NewNop().Sugar()

	expectedError := errors.New("something went wrong")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, value FROM wallet WHERE id = $1")).
		WithArgs(id).
		WillReturnError(expectedError)

	_, err := walletBusiness.GetWallet(id, db, log)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
}
