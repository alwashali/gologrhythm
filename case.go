package gologrhythm

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (lr *LogRhtyhm) Cases(filters *CasesFilters) ([]*Case, error) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	cases_api := fmt.Sprintf("https://%s/lr-case-api/cases/", lr.Host)
	cases := []*Case{}

	client := &http.Client{}
	req, err := http.NewRequest("GET", cases_api, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Authorization", "Bearer "+lr.Token)

	if filters.Count != "" {
		req.Header.Set("count", filters.Count)
	}

	if filters.CreatedAfter != "" {
		req.Header.Set("createdAfter", filters.CreatedAfter)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Can't request cases from:", lr.Name, err)
		return nil, errors.New("Error fetching cases")
	}

	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &cases)

		if err != nil {
			log.Println(err)
		}

		return cases, nil
	} else if resp.StatusCode == 401 || resp.StatusCode == 403 {
		return nil, errors.New("Unauthorized")
	} else if resp.StatusCode == 400 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(string(body))
	}

	return nil, errors.New("Error fetching cases")

}

func (lr *LogRhtyhm) Case(ID int) (*Case, error) {

	c := Case{}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	case_api := fmt.Sprintf("https://%s/lr-case-api/cases/%d", lr.Host, ID)

	client := &http.Client{}
	req, err := http.NewRequest("GET", case_api, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Authorization", "Bearer "+lr.Token)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Can't request cases from:", lr.Name, err)
	}

	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &c)

		if err != nil {
			log.Println(err)
		}

		return &c, nil

	} else if resp.StatusCode == 401 || resp.StatusCode == 403 {
		return nil, errors.New("Unauthorized")
	} else if resp.StatusCode == 400 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(string(body))
	}

	return nil, errors.New("Error fetching cases")

}

func (lr *LogRhtyhm) SetCaseStatusCreated(ID int) (string, error) {

	return updateCaseStatus(lr, ID, 1)
}
func (lr *LogRhtyhm) SetCaseStatusCompleted(ID int) (string, error) {
	return updateCaseStatus(lr, ID, 2)
}

func (lr *LogRhtyhm) SetCaseStatusIncident(ID int) (string, error) {
	return updateCaseStatus(lr, ID, 3)
}

func (lr *LogRhtyhm) SetCaseStatusMitigated(ID int) (string, error) {
	return updateCaseStatus(lr, ID, 4)
}

func (lr *LogRhtyhm) SetCaseStatusResolved(ID int) (string, error) {
	return updateCaseStatus(lr, ID, 5)
}

func updateCaseStatus(lr *LogRhtyhm, id, statusID int) (string, error) {
	changeCaseAPI := fmt.Sprintf("https://%s/lr-case-api/cases/%d/actions/changeStatus/", lr.Host, id)
	reqJson := []byte(fmt.Sprintf(`{"statusNumber": %d}`, statusID))

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}

	req, err := http.NewRequest("PUT", changeCaseAPI, bytes.NewBuffer(reqJson))

	if err != nil {
		log.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+lr.Token)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error updating case %d \n%s", id, err)
		return "Failed", nil
	}

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return "Failed", errors.New(string(body))
	}

	return "Success", nil
}
