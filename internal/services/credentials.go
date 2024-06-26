package services

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/broodbear/cave/internal/models"
	"github.com/jedib0t/go-pretty/v6/table"
)

var ErrMissingFields = errors.New("not enough fields in file")

const minNumFields = 4

type Credentials struct {
	project          string
	credentialsStore credentialsStorer
}

type credentialsStorer interface {
	Save(project, target, username, password string) error
	All() ([]models.Record, error)
}

func NewCredentials(project string, cs credentialsStorer) Credentials {
	return Credentials{
		project:          project,
		credentialsStore: cs,
	}
}

func (c Credentials) Add() error {
	var (
		target   string
		username string
		password string
	)

	fmt.Println("All fields are optional")

	fmt.Print("Target: ")
	fmt.Scan(&target)

	fmt.Print("Username: ")
	fmt.Scan(&username)

	fmt.Print("Password: ")
	fmt.Scan(&password)

	return c.credentialsStore.Save(c.project, target, username, password)
}

func (c Credentials) Export(filename, sep string) error {
	records, err := c.credentialsStore.All()
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, record := range records {
		_, err = file.WriteString(record.Project + sep + record.Target + sep + record.Username + sep + record.Password + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func (c Credentials) Import(filename, sep string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, sep)
		if len(parts) < minNumFields {
			return ErrMissingFields
		}

		err = c.credentialsStore.Save(
			parts[0],
			parts[1],
			parts[2],
			parts[3],
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c Credentials) Print() error {
	records, err := c.credentialsStore.All()
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Project", "Target", "Username", "Password", "Created"})

	for _, record := range records {
		t.AppendRow([]interface{}{
			record.ID,
			record.Project,
			record.Target,
			record.Username,
			record.Password,
			record.CreatedAt,
		})
	}

	t.Render()

	return nil
}
