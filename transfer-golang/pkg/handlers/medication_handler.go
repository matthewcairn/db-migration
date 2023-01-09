package handlers

import (
	"encoding/json"
	"log"
	"transfer/pkg/model"
	"transfer/pkg/repo"
	"transfer/pkg/utils"
)

type MedicationHandlers struct {
	kafkaProducer *repo.KafkaProducer
}

func NewMedicationHandlers(kafkaProducer *repo.KafkaProducer) *MedicationHandlers {
	return &MedicationHandlers{kafkaProducer: kafkaProducer}
}

func (h *MedicationHandlers) TransformData(Key []byte, Value []byte) (err error) {
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
		afterData["patient"] = utils.GetKeyOrNil(parsedTempAfterData, "patients_id")
		afterData["drug"] = utils.GetKeyOrNil(parsedTempAfterData, "drug_id")
		afterData["startDate"] = utils.GetKeyOrNil(parsedTempAfterData, "pssttimestamp")
		afterData["endDate"] = utils.GetKeyOrNil(parsedTempAfterData, "psedtimestamp")
		afterData["note"] = utils.GetKeyOrNil(parsedTempAfterData, "comment")
		afterDataPointer = &afterData
	} else {
		afterDataPointer = nil
	}

	beforeData := utils.GetKeyOrNil(rawPayload, "before")
	sourceData := utils.GetKeyOrNil(rawPayload, "source")
	opData := utils.GetKeyOrNil(rawPayload, "op")
	tsMsData := utils.GetKeyOrNil(rawPayload, "ts_ms")
	transactionData := utils.GetKeyOrNil(rawPayload, "transaction")

	var medicationRequest model.MedicationRequest
	json.Unmarshal(Value, &medicationRequest)
	medicationRequest.Schema = model.MEDICATION_SCHEMA
	medicationRequest.Payload.After = afterDataPointer
	medicationRequest.Payload.Before = beforeData
	medicationRequest.Payload.Source = &sourceData
	medicationRequest.Payload.Op = opData
	medicationRequest.Payload.TsMs = tsMsData
	medicationRequest.Payload.Transaction = transactionData

	byteValue, err := json.Marshal(medicationRequest)
	if err != nil {
		log.Println("Got error " + err.Error())
	}
	h.kafkaProducer.PublishSimple("mysql_server_name.subaid_01.Prescriptions_SINK", string(Key), string(byteValue))
	return nil
}
