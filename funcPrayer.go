package iotmakerdockerbuilder

import (
	"log"
)

// Prayer (english): Programmer prayer
//
// Prayer (português): Oração do programador
func (e *ContainerBuilder) Prayer() {
	log.Print("Português:")
	log.Print("Código nosso que estais em Golang\nSantificado seja Vós, Console\nVenha a nós a Vossa Reflexão\nE seja feita a {Vossa chave}\nAssim no if(){}\nComo no else{}\nO for (nosso; de cada dia; nos dai hoje++)\nDebugai as nossas sentenças\nAssim como nós colocamos\nO ponto e vírgula esquecido;\nE não nos\n\tdeixe errar\n\t\tindentação\nMas, livreiros das funções recursivas\nA main ()")
	log.Print("")
	log.Print("English:")
	log.Print("Our program, who art in memory,\ncalled by thy name;\nthy operating system run;\nthy function be done at runtime\nas it was on development.\nGive us this day our daily output.\nAnd forgive us our code duplication,\nas we forgive those who\nduplicate code against us.\nAnd lead us not into frustration;\nbut deliver us from GOTOs.\nFor thine is algorithm,\nthe computation, and the solution,\nlooping forever and ever.\nReturn;")
	log.Print("")
}
