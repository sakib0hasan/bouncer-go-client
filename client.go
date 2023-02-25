package bouncer

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strconv"
)

type Client interface {
	GetAvailableCredit() (int, error)
	VerifyEmail(email string, timeout int) (VerifyEmailResponse, error)
}

type bouncerService struct {
	restyClient *resty.Client
}

const DefaultTimeoutSecond = 10

func NewClient(apiKey string) Client {
	_client := bouncerService{
		restyClient: resty.New(),
	}
	_client.restyClient.SetBaseURL("https://api.usebouncer.com/v1/")
	_client.restyClient.SetHeader("x-api-key", apiKey)
	_client.restyClient.SetHeader("Content-Type", "application/json")
	return &_client
}

func (service *bouncerService) GetAvailableCredit() (int, error) {
	var result GetCreditResponse
	resp, err := service.restyClient.R().
		SetResult(&result).
		Get("customer/user/current/credit")
	if err != nil {
		return 0, err
	}
	if resp.StatusCode() != http.StatusOK {
		return 0, errors.New(string(resp.Body()))
	}
	return result.Credits, nil
}

func (service *bouncerService) VerifyEmail(email string, timeout int) (VerifyEmailResponse, error) {
	var result VerifyEmailResponse
	resp, err := service.restyClient.R().
		SetResult(&result).
		Get("email/verify?email=" + email + "&timeout=" + strconv.Itoa(timeout))
	if err != nil {
		return result, err
	}
	if resp.StatusCode() != http.StatusOK {
		return result, errors.New(string(resp.Body()))
	}
	return result, nil
}
