package adapter

import (
	"io/ioutil"
	"log"
	"marvel/controller"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dnaeon/go-vcr/v2/recorder"
	"github.com/gin-gonic/gin"
)

func getTs() int64 {
	return 1632146917
}

func TestGetCharacter(t *testing.T) {
	// NOTE: Setup
	originalPubKey := os.Getenv(controller.ENV_MARVEL_PUBLIC_KEY)
	os.Setenv(controller.ENV_MARVEL_PUBLIC_KEY, "public")
	originalPrivKey := os.Getenv(controller.ENV_MARVEL_PRIVATE_KEY)
	os.Setenv(controller.ENV_MARVEL_PRIVATE_KEY, "private")

	t.Run("Success", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		r, err := recorder.New("../fixture/character200")
		if err != nil {
			log.Fatal(err)
		}
		defer r.Stop()

		client := &http.Client{
			Transport: r,
		}

		lambda := GetCharacter(getTs, client)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Params = gin.Params{gin.Param{Key: "id", Value: "1017100"}}

		lambda(c)

		if w.Code != 200 {
			t.Errorf("Expected 200, got %d", w.Code)
		}
	})

	t.Run("Not Found", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		r, err := recorder.New("../fixture/character404")
		if err != nil {
			log.Fatal(err)
		}
		defer r.Stop()

		client := &http.Client{
			Transport: r,
		}

		lambda := GetCharacter(getTs, client)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Params = gin.Params{gin.Param{Key: "id", Value: "101710"}}

		lambda(c)

		if w.Code != 404 {
			t.Errorf("Expected 404, got %d", w.Code)
		}
	})

	t.Run("Failure returns 500", func(t *testing.T) {
		log.SetOutput(ioutil.Discard)
		gin.SetMode(gin.TestMode)

		r, err := recorder.New("../fixture/character401")
		if err != nil {
			log.Fatal(err)
		}
		defer r.Stop()

		client := &http.Client{
			Transport: r,
		}

		lambda := GetCharacter(getTs, client)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Params = gin.Params{gin.Param{Key: "id", Value: "1017100"}}

		lambda(c)

		if w.Code != 500 {
			t.Errorf("Expected 500, got %d", w.Code)
		}
	})

	// NOTE: teardown
	os.Setenv(controller.ENV_MARVEL_PUBLIC_KEY, originalPubKey)
	os.Setenv(controller.ENV_MARVEL_PRIVATE_KEY, originalPrivKey)
}
