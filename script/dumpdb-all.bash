PDB_DUMP_ALL_CONTEXT_DIR=${PDB_DUMP_ALL_CONTEXT_DIR:-/bitnami/postgresql/dump}
PDB_DUMP_ALL_DB_LIST=${PDB_DUMP_ALL_DB_LIST:-db-list.txt}
PDB_DUMP_ALL_SCRIPT_NAME=${PDB_DUMP_ALL_SCRIPT_NAME:-dumpdb.bash}

export PGUSER=${PGUSER:-postgres}
export PGPASSWORD=${PGPASSWORD:-postgres}
export PGHOST=${PGHOST:-postgresql.authority.svc.cluster.local}

pushd "$PDB_DUMP_ALL_CONTEXT_DIR"

echo "Context directory: $PDB_DUMP_ALL_CONTEXT_DIR"
cat "$PDB_DUMP_ALL_DB_LIST" | bash "$PDB_DUMP_ALL_SCRIPT_NAME"

popd
