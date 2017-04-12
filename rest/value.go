package rest

import (
	"fmt"
	"errors"
)

func (cli *ConfHubClient) SetValue(appName, appVersion, key, value string) error {
	resp, err := cli.R().SetFormData(map[string]string{
		"value": value,
	}).Put(cli.Url + "/api/values/" + appName + "/" + appVersion + "/" + key)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 204 {
		return errors.New(fmt.Sprintf("Error putting application value:[%d] %s", resp.StatusCode(), resp.Body()))
	}
	return nil
}

func (cli *ConfHubClient) DeleteValue(appName, appVersion, key string) error {
	resp, err := cli.R().Delete(cli.Url + "/api/values/" + appName + "/" + appVersion + "/" + key)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 204 {
		return errors.New(fmt.Sprintf("Error deleting application value:[%d] %s", resp.StatusCode(), resp.Body()))
	}
	return nil
}