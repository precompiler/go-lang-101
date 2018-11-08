package main

import (
	"github.com/precompiler/go-lang-101/dp/builder"
	"github.com/precompiler/go-lang-101/dp/singleton"
)

func main() {

	singleton.GetInstance().CreateConnection()
	singleton.GetInstance().CloseConnection()

	mybuilder := new(builder.ConnectionBuilder)
	mybuilder.SetDBUrl("jdbc:dummy").SetUsername("rich").SetPassword("p@55w0rd")
	mybuilder.Build()

}
