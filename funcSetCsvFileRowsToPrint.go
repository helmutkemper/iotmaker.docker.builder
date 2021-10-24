package iotmakerdockerbuilder

// SetCsvFileRowsToPrint
//
//
//
// Português: Define quais colunas vão ser impressas no log, na forma de arquivo CSV, com informações de desempenho do container, com indicadores de consumo de memória e tempos de acesso.
//   Entrada:
//     value: Lista das colunas impressas no arquivo CSV. Ex.: KLogColumnMacOsLog, KLogColumnWindows, KLogColumnAll ou qualquer combinação de KLogColumn... concatenado com pipe, KLogColumnReadingTime | KLogColumnCurrentNumberOfOidsInTheCGroup | KLogColumnTotalCPUTimeConsumed
//
//   Nota: - Para vê a lista completa de colunas, use SetCsvFileRowsToPrint(KLogColumnAll) e SetCsvFileReader(true).
//           Isto irá imprimir os nomes das constantes em cima de cada coluna do log.
func (e *ContainerBuilder) SetCsvFileRowsToPrint(value int64) {
	e.rowsToPrint = value
}
