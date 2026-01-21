package migrations

import "embed"

// for when we compile to binary - file structure for SQL files
//go:embed *.sql
var FS embed.FS
