package service

import (
	"os"
	"testing"
)

func getTs() int64 { return 1632146917 }

func TestBuildUrl(t *testing.T) {
	// NOTE: Setup
	originalPubKey := os.Getenv(ENV_MARVEL_PUBLIC_KEY)
	os.Setenv(ENV_MARVEL_PUBLIC_KEY, "public")
	originalPrivKey := os.Getenv(ENV_MARVEL_PRIVATE_KEY)
	os.Setenv(ENV_MARVEL_PRIVATE_KEY, "private")

	t.Run("Build character URL", func(t *testing.T) {
		expected := "http://gateway.marvel.com/v1/public/characters/1017100?ts=1632146917&apikey=public&hash=1579e01b1726d59e140d15d3a39443f3"

		url := BuildCharacterUrl(getTs(), 1017100)

		if url != expected {
			t.Errorf("Unexpected character url")
		}
	})

	t.Run("Build all character URL offset=0 limit=100", func(t *testing.T) {
		expected := "http://gateway.marvel.com/v1/public/characters?orderBy=name&offset=0&limit=100&ts=1632146917&apikey=public&hash=1579e01b1726d59e140d15d3a39443f3"

		url := BuildAllCharacterUrl(getTs(), 0, 100)

		if url != expected {
			t.Errorf("Unexpected character url: %s", url)
		}
	})

	t.Run("Build all character URL offset=100 limit=100", func(t *testing.T) {
		expected := "http://gateway.marvel.com/v1/public/characters?orderBy=name&offset=100&limit=100&ts=1632146917&apikey=public&hash=1579e01b1726d59e140d15d3a39443f3"

		url := BuildAllCharacterUrl(getTs(), 100, 100)

		if url != expected {
			t.Errorf("Unexpected character url: %s", url)
		}
	})

	// NOTE: teardown
	os.Setenv(ENV_MARVEL_PUBLIC_KEY, originalPubKey)
	os.Setenv(ENV_MARVEL_PRIVATE_KEY, originalPrivKey)
}
