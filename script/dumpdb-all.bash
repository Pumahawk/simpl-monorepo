CONTEXT_DIR=/bitnami/postgresql/dump
DB_LIST=db-list.txt
SCRIPT_NAME=dumpdb.bash

export PGUSER=postgres
export PGPASSWORD=postgres
export PGHOST=postgresql.authority.svc.cluster.local

pushd "$CONTEXT_DIR"

echo "Context directory: $CONTEXT_DIR"
cat "$DB_LIST" | bash "$SCRIPT_NAME"

popd
