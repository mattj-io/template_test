package main

import (
    "os"
    "log"
    "html/template"
    "github.com/Masterminds/sprig"
)

func main()  {

    type Params struct {
        Count    string
        Name     string
    }

    tpl, err := template.New("test").Funcs(sprig.FuncMap()).Parse(test_var)
    if err != nil {
        log.Fatalln("Could not create Template: " + err.Error())
    }

    params := Params{
        Count: "3",
        Name: "test",
    }

    if err = tpl.Execute(os.Stdout, params); err != nil {
        log.Fatalln("Could not execute template" + err.Error())
    }
      
}

