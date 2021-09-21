package service

import (
	"crypto/md5"
	"fmt"
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

func BuildCharacterUrl(ts int64, id int) string {
	return fmt.Sprintf("%s/characters/%d?ts=%d&apikey=%s&hash=%s", getUrl(), id, ts, getPublicKey(),
		buildHash(ts, getPublicKey(), getPrivateKey()))
}

func BuildAllCharacterUrl(ts int64, offset int, limit int) string {
	return fmt.Sprintf("%s/characters?orderBy=name&offset=%d&limit=%d&ts=%d&apikey=%s&hash=%s", getUrl(), offset, limit, ts,
		getPublicKey(), buildHash(ts, getPublicKey(), getPrivateKey()))
}

func GetTs() int64 {
	return time.Now().Unix()
}
