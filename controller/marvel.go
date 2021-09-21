package controller

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type MarvelCharacterDataWrapper struct {
	Data MarvelCharacterDataContainer `json:"data"`
	Etag string                       `json:"etag"`
}

type MarvelCharacterDataContainer struct {
	Results []MarvelCharacter `json:"results"`
}

type MarvelCharacter struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Character struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Result struct {
	Character  Character
	StatusCode int
}

const ENV_MARVEL_URL string = "MARVEL_URL"
const ENV_MARVEL_PRIVATE_KEY string = "MARVEL_PRIVATE_KEY"
const ENV_MARVEL_PUBLIC_KEY string = "MARVEL_PUBLIC_KEY"

func getUrl() string {
	return os.Getenv(ENV_MARVEL_URL)
}

func getPrivateKey() string {
	return os.Getenv(ENV_MARVEL_PRIVATE_KEY)
}

func getPublicKey() string {
	return os.Getenv(ENV_MARVEL_PUBLIC_KEY)
}

func buildHash(ts int64, publicKey string, privateKey string) string {
	data := []byte(fmt.Sprintf("%d%s%s", ts, privateKey, publicKey))

	return fmt.Sprintf("%x", md5.Sum(data))
}

func buildCharacterUrl(ts int64, id int) string {
	return fmt.Sprintf("%s/characters/%d?ts=%d&apikey=%s&hash=%s", getUrl(), id, ts, getPublicKey(),
		buildHash(ts, getPublicKey(), getPrivateKey()))
}

func GetTs() int64 {
	return time.Now().Unix()
}

func GetMarvelCharacter(getTs func() int64, id int, client *http.Client) Result {
	resp, err := client.Get(buildCharacterUrl(getTs(), id))

	if err != nil {
		log.Println("Error:", err)

		return Result{StatusCode: 500, Character: Character{}}
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return Result{StatusCode: 404, Character: Character{}}
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error:", err)

		return Result{StatusCode: 500, Character: Character{}}
	}

	if resp.StatusCode != 200 {
		log.Println("Error:", resp.StatusCode, string(body))

		return Result{StatusCode: 500, Character: Character{}}
	}

	wrapper := MarvelCharacterDataWrapper{}

	err = json.Unmarshal(body, &wrapper)

	if err != nil {
		log.Println("Error:", err)

		return Result{StatusCode: 500, Character: Character{}}
	}

	if len(wrapper.Data.Results) == 0 {
		log.Println("Error: received no data")

		return Result{StatusCode: 500, Character: Character{}}
	}

	mc := wrapper.Data.Results[0]

	return Result{StatusCode: 200, Character: Character{Id: mc.Id, Name: mc.Name, Description: mc.Description}}
}
