package utils

func ValuePayloadSourceGenerator(database string, schema string, table string) map[string]interface{} {
	return map[string]interface{}{
		"version":   "1.2.5.Final",
		"connector": "postgresql",
		"name":      "dbserver2",
		"ts_ms":     1653809880207,
		"snapshot":  "last",
		"db":        database,
		"schema":    schema,
		"table":     table,
		"txId":      491,
		"lsn":       24603048,
		"xmin":      nil,
	}
}

func KeyIDGenerator(primaryKey string, primaryValue interface{}, schema string, table string) map[string]interface{} {
	return map[string]interface{}{
		"schema": map[string]interface{}{
			"type": "struct",
			"fields": []map[string]interface{}{
				{
					"type":     "string",
					"optional": false,
					"field":    primaryKey,
					"name":     "io.debezium.data.Uuid",
					"version":  1,
				},
			},
			"optional": false,
			"name":     "dbserver2." + schema + "." + table + ".Key",
		},
		"payload": map[string]interface{}{
			primaryKey: primaryValue,
		},
	}
}

func KeyStringGenerator(primaryKey string, primaryValue interface{}, schema string, table string) map[string]interface{} {
	return map[string]interface{}{
		"schema": map[string]interface{}{
			"type": "struct",
			"fields": []map[string]interface{}{
				{
					"type":     "string",
					"optional": false,
					"field":    primaryKey,
					"version":  1,
				},
			},
			"optional": false,
			"name":     "dbserver2." + schema + "." + table + ".Key",
		},
		"payload": map[string]interface{}{
			primaryKey: primaryValue,
		},
	}
}

func ValueSchemaDetailGenerator(Field []map[string]interface{}, Topic string, Type string) map[string]interface{} {
	return map[string]interface{}{
		"type":     "struct",
		"fields":   Field,
		"optional": true,
		"name":     Topic + ".Value",
		"field":    Type,
	}
}

func ValueGenerator(topic string, schema string, table string, field []map[string]interface{}, data map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"schema": map[string]interface{}{
			"type": "struct",
			"fields": []map[string]interface{}{
				ValueSchemaDetailGenerator(field, topic, "before"),
				ValueSchemaDetailGenerator(field, topic, "after"),
				SourceValue,
				OpValue,
				TsMsValue,
				Transactionvalue,
			},
			"optional": false,
			"name":     topic + ".Envelope",
		},
		"payload": map[string]interface{}{
			"before":      nil,
			"after":       data,
			"source":      ValuePayloadSourceGenerator("postgres", schema, table),
			"op":          "r",
			"ts_ms":       1653809880212,
			"transaction": nil,
		},
	}
}
