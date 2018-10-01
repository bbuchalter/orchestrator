#!/bin/bash

if [ -n "$DEBUG" ]; then
  set -x
  export PS4='+(${BASH_SOURCE}:${LINENO}): ${FUNCNAME[0]:+${FUNCNAME[0]}(): }'
fi

MYSQL_VERSION=5.6.41
MYSQL_USER=orch_dev
MYSQL_PASS=orch_dev_pass
MYSQL_HOST=127.0.0.1
MYSQL_PORT=33307
SANDBOX_DESCRIPTION="dev"

BASE_PATH=$(dirname $0)
DBD_PATH=$BASE_PATH/../../dbdeployer
DBD_SETUP_LOG=$BASE_PATH/mysql_frontend_setup.log
CONF_PATH=conf/orchestrator.conf.json

ORCH_USER=$(cat $CONF_PATH | grep MySQLTopologyUser | tr -d \",: | awk '{print $2}')
if [ -z "$ORCH_USER" ]; then
    echo "Could not find MySQLTopologyUser in $CONF_PATH"
    exit 1
fi

ORCH_PASS=$(cat $CONF_PATH | grep MySQLTopologyPassword | tr -d \",: | awk '{print $2}')
if [ -z "$ORCH_PASS" ]; then
    echo "Could not find MySQLTopologyPassword in $CONF_PATH"
    exit 1
fi

echo "Setup MySQL with DB Deployer..."
echo "Step 1: Prepare DB Deployer..."
echo "Logs for this are stored in $DBD_SETUP_LOG"
MYSQL_VERSION=$MYSQL_VERSION DESCRIPTION=$SANDBOX_DESCRIPTION $DBD_PATH/setup.sh > $DBD_SETUP_LOG

if [ $? -ne 0 ]; then
    echo "Preparing DB Deployer failed"
    exit 1
fi

echo "Step 2: Deploy MySQL instances..."
DBD_CMD=$(grep -A 1 "Use the following command to work with this setup" $DBD_SETUP_LOG | tail -n 1)
if [ -z "$DBD_CMD" ]; then
    echo "Could not find the DB deployer command in $DBD_SETUP_LOG"
    exit 1
fi
$DBD_CMD deploy replication $MYSQL_VERSION \
    --nodes 3 \
    --force \
    --db-user $MYSQL_USER \
    --db-password $MYSQL_PASS \
    --port $MYSQL_PORT \
    --bind-address $MYSQL_HOST \
    --post-grants-sql "CREATE USER '$ORCH_USER'@'%' IDENTIFIED BY '$ORCH_PASS';" \
    --post-grants-sql "GRANT SUPER, PROCESS, REPLICATION SLAVE, RELOAD ON *.* TO '$ORCH_USER'@'%';" \
    --bind-address 0.0.0.0 \
    --skip-report-host \
    --my-cnf-options "report-host=localhost"

if [ $? -ne 0 ]; then
    echo "Deploying MySQL with dbdeployer failed"
    exit 1
fi

STATUS=$($DBD_CMD global status)
if [ $? -ne 0 ]; then
    echo "Could not get the status of the deployed MySQL instance"
    exit 1
fi

MASTER_PORT=$(echo "${STATUS}" | grep "master on" | awk '{print $7}')
if [ -z "$MASTER_PORT" ]; then
    echo "Could not find the port of the master node. Is it running?"
    exit 1
fi
echo $MASTER_PORT

echo "You are ready to discover your nodes!"
echo "Start Orchestrator and visit http://localhost:3000/web/discover."
echo "Enter 'localhost' as the hostname and and '${MASTER_PORT}' as the port."