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
	Name      string                `json:"name"`
	UUID      string                `json:"uuid"`
	Session   string                `json:"session"`
	Timestamp int                   `json:"timestamp"`
	Sign      string                `json:"sign"`
	Data      PersonListRequestData `json:"data"`
}

type PersonListRequestData struct {
	Action      string       `json:"action"`
	PersonType  int          `json:"persontype"`
	FeatureInfo FeatureSub   `json:"featureinfo"`
	PersonInfo  []PersonInfo `json:"personinfolist"`
}

type FeatureSub struct {
	FeatureType int `json:"featuretype"`
}

type PersonInfo struct {
	PersonCover     int             `json:"personcover"`
	PersonId        string          `json:"personid"`
	PersonName      string          `json:"personname"`
	Sex             int             `json:"sex"`
	IDCard          string          `json:"idcard"`
	Nation          string          `json:"nation"`
	Birthday        string          `json:"birthday"`
	Phone           string          `json:"phone"`
	Address         string          `json:"address"`
	LimitTime       int             `json:"limittime"`
	StartTime       string          `json:"starttime"`
	EndTime         string          `json:"endtime"`
	Label           string          `json:"label"`
	PersonExtension PersonExtension `json:"personextension"`
	PersonPhoto     string          `json:"personphoto"`
	FeatureValue    string          `json:"featurevalue"`
}

type PersonExtension struct {
	PersonCode1       string `json:"personcode1"`
	PersonCode2       string `json:"personcode2"`
	PersonCode3       string `json:"personcode3"`
	PersonReserveName string `json:"personreservename"`
	PersonParam1      int    `json:"personparam1"`
	PersonParam2      int    `json:"personparam2"`
	PersonParam3      int    `json:"personparam3"`
	PersonParam4      int    `json:"personparam4"`
	PersonParam5      int    `json:"personparam5"`
	PersonData1       string `json:"persondata1"`
	PersonData2       string `json:"persondata2"`
	PersonData3       string `json:"persondata3"`
	PersonData4       string `json:"persondata4"`
	PersonData5       string `json:"persondata5"`
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
