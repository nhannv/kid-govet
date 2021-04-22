package main

import (
	"github.com/nhannv/kid-govet/apiAuditLogs"
	"github.com/nhannv/kid-govet/configtelemetry"
	"github.com/nhannv/kid-govet/emptyStrCmp"
	"github.com/nhannv/kid-govet/equalLenAsserts"
	"github.com/nhannv/kid-govet/errorAssertions"
	"github.com/nhannv/kid-govet/errorVars"
	"github.com/nhannv/kid-govet/errorVarsName"
	"github.com/nhannv/kid-govet/immut"
	"github.com/nhannv/kid-govet/inconsistentReceiverName"
	"github.com/nhannv/kid-govet/license"
	"github.com/nhannv/kid-govet/openApiSync"
	"github.com/nhannv/kid-govet/rawSql"
	"github.com/nhannv/kid-govet/structuredLogging"
	"github.com/nhannv/kid-govet/tFatal"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(
		license.Analyzer,
		license.EEAnalyzer,
		structuredLogging.Analyzer,
		// appErrorWhere.Analyzer,
		tFatal.Analyzer,
		equalLenAsserts.Analyzer,
		openApiSync.Analyzer,
		rawSql.Analyzer,
		inconsistentReceiverName.Analyzer,
		apiAuditLogs.Analyzer,
		immut.Analyzer,
		emptyStrCmp.Analyzer,
		configtelemetry.Analyzer,
		errorAssertions.Analyzer,
		errorVarsName.Analyzer,
		errorVars.Analyzer,
	)
}
