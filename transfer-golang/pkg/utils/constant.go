package utils

const (
	HeaderXRequestID = "x-request-id"
	HeaderUserID     = "x-user-id"
	HeaderUserMeta   = "x-user-type"
)

const (
	OBJECT_USER       = "user"
	OBJECT_ROLE       = "role"
	OBJECT_PERMISSION = "permission"
)

var SourceValue = map[string]interface{}{
	"type": "struct",
	"fields": []map[string]interface{}{
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
			"optional": true,
			"field":    "ts_ms",
		},
		{
			"type":     "string",
			"optional": true,
			"name":     "io.debezium.data.Enum",
			"version":  1,
			"parameters": map[string]string{
				"allowed": "true,last,false",
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
			"optional": false,
			"field":    "schema",
		},
		{
			"type":     "string",
			"optional": false,
			"field":    "table",
		},
		{
			"type":     "int64",
			"optional": true,
			"field":    "txId",
		},
		{
			"type":     "int64",
			"optional": true,
			"field":    "lsn",
		},
		{
			"type":     "int64",
			"optional": true,
			"field":    "xmin",
		},
	},
	"optional": false,
	"name":     "io.debezium.connector.postgresql.Source",
	"field":    "source",
}

var OpValue = map[string]interface{}{
	"type":     "string",
	"optional": false,
	"field":    "op",
}

var TsMsValue = map[string]interface{}{
	"type":     "int64",
	"optional": true,
	"field":    "ts_ms",
}

var Transactionvalue = map[string]interface{}{
	"type": "struct",
	"fields": []map[string]interface{}{
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
	},
	"optional": true,
	"field":    "transaction",
}
