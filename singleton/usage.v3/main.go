package main

import (
	appV3 "github.com/daniloanp/gophers-rj-meetup/singleton/app.v3.extending"
	appV2 "github.com/daniloanp/gophers-rj-meetup/singleton/app.v2"
	"fmt"
)

func main() {
	v2  := appV2.Instance()
	v3  := appV3.Instance()
	fmt.Println("Esta é a segunda versão, chamando calls pela primeira vez", v2.Calls())
	fmt.Println("Esta é a última, mas já foi chamada um vez:", v3.Calls())
	//Resultado: v3 compartilha estados de v2.
	//Não parece uma boa prática.
}
