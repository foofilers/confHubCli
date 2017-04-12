package rest

import (
	"gopkg.in/resty.v0"
	"github.com/Sirupsen/logrus"
)

type ConfHubClient struct {
	Url           string
	Username      string
	Password      string
	authorization string
}

func NewConfHubClient(url, username, password string) *ConfHubClient {
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



