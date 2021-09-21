package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	m "marvel/model"
	s "marvel/service"
	"net/http"
)

func GetMarvelCharacter(getTs func() int64, id int, client *http.Client) m.CharacterResult {
	resp, err := client.Get(s.BuildCharacterUrl(getTs(), id))

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
