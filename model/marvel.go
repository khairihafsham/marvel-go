package model

type MarvelCharacterDataWrapper struct {
	Data MarvelCharacterDataContainer `json:"data"`
	Etag string                       `json:"etag"`
}

type MarvelCharacterDataContainer struct {
	Offset  int               `json:"offset"`
	Limit   int               `json:"limit"`
	Count   int               `json:"count"`
	Total   int               `json:"total"`
	Results []MarvelCharacter `json:"results"`
}

type MarvelCharacter struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Character struct {
	Id          int    `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type CharacterResult struct {
	Character  Character
	StatusCode int
}
