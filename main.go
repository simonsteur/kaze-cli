package main

var (
	config = Cfg()

	// flag vars
	pxyClientName string
)

func main() {

	// app := cli.NewApp()
	// app.Name = "sensuamplo"
	// app.Version = "0.1"
	// app.Usage = "control sensu from a cli"
	// app.EnableBashCompletion = true

	// app.Commands = []cli.Command{
	// 	{
	//     Name: "proxy-client"
	//     Usage: "use to add a proxy client to sensu"
	//     Flags: []cli.Flag{
	//       cli.StringFlag{
	//         Name: "name"
	//         Usage: "name of the proxy client"
	//         Destination &pxyClientName
	//       }
	//     }
	//   }

	createNewProxyClient()
}
