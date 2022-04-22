#!/bin/bash

. "${BASH_SOURCE[0]%/*}/config.sh"

mysql -h "$MYSQL_SERVICE_HOST" -u "$DB_USER" -p"$DB_PASS" <<EOF
DROP DATABASE IF EXISTS $DB_NAME;
CREATE DATABASE $DB_NAME;
USE $DB_NAME;
CREATE TABLE $DB_COLLECTION(
    uuid VARCHAR(100) primary key NOT NULL,
    content TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)
EOF
