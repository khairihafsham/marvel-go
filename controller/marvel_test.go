package controller

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/dnaeon/go-vcr/v2/recorder"
)

func TestGetCharacter(t *testing.T) {
	// NOTE: Setup
	originalPubKey := os.Getenv(ENV_MARVEL_PUBLIC_KEY)
	os.Setenv(ENV_MARVEL_PUBLIC_KEY, "public")
	originalPrivKey := os.Getenv(ENV_MARVEL_PRIVATE_KEY)
	os.Setenv(ENV_MARVEL_PRIVATE_KEY, "private")

	t.Run("Success", func(t *testing.T) {
		r, err := recorder.New("../fixture/character200")
		if err != nil {
			log.Fatal(err)
		}
		defer r.Stop()

		client := &http.Client{
			Transport: r,
		}

		result := GetMarvelCharacter(func() int64 { return 1632146917 }, 1017100, client)

		if result.StatusCode != 200 {
			t.Errorf("Expected 200, got %d", result.StatusCode)
		}

		if result.Character.Id != 1017100 {
			t.Errorf("Unexpected Character.Id: %d", result.Character.Id)
		}
		if result.Character.Name != "A-Bomb (HAS)" {
			t.Errorf("Unexpected Character.Name: %s", result.Character.Name)
		}

		if !strings.HasPrefix(result.Character.Description, "Rick Jones has been Hulk's") {
			t.Errorf("Unexpected Character.Description prefix: %s", result.Character.Description)
		}
	})

	t.Run("Not Found", func(t *testing.T) {
		r, err := recorder.New("../fixture/character404")
		if err != nil {
			log.Fatal(err)
		}
		defer r.Stop()

		client := &http.Client{
			Transport: r,
		}

		result := GetMarvelCharacter(func() int64 { return 1632146917 }, 101710, client)

		if result.StatusCode != 404 {
			t.Errorf("Expected 404, got %d", result.StatusCode)
		}

		if result.Character.Id != 0 {
			t.Errorf("Unexpected Character.Id: %d", result.Character.Id)
		}
		if result.Character.Name != "" {
			t.Errorf("Unexpected Character.Name: %s", result.Character.Name)
		}

		if result.Character.Description != "" {
			t.Errorf("Unexpected Character.Description: %s", result.Character.Description)
		}
	})

	t.Run("Failed Authentication", func(t *testing.T) {
		log.SetOutput(ioutil.Discard)

		r, err := recorder.New("../fixture/character401")
		if err != nil {
			log.Fatal(err)
		}
		defer r.Stop()

		client := &http.Client{
			Transport: r,
		}

		result := GetMarvelCharacter(func() int64 { return 1632146917 }, 1017100, client)

		if result.StatusCode != 500 {
			t.Errorf("Expected 500, got %d", result.StatusCode)
		}

		if result.Character.Id != 0 {
			t.Errorf("Unexpected Character.Id: %d", result.Character.Id)
		}
		if result.Character.Name != "" {
			t.Errorf("Unexpected Character.Name: %s", result.Character.Name)
		}

		if result.Character.Description != "" {
			t.Errorf("Unexpected Character.Description: %s", result.Character.Description)
		}
	})

	// NOTE: teardown
	os.Setenv(ENV_MARVEL_PUBLIC_KEY, originalPubKey)
	os.Setenv(ENV_MARVEL_PRIVATE_KEY, originalPrivKey)
}
