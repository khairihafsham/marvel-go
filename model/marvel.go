package model

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

type CharacterResult struct {
	Character  Character
	StatusCode int
}
