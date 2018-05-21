package env

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

var mutex = &sync.Mutex{}
var env = map[string]string{}

func init() {
	Load()
}

func loadEnv() {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		env[pair[0]] = os.Getenv(pair[0])
	}
}

// Reload the ENV variables. Useful if
// an external ENV manager has been used
func Reload() {
	env = map[string]string{}
	loadEnv()
}

// Load .env files. Files will be loaded in the same order that are received.
// Redefined vars will override previously existing values.
// IE: envy.Load(".env", "test_env/.env") will result in DIR=test_env
// If no arg passed, it will try to load a .env file.
func Load(files ...string) error {
	if len(files) == 0 {
		Reload()
		return nil
	}

	// We received a list of files
	for _, file := range files {

		// Check if it exists or we can access
		if _, err := os.Stat(file); err != nil {
			// It does not exist or we can not access.
			// Return and stop loading
			return err
		}

		// It exists and we have permission. Load it
		if err := godotenv.Overload(file); err != nil {
			return err
		}

		// Reload the env so all new changes are noticed
		Reload()

	}
	return nil
}

// Get a value from the ENV. If it doesn't exist the
// default value will be returned.
func Get(key, defaultValue string) string {
	mutex.Lock()
	defer mutex.Unlock()
	if val, ok := env[key]; ok {
		return val
	}
	return defaultValue
}

// MustGet a value from the ENV. If it doesn't exist
// an error will be returned
func MustGet(key string) (string, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if val, ok := env[key]; ok {
		return val, nil
	}
	return "", fmt.Errorf("could not find ENV var with KEY: %s", key)
}

// Set a value into the ENV. This is NOT permanent. It will
// only affect values accessed through env package.
func set(key, value string) {
	mutex.Lock()
	defer mutex.Unlock()
	env[key] = value
}

// MustSet the value into the underlying ENV, as well as env package.
// This may return an error if there is a problem setting the
// underlying ENV value.
func MustSet(key, value string) error {
	mutex.Lock()
	defer mutex.Unlock()
	err := os.Setenv(key, value)
	if err != nil {
		return err
	}
	env[key] = value
	return nil
}

// Map gets all of the keys/values from env.
func Map() map[string]string {
	return env
}

// Environ loads enviroment variables as KEY=VALUE pairs
func Environ() []string {
	mutex.Lock()
	defer mutex.Unlock()
	var e []string
	for k, v := range env {
		e = append(e, fmt.Sprintf("%s=%s", k, v))
	}
	return e
}
