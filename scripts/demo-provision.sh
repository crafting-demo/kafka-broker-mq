#!/bin/bash

CURRENT_DATE=$(date '+%Y%m%d')

[[ -n "$ORG_NAME" ]] || {
	fatal "Missing org name, eg: ORG_NAME=\"xyz\" ./demo-provision.sh"
	exit 1
}
[[ -n "$SANDBOX_NAME" ]] || {
	fatal "Missing sandbox name, eg: SANDBOX_NAME=\"xyz\" ./demo-provision.sh"
	exit 1
}

# Prepare workspaces via setup scripts
cs ssh 'kafka/scripts/setup-workspace.sh' -O "$ORG_NAME" -W "$SANDBOX_NAME/kafka"
cs ssh 'backend/scripts/setup-workspace.sh' -O "$ORG_NAME" -W "$SANDBOX_NAME/django"
cs ssh 'backend/scripts/setup-workspace.sh' -O "$ORG_NAME" -W "$SANDBOX_NAME/rails"

# Restart all daemons
cs restart -O "$ORG_NAME" -W "$SANDBOX_NAME/kafka"
cs restart -O "$ORG_NAME" -W "$SANDBOX_NAME/django"
cs restart -O "$ORG_NAME" -W "$SANDBOX_NAME/rails"

# Prepare home snapshots files/folders to include
cs ssh 'mkdir -p /home/owner/.snapshot; touch /home/owner/.snapshot/includes.txt; echo ".cache/yarn" > /home/owner/.snapshot/includes.txt' -O "$ORG_NAME" -W "$SANDBOX_NAME/react"
cs ssh 'mkdir -p /home/owner/.snapshot; touch /home/owner/.snapshot/includes.txt; echo ".gems" > /home/owner/.snapshot/includes.txt' -O "$ORG_NAME" -W "$SANDBOX_NAME/rails"
cs ssh 'mkdir -p /home/owner/.snapshot; touch /home/owner/.snapshot/includes.txt; echo "kafka/broker" > /home/owner/.snapshot/includes.txt' -O "$ORG_NAME" -W "$SANDBOX_NAME/kafka"

# Only create snapshots in org if explicitly specified
[[ -n "$WITH_SNAPSHOTS" ]] || exit 0

# Create base snapshots
cs snapshot create base/demo-rails/"$CURRENT_DATE" -O "$ORG_NAME" -W "$SANDBOX_NAME/rails"
cs snapshot create base/demo-django/"$CURRENT_DATE" -O "$ORG_NAME" -W "$SANDBOX_NAME/django"

# Create home snapshots
cs snapshot create --home home/demo-react/"$CURRENT_DATE" -O "$ORG_NAME" -W "$SANDBOX_NAME/react"
cs snapshot create --home home/demo-rails/"$CURRENT_DATE" -O "$ORG_NAME" -W "$SANDBOX_NAME/rails"
cs snapshot create --home home/demo-kafka/"$CURRENT_DATE" -O "$ORG_NAME" -W "$SANDBOX_NAME/kafka"

# Create dependency snapshots
cs snapshot create mysql/demo/"$CURRENT_DATE" -O "$ORG_NAME" -W "$SANDBOX_NAME/mysql"
