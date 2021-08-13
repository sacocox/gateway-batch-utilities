package restclient

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const furyToken = "2e974946b51a860152fa87c00250215d47449aecc2256261e8244167c80374dd"

func DoGet(url string) (string, error){
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Errorf("error %w", err)
		return "", err
	}

	req.Header.Add("X-Auth-Token", furyToken)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Errorf("error %w", err)
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("Codigo de respuesta invalido: " + resp.Status)
	}

	var response = ""
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan(); i++ {
		response +=  scanner.Text()
	}
	return response, nil
}


func DoPost(url string, body []byte) (string, error){
	return "", nil
}

func DoPut(url string, json string) error{

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(json))

	if err != nil {
		return fmt.Errorf("Error updating transaction ->: %s", err.Error())
	}
	req.Header.Add("x-caller-scopes", "admin")
	req.Header.Add("x-auth-token", furyToken)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error updating transaction ->: %s", err.Error())
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error parsing response: %s", err.Error())
	}
	return nil
}

func DoDelete(url string, body []byte) error{
	return nil
}