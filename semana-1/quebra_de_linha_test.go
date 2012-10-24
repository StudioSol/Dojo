package main

import (
	"github.com/orfjackal/gospec/src/gospec"
	. "github.com/orfjackal/gospec/src/gospec"
)

func QuebraLinhaTest(c gospec.Context){
    c.Specify("Caso texto seja menor q tamanho maximo", func(){
        c.Expect(QuebraLinha("ola mundo", 14), Equals, "ola mundo")
    })

    c.Specify("Caso texto seja maior q tamanho maximo", func(){
        c.Expect(QuebraLinha("lavarini lindo demais", 15), Equals, "lavarini lindo\ndemais")
    })
}
