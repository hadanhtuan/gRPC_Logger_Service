package api

import (
	"LOGGER-SERVICE/common"
	"LOGGER-SERVICE/model"
	"bytes"
	"encoding/json"
	"net/http"
)

func CreateLogEntry(w http.ResponseWriter, r *http.Request) {
	var logEntry model.LogEntry

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&logEntry)
	if err != nil {
		common.WriteJSON(w, 400, &common.APIResponse{
			Status:  common.APIStatus.Invalid,
			Message: err.Error(),
		}, nil)
		return
	}

	createResult := model.LogEntryDB.Create(logEntry)

	if createResult.Status != common.APIStatus.Ok {
		common.WriteJSON(w, 400, createResult, nil)
		return
	}

	result := &common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Tạo biểu phí thành công",
	}
	common.WriteJSON(w, 400, result, nil)
}

func logRequest(name, data string) error {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}
