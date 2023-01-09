USE subaid_01;
CREATE USER 'debezium'@'%' IDENTIFIED WITH mysql_native_password BY 'debezium_password';
GRANT ALL PRIVILEGES on *.* TO 'debezium'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
;

;
SET GLOBAL enforce_gtid_consistency=ON;
SET GLOBAL gtid_mode=OFF_PERMISSIVE;
SET GLOBAL gtid_mode=ON_PERMISSIVE;
SET GLOBAL gtid_mode=ON;
SET GLOBAL binlog_rows_query_log_events=ON;