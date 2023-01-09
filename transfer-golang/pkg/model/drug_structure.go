package model

var DRUG_FROM_FIELD = []map[string]interface{}{
	{
		"type":     "int32",
		"optional": false,
		"field":    "id",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "alias",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "shortName",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "name",
	},
	{
		"type":     "int32",
		"optional": true,
		"field":    "stock",
	},
}

var DRUG_FROM_STRUCT = map[string]interface{}{
	"type":     "struct",
	"fields":   DRUG_FROM_FIELD,
	"optional": true,
	"name":     "mysql_server_name.subaid_01.drug.Value",
	"field":    "before",
}

var DRUG_TO_FIELD = []map[string]interface{}{
	{
		"type":     "int32",
		"optional": false,
		"field":    "id",
	},

	{
		"type":     "string",
		"optional": true,
		"field":    "alias",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "shortName",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "name",
	},
	{
		"type":     "int32",
		"optional": true,
		"field":    "stock",
	},
}

var DRUG_TO_STRUCT = map[string]interface{}{
	"type":     "struct",
	"fields":   DRUG_TO_FIELD,
	"optional": true,
	"name":     "mysql_server_name.subaid_01.drug.Value",
	"field":    "after",
}

var DRUG_SOURCE_FIELD = []map[string]interface{}{
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

var DRUG_SOURCE_STRUCT = map[string]interface{}{
	"type":     "struct",
	"fields":   DRUG_SOURCE_FIELD,
	"name":     "io.debezium.connector.mysql.Source",
	"optional": false,
	"field":    "source",
}

var DRUG_OP_STRUCT = map[string]interface{}{
	"type":     "string",
	"optional": false,
	"field":    "op",
}

var DRUG_TSMS_STRUCT = map[string]interface{}{
	"type":     "int64",
	"optional": true,
	"field":    "ts_ms",
}

var DRUG_TRANSACTION_FIELD = []map[string]interface{}{
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

var DRUG_TRANSACTION_STRUCT = map[string]interface{}{
	"type":     "struct",
	"fields":   DRUG_TRANSACTION_FIELD,
	"optional": true,
	"field":    "transaction",
}

var DRUG_SCHEMA = map[string]interface{}{
	"type": "struct",
	"fields": []map[string]interface{}{
		DRUG_FROM_STRUCT,
		DRUG_TO_STRUCT,
		DRUG_SOURCE_STRUCT,
		DRUG_OP_STRUCT,
		DRUG_TSMS_STRUCT,
		DRUG_TRANSACTION_STRUCT,
	},
	"optional": false,
	"name":     "mysql_server_name.subaid_01.drug.Envelope",
}

type DRUG_PAYLOAD struct {
	Before      interface{}             `json:"before"`
	After       *map[string]interface{} `json:"after"`
	Source      interface{}             `json:"source"`
	Op          interface{}             `json:"op"`
	TsMs        interface{}             `json:"ts_ms"`
	Transaction interface{}             `json:"transaction"`
}

type DrugRequest struct {
	Schema  interface{}  `json:"schema"`
	Payload DRUG_PAYLOAD `json:"payload"`
}
