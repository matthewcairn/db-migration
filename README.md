# Data migration from MySQL to MongoDB


### Turn on Binlog
#### MySQL version 8
```
SET GLOBAL enforce_gtid_consistency=ON;
SET GLOBAL gtid_mode=OFF_PERMISSIVE;
SET GLOBAL gtid_mode=ON_PERMISSIVE;
SET GLOBAL gtid_mode=ON;
SET GLOBAL binlog_rows_query_log_events=ON;
```

#### MySQL version 5 (Window)
1. Press Window + R and type `services.msc`
2. Identify `MySQL` service, then right click and choose `Properties`
- Find the config file of mySQL 
- The path of config file is below the section of `Path to executable`
- The file of config should be `my.ini` or `my.conf`
3. Update that file content
- Adding below the section `[mysqld]`
```
log-bin=bin.log
log-bin-index=bin-log.index
max_binlog_size=100M
binlog_format=row
socket=mysql.sock
```
4. Restart the mySQL service
At step `2` location, choose `MySQL` service, then the will be an option of restart on the left pannel


### Start Kafka CDC service
Run `docker-compose up --build`

### Then call API to Kafka to start Kafka CDC
All APIs are included in `Metaroom.postman_collection.json`