package tmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SearchMovie(client *http.Client, q string) (ResponseSearch, error) {
	url := UrlBase + "/search/movie" + AuthToken + "&query=" + q

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return ResponseSearch{}, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ResponseSearch{}, err
	}

	if err != nil {
		fmt.Println(err)
		return ResponseSearch{}, err
	}
	fmt.Println(resp.Status)
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ResponseSearch{}, err
	}
	var responseObject ResponseSearch
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		fmt.Println(err.Error())
		return ResponseSearch{}, err
	}
	return responseObject, nil
}

func RetrieveMovieInfo(client *http.Client, id string) (ResultFull, error) {
	url := UrlBase + "/movie/" + id + AuthToken

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return ResultFull{}, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ResultFull{}, err
	}

	fmt.Println(resp.Status)
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ResultFull{}, err
	}
	var responseObject ResultFull
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		fmt.Println(err.Error())
		return ResultFull{}, err
	}
	return responseObject, nil
}
