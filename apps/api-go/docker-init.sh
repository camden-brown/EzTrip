#!/bin/sh
set -e

# Run migrations
echo "Running migrations..."
for file in /docker-entrypoint-initdb.d/migrations/*.sql; do
  if [ -f "$file" ]; then
    echo "Executing $file"
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" -f "$file"
  fi
done

# Run seeds (optional, only in development)
if [ "$SEED_DATA" = "true" ]; then
  echo "Running seeds..."
  for file in /docker-entrypoint-initdb.d/seeds/*.sql; do
    if [ -f "$file" ]; then
      echo "Executing $file"
      psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" -f "$file"
    fi
  done
fi

echo "Database initialization complete"
