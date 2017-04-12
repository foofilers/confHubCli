package rest

import (
	"fmt"
	"errors"
	"encoding/json"
)

func (cli *ConfHubClient) SetDefaultVersion(appName, appVersion string) error {
	resp, err := cli.R().Put(cli.Url + "/api/versions/" + appName + "/" + appVersion)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 204 {
		return errors.New(fmt.Sprintf("Error setting default version:[%d] %s", resp.StatusCode(), resp.Body()))
	}
	return nil
}

func (cli *ConfHubClient) AddVersion(appName, appVersion string) error {
	resp, err := cli.R().SetFormData(map[string]string{
		"version":appVersion,
	}).Post(cli.Url + "/api/versions/" + appName)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 201 {
		return errors.New(fmt.Sprintf("Error setting default version:[%d] %s", resp.StatusCode(), resp.Body()))
	}
	return nil
}

func (cli *ConfHubClient)CopyVersion(appName, srcVersion, dstVersion string) error {
	resp, err := cli.R().SetFormData(map[string]string{
		"version":dstVersion,
	}).Put(cli.Url + "/api/versions/" + appName + "/" + srcVersion + "/copy")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New(fmt.Sprintf("Error copy version from %s to %s:[%d] %s", srcVersion, dstVersion, resp.StatusCode(), resp.Body()))
	}
	return nil
}

func (cli *ConfHubClient) GetVersions(appName string) ([]string, error) {
	resp, err := cli.R().Get(cli.Url + "/api/versions/" + appName)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, errors.New(fmt.Sprintf("Error getting versions:[%d] %s", resp.StatusCode(), resp.Body()))
	}
	versions := make([]string, 0)
	err = json.Unmarshal(resp.Body(), &versions)
	return versions, err
}
