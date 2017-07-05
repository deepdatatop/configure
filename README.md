# configure
read key=value configure file

import(
	"deepdata.top/config"
)
func main(){
	file, _ := exec.LookPath(os.Args[0])
  path, _ := filepath.Abs(file)
	dir,filename := filepath.Split(path)
	
	println("path:",dir)
	println("filename:",filename)

	dispatcher_server := ""
	dirWorking := ""	
	params := config.LoadConfigure(dir+"stacker.conf")
	if params!=nil {
		dispatcher_server = params["dispatcher"]
		dirWorking 		  = params["workingdir"]
	}
	fmt.Println( dispatcher_server,dirWorking )
}


# configure file sample

#user manual
#
workingdir	= /dat/crawler
dispatcher	=	"10.100.16.20:7777"
