package _http

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func GetPersonListFromDevice(deviceInfo _models.DEVICE) _models.PersonListResponse {
	url := "http://" + deviceInfo.IP + ":8011/Request"
	uuid := deviceInfo.UUID
	var personList _models.PersonListResponse

	UserName := "admin"
	PassWord := "admin"
	ts := strconv.Itoa(deviceInfo.TimeStamp)
	strData := []byte(uuid + ":" + UserName + ":" + PassWord + ":" + ts)
	hasher := md5.New()
	hasher.Write([]byte(strData))
	signedData := hex.EncodeToString(hasher.Sum(nil))

	var jsonBody = []byte(`
	{
	"Name": "personListRequest",
	"UUID": "` + uuid + `",
	"Session": "` + uuid + `_` + ts + `",
	"TimeStamp": ` + ts + `,
	"Sign": "` + signedData + `",
	  	"Data": {
		      "Action": "getPersonList", 
					"PersonType": 2,
					"PageNo": 1,
				  "PageSize": 1000
		        }
	}
	 `)
	fmt.Println(string(jsonBody))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	// 	"Sign": "` + signedData + `",
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Printf("%+v", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	json.Unmarshal([]byte(body), &personList)

	return personList
}

func AddPersonListToDevice(deviceInfo _models.DEVICE, plist _models.PersonListRequest) string {
	url := "http://" + deviceInfo.IP + ":8011/Request"

	jsonBody, err := json.Marshal(&plist)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBody))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Printf("%+v", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "\t")
	if error != nil {
		fmt.Println("JSON parse error: ", error)

	}

	log.Println(string(prettyJSON.Bytes()))
	return string(prettyJSON.Bytes())
}
