# cave

Use a bear cave to keep track of your target.

## installation

```bash
wget https://github.com/broodbear/cave/releases/latest/download/cave_$(uname)_$(uname -m).tar.gz
tar -xf cave_$(uname)_$(uname -m).tar.gz
sudo install ./cave /usr/local/bin
```

## usage instructions

### initialize

Initialize a new database.

#### Initialize database

```bash
cave i
```

#### Manually create database

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

### print

The following command will print out all the credentials from the database to the
terminal.

```bash
cave credentials print
```

### import

You may already have a file containing some data you want to import.

The import command will look for a file named `cave.csv` and use a comma as a
separator.

```bash
cave credentials import
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

## Contributing

Contributions are welcomed. You will need to install the following.

- [gofumpt](https://github.com/mvdan/gofumpt)
- [golangci-lint](https://golangci-lint.run/)
- [sqlite3](https://www.sqlite.org/download.html)
- [goreleaser](https://goreleaser.com/)

```bash
go install mvdan.cc/gofumpt@latest
go install github.com/goreleaser/goreleaser/v2@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1

sudo apt install sqlite3
```
