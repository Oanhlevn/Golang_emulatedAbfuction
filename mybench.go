
package main
import ( "fmt" 
 "flag" 
 "os"
 "io"
 "io/ioutil"
 "net/http"
 "time"
)
type responseInfor struct {
	status int 
	bytes int64
	duration time.Duration 
}
func main () {
	fmt.Println ("Hello from my app")
	requests:= flag.Int64( "n", 1, "Number of requests to perform")
	concurrency:= flag.Int64("c", 1, "Number of multiple requests to make at a time" )
	fmt.Println(requests,concurrency)
	flag.Parse()
	if  flag.NArg()==0 || *requests==0 || *requests < *concurrency {
	flag.PrintDefaults()
	os.Exit(-1)
	}
	link := flag.Arg(0)
	c:= make (chan responseInfor)
	for i:= int64(0); i< *concurrency; i++ {
		go checkLink(link, c)
	}
	for response:= range c {
		fmt.Println(response ) 
	}
	
}

func checkLink(Link string, c chan responseInfor)  {
	start:= time.Now()
	res, err := http.Get(Link)
	if err != nil {
		panic(err)			
	}
	read, _ := io.Copy(ioutil.Discard, res.Body )

	c <- responseInfor{
		status : res.StatusCode,
		bytes : read,
		duration : time.Now().Sub(start),
	}
}

