package repositories

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kashyapkrlucky/ff-go-src/internal/models"
	"github.com/redis/go-redis/v9"
)

// FeatureFlagRepo handles DB operations
type FeatureFlagRepo struct {
	DB    *sqlx.DB
	Redis *redis.Client
}

// NewFeatureFlagRepo creates a new repository
func NewFeatureFlagRepo(db *sqlx.DB, redis *redis.Client) *FeatureFlagRepo {
	return &FeatureFlagRepo{DB: db, Redis: redis}
}

// GetAll retrieves all feature flags, using Redis cache
func (r *FeatureFlagRepo) GetAll() ([]models.FeatureFlag, error) {
	cacheKey := "feature_flags"
	ctx := context.Background()

	// Check Redis cache
	cachedData, err := r.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var flags []models.FeatureFlag
		err = json.Unmarshal([]byte(cachedData), &flags)
		if err == nil {
			log.Println("Serving feature flags from Redis cache")
			return flags, nil
		}
	}

	// If cache miss, fetch from DB
	var flags []models.FeatureFlag
	err = r.DB.Select(&flags, "SELECT * FROM feature_flags")
	if err != nil {
		log.Println("Error fetching feature flags from DB:", err)
		return nil, err
	}

	// Store result in Redis for 1 minute
	data, _ := json.Marshal(flags)
	r.Redis.Set(ctx, cacheKey, data, 1*time.Minute)

	log.Println("Serving feature flags from DB and updating Redis cache")
	return flags, nil
}

// Create adds a new feature flag
func (r *FeatureFlagRepo) Create(flag models.FeatureFlag) error {
	_, err := r.DB.Exec("INSERT INTO feature_flags (name, enabled, created_at) VALUES ($1, $2, NOW())",
		flag.Name, flag.Enabled)
	return err
}

// Update modifies an existing feature flag and invalidates Redis cache
func (r *FeatureFlagRepo) Update(flag models.FeatureFlag) error {
	query := `UPDATE feature_flags SET name=$1, enabled=$2 WHERE id=$3`
	_, err := r.DB.Exec(query, flag.Name, flag.Enabled, flag.ID)
	if err == nil {
		// Invalidate cache
		_ = r.Redis.Del(context.Background(), "feature_flags")
	}
	return err
}

// Delete removes a feature flag and invalidates Redis cache
func (r *FeatureFlagRepo) Delete(id int) error {
	query := `DELETE FROM feature_flags WHERE id=$1`
	_, err := r.DB.Exec(query, id)
	if err == nil {
		// Invalidate cache
		_ = r.Redis.Del(context.Background(), "feature_flags")
	}
	return err
}
