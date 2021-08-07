package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//general response function
func response(responseData interface{}, statusCode int, w http.ResponseWriter) error {
	//response data
	var respData struct{
		Success		bool			//is api finished succesfully?
		Data 		interface{}		//spesific response data
	}

	//status code check for respData.Success
	if statusCode == 200 {
		respData.Success = true
	} else {
		respData.Success = false
	}
	respData.Data = responseData

	//marshall the spesific response data
	jsonResponse, err := json.Marshal(respData)
	if err != nil {
		return err
	}

	//set response header and body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(jsonResponse)
	return err
}

//for read request body and set to spesific struct
func BodyToJson(r *http.Request, data interface{}) error {

	//read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	//unmarshall the bıdy into our spesific reqBody
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	return nil
}


func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}