#!/bin/bash

# Script to restore PostgreSQL database from dump file

# Set variables
PG_USER="postgres"  # Change this if necessary
PG_DB="cloudbridge" # Change this if necessary
DUMP_FILE="cloudbridge_dump.sql" # Change this if necessary

# Check if dump file exists
if [ ! -f "$DUMP_FILE" ]; then
    echo "Dump file '$DUMP_FILE' not found."
    exit 1
fi

# Restore the database
echo "Restoring database '$PG_DB'..."
psql -U "$PG_USER" -d "$PG_DB" -f "$DUMP_FILE"

# Check if restoration was successful
if [ $? -eq 0 ]; then
    echo "Database restored successfully."
else
    echo "Error: Database restoration failed."
fi
