package rest

import (
	"gopkg.in/resty.v0"
	"github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"fmt"
	"encoding/json"
	"bytes"
)

type ConfHubClient struct {
	Url           string
	Username      string
	Password      string
	authorization string
}

func New(url, username, password string) *ConfHubClient {
	return &ConfHubClient{Url:url, Username:username, Password:password}
}

func (cli *ConfHubClient) Login() *ConfHubClient {
	resp, err := resty.R().SetFormData(map[string]string{
		"username":cli.Username,
		"password":cli.Password,
	}).Post(cli.Url + "/api/auth/login")
	if err != nil {
		logrus.Fatalf("Error during authentication:%s", err)
	}
	if resp.StatusCode() != 200 {
		logrus.Fatalf("Error during authentication:%s", string(resp.Body()))
	}
	cli.authorization = string(resp.Body())
	logrus.Debug(cli.authorization)
	return cli
}

func (cli *ConfHubClient) getAuthorization() string {
	if len(cli.authorization) == 0 {
		cli.Login()
	}
	return "Bearer " + cli.authorization

}

func (cli *ConfHubClient) R() *resty.Request {
	return resty.R().SetHeader("Authorization", cli.getAuthorization())
}

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

func (cli *ConfHubClient) GetConfigs(appName, appVersion, format string) (string, error) {
	resp, err := cli.R().Get(cli.Url + "/api/configs/" + appName + "/" + appVersion)
	if err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		return "", errors.New(fmt.Sprintf("Error retrieving application configs:[%d] %s", resp.StatusCode(), resp.Body()))
	}
	//TODO implements the other formats
	prettyJson:=new(bytes.Buffer)
	err = json.Indent(prettyJson,resp.Body(),"","  ")
	if err != nil {
		return "", err
	}
	return prettyJson.String(), nil
}

func (cli *ConfHubClient) SetValue(appName, appVersion, key, value string) error {
	resp, err := cli.R().SetFormData(map[string]string{
		"value": value,
	}).Put(cli.Url + "/api/values/" + appName + "/" + appVersion + "/" + key)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 204 {
		return errors.New(fmt.Sprintf("Error putting application config:[%d] %s", resp.StatusCode(), resp.Body()))
	}
	return nil

}