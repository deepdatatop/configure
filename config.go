//	read key/value parameter from .conf file
//	params := LoadConfigure("c:\\crawler\\stacker.conf")
//	if params!=nil {
//		dispatcher_server = params["dispatcher"]
//		dirWorking 		  = params["workingdir"]
//	}

	
package config
import(
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

var BOM []byte = []byte{'\xEF','\xBB','\xBF'}  //BOM(Byte Order Mark)
func LoadConfigure(filename string) map[string]string{
	mapParam := make( map[string]string )
	file, err := os.Open(filename)
	if err != nil { return nil }
	defer file.Close()
	
	buf := bufio.NewReader(file)

	for{
		line, _, err := buf.ReadLine()
		ne := bytes.Trim( line," \t" )
		ln := bytes.Trim( ne,string(BOM) )
		if err == io.EOF { break }
		if !bytes.HasPrefix( ln,[]byte{'#'} ) && !bytes.Equal( ln,[]byte{} ) {
			parts := bytes.SplitN(ln, []byte{'='}, 2)
			if len(parts)==2 {
				key := strings.TrimSpace(string(parts[0]))
				val := strings.Trim(strings.TrimSpace(string(parts[1])),"\"")
				mapParam[key] = val
			}
		}
	}
	return mapParam
}

/*
import(
	"deepdata.top/config"
)
func main(){
	file, _ := exec.LookPath(os.Args[0])
    path, _ := filepath.Abs(file)
	dir,filename := filepath.Split(path)
	
//	println("path:",dir)
//	println("filename:",filename)

	dispatcher_server := ""
	dirWorking := ""	
	params := config.LoadConfigure(dir+"stacker.conf")
	if params!=nil {
		dispatcher_server = params["dispatcher"]
		dirWorking 		  = params["workingdir"]
	}
	fmt.Println( dispatcher_server,dirWorking )
}*/