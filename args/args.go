package args

import "flag"

type CreateArgs struct {
    data_file string
    template int
}

func ParseCreateArgs() CreateArgs {
    data_file := flag.String("file", "", "The name to greet: ")    
    template := flag.Int("template", 0, "CV template id: ")
    flag.Parse()
    return CreateArgs{data_file: *data_file, template: *template}
}
