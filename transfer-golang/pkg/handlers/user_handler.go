package handlers

import (
	"encoding/json"
	"log"
	"transfer/pkg/model"
	"transfer/pkg/repo"
	"transfer/pkg/utils"
)

type UserHandlers struct {
	kafkaProducer *repo.KafkaProducer
}

func NewUserHandlers(kafkaProducer *repo.KafkaProducer) *UserHandlers {
	return &UserHandlers{kafkaProducer: kafkaProducer}
}

func (h *UserHandlers) TransformData(Key []byte, Value []byte) (err error) {
	var rawInput map[string]interface{}
	json.Unmarshal(Value, &rawInput)
	var rawPayload map[string]interface{}
	rawPayload = rawInput["payload"].(map[string]interface{})

	tempAfterData := utils.GetKeyOrNil(rawPayload, "after")
	var afterDataPointer *map[string]interface{}
	if tempAfterData != nil {
		parsedTempAfterData := tempAfterData.(map[string]interface{})
		afterData := make(map[string]interface{})
		afterData["id"] = utils.GetKeyOrNil(parsedTempAfterData, "id")
		afterData["street"] = utils.GetKeyOrNil(parsedTempAfterData, "street")
		afterData["city"] = utils.GetKeyOrNil(parsedTempAfterData, "town")
		afterData["zip"] = utils.GetKeyOrNil(parsedTempAfterData, "postcode")
		afterData["birthday"] = utils.GetKeyOrNil(parsedTempAfterData, "dob")
		afterData["email"] = utils.GetKeyOrNil(parsedTempAfterData, "email")
		afterData["emergencyContactPhoneNumber"] = utils.GetKeyOrNil(parsedTempAfterData, "mobile")
		afterDataPointer = &afterData
	} else {
		afterDataPointer = nil
	}

	beforeData := utils.GetKeyOrNil(rawPayload, "before")
	sourceData := utils.GetKeyOrNil(rawPayload, "source")
	opData := utils.GetKeyOrNil(rawPayload, "op")
	tsMsData := utils.GetKeyOrNil(rawPayload, "ts_ms")
	transactionData := utils.GetKeyOrNil(rawPayload, "transaction")

	var userRequest model.UserRequest
	json.Unmarshal(Value, &userRequest)
	userRequest.Schema = model.USER_SCHEMA
	userRequest.Payload.After = afterDataPointer
	userRequest.Payload.Before = beforeData
	userRequest.Payload.Source = &sourceData
	userRequest.Payload.Op = opData
	userRequest.Payload.TsMs = tsMsData
	userRequest.Payload.Transaction = transactionData

	byteValue, err := json.Marshal(userRequest)
	if err != nil {
		log.Println("Got error " + err.Error())
	}
	h.kafkaProducer.PublishSimple("mysql_server_name.subaid_01.Patients_SINK", string(Key), string(byteValue))
	return nil
}
