package dup_checker

import (
	"encoding/json"
	"errors"
	"fmt"
	restclient "github.com/mercadolibre/gateway_batch_utilities/internal/resclient"
)

type Key struct {
	ID    string `json:"id"`
	Value struct {
		Token string `json:"token"`
	} `json:"value"`
	CreationDate string `json:"creation_date"`
}

type DupCheckerServiceI interface {
	CreateKey(key Key) (string, error)
	GetKey(key string) (string, error)
	DeleteKey(key Key) (string, error)
}

type DupCheckerService struct {
	urlApi string
}

func NewDupCheckerService(urlApi string) DupCheckerService {
	return DupCheckerService{urlApi: urlApi}
}

func (dcs DupCheckerService) CreateKey(key Key) (string, error) {
	return "", nil
}

func (dcs DupCheckerService) GetKey(key string) (string, error) {

	url := dcs.buildUrl(key)

	response, err := restclient.DoGet(url)
	if err != nil {
		fmt.Errorf("error obteniendo el key: %w", err)
		return "", err
	}

	return response, nil
}

func (dcs DupCheckerService) DeleteKey(key Key) (string, error) {
	url := dcs.buildUrl(key.ID)

	body, err := json.Marshal(key.Value)

	if err != nil {
		fmt.Errorf("error formateando el body: %w", err)
		return "", err
	}

	err = restclient.DoDelete(url, body)

	if err != nil {
		return "", errors.New("No pudo ser borrada el key: " + err.Error())
	}

	return "", nil
}

func (dcs DupCheckerService) buildUrl(key string) string{
	return dcs.urlApi + "/v1/key/" + key
}