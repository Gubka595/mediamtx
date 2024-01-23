package storage

import (
	"github.com/bluenviron/mediamtx/internal/conf"
	"github.com/bluenviron/mediamtx/internal/storage/psql"
)

type Storage struct {
	Use              bool
	Req              psql.Requests
	DbDrives         bool
	DbUseCodeMP      bool
	UseDbPathStream  bool
	UseUpdaterStatus bool
	Sql              conf.Sql
}
