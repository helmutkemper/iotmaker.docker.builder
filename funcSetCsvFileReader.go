package iotmakerdockerbuilder

// SetCsvFileReader
//
// English:
//
//  Prints in the header of the file the name of the constant responsible for printing the column in
//  the log.
//
//   Input:
//     value: true to print the name of the constant responsible for printing the column in the log
//       in the header of the file.
//
// Português:
//
//  Imprime no cabeçalho do arquivo o nome da constante responsável por imprimir a coluna no log.
//
//   Entrada:
//     value: true para imprimir no cabeçalho do arquivo o nome da constante responsável por imprimir
//       a coluna no log.
func (e *ContainerBuilder) SetCsvFileReader(value bool) {
	e.csvConstHeader = value
}
