package model

var MEDICATION_FROM_FIELD = []map[string]interface{}{
	{
		"type":     "int32",
		"optional": false,
		"field":    "id",
	},
	{
		"type":     "int32",
		"optional": true,
		"field":    "patient",
	},
	{
		"type":     "int32",
		"optional": true,
		"field":    "drug",
	},
	{
		"type":     "int32",
		"optional": false,
		"name":     "io.debezium.time.Date",
		"version":  1,
		"field":    "startDate",
	},
	{
		"type":     "int32",
		"optional": false,
		"name":     "io.debezium.time.Date",
		"version":  1,
		"field":    "endDate",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "note",
	},
}

var MEDICATION_FROM_STRUCT = map[string]interface{}{
	"type":     "struct",
	"fields":   MEDICATION_FROM_FIELD,
	"optional": true,
	"name":     "mysql_server_name.subaid_01.Prescriptions.Value",
	"field":    "before",
}

var MEDICATION_TO_FIELD = []map[string]interface{}{
	{
		"type":     "int32",
		"optional": false,
		"field":    "id",
	},
	{
		"type":     "int32",
		"optional": true,
		"field":    "patient",
	},
	{
		"type":     "int32",
		"optional": true,
		"field":    "drug",
	},
	{
		"type":     "int32",
		"optional": false,
		"name":     "io.debezium.time.Date",
		"version":  1,
		"field":    "startDate",
	},
	{
		"type":     "int32",
		"optional": false,
		"name":     "io.debezium.time.Date",
		"version":  1,
		"field":    "endDate",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "note",
	},
}

var MEDICATION_TO_STRUCT = map[string]interface{}{
	"type":     "struct",
	"fields":   MEDICATION_TO_FIELD,
	"optional": true,
	"name":     "mysql_server_name.subaid_01.Prescriptions.Value",
	"field":    "after",
}

var MEDICATION_SOURCE_FIELD = []map[string]interface{}{
	{
		"type":     "string",
		"optional": false,
		"field":    "version",
	},
	{
		"type":     "string",
		"optional": false,
		"field":    "connector",
	},
	{
		"type":     "string",
		"optional": false,
		"field":    "name",
	},
	{
		"type":     "int64",
		"optional": false,
		"field":    "ts_ms",
	},
	{
		"type":     "string",
		"optional": true,
		"name":     "io.debezium.data.Enum",
		"version":  1,
		"parameters": map[string]interface{}{
			"allowed": "true,last,false,incremental",
		},
		"default": "false",
		"field":   "snapshot",
	},
	{
		"type":     "string",
		"optional": false,
		"field":    "db",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "sequence",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "table",
	},
	{
		"type":     "int64",
		"optional": false,
		"field":    "server_id",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "gtid",
	},
	{
		"type":     "string",
		"optional": false,
		"field":    "file",
	},
	{
		"type":     "int64",
		"optional": false,
		"field":    "pos",
	},
	{
		"type":     "int32",
		"optional": false,
		"field":    "row",
	},
	{
		"type":     "int64",
		"optional": true,
		"field":    "thread",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "query",
	},
}

var MEDICATION_SOURCE_STRUCT = map[string]interface{}{
	"type":     "struct",
	"fields":   MEDICATION_SOURCE_FIELD,
	"name":     "io.debezium.connector.mysql.Source",
	"optional": false,
	"field":    "source",
}

var MEDICATION_OP_STRUCT = map[string]interface{}{
	"type":     "string",
	"optional": false,
	"field":    "op",
}

var MEDICATION_TSMS_STRUCT = map[string]interface{}{
	"type":     "int64",
	"optional": false,
	"field":    "ts_ms",
}

var MEDICATION_TRANSACTION_FIELD = []map[string]interface{}{
	{
		"type":     "string",
		"optional": false,
		"field":    "id",
	},
	{
		"type":     "int64",
		"optional": false,
		"field":    "total_order",
	},
	{
		"type":     "int64",
		"optional": false,
		"field":    "data_collection_order",
	},
}

var MEDICATION_TRANSACTION_STRUCT = map[string]interface{}{
	"type":     "struct",
	"fields":   MEDICATION_TRANSACTION_FIELD,
	"optional": true,
	"field":    "transaction",
}

var MEDICATION_SCHEMA = map[string]interface{}{
	"type": "struct",
	"fields": []map[string]interface{}{
		MEDICATION_FROM_STRUCT,
		MEDICATION_TO_STRUCT,
		MEDICATION_SOURCE_STRUCT,
		MEDICATION_OP_STRUCT,
		MEDICATION_TSMS_STRUCT,
		MEDICATION_TRANSACTION_STRUCT,
	},
	"optional": false,
	"name":     "mysql_server_name.subaid_01.Prescriptions.Envelope",
}

type MEDICATION_PAYLOAD struct {
	Before      interface{}             `json:"before"`
	After       *map[string]interface{} `json:"after"`
	Source      interface{}             `json:"source"`
	Op          interface{}             `json:"op"`
	TsMs        interface{}             `json:"ts_ms"`
	Transaction interface{}             `json:"transaction"`
}

type MedicationRequest struct {
	Schema  interface{}        `json:"schema"`
	Payload MEDICATION_PAYLOAD `json:"payload"`
}
