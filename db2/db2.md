# Using DB2

Ssh into the db2 host
Change to the db2 instance user and setup profile
```
su - {database_instance_name}
. sqllib/db2profile
```
Start db2:
```
db2start
```
Create the ODRDERDB database:
```
db2 create database ORDERDB
# on specific path
db2 create database INDB on /db2
```
Connect to the ORDERDB database:
```
# connect:
db2 connect to ORDERDB

# disconnect:
db2 connect reset
```

## useful commands
```
# list all dbs:
db2 list db directory

# List all tables:
db2 list tables for all

# To list all tables in selected schema, use:
db2 list tables for schema <schema-name>
db2 list tables for schema CSUSER

# To describe a table, type:
db2 describe table <table-schema.table-name>
// For example:
db2 describe table CSUSER.WIKI_PAGES

# List information about databases
# List databases:
db2 list database directory show detail | grep -B6 -i indirect | grep "Database name"

# If you just want the database names, without the titles, use:
db2 list database directory show detail | grep -B6 -i indirect | grep "Database name" | sed "s/.*= //"

# queries:
db2 "select * from DB2INST1.PRODUCT WHERE PRODUCT_ID > 1200"
db2 "DELETE from DB2INST1.PRODUCT WHERE PRODUCT_ID > 500"
```

## Teardown database:
```
db2 connect to ODERDB
db2 force application all
db2 terminate
db2 deactivate db ORDERDB
db2 uncatalog  db ORDERDB
db2 connect to INDB
db2 force application all
db2 terminate
db2 deactivate db INDB
db2 uncatalog  db INDB

db2stop
```

## Bring DB2 up from mount point /db2
```
db2 catalog db ORDERDB on /db2
db2 catalog db INDB on /db2
db2start
```

## Find the DB2 port
https://www.ibm.com/support/knowledgecenter/en/SSPQ7D_4.1.0/com.ibm.saam.doc_4.1/tatune_determinedb2port.html

Issue the command
```
db2 get dbm cfg
# Go to the following line:

TCP/IP Service name                          (SVCENAME) = db2c_db2inst1

# then find the port:
cat /etc/services | grep db2c_db2inst1
db2c_db2inst1      50001/tcp
```

## Run the shell script invoking DB2

```sh
#!/bin/bash
# run some stuff as root
# then switch to db2inst1 instance

su - db2inst1 -c '. /home/db2inst1/sqllib/db2profile; db2 catalog db ORDERDB on /db2;db2 catalog db INDB on /db2;db2start'
```
