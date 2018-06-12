package handlers

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

//Send joke
func SendJoke (w http.ResponseWriter, r *http.Request) {
	// Get a Chuck Norris joke
	joke, err := getJoke()
	if err != nil {
		return
	}
	// Send
	if err != nil {
		fmt.Println("Couldn't send message")
		fmt.Println(err)
		return
	} else {
		fmt.Fprintln(w,joke)
		return
	}

	return
}

//Fetch Chuck Norris Joke
func getJoke() (string, error) {
	resp, err := http.Get("http://api.icndb.com/jokes/random")
	if err != nil {
		fmt.Println("Could not fetch joke")
		return "nil", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Unknown response body")
		return "nil", err
	}

	var jokeResp JokeApiResponse
	json.Unmarshal(body, &jokeResp)
	fmt.Println(jokeResp)
	return jokeResp.Value.Joke, nil
}