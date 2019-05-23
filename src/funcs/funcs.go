package funcs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "strconv"
	// "reflect"

	"db"
)

type Dialogue struct {
	Character string `json:"character"`
	Dialogues string `json:"dialogue"`
	// Keywords  []string `json:"keyword"`
}

type Character struct {
	Name        string `json:"name"`
	BaseImg     string `json:"baseimg"`
	Weapon      string `json:"weapon"`
	Description string `json:"description"`
	Speciality  string `json:"speciality"`
	Defeated    string `json:"defeated"`
	Dialogues   string `json:"dialogue"`
}

func AddCharacter(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var d Character
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Body reading error", err)
		}
		err = json.Unmarshal(reqBody, &d)
		if err != nil {
			fmt.Println("unmarshaling err", err)
		}
		fmt.Println("d-->", d)
		db.HSet("h:"+d.Name, "name", d.Name)
		db.HSet("h:"+d.Name, "baseimg", d.BaseImg)
		db.HSet("h:"+d.Name, "weapon", d.Weapon)
		db.HSet("h:"+d.Name, "description", d.Description)
		db.HSet("h:"+d.Name, "speciality", d.Speciality)
		db.HSet("h:"+d.Name, "defeated", d.Defeated)
		db.HSet("h:"+d.Name, "dialogue", d.Dialogues)
	}

	if r.Method == "GET" {
		character := r.FormValue("name")
		fmt.Println("char-->", character)
		dialogue, err := db.HGetAll("h:" + character)
		if err != nil {
			fmt.Println("error in getting dialogues", err)
		}
		// fmt.Println("dialoues", dialogue)
		jsonString, _ := json.Marshal(dialogue)
		w.Write(jsonString)
	}

	if r.Method == "DELETE" {
		character := r.FormValue("character")
		db.Del("h:" + character)
	}

}

// func AddCharacter(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == "POST" {
// 		var c Character
// 		characters, err := ioutil.ReadAll(r.Body)
// 		err = json.Unmarshal(characters, &c)
// 		if err != nil {
// 			fmt.Println("Error in reading character Array", err)
// 		}
// 		db.HSet("h:"+c.Characters, "baseimg", c.BaseImg)
// 		db.HSet("h:"+c.Characters, "character", c.Characters)

// 	}

// 	if r.Method == "GET" {
// 		// var c []Character
// 		list := db.LRange("mylist")
// 		fmt.Println("list>>", list)
// 		byteArr, _ := json.Marshal(list)
// 		w.Write(byteArr)
// 	}
// }
