package restclient

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const authHeader = "2e974946b51a860152fa87c00250215d47449aecc2256261e8244167c80374dd"

func DoGet(url string) (string, error){
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Errorf("error %w", err)
		return "", err
	}

	req.Header.Add("X-Auth-Token", authHeader)
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

func DoPut(url string, body []byte) (string, error){
	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req.Header.Add("Content-Type","application/json")
	req.Header.Add("Accept","application/json")
	req.Header.Add("X-Auth-Token", authHeader)

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	// Read Response Body
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	/*if resp.StatusCode != 200 {
		return "", errors.New("Codigo de respuesta invalido: " + resp.Status)
	}*/
	/*var response = ""
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan(); i++ {
		response +=  scanner.Text()
	}*/
	return string(response), nil
}

func DoDelete(url string, body []byte) error{

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return err
	}

	req.Header.Add("X-Auth-Token", authHeader)

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	// Read Response Body
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("Codigo de respuesta invalido: " + resp.Status)
	}
	return nil

}