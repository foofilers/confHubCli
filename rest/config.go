package rest

import (
	"fmt"
	"encoding/json"
	"errors"
)

func (cli *ConfHubClient) GetConfigs(appName, appVersion string) (ConfigMap, error) {
	if len(appVersion) > 0 {
		appVersion = "/" + appVersion
	}
	resp, err := cli.R().Get(cli.Url + "/api/configs/" + appName + appVersion)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, errors.New(fmt.Sprintf("Error retrieving application configs:[%d] %s", resp.StatusCode(), resp.Body()))
	}
	res := make(ConfigMap)
	err = json.Unmarshal(resp.Body(), &res)
	return res, err
}

func (cli *ConfHubClient) GetFormattedConfigs(appName, appVersion, format string) (string, error) {
	if len(appVersion) > 0 {
		appVersion = "/" + appVersion
	}
	resp, err := cli.R().SetQueryParam("format", format).Get(cli.Url + "/api/configs/" + appName + appVersion)
	if err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		return "", errors.New(fmt.Sprintf("Error retrieving application configs:[%d] %s", resp.StatusCode(), resp.Body()))
	}
	return string(resp.Body()), nil
}
