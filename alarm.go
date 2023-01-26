package gologrhythm

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (lr *LogRhtyhm) Alarms(filters *AlarmsFilters) (Alarms, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	alarms_api := fmt.Sprintf("https://%s/lr-alarm-api/alarms", lr.Host)
	alarms := Alarms{}

	client := &http.Client{}
	req, err := http.NewRequest("GET", alarms_api, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Authorization", "Bearer "+lr.Token)

	q := req.URL.Query()

	//fmt.Println(filters)

	if filters.Count != "" {

		q.Set("count", filters.Count)
	}

	if filters.EntityName != "" {

		q.Set("entityName", filters.EntityName)
	}

	if filters.AlarmStatus != "" {

		q.Set("count", filters.Count)
	}

	if filters.AlarmRuleName != "" {

		q.Set("alarmRuleName", filters.AlarmRuleName)
	}

	if filters.offset != "" {

		q.Set("count", filters.offset)
	}

	if filters.DateInserted != "" {

		q.Set("dateInserted", filters.DateInserted)
	}

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Can't request alarms from:", lr.Name, err)
	}

	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &alarms)

		if err != nil {
			log.Println(err)
		}

		return alarms, nil
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return alarms, errors.New("Error fetching alarms" + string(body))
}

func (lr *LogRhtyhm) Alarm(id int64) (*Alarm, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	alarm_api := fmt.Sprintf("https://%s/lr-alarm-api/alarms/%d", lr.Host, id)
	alarm := Alarm{}

	client := &http.Client{}
	req, err := http.NewRequest("GET", alarm_api, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Authorization", "Bearer "+lr.Token)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Can't request alarms from:", lr.Name, err)
	}

	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &alarm)

		if err != nil {
			log.Println(err)
		}

		return &alarm, nil
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return nil, errors.New("Error fetching alarms" + string(body))
}

func (lr *LogRhtyhm) AlarmEvents(id int64) (*AlarmEvent, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	alarm_api := fmt.Sprintf("https://%s/lr-alarm-api/alarms/%d/events", lr.Host, id)
	alarmEvent := AlarmEvent{}

	client := &http.Client{}
	req, err := http.NewRequest("GET", alarm_api, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Authorization", "Bearer "+lr.Token)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Can't request alarms from:", lr.Name, err)
	}

	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(body, &alarmEvent)

		if err != nil {
			log.Println(err)
		}

		return &alarmEvent, nil
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return nil, errors.New("Error fetching alarms" + string(body))
}
