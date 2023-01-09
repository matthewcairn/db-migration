package model

var USER_FROM_FIELD = []map[string]interface{}{
	{
		"type":     "int32",
		"optional": false,
		"field":    "id",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "street",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "city",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "zip",
	},
	{
		"type":     "int32",
		"optional": false,
		"name":     "io.debezium.time.Date",
		"version":  1,
		"field":    "birthday",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "email",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "emergencyContactPhoneNumber",
	},
}

var USER_FROM_STRUCT = map[string]interface{}{
	"type":     "struct",
	"fields":   USER_FROM_FIELD,
	"optional": true,
	"name":     "mysql_server_name.subaid_01.Patients.Value",
	"field":    "before",
}

var USER_TO_FIELD = []map[string]interface{}{
	{
		"type":     "int32",
		"optional": false,
		"field":    "id",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "street",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "city",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "zip",
	},
	{
		"type":     "int32",
		"optional": false,
		"name":     "io.debezium.time.Date",
		"version":  1,
		"field":    "birthday",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "email",
	},
	{
		"type":     "string",
		"optional": true,
		"field":    "emergencyContactPhoneNumber",
	},
}

var USER_TO_STRUCT = map[string]interface{}{
	"type":     "struct",
	"fields":   USER_TO_FIELD,
	"optional": true,
	"name":     "mysql_server_name.subaid_01.Patients.Value",
	"field":    "after",
}

var USER_SOURCE_FIELD = []map[string]interface{}{
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

var USER_SOURCE_STRUCT = map[string]interface{}{
	"type":     "struct",
	"fields":   USER_SOURCE_FIELD,
	"name":     "io.debezium.connector.mysql.Source",
	"optional": false,
	"field":    "source",
}

var USER_OP_STRUCT = map[string]interface{}{
	"type":     "string",
	"optional": false,
	"field":    "op",
}

var USER_TSMS_STRUCT = map[string]interface{}{
	"type":     "int64",
	"optional": false,
	"field":    "ts_ms",
}

var USER_TRANSACTION_FIELD = []map[string]interface{}{
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

var USER_TRANSACTION_STRUCT = map[string]interface{}{
	"type":     "struct",
	"fields":   USER_TRANSACTION_FIELD,
	"optional": true,
	"field":    "transaction",
}

var USER_SCHEMA = map[string]interface{}{
	"type": "struct",
	"fields": []map[string]interface{}{
		USER_FROM_STRUCT,
		USER_TO_STRUCT,
		USER_SOURCE_STRUCT,
		USER_OP_STRUCT,
		USER_TSMS_STRUCT,
		USER_TRANSACTION_STRUCT,
	},
	"optional": false,
	"name":     "mysql_server_name.subaid_01.Patients.Envelope",
}

type USER_PAYLOAD struct {
	Before      interface{}             `json:"before"`
	After       *map[string]interface{} `json:"after"`
	Source      interface{}             `json:"source"`
	Op          interface{}             `json:"op"`
	TsMs        interface{}             `json:"ts_ms"`
	Transaction interface{}             `json:"transaction"`
}

type UserRequest struct {
	Schema  interface{}  `json:"schema"`
	Payload USER_PAYLOAD `json:"payload"`
}
