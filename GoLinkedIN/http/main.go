package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type User struct {
	Login    string
	Name     string
	NumRepos int `json:"public_repos"`
}

func userInfo(login string) (*User, error) {
	u := fmt.Sprintf("https://api.github.com/users/" + url.PathEscape(login))
	res, err := http.Get(u)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: %s", res.Status)
	}
	usr := User{Login: login}
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&usr); err != nil {
		return nil, err
	}
	return &usr, nil
}

func main() {

	usr, err := userInfo("Vivek-Raj-IITB")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%+v\n", usr)

}
