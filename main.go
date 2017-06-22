package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	//load config
	config = Cfg()

	// api endpoints
	apibase          = "http://" + config.Sensu + ":" + config.Port
	clientsapi       = apibase + "/clients"
	checksapi        = apibase + "/checks"
	requestapi       = apibase + "/request"
	resultsapi       = apibase + "/results"
	aggregatesapi    = apibase + "/aggregates"
	eventsapi        = apibase + "/events"
	silencedapi      = apibase + "/silenced"
	silencedapiclear = apibase + "/silenced/clear"
	stashesapi       = apibase + "/stashes"
	healthapi        = apibase + "/health"
	infoapi          = apibase + "/info"
	resolveapi       = apibase + "/resolve"

	// generic
	client    bool
	check     bool
	event     bool
	silence   bool
	result    bool
	aggregate bool
	stash     bool
	name      stringArray
	file      string
	all       bool
	// config command specific
	address string
	port    string
	// createClient command specific
	clientAddress       string
	clientSubscriptions stringArray
	clientEnvironment   string
	// delete command specific
	deleteCheckName string
	// createResult command specific
	resultSource string
	resultOutput string
	resultStatus int
	// create stash command specific
	stashPath   string
	stashExpire int
	// silence command specific
	silenceClear           bool
	silenceList            bool
	silenceSubscription    bool
	silenceCheckName       string
	silenceExpire          int
	silenceExpireOnResolve bool
	silenceReason          string
	silenceCreator         string
	// check & resolve command specific
	checkTarget      stringArray
	checkAll         bool
	resolveCheckName string
)

type stringArray []string

func (array *stringArray) String() string {
	return fmt.Sprintf("%v", *array)
}

func (array *stringArray) Set(value string) error {
	*array = strings.Split(value, ",")
	return nil
}

