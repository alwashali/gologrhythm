package gologrhythm

// filter options for cases
type CasesFilters struct {
	Name          string
	Count         string
	Priority      []int
	DueBefore     string
	Text          string
	OrderBy       string
	EvidenceType  string
	ReferenceId   string
	UpdatedBefore string
	CreatedAfter  string
	CreatedBefore string
	Direction     string
}

// filter options for alarms

type AlarmsFilters struct {
	Name           string
	AlarmRuleName  string
	AlarmStatus    string
	AseAssociation string
	DateInserted   string
	Count          string
	offset         string
	Orderby        string
	EntityName     string
}

type Case struct {
	Id            string `json:"id"`
	Number        int    `json:"number"`
	DateCreated   string `json:"dateCreated"`
	DateUpdated   string `json:"dateUpdated"`
	DateClosed    string `json:"dateClosed"`
	Owner         owner  `json:"owner"`
	LastUpdatedBy lastupdatedby
	DueDate       string `json:"dueDate"`
	Resolution    string `json:"resolution"`
	Name          string `json:"name"`
	Status        status `json:"status"`
	Priority      int    `json:"priority"`
	Summary       string `json:"summary"`
	Collaborators []collaborator
}

type status struct {
	Name   string `json:"name"`
	Number int    `json:"numbers"`
}

type owner struct {
	Number int    `json:"number"`
	Nmae   string `json:"Name"`
}

type collaborator struct {
	Number int    `json:"number"`
	Nmae   string `json:"Name"`
}

type lastupdatedby struct {
	Number   int    `json:"number"`
	Nmae     string `json:"name"`
	Disabled bool   `json:"disabled"`
}

type Alarm struct {
	AlarmDetails    alarmDetails `json:"alarmDetails"`
	StatusCode      int          `json:"statusCode"`
	StatusMessage   string       `json:"statusMessage"`
	ResponseMessage string       `json:"responseMessage"`
}

type alarmDetails struct {
	AlarmId          int      `json:"alarmid"`
	PersonId         int      `json:"personid"`
	EntityId         int      `json:"entityid"`
	EntityName       string   `json:"entityName"`
	AlarmDate        string   `json:"alarmDate"`
	AlarmRuleID      int      `json:"alarmRuleID"`
	AlarmRuleName    string   `json:"alarmRuleName"`
	AlarmStatus      int      `json:"alarmStatus"`
	AlarmStatusName  string   `json:"alarmStatusName"`
	LastUpdatedID    int      `json:"lastUpdatedID"`
	LastUpdatedName  string   `json:"lastUpdatedName"`
	DateInserted     string   `json:"dateInserted"`
	DateUpdated      string   `json:"dateUpdated"`
	Associatedalarms []string `json:"associatedCases"`
}

type AlarmEvent struct {
	AlarmDetails    []alarmEventsDetails `json:"alarmEventsDetails"`
	StatusCode      int                  `json:"statusCode"`
	StatusMessage   string               `json:"statusMessage"`
	ResponseMessage string               `json:"responseMessage"`
}

