package main

import (
    "flag"
    "strings"
    "io/ioutil"
)

func QuebraLinha(texto string, tamanho int) string{
    if len(texto) < tamanho {
        return texto
    }

    linhas := ""
    cp := ""
    tamAtual := 0
    palavras := strings.Split(texto, " ")

    for _, palavra := range(palavras) {
        tamAtual += len(palavra)
        if tamAtual > tamanho {
            linhas += "\n"
            linhas += palavra
            tamAtual = 0
        } else {
            linhas += cp
            linhas += palavra
            tamAtual += 1
            cp = " "
        }
    }

    return linhas
}

func main() {
    flag.Parse()
    arquivo := flag.Args()[0]

    conteudo, _ := ioutil.ReadFile(arquivo)

    println(QuebraLinha(string(conteudo), 12))
}
