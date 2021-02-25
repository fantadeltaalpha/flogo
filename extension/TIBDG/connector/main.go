package main

import (
	"fmt"

	"github.com/fantadeltaalpha/flogo/extension/TIBDG/connector/datagrid"
)

func main() {
	var c interface{}

	c = &datagrid.ClientManager{GridName: "test",RealmURL: "http://:8080"}

	fmt.Printf("%T\n", c) 

	cm,ok := c.(*datagrid.ClientManager)
	
	fmt.Println(ok)
	fmt.Println(cm.GridName)
}