package auth

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"

	"github.com/LeoAdamek/eventwrite/internal/logging"
	"go.uber.org/zap"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/spf13/viper"
)

type APIKey struct {
	PK       string `json:"-" dynamo:"pk"`
	SK       string `json:"-" dynamo:"sk"`
	Key      string `json:"key,omitempty" dynamo:"-"`
	SourceID string `json:"source_id" dynamo:"source_id"`
}

var cursor *dynamo.DB

func GetAPIKey(key string) (*APIKey, error) {

	log := logging.GetLogger().Named("auth")

	pk := "api_key:" + hashApiKey(key)

	c := getCursor()

	apiKey := APIKey{}

	err := c.Table(viper.GetString("dynamo_table")).
		Get("pk", pk).
		Range("sk", dynamo.Equal, "api_key").
		One(&apiKey)

	if err != nil {
		log.Warn(
			"Unable to find API Key with given pk",
			zap.String("pk", pk),
			zap.Error(err),
		)
	}

	return &apiKey, err
}

func getCursor() *dynamo.DB {
	if cursor == nil {
		cursor = createCursor()
	}

	return cursor
}

func createCursor() *dynamo.DB {
	s, err := session.NewSession()

	if err != nil {
		panic(err)
	}

	return dynamo.New(s)
}

// hashApiKey hashes the given API Key
func hashApiKey(key string) string {
	var st string

	buffer := bytes.NewBufferString(st)
	enc := base64.NewEncoder(base64.StdEncoding, buffer)

	enc.Write(sha256.New().Sum([]byte(key)))
	enc.Close()

	return buffer.String()
}
