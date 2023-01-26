# GoLogRhythm

Go API Client for LogRhythm SIEM. 

The client was built to have an App for each entity API if logRhythm is used in a multi-tenancy environment.


## Get started 
---

**Fetch Cases** 

```go
token := "Token"

app := gologrhythm.Create_App("DD", token, "10.10.10.10:8501")

cf := &gologrhythm.CasesFilters{
  Count: "10",
}

cases, err := app.Cases(cf)
if err != nil {
  fmt.Println(err)
}

for _, c := range cases {
  fmt.Println(c.Name)
}

```


**Fetch Alarms** 

```go

token := "Token"

app := gologrhythm.Create_App("DD", token, "10.10.10.10:8501")

af := gologrhythm.AlarmsFilters{
  Count:      "50",
  EntityName: "Entity1",
}

alarms, err := app.Alarms(&af)
if err != nil {
  fmt.Println(err)
}

for _, alarm := range alarms.AlarmsSearchDetails {
		fmt.Println(alarm.EntityName, alarm.AlarmRuleName )
}

```

**Fetch Alarm or Case by ID**


```go

token := "token"

app := gologrhythm.Create_App("DD", token, "10.10.10.10:8501")

c, err := app.Case(123)
if err != nil {
  fmt.Println(err)
}
fmt.Println(c.Name)

alarm, err := app.Alarm(66123)
if err != nil {
  fmt.Println(err)
}
fmt.Println(alarm.AlarmDetails.AlarmRuleName)

```
