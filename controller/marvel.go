package controller

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	m "marvel/model"
	"net/http"
	"os"
	"time"
)

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

func GetMarvelCharacter(getTs func() int64, id int, client *http.Client) m.CharacterResult {
	resp, err := client.Get(buildCharacterUrl(getTs(), id))

	if err != nil {
		log.Println("Error:", err)

		return m.CharacterResult{StatusCode: 500, Character: m.Character{}}
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return m.CharacterResult{StatusCode: 404, Character: m.Character{}}
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error:", err)

		return m.CharacterResult{StatusCode: 500, Character: m.Character{}}
	}

	if resp.StatusCode != 200 {
		log.Println("Error:", resp.StatusCode, string(body))

		return m.CharacterResult{StatusCode: 500, Character: m.Character{}}
	}

	wrapper := m.MarvelCharacterDataWrapper{}

	err = json.Unmarshal(body, &wrapper)

	if err != nil {
		log.Println("Error:", err)

		return m.CharacterResult{StatusCode: 500, Character: m.Character{}}
	}

	if len(wrapper.Data.Results) == 0 {
		log.Println("Error: received no data")

		return m.CharacterResult{StatusCode: 500, Character: m.Character{}}
	}

	mc := wrapper.Data.Results[0]

	return m.CharacterResult{StatusCode: 200, Character: m.Character{Id: mc.Id, Name: mc.Name, Description: mc.Description}}
}
