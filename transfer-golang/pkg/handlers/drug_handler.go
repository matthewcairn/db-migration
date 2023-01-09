package handlers

import (
	"encoding/json"
	"log"
	"transfer/pkg/model"
	"transfer/pkg/repo"
	"transfer/pkg/utils"
)

type DrugHandlers struct {
	kafkaProducer *repo.KafkaProducer
}

func NewDrugHandlers(kafkaProducer *repo.KafkaProducer) *DrugHandlers {
	return &DrugHandlers{kafkaProducer: kafkaProducer}
}

func (h *DrugHandlers) TransformData(Key []byte, Value []byte) (err error) {
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
		afterData["name"] = utils.GetKeyOrNil(parsedTempAfterData, "name")
		afterData["alias"] = utils.GetKeyOrNil(parsedTempAfterData, "aliasname")
		afterData["shortName"] = utils.GetKeyOrNil(parsedTempAfterData, "shortname")
		afterData["stock"] = utils.GetKeyOrNil(parsedTempAfterData, "maxstock")
		afterDataPointer = &afterData
	} else {
		afterDataPointer = nil
	}

	beforeData := utils.GetKeyOrNil(rawPayload, "before")
	sourceData := utils.GetKeyOrNil(rawPayload, "source")
	opData := utils.GetKeyOrNil(rawPayload, "op")
	tsMsData := utils.GetKeyOrNil(rawPayload, "ts_ms")
	transactionData := utils.GetKeyOrNil(rawPayload, "transaction")

	var drugRequest model.DrugRequest
	json.Unmarshal(Value, &drugRequest)
	drugRequest.Schema = model.DRUG_SCHEMA
	drugRequest.Payload.After = afterDataPointer
	drugRequest.Payload.Before = beforeData
	drugRequest.Payload.Source = &sourceData
	drugRequest.Payload.Op = opData
	drugRequest.Payload.TsMs = tsMsData
	drugRequest.Payload.Transaction = transactionData

	byteValue, err := json.Marshal(drugRequest)
	if err != nil {
		log.Println("Got error " + err.Error())
	}
	h.kafkaProducer.PublishSimple("mysql_server_name.subaid_01.Drug_SINK", string(Key), string(byteValue))
	return nil
}
