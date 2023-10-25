package resth_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	resth "github.com/mrpandey/gobp/src/delivery/rest/handlers"
	healthdom "github.com/mrpandey/gobp/src/internal/core/domain/health"
	usecase "github.com/mrpandey/gobp/src/internal/core/usecase"
	"github.com/mrpandey/gobp/src/internal/repo"
	healthrepo "github.com/mrpandey/gobp/src/internal/repo/health"
	"github.com/mrpandey/gobp/src/util"
	"github.com/mrpandey/gobp/src/util/database"

	"github.com/stretchr/testify/assert"
	glogger "gorm.io/gorm/logger"
)

func TestRestHandler_Ping(t *testing.T) {
	cfg := util.NewConfig()
	disableCallerLog := false
	logger := util.NewLogger(cfg.LogLevel, cfg.DGN, disableCallerLog)

	db := database.NewGormPostgresClient(context.TODO(), cfg, logger, glogger.Warn)
	repos := repo.InitRepo(cfg, logger, db)

	useCases := usecase.InitUseCases(cfg, logger, repos)

	req, err := http.NewRequest("GET", "/ping/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	validator := util.NewRequestBodyValidator(logger)
	server := resth.NewRestServer(logger, validator, useCases)
	handler := http.HandlerFunc(server.PingHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v , expected %v", status, http.StatusOK)
	}

	var response healthdom.Ping
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	expected := healthdom.Ping{
		Message: "pong",
	}
	assert.Equal(t, expected, response)
}

func TestRestHandler_HealthSuccess(t *testing.T) {
	ctx := context.Background()
	cfg := util.NewConfig()
	disableCallerLog := false
	logger := util.NewLogger(cfg.LogLevel, cfg.DGN, disableCallerLog)

	db := database.NewGormPostgresClient(ctx, cfg, logger, glogger.Warn)
	sqlDB := database.GetGormConnPool(logger, db)
	defer sqlDB.Close()

	repos := repo.InitRepo(cfg, logger, db)

	useCases := usecase.InitUseCases(cfg, logger, repos)

	req := httptest.NewRequest("GET", "/health/", nil)

	rr := httptest.NewRecorder()
	validator := util.NewRequestBodyValidator(logger)
	server := resth.NewRestServer(logger, validator, useCases)
	handler := http.HandlerFunc(server.HealthHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp healthdom.DatabaseHealth
	err := json.NewDecoder(bytes.NewReader(rr.Body.Bytes())).Decode(&resp)
	if err != nil {
		t.Fatalf("error decoding response body: %v", err)
	}

	assert.Equal(t, true, resp.PostgresReachable)
}

func TestRestHandler_HealthFailure(t *testing.T) {
	cfg := util.NewConfig()
	disableCallerLog := false
	logger := util.NewLogger(cfg.LogLevel, cfg.DGN, disableCallerLog)

	repos := &repo.Repo{
		Health: healthrepo.NewHealthRepo(cfg, logger, nil),
	}

	useCases := usecase.InitUseCases(cfg, logger, repos)

	req := httptest.NewRequest("GET", "/health/", nil)
	rr := httptest.NewRecorder()
	validator := util.NewRequestBodyValidator(logger)
	server := resth.NewRestServer(logger, validator, useCases)
	handler := http.HandlerFunc(server.HealthHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusServiceUnavailable, rr.Code)

	var resp healthdom.DatabaseHealth
	err := json.NewDecoder(bytes.NewReader(rr.Body.Bytes())).Decode(&resp)
	if err != nil {
		t.Fatalf("error decoding response body: %v", err)
	}

	assert.Equal(t, false, resp.PostgresReachable)
}
