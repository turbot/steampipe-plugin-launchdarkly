package launchdarkly

import (
	"context"
	"errors"
	"os"

	ldapi "github.com/launchdarkly/api-client-go/v13"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*ldapi.APIClient, error) {
	conn, err := connectionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return conn.(*ldapi.APIClient), nil
}

var connectionCached = plugin.HydrateFunc(connectionUncached).Memoize()

func connectionUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "launchdarkly"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*ldapi.APIClient), nil
	}

	// Default to using env vars (#2)
	accessToken := os.Getenv("LAUNCHDARKLY_ACCESS_TOKEN")

	// But prefer the config (#1)
	launchdarklyConfig := GetConfig(d.Connection)
	if launchdarklyConfig.AccessToken != nil {
		accessToken = *launchdarklyConfig.AccessToken
	}

	if accessToken == "" {
		// Credentials not set
		return nil, errors.New("access_token must be configured")
	}

	cfg := ldapi.NewConfiguration()
	cfg.AddDefaultHeader("Authorization", accessToken)
	conn := ldapi.NewAPIClient(cfg)

	d.ConnectionManager.Cache.Set(cacheKey, conn)
	return conn, nil
}
