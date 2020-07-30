package model

import (
	"time"

	v1 "github.com/anchore/grype-db/pkg/db/v1"
)

const (
	IDTableName = "id"
)

type IDModel struct {
	BuildTimestamp string `gorm:"column:build_timestamp"`
	SchemaVersion  int    `gorm:"column:schema_version"`
}

func NewIDModel(id v1.ID) IDModel {
	return IDModel{
		BuildTimestamp: id.BuildTimestamp.Format(time.RFC3339Nano),
		SchemaVersion:  id.SchemaVersion,
	}
}

func (IDModel) TableName() string {
	return IDTableName
}

func (m *IDModel) Inflate() v1.ID {
	buildTime, err := time.Parse(time.RFC3339Nano, m.BuildTimestamp)
	if err != nil {
		// TODO: just no...
		panic(err)
	}
	return v1.ID{
		BuildTimestamp: buildTime,
		SchemaVersion:  m.SchemaVersion,
	}
}