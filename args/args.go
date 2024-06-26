package args

import "flag"

type CreateArgs struct {
    DataFile string
    Template uint
}

func ParseCreateArgs() *CreateArgs {
    data_file := flag.String("file", "", "The name to greet: ")    
    template := flag.Uint("template", 0, "CV template id: ")
    flag.Parse()
    return &CreateArgs{DataFile: *data_file, Template: *template}
}
