// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

type LogCustom struct {
	*LogRole
	TableName string
}

func NewLogCustom(logRole *LogRole, tableName string) *LogCustom {
	return &LogCustom{LogRole: logRole, TableName: tableName}
}
