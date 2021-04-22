package main

import (
	"github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/util"
)

func New() (referenceInicialized *SQLiteMenu, err error) {
	referenceInicialized = &SQLiteMenu{}
	err = referenceInicialized.Connect(constants.KSQLiteConnectionString)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = referenceInicialized.Install()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
