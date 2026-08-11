#!/bin/bash

set -e

# cari Path .env
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"

ENV_FILE="$PROJECT_ROOT/.env"
MIGRATION_DIR="$PROJECT_ROOT/migration"


# load .env
if [ ! -f "$ENV_FILE" ]; then
    echo "Error: .env file not found!"
    echo "Expected: $ENV_FILE"
    exit 1
fi

source "$ENV_FILE"

# validasi .env variable
# : "${DB_HOST:?DB_HOST is not set}" # Karna pake docker, jadi tidak perlu 
# : "${DB_PORT:?DB_PORT is not set}"
: "${DB_NAME:?DB_NAME is not set}"
: "${DB_USERNAME:?DB_USERNAME is not set}"
: "${DB_PASSWORD:?DB_PASSWORD is not set}"

# command
mysql_cmd=(
    sudo docker compose exec -T mysql mysql
    # --host="$DB_HOST" # Krna pake docker, jadi tidak perlu
    # --port="$DB_PORT"
    --user="$DB_USERNAME"
    --password="$DB_PASSWORD"
)


echo "Checking database: $DB_NAME"

"${mysql_cmd[@]}" <<SQL
CREATE DATABASE IF NOT EXISTS \`$DB_NAME\`
    CHARACTER SET utf8mb4
    COLLATE utf8mb4_unicode_ci;
SQL

echo "[SUCCESS] Database is ready."


mysql_db_cmd=(
    sudo docker compose exec -T mysql mysql
    # --host="$DB_HOST"
    # --port="$DB_PORT"
    --user="$DB_USERNAME"
    --password="$DB_PASSWORD"
    "$DB_NAME"
)


# Create Migration
echo "[Log] Checking migration table..."

"${mysql_db_cmd[@]}" < \
    "$MIGRATION_DIR/000_create_schema_migrations.sql"


echo
echo "Running migrations..."
echo

for file in "$MIGRATION_DIR"/*.sql; do

    migration="$(basename "$file")"

    # Skip migration table
    if [ "$migration" = "000_create_schema_migrations.sql" ]; then
        continue
    fi


    # Checking : Migration sudah di execute?
    exists=$(
        "${mysql_db_cmd[@]}" \
            --batch \
            --skip-column-names \
            -e "
                SELECT COUNT(*)
                FROM schema_migrations
                WHERE migration = '$migration';
            "
    )

    if [ "$exists" -eq 1 ]; then
        echo "Already migrated: $migration"
        continue
    fi


    # Execute migration
    echo "Migrating: $migration"

    "${mysql_db_cmd[@]}" < "$file"


    # Record migration
    "${mysql_db_cmd[@]}" -e "
        INSERT INTO schema_migrations (migration)
        VALUES ('$migration');
    "


    echo "Completed: $migration"
done

echo
echo "[SUCCESS] All migrations completed successfully."