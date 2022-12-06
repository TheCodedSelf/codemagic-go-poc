package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	apps := FetchApps()
	fmt.Println(apps)
	firstApp := apps[0]
	builds := FetchBuilds(firstApp.ID)
	fmt.Println(builds)
}

type App struct {
	ID          string `json:"_id"`
	Name        string `json:"appName"`
	ProjectType string
}

type AppsResponse struct {
	Applications []App `json:"applications"`
}

type BuildsResponse struct {
	Applications []App
	Builds       []Build
}

type Build struct {
	ID         string `json:"_id"`
	AppID      string `json:"appId"`
	WorkflowID string `json:"workflowId"`
	Branch     string
	Tag        string
	Status     string
	Artefacts  []Artefact
}

type Artefact struct {
	VersionName string
}

func FetchApps() []App {
	req, err := http.NewRequest("GET", "https://api.codemagic.io/apps", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("x-auth-token", "3s8X0R28KXch3w3WKgkEoiDqkQcWwhlOX-dBLuMosfQ")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var appsResp AppsResponse
	err = json.Unmarshal(body, &appsResp)
	if err != nil {
		log.Fatal(err)
	}

	return appsResp.Applications
}

func FetchBuilds(app string) []Build {
	req, err := http.NewRequest("GET", "https://api.codemagic.io/builds", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("x-auth-token", "3s8X0R28KXch3w3WKgkEoiDqkQcWwhlOX-dBLuMosfQ")

	urlValues := url.Values{}
	urlValues.Set("appId", app)
	req.URL.RawQuery = urlValues.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var buildsResp BuildsResponse
	err = json.Unmarshal(body, &buildsResp)
	if err != nil {
		log.Fatal(err)
	}

	return buildsResp.Builds
}
