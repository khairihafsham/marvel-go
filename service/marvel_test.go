package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/dnaeon/go-vcr/v2/recorder"
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

func TestGetAllCharacter(t *testing.T) {
	// NOTE: Setup
	originalPubKey := os.Getenv(ENV_MARVEL_PUBLIC_KEY)
	os.Setenv(ENV_MARVEL_PUBLIC_KEY, "public")
	originalPrivKey := os.Getenv(ENV_MARVEL_PRIVATE_KEY)
	os.Setenv(ENV_MARVEL_PRIVATE_KEY, "private")
	log.SetOutput(ioutil.Discard)

	t.Run("Success offset=0 limit=10", func(t *testing.T) {
		r, err := recorder.New("../fixture/allcharacters-offset0-limit10")

		if err != nil {
			log.Fatal(err)
		}
		defer r.Stop()

		client := &http.Client{
			Transport: r,
		}

		result, err := GetAllCharacter(func() int64 { return 1632146917 }, 0, 10, client)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if result.Data.Offset != 0 {
			t.Errorf("Expected offset 0, got %d", result.Data.Offset)
		}

		if result.Data.Limit != 10 {
			t.Errorf("Expected limit 10, got %d", result.Data.Limit)
		}

		if result.Data.Total != 1559 {
			t.Errorf("Expected total 1559, got %d", result.Data.Total)
		}

		if result.Data.Count != 10 {
			t.Errorf("Expected count 10, got %d", result.Data.Count)
		}

		if len(result.Data.Results) != result.Data.Count {
			t.Errorf("Expected 10 results, got %d", len(result.Data.Results))
		}
	})

	t.Run("Success offset=1550 limit=10", func(t *testing.T) {
		r, err := recorder.New("../fixture/allcharacters-offset1550-limit10")

		if err != nil {
			log.Fatal(err)
		}
		defer r.Stop()

		client := &http.Client{
			Transport: r,
		}

		result, err := GetAllCharacter(func() int64 { return 1632146917 }, 1550, 10, client)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if result.Data.Offset != 1550 {
			t.Errorf("Expected offset 1550, got %d", result.Data.Offset)
		}

		if result.Data.Limit != 10 {
			t.Errorf("Expected limit 10, got %d", result.Data.Limit)
		}

		if result.Data.Total != 1559 {
			t.Errorf("Expected total 1559, got %d", result.Data.Total)
		}

		if result.Data.Count != 9 {
			t.Errorf("Expected count 9, got %d", result.Data.Count)
		}

		if len(result.Data.Results) != result.Data.Count {
			t.Errorf("Expected 9 results, got %d", len(result.Data.Results))
		}
	})

	t.Run("Success offset=1560 limit=10", func(t *testing.T) {
		r, err := recorder.New("../fixture/allcharacters-offset1560-limit10")

		if err != nil {
			log.Fatal(err)
		}
		defer r.Stop()

		client := &http.Client{
			Transport: r,
		}

		result, err := GetAllCharacter(func() int64 { return 1632146917 }, 1560, 10, client)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if result.Data.Offset != 1560 {
			t.Errorf("Expected offset 1550, got %d", result.Data.Offset)
		}

		if result.Data.Limit != 10 {
			t.Errorf("Expected limit 10, got %d", result.Data.Limit)
		}

		if result.Data.Total != 1559 {
			t.Errorf("Expected total 1559, got %d", result.Data.Total)
		}

		if result.Data.Count != 0 {
			t.Errorf("Expected count 0, got %d", result.Data.Count)
		}

		if len(result.Data.Results) != result.Data.Count {
			t.Errorf("Expected 0 results, got %d", len(result.Data.Results))
		}
	})

	t.Run("Fail offset=0 limit=101", func(t *testing.T) {
		r, err := recorder.New("../fixture/allcharacters-offset0-limit101")

		if err != nil {
			log.Fatal(err)
		}
		defer r.Stop()

		client := &http.Client{
			Transport: r,
		}

		_, err = GetAllCharacter(func() int64 { return 1632146917 }, 0, 101, client)

		if err == nil {
			t.Fatalf("Expected error, got nil")
		}
	})

	// NOTE: teardown
	os.Setenv(ENV_MARVEL_PUBLIC_KEY, originalPubKey)
	os.Setenv(ENV_MARVEL_PRIVATE_KEY, originalPrivKey)
}

func TestGetAllCharacterId(t *testing.T) {
	// NOTE: Setup
	originalPubKey := os.Getenv(ENV_MARVEL_PUBLIC_KEY)
	os.Setenv(ENV_MARVEL_PUBLIC_KEY, "public")
	originalPrivKey := os.Getenv(ENV_MARVEL_PRIVATE_KEY)
	os.Setenv(ENV_MARVEL_PRIVATE_KEY, "private")
	log.SetOutput(ioutil.Discard)

	t.Run("Success limit=10 total=19", func(t *testing.T) {
		r, err := recorder.New("../fixture/allcharacters-3-complete-requests")

		if err != nil {
			log.Fatal(err)
		}
		defer r.Stop()

		client := &http.Client{
			Transport: r,
		}

		result, err := GetAllCharacterId(func() int64 { return 1632146917 }, 10, client)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(result) != 19 {
			t.Errorf("Expected 19 results, got %d", len(result))
		}
	})

	// NOTE: teardown
	os.Setenv(ENV_MARVEL_PUBLIC_KEY, originalPubKey)
	os.Setenv(ENV_MARVEL_PRIVATE_KEY, originalPrivKey)
}
