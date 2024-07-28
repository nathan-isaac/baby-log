package schema

import "embed"

//go:embed *.sql
var EmbedMigrations embed.FS

func ListMigrations() ([]string, error) {
	files, err := EmbedMigrations.ReadDir(".")
	if err != nil {
		return nil, err
	}

	var names []string

	for _, file := range files {
		names = append(names, file.Name())
	}

	return names, nil
}