type alarmEventsDetails struct {
	Account                  string `json:"account"`
	Action                   string `json:"action"`
	Amount                   int    `json:"amount"`
	BytesIn                  string `json:"bytesIn"`
	BytesOut                 string `json:"bytesOut"`
	ClassificationId         int    `json:"classificationId"`
	ClassificationName       string `json:"classificationName"`
	ClassificationTypeName   string `json:"classificationTypeName"`
	Command                  string `json:"command"`
	CommonEventId            int    `json:"commonEventId"`
	Cve                      string `json:"cve"`
	CommonEventName          string `json:"commonEventName"`
	Count                    int    `json:"count"`
	DirectionId              int    `json:"directionId"`
	DirectionName            string `json:"directionName"`
	Domain                   string `json:"domain"`
	Duration                 int    `json:"duration"`
	EntityId                 int    `json:"entityId"`
	EntityName               string `json:"entityName"`
	Group                    string `json:"group"`
	ImpactedEntityId         int    `json:"impactedEntityId"`
	ImpactedEntityName       string `json:"impactedEntityName"`
	ImpactedHostId           int    `json:"impactedHostId"`
	ImpactedHostName         string `json:"impactedHostName"`
	ImpactedInterface        string `json:"impactedInterface"`
	ImpactedIP               string `json:"impactedIP"`
	ImpactedMAC              string `json:"impactedMAC"`
	ImpactedName             string `json:"impactedName"`
	ImpactedNATIP            string `json:"impactedNATIP"`
	ImpactedNATPort          string `json:"impactedNATPort"`
	ImpactedPort             int    `json:"impactedPort"`
	ImpactedZone             string `json:"impactedZone"`
	ItemsPacketsIn           int    `json:"itemsPacketsIn"`
	ItemsPacketsOut          int    `json:"itemsPacketsOut"`
	LogDate                  string `json:"logDate"`
	Login                    string `json:"login"`
	LogMessage               string `json:"logMessage"`
	LogSourceHostId          int    `json:"logSourceHostId"`
	LogSourceHostName        string `json:"logSourceHostName"`
	LogSourceName            string `json:"logSourceName"`
	LogSourceTypeName        string `json:"logSourceTypeName"`
	MessageId                int    `json:"messageId"`
	MpeRuleId                int    `json:"mpeRuleId"`
	MpeRuleName              string `json:"mpeRuleName"`
	NormalDateMax            string `json:"normalDateMax"`
	ObjectName               string `json:"objectName"`
	ObjectType               string `json:"objectType"`
	OriginEntityId           int    `json:"originEntityId"`
	OriginEntityName         string `json:"originEntityName"`
	OriginHostId             int    `json:"originHostId"`
	OriginHostName           string `json:"originHostName"`
	OriginInterface          string `json:"originInterface"`
	OriginIP                 string `json:"originIP"`
	OriginMAC                string `json:"originMAC"`
	OriginName               string `json:"originName"`
	OriginNATIP              string `json:"originNATIP"`
	OriginNATPort            string `json:"originNATPort"`
	OriginPort               int    `json:"originPort"`
	OriginZone               string `json:"originZone"`
	ParentProcessId          string `json:"parentProcessId"`
	ParentProcessName        string `json:"parentProcessName"`
	ParentProcessPath        string `json:"parentProcessPath"`
	Policy                   string `json:"policy"`
	Priority                 int    `json:"priority"`
	Process                  string `json:"process"`
	ProcessId                int    `json:"processId"`
	ProtocolId               int    `json:"protocolId"`
	ProtocolName             string `json:"protocolName"`
	Quantity                 int    `json:"quantity"`
	Rate                     int    `json:"rate"`
	Reason                   string `json:"reason"`
	Recipient                string `json:"recipient"`
	Result                   string `json:"result"`
	ResponseCode             string `json:"responseCode"`
	Sender                   string `json:"sender"`
	Session                  string `json:"session"`
	SessionType              string `json:"sessionType"`
	Serial                   string `json:"serial"`
	ServiceId                int    `json:"serviceId"`
	ServiceName              string `json:"serviceName"`
	Severity                 string `json:"severity"`
	Status                   string `json:"status"`
	Size                     int    `json:"size"`
	Subject                  string `json:"subject"`
	ThreatId                 string `json:"threatId"`
	ThreatName               string `json:"threatName"`
	Url                      string `json:"url"`
	UserAgent                string `json:"userAgent"`
	VendorInfo               string `json:"vendorInfo"`
	VendorMsgId              string `json:"vendorMsgId"`
	Version                  string `json:"version"`
	OriginUserIdentityName   string `json:"originUserIdentityName"`
	ImpactedUserIdentityName string `json:"impactedUserIdentityName"`
	OriginUserIdentityId     int    `json:"originUserIdentityId"`
	ImpactedUserIdentityId   int    `json:"impactedUserIdentityId"`
	SenderIdentityId         int    `json:"senderIdentityId"`
	SenderIdentityName       string `json:"senderIdentityName"`
	RecipientIdentityId      int    `json:"recipientIdentityId"`
	RecipientIdentityName    string `json:"recipientIdentityName"`
	StatusCode               string `json:"statusCode"`
	StatusMessage            string `json:"statusMessage"`
}

type Alarms struct {
	AlarmsSearchDetails []alarmsSearchDetails `json:"alarmsSearchDetails"`
	AlarmsCount         int                   `json:"alarmsCount"`
	StatusCode          int                   `json:"statusCode"`
	StatusMessage       string                `json:"statusMessage"`
	ResponseMessage     string                `json:"responseMessage"`
}

type alarmsSearchDetails struct {
	AlarmId         int      `json:"alarmId"`
	AlarmRuleName   string   `json:"alarmRuleName"`
	AlarmStatus     int      `json:"alarmStatus"`
	AlarmDataCached string   `json:"alarmDataCached"`
	AssociatedCases []string `json:"associatedCases"`
	EntityName      string   `json:"entityName"`
	DateInserted    string   `json:"dateInserted"`
}
