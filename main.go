package main

import (
	"flag"
	"os"
)

var (
	//load config
	config = Cfg()

	// api endpoints
	apibase       = "http://" + config.Sensu + ":" + config.Port
	clientsapi    = apibase + "/clients"
	checksapi     = apibase + "/checks"
	resultsapi    = apibase + "/results"
	aggregatesapi = apibase + "/aggregates"
	eventsapi     = apibase + "/events"
	silencedapi   = apibase + "/silenced"
	stashesapi    = apibase + "/stashes"
	healthapi     = apibase + "/health"
	infoapi       = apibase + "/info"

	// generic
	client    bool
	check     bool
	event     bool
	silence   bool
	result    bool
	aggregate bool
	stash     bool
	name      string
	file      string
	// create command specific
	createAddress       string
	createSubscriptions string
	createEnvironment   string
	createContent       string
	createSource        string
	createOutput        string
	createStatus        int
	// silence command specific
	silenceClear        bool
	silenceList         bool
	silenceSubscription bool
	// check & resolve command specific
	checkName string
	checkAll  bool
)

func main() {

	// list subcommand
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	// list flags
	listCmd.BoolVar(&client, "client", false, "use to list client(s)")
	listCmd.BoolVar(&check, "check", false, "use to list check(s)")
	listCmd.BoolVar(&event, "event", false, "use to list event(s)")
	listCmd.BoolVar(&silence, "silence", false, "use to list silence entr(y)(ies)")
	listCmd.BoolVar(&result, "result", false, "use to list result(s)")
	listCmd.BoolVar(&aggregate, "aggregate", false, "use to ist aggregate(s)")
	listCmd.BoolVar(&stash, "stash", false, "use to list stash(es)")
	listCmd.StringVar(&name, "name", "", "specify the name(s) of the object(s) to list")

	// create subcommand
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	// create flags
	createCmd.BoolVar(&client, "client", false, "use to create client(s)")
	createCmd.BoolVar(&result, "result", false, "use to create result(s)")
	createCmd.BoolVar(&stash, "stash", false, "use to create stash(es)")
	createCmd.StringVar(&file, "file", "", "a valid json file for creation of objects, for bulk operations")
	createCmd.StringVar(&name, "name", "", "name of client/stash/result to create")
	createCmd.StringVar(&createAddress, "client-address", "", "address of the client to create")
	createCmd.StringVar(&createSubscriptions, "client-subscriptions", "", "subscriptions of the client to create, comma sperated")
	createCmd.StringVar(&createContent, "stash-content", "", "content of the stash to create, json formatted")
	createCmd.StringVar(&createSource, "result-source", "", "source of the result")
	createCmd.StringVar(&createOutput, "result-output", "", "output of the result")
	createCmd.IntVar(&createStatus, "result-status", 0, "status of the result")

	// delete subcommand
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	// delete flags
	deleteCmd.BoolVar(&client, "client", false, "use to delete client(s)")
	deleteCmd.BoolVar(&event, "event", false, "use to delete event(s)")
	deleteCmd.BoolVar(&result, "result", false, "use to delete result(s)")
	deleteCmd.BoolVar(&aggregate, "aggregate", false, "use to delete aggregate(s)")
	deleteCmd.BoolVar(&stash, "stash", false, "use to delete stash(es)")
	deleteCmd.StringVar(&name, "name", "", "specify the name(s) of the object(s) to delete")

	// silence subcommand
	silenceCmd := flag.NewFlagSet("silence", flag.ExitOnError)
	// silence flags
	silenceCmd.BoolVar(&silenceClear, "clear", false, "use to clear silenced entr(y)(ies)")
	silenceCmd.BoolVar(&silenceList, "list", false, "use to list silenced entr(y(ies)")
	silenceCmd.BoolVar(&client, "client", false, "use to target client(s)")
	silenceCmd.BoolVar(&silenceSubscription, "subscription", false, "use to target subscription(s)")
	silenceCmd.StringVar(&name, "name", "", "specify the name(s) of the client(s) or subscription(s)")

	//check subcommand
	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)
	//check flags
	checkCmd.StringVar(&name, "client-name", "", "specify the name of the client")
	checkCmd.StringVar(&checkName, "check-name", "", "specify the name of the check")
	checkCmd.BoolVar(&checkAll, "all", false, "use to target all checks")
	checkCmd.BoolVar(&result, "result", false, "use to get the result back from the requested check")

	//resolve subcommand
	resolveCmd := flag.NewFlagSet("check", flag.ExitOnError)
	//resolve flags
	resolveCmd.StringVar(&name, "client-name", "", "specify the name of the client")
	resolveCmd.StringVar(&checkName, "check-name", "", "specify the name of the check")
	resolveCmd.BoolVar(&checkAll, "all", false, "use to target all events")

	switch os.Args[1] {
	case "list":
		listCmd.Parse(os.Args[2:])
	case "create":
		createCmd.Parse(os.Args[2:])
	case "delete":
		deleteCmd.Parse(os.Args[2:])
	case "silence":
		silenceCmd.Parse(os.Args[2:])
	case "check":
		checkCmd.Parse(os.Args[2:])
	case "resolve":
		resolveCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if listCmd.Parsed() {
		if listCmd.NFlag() <= 2 {
			if listCmd.NFlag() == 2 && name == "" {
				trowError("no name specified or too many arguments given. Only select one type ( e.g. --client ) or specify a name with --name")
			}
			cmdControllerList()
		} else {
			trowError("too many arguments given, expecting 2 or less.")
		}
	}

	if createCmd.Parsed() {
		if file == "" {
			trowError("no file specified")
		}
		if createCmd.NFlag() == 2 && file != "" {
			cmdControllerCreate()
		}
		if createCmd.NFlag() < 2 {
			trowError("too few arguments given, expecting 2 ( e.g. --client and --file ).")
		}
		if createCmd.NFlag() > 2 {
			trowError("too many arguments given, expecting 2.")
		}
	}

	// if createCmd.Parsed() {
	// 	cmdControllerCreate()
	// }
	// if deleteCmd.Parsed() {
	// 	cmdControllerDelete()
	// }
	// if silenceCmd.Parsed() {
	// 	cmdControllerSilence()
	// }
	// if checkCmd.Parsed() {
	// 	cmdControllerCheck()
	// }
	// if checkCmd.Parsed() {
	// 	cmdControllerResolve()
	// }

}
