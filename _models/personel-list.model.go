package _models

type PersonListResponse struct {
	Name    string         `json:"name"`
	Data    PersonListData `json:"data"`
	Code    int            `json:"code"`
	Message string         `json:"message"`
}

type PersonListData struct {
	PersonType    int          `json:"persontype"`
	Action        string       `json:"action"`
	PageNo        int          `json:"pageno"`
	PageSize      int          `json:"pagesize"`
	PageCount     int          `json:"pagecount"`
	PersonCount   int          `json:"personcount"`
	PersonListNum int          `json:"personlistnum"`
	PersonList    []PersonList `json:"personlist"`
}

type PersonList struct {
	PersonId   string `json:"personid"`
	PersonName string `json:"personname"`
}

type PersonListRequest struct {
	Name      string                `json:"Name"`
	UUID      string                `json:"UUID"`
	Session   string                `json:"Session"`
	Timestamp int                   `json:"TimeStamp"`
	Sign      string                `json:"Sign"`
	Data      PersonListRequestData `json:"Data"`
}

type PersonListRequestData struct {
	Action     string     `json:"Action"`
	PersonType int        `json:"PersonType"`
	PersonInfo PersonInfo `json:"PersonInfo"`
}

type FeatureSub struct {
	FeatureType int `json:"FeatureType"`
}

type PersonInfo struct {
	PersonCover     int             `json:"PersonCover"`
	PersonId        string          `json:"PersonId"`
	PersonName      string          `json:"PersonName"`
	Sex             int             `json:"Sex"`
	IDCard          string          `json:"IDCard"`
	Nation          string          `json:"Nation"`
	Birthday        string          `json:"Birthday"`
	Phone           string          `json:"Phone"`
	Address         string          `json:"Address"`
	LimitTime       int             `json:"LimitTime"`
	StartTime       string          `json:"StartTime"`
	EndTime         string          `json:"EndTime"`
	Label           string          `json:"Label"`
	PersonExtension PersonExtension `json:"PersonExtension"`
	PersonPhoto     string          `json:"PersonPhoto"`
	FeatureValue    string          `json:"FeatureValue"`
}

type PersonExtension struct {
	PersonCode1       string `json:"PersonCode1"`
	PersonCode2       string `json:"PersonCode2"`
	PersonCode3       string `json:"PersonCode3"`
	PersonReserveName string `json:"PersonReserveName"`
	PersonParam1      int    `json:"PersonParam1"`
	PersonParam2      int    `json:"PersonParam2"`
	PersonParam3      int    `json:"PersonParam3"`
	PersonParam4      int    `json:"PersonParam4"`
	PersonParam5      int    `json:"PersonParam5"`
	PersonData1       string `json:"PersonData1"`
	PersonData2       string `json:"PersonData2"`
	PersonData3       string `json:"PersonData3"`
	PersonData4       string `json:"PersonData4"`
	PersonData5       string `json:"PersonData5"`
}

type DEVICE struct {
	TimeStamp   int    `json:"timestamp"`
	ID          string `json:"deviceid"`
	UUID        string `json:"deviceuuid"`
	IP          string `json:"deviceip"`
	MAC         string `json:"devicemac"`
	CoreVersion string `json:"coreversion"`
	WebVersion  string `json:"webversion"`
	DeviceType  int    `json:"devicetype"`
	ChannelNo   int    `json:"channelno"`
	VersionDate string `json:"versiondate"`
	Created     string `json:"created"`
}

type FeatureValueResponse struct {
	Name    string           `json:"name"`
	Data    FeatureValueData `json:"data"`
	Code    int              `json:"code"`
	Message string           `json:"message"`
}

type FeatureValueData struct {
	FeatureType  int    `json:"featuretype"`
	FeatureValue string `json:"featurevalue"`
}
