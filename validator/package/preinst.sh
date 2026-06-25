#!/bin/bash

set -e

SERVICE_USER=sila-validator

# Create the service account, if needed
getent passwd $SERVICE_USER > /dev/null || useradd -s /bin/false --no-create-home --system --user-group $SERVICE_USER

# Create directories
mkdir -p /etc/sila
mkdir -p /var/lib/sila
install -d -m 0700 -o $SERVICE_USER -g $SERVICE_USER /var/lib/sila/validator