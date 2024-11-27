# PostgreSQL Configuration Environment Variables:
# - PGUSER: Username to connect to the PostgreSQL database.
# - PGPASSWORD: Password for the specified PostgreSQL user.
# - PGHOST: Host address of the PostgreSQL server (e.g., localhost or an IP address).

TIME=$(date -Iminutes)

while read DBNAME; do
	FILE_NAME="$DBNAME-$TIME.dump"
	echo "Dump database $DBNAME, File: $FILE_NAME"
	pg_dump -Fc -d $DBNAME -f $FILE_NAME
done
