# cave

Use a bear cave to keep track of your target.

## installation

```bash
go install github.com/broodbear/cave/cmd/cave@latest
```

## usage instructions

### initialize

Initialize a new database.

```bash
cat <<EOT > ./migration.sql
create table
  'cave' (
    'id' integer not null primary key autoincrement,
    'created_at' datetime not null default CURRENT_TIMESTAMP,
    'project' varchar(255) null,
    'target' varchar(255) null,
    'username' varchar(255) null,
    'password' varchar(255) null,
    unique ('id')
  );
EOT

sqlite3 default.db < migration.sql
```

### add

The following command will prompt you for the fields then add those to the database.

```bash
cave credentials add
```

### import

You may already have a file containing some data you want to import.

The import command will look for a file named `cave.csv` and use a comma as a
separator.

```bash
./bin/cave credentials import
```

You can change the input filename and the separator.

```bash
cave credentials import --separator "^" --filename "export.csv"
```

### export

The export command will export all the credentials to a file named `cave.csv`
using a comma as the separator.

```bash
cave credentials export
```

You can change the destination filename and the separator.

```bash
cave credentials export --separator "^" --filename "export.csv"
```
