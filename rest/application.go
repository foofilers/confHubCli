package rest

import (
	"fmt"
	"encoding/json"
	"errors"
)

func (cli *ConfHubClient) AddApplication(name string) error {
	resp, err := cli.R().SetFormData(map[string]string{
		"name": name,
	}).Post(cli.Url + "/api/apps")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 201 {
		return errors.New(fmt.Sprintf("Application not created:[%d] %s", resp.StatusCode(), resp.Body()))
	}
	return nil
}

func (cli *ConfHubClient) DeleteApplication(name string) error {
	resp, err := cli.R().Delete(cli.Url + "/api/apps/" + name)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 204 {
		return errors.New(fmt.Sprintf("Application not deleted:[%d] %s", resp.StatusCode(), resp.Body()))
	}
	return nil
}

func (cli *ConfHubClient) RenameApplication(name, newName string) error {
	resp, err := cli.R().SetFormData(map[string]string{
		"name": newName,
	}).Put(cli.Url + "/api/apps/" + name)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 204 {
		return errors.New(fmt.Sprintf("Application not renamed:[%d] %s", resp.StatusCode(), resp.Body()))
	}
	return nil
}

func (cli *ConfHubClient) ListApplication() ([]map[string]interface{}, error) {
	resp, err := cli.R().Get(cli.Url + "/api/apps")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, errors.New(fmt.Sprintf("Error retrieving application list:[%d] %s", resp.StatusCode(), resp.Body()))
	}
	apps := make([]map[string]interface{}, 0)
	if err := json.Unmarshal(resp.Body(), &apps); err != nil {
		return nil, err
	}
	return apps, nil
}
