package frepo_test

import (
	"context"
	"os"
	"testing"

	fdom "github.com/mrpandey/gobp/src/internal/core/domain/furniture"
	frepo "github.com/mrpandey/gobp/src/internal/repo/furniture"
	testrepo "github.com/mrpandey/gobp/src/internal/repo/testutil"
	"github.com/mrpandey/gobp/src/util"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var fr fdom.FurnitureRepoInterface
var testDBClient *gorm.DB

func TestMain(m *testing.M) {
	exitCode := 0
	defer func() { os.Exit(exitCode) }()

	db, cfg, teardownTestDB := testrepo.SetupTestDB()
	defer teardownTestDB()
	testDBClient = db

	disableCallerLog := false
	logger := util.NewLogger(cfg.LogLevel, cfg.DGN, disableCallerLog)
	// set global var
	fr = frepo.NewFurnitureRepo(logger, testDBClient)

	exitCode = m.Run() // run tests
}

func TestAutoIncrement(t *testing.T) {
	id1, err := fr.Create(context.TODO(), "Table 1", "table")
	assert.NoError(t, err)

	id2, err := fr.Create(context.TODO(), "Chair 1", "chair")
	assert.NoError(t, err)

	assert.Equal(t, id1+1, id2)

	deleteAllFurnitres()
}

func TestCreateFurniture(t *testing.T) {
	_, err := fr.Create(context.TODO(), "Sample Table", "table")
	assert.NoError(t, err)

	deleteAllFurnitres()
}

func TestGetFurniture(t *testing.T) {
	furnitureID, err := fr.Create(context.TODO(), "Some Table", "table")
	assert.NoError(t, err)

	record, err := fr.GetByID(context.TODO(), furnitureID)
	assert.NoError(t, err)

	expectedRecord := &fdom.FurnitureRecord{
		ID:   record.ID,
		Type: fdom.Table,
		Name: "Some Table",
	}
	assert.Equal(t, expectedRecord, record)

	deleteAllFurnitres()
}

func deleteAllFurnitres() {
	testDBClient.Exec("DELETE FROM furnitures;")
}
