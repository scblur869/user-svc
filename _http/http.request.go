package _http

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"local/user-svc/_models"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetPersonListFromDevice(payload _models.DeviceAuth) _models.PersonListResponse {
	url := "http://" + payload.DeviceIP + ":8011/Request"
	var personList _models.PersonListResponse

	UserName := payload.User
	PassWord := payload.Password
	uuid := payload.UUID
	ts := int(time.Now().Unix())
	strData := []byte(uuid + ":" + UserName + ":" + PassWord + ":" + strconv.Itoa(ts))
	hasher := md5.New()
	hasher.Write([]byte(strData))
	signedData := hex.EncodeToString(hasher.Sum(nil))

	var jsonBody = []byte(`
	{
	"Name": "personListRequest",
	"UUID": "` + uuid + `",
	"Session": "` + uuid + `_` + strconv.Itoa(ts) + `",
	"TimeStamp": ` + strconv.Itoa(ts) + `,
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

func AddPersonListToDevice(payload _models.ListPayload, plist _models.PersonListRequest) string {
	url := "http://" + payload.Auth.DeviceIP + ":8011/Request"

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

func GetFeatureValuesFromImage(ip string, user string, pass string, _uuid string, photo string) _models.FeatureValueResponse {
	url := "http://" + ip + ":8011/Request"
	var featureValue _models.FeatureValueResponse

	UserName := user
	PassWord := pass
	uuid := _uuid
	ts := int(time.Now().Unix())
	strData := []byte(uuid + ":" + UserName + ":" + PassWord + ":" + strconv.Itoa(ts))
	hasher := md5.New()
	hasher.Write([]byte(strData))
	signedData := hex.EncodeToString(hasher.Sum(nil))

	var jsonBody = []byte(`
	{
	"Name": "faceFeatureValueRequest",
	"UUID": "` + uuid + `",
	"Session": "` + uuid + `_` + strconv.Itoa(ts) + `",
	"TimeStamp": ` + strconv.Itoa(ts) + `,
	"Sign": "` + signedData + `",
	  	"Data": {
		        "FacePicture": "` + photo + `"
		        }
	}
	 `)
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
	fmt.Println("response Body:", string(body))

	json.Unmarshal([]byte(body), &featureValue)

	return featureValue
}
