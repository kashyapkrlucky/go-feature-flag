package repositories

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/jmoiron/sqlx"
	"github.com/kashyapkrlucky/ff-go-src/db"
	"github.com/kashyapkrlucky/ff-go-src/internal/models"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

// Test GetAll with Redis caching
func TestGetAll(t *testing.T) {
	// Start a mini Redis instance for testing
	mr, err := miniredis.Run()
	assert.NoError(t, err)
	defer mr.Close()

	// Mock Redis client
	redisClient := db.NewMockRedisClient(mr.Addr())

	// Mock database
	db, _ := sqlx.Open("postgres", "user=lkadmin password=secretpia dbname=feature_flags sslmode=disable")
	defer db.Close()

	repo := NewFeatureFlagRepo(db, redisClient)

	// Insert test data into DB
	db.Exec("INSERT INTO feature_flags (name, enabled) VALUES ($1, $2)", "test-feature", true)

	// Call GetAll()
	flags, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, flags, 1)

	// Check if cache is set
	cachedData, _ := redisClient.Get(context.Background(), "feature_flags").Result()
	var cachedFlags []models.FeatureFlag
	_ = json.Unmarshal([]byte(cachedData), &cachedFlags)
	assert.Len(t, cachedFlags, 1)
}