func main() {

	// help subcommand
	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)
	// help flags
	helpCmd.Bool("help", false, "get help on the help command?")

	// configure subcommand
	configCmd := flag.NewFlagSet("configure", flag.ExitOnError)
	// configure flags
	configCmd.StringVar(&address, "address", "", "specify the address of the sensu-api")
	configCmd.StringVar(&port, "port", "", "specify the port used by the sensu-api (defaults to 4567 when empty)")

	// list subcommand
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	// list flags
	listCmd.BoolVar(&client, "client", false, "use to list client(s)")
	listCmd.BoolVar(&check, "check", false, "use to list check(s)")
	listCmd.BoolVar(&event, "event", false, "use to list event(s)")
	listCmd.BoolVar(&silence, "silence", false, "use to list silence entr(y)(ies)")
	listCmd.BoolVar(&result, "result", false, "use to list result(s)")
	listCmd.BoolVar(&aggregate, "aggregate", false, "use to list aggregate(s)")
	listCmd.BoolVar(&stash, "stash", false, "use to list stash(es)")
	listCmd.Var(&name, "name", "specify the name(s) of the object(s) to list")

	// createClient subcommand
	createClientCmd := flag.NewFlagSet("create-client", flag.ExitOnError)
	// createClient flags
	createClientCmd.StringVar(&file, "file", "", "a valid json file for creation of objects, for bulk operations. if specified it will override all other arguments")
	createClientCmd.Var(&name, "name", "name of client to create")
	createClientCmd.StringVar(&clientAddress, "address", "", "address of the client to create")
	createClientCmd.Var(&clientSubscriptions, "subscriptions", "subscriptions of the client to create, comma sperated")
	createClientCmd.StringVar(&clientEnvironment, "environment", "", "content of the stash to create, json formatted")

	// createResult subcommand
	createResultCmd := flag.NewFlagSet("create-result", flag.ExitOnError)
	// createClient flags
	createResultCmd.StringVar(&file, "file", "", "a valid json file for creation of objects, for bulk operations. if specified it will override all other arguments")
	createResultCmd.Var(&name, "name", "name of the result check to create")
	createResultCmd.StringVar(&resultSource, "source", "", "source of the result")
	createResultCmd.StringVar(&resultOutput, "output", "", "output of the result")
	createResultCmd.IntVar(&resultStatus, "status", 0, "statuscode of the result")

	// createResult subcommand
	createStashCmd := flag.NewFlagSet("create-stash", flag.ExitOnError)
	// createClient flags
	createStashCmd.StringVar(&file, "file", "", "a valid json file for creation of objects, for bulk operations. if specified it will override all other arguments")
	createStashCmd.StringVar(&stashPath, "path", "", "path of the stash to create/update")
	createStashCmd.IntVar(&stashExpire, "expire", -1, "TTL of the stash in seconds")

	// delete subcommand
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	// delete flags
	deleteCmd.BoolVar(&client, "client", false, "use to delete client(s)")
	deleteCmd.BoolVar(&event, "event", false, "use to delete event(s)")
	deleteCmd.BoolVar(&result, "result", false, "use to delete result(s)")
	deleteCmd.BoolVar(&aggregate, "aggregate", false, "use to delete aggregate(s)")
	deleteCmd.BoolVar(&stash, "stash", false, "use to delete stash(es)")
	deleteCmd.Var(&name, "name", "specify the name(s) of the object(s) to delete")
	deleteCmd.StringVar(&deleteCheckName, "check-name", "", "use for specifying the check when deleting a result entry. In this scenario also use the -name flag to specify the client to delete the result from")

	// silence subcommand
	silenceCmd := flag.NewFlagSet("silence", flag.ExitOnError)
	// silence flags
	silenceCmd.BoolVar(&client, "client", false, "use to target client(s)")
	silenceCmd.BoolVar(&silenceSubscription, "subscription", false, "use to target subscription(s)")
	silenceCmd.StringVar(&silenceCheckName, "check-name", "", "specify the name of the check you want to silence")
	silenceCmd.Var(&name, "name", "specify the name(s) of the client(s) or subscription(s)")
	silenceCmd.BoolVar(&all, "all", false, "use to target all silenced entries. if specified it will override other arguments.")
	silenceCmd.BoolVar(&silenceExpireOnResolve, "expire-on-resolve", false, "sets the silenced entry to expire once check is resolved")
	silenceCmd.StringVar(&silenceReason, "reason", "", "specify reason for silencing")
	silenceCmd.StringVar(&silenceCreator, "creator", "", "specify the creator of the silence entry")
	silenceCmd.IntVar(&silenceExpire, "expire", 0, "specify when the silence should expire (seconds)")

	// silence-clear subcommand
	silenceClearCmd := flag.NewFlagSet("silence-clear", flag.ExitOnError)
	// silence-clear flags
	silenceClearCmd.BoolVar(&client, "client", false, "use to target client(s)")
	silenceClearCmd.BoolVar(&silenceSubscription, "subscription", false, "use to target subscription(s)")
	silenceClearCmd.StringVar(&silenceCheckName, "check-name", "", "specify the name of the check you want to target")
	silenceClearCmd.Var(&name, "name", "specify the name(s) of the client(s) or subscription(s)")
	silenceClearCmd.BoolVar(&all, "all", false, "use to target all silenced entries. if specified it will override other arguments.")

	//check subcommand
	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)
	//check flags
	checkCmd.Var(&checkTarget, "target", "specify a target (if targeting a client specify as 'client:<name>'). if not specfied the check will target its default subscribers")
	checkCmd.Var(&name, "name", "specify the name of the check")
	checkCmd.BoolVar(&all, "all", false, "use to target all checks, overrides other flags")
	checkCmd.BoolVar(&result, "result", false, "use to get the result back from the requested check")

	//resolve subcommand
	resolveCmd := flag.NewFlagSet("resolve", flag.ExitOnError)
	//resolve flags
	resolveCmd.Var(&name, "client-name", "specify the name of the client")
	resolveCmd.StringVar(&resolveCheckName, "check-name", "", "specify the name of the check you want to resolve")
	//resolveCmd.StringVar(&checkName, "check-name", "", "specify the name of the check")
	resolveCmd.BoolVar(&all, "all", false, "use to target all events")

	if len(os.Args) < 2 {
		help()
	}

	//switch on subcommand
	switch os.Args[1] {
	case "configure":
		configCmd.Parse(os.Args[2:])
	case "list":
		listCmd.Parse(os.Args[2:])
	case "create-client":
		createClientCmd.Parse(os.Args[2:])
	case "create-result":
		createResultCmd.Parse(os.Args[2:])
	case "create-stash":
		createStashCmd.Parse(os.Args[2:])
	case "delete":
		deleteCmd.Parse(os.Args[2:])
	case "silence":
		silenceCmd.Parse(os.Args[2:])
	case "clear-silence":
		silenceClearCmd.Parse(os.Args[2:])
	case "check":
		checkCmd.Parse(os.Args[2:])
	case "resolve":
		resolveCmd.Parse(os.Args[2:])
	case "help":
		helpCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	//required flag check and function handling.

	if configCmd.Parsed() {
		if configCmd.NFlag() < 1 {
			usagePrint()
			configCmd.PrintDefaults()
		}
		if configCmd.NFlag() >= 1 && address != "" {
			cmdControllerConfigure()
		}
	}

	if listCmd.Parsed() {
		if listCmd.NFlag() < 1 {
			usagePrint()
			listCmd.PrintDefaults()
		}
		if listCmd.NFlag() <= 2 {
			if listCmd.NFlag() == 2 && len(name) == 0 {
				trowError("no name specified or too many arguments given. Only select one type ( e.g. --client ) or specify a name with --name")
			}
			cmdControllerList()
		} else {
			trowError("too many arguments given, expecting 2 or less.")
		}
	}

	if createClientCmd.Parsed() {
		if createClientCmd.NFlag() < 1 {
			usagePrint()
			createClientCmd.PrintDefaults()
		}
		if createClientCmd.NFlag() == 1 && file == "" {
			trowError("not enough arguments given, if -file flag is not used then 4 arguments are expected.")
		}
		if createClientCmd.NFlag() == 1 && file != "" {
			cmdControllerCreateClient()
		}
		if createClientCmd.NFlag() == 4 {
			cmdControllerCreateClient()
		}
	}

	if createResultCmd.Parsed() {
		if createResultCmd.NFlag() < 1 {
			usagePrint()
			createResultCmd.PrintDefaults()
		}
		if createResultCmd.NFlag() >= 1 {
			cmdControllerCreateResult()
		}
	}

	if createStashCmd.Parsed() {
		if createStashCmd.NFlag() < 1 {
			usagePrint()
			createStashCmd.PrintDefaults()
		}
		if createStashCmd.NFlag() >= 1 {
			cmdControllerCreateStash()
		}
	}

	if deleteCmd.Parsed() {
		if deleteCmd.NFlag() < 1 {
			usagePrint()
			deleteCmd.PrintDefaults()
		}
		if deleteCmd.NFlag() <= 3 {
			if deleteCmd.NFlag() == 3 && len(name) == 0 && deleteCheckName == "" {
				trowError("no name or check-name specified or too many arguments given. Only select one type ( e.g. -client ) or specify a name with -name. incase of deleting a result use 3 arguments ( e.g. -result, -name & -check-name.")
			}
			if deleteCmd.NFlag() == 2 && len(name) == 0 {
				trowError("no name specified or too many arguments given. Only select one type ( e.g. -client ) or specify a name with -name. incase of deleting a result use 3 arguments ( e.g. -result, -name & -check-name.")
			}
			if deleteCmd.NFlag() == 2 && result {
				trowError("incase of deleting a result use 3 arguments (-result, -name and -check-name).")
			}
			if deleteCmd.NFlag() == 2 && len(name) != 0 {
				cmdControllerDelete()
			}
			if deleteCmd.NFlag() == 1 && len(name) == 0 && deleteCheckName == "" {
				cmdControllerDelete()
			}
			if deleteCmd.NFlag() == 3 && len(name) != 0 && deleteCheckName != "" {
				cmdControllerDelete()
			}
		}
		if deleteCmd.NFlag() >= 4 {
			trowError("too many arguments given, expecting 3 or less.")
		}
	}

	if silenceCmd.Parsed() {
		if silenceCmd.NFlag() < 1 {
			usagePrint()
			silenceCmd.PrintDefaults()
		}
		if silenceCmd.NFlag() >= 1 {
			if client && silenceSubscription {
				trowError("cannot combine -client and -subscription flag.")
			}
			if len(name) != 0 && client || silenceSubscription {
				cmdControllerSilence()
			}
			if all && !client && !silenceSubscription {
				cmdControllerSilence()
			}
		}
	}

	if silenceClearCmd.Parsed() {
		if silenceClearCmd.NFlag() < 1 {
			usagePrint()
			silenceClearCmd.PrintDefaults()
		}
		if silenceClearCmd.NFlag() >= 1 {
			if client && silenceSubscription {
				trowError("cannot combine -client and -subscription flag.")
			}
			if len(name) != 0 && client || silenceSubscription {
				cmdControllerSilenceClear()
			}
			if all {
				cmdControllerSilenceClear()
			}
		}
	}

	if checkCmd.Parsed() {
		if checkCmd.NFlag() < 1 {
			usagePrint()
			checkCmd.PrintDefaults()
		}
		if checkCmd.NFlag() >= 1 {
			if len(name) == 0 && all == false {
				trowError("flag -name is required.")
			}
			cmdControllerCheck()
		}
	}

	if resolveCmd.Parsed() {
		if checkCmd.NFlag() < 1 {
			usagePrint()
			resolveCmd.PrintDefaults()
		}
		if resolveCmd.NFlag() > 2 {
			trowError("too many arguments given, expecting 3 or less.")
		}
		if resolveCmd.NFlag() == 2 && len(name) != 0 && resolveCheckName != "" {
			cmdControllerResolve()
		}
		if resolveCmd.NFlag() == 1 && all {
			cmdControllerResolve()
		}
	}

	if helpCmd.Parsed() {
		if helpCmd.NFlag() < 1 {
			help()
		} else {
			fmt.Printf("We're no strangers to love\nYou know the rules and so do I\nA full commitment's what I'm thinking of\nYou wouldn't get this from any other guy\nI just want to tell you how I'm feeling\nGotta make you understand\n\n")
			fmt.Printf("Never gonna give you up\nNever gonna let you down\nNever gonna run around and desert you\nNever gonna make you cry\nNever gonna say goodbye\nNever gonna tell a lie and hurt you\n")
		}
	}
}
