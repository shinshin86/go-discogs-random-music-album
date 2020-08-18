package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Album struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	ResourceUrl string `json:"resource_url"`
	ArtistName  string `json:"artists_sort"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	releaseId := rand.Intn(9999999)
	url := "https://api.discogs.com/releases/" + strconv.Itoa(releaseId)
	fmt.Println("Fetch release url: ", url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("status:", resp.Status)

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var album Album
	if err := json.Unmarshal(bytes, &album); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("Album Title: %s\n", album.Title)
	fmt.Printf("Artist Name: %s\n", album.ArtistName)
	fmt.Printf("Discogs ID: %d \n", album.Id)
	fmt.Printf("Discogs Resource URL: %s\n", album.ResourceUrl)
}
