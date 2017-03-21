package main

func cmdControllerList() {

	if client {
		kazeList(clientsapi, name)
	}
	if check {
		kazeList(checksapi, name)
	}
	if event {
		kazeList(eventsapi, name)
	}
	if silence {
		kazeList(silencedapi, name)
	}
	if result {
		kazeList(resultsapi, name)
	}
	if aggregate {
		kazeList(aggregatesapi, name)
	}
	if stash {
		kazeList(stashesapi, name)
	}
}

func cmdControllerCreateClient() {
	kazeCreateClient()
}

func cmdControllerCreateResult() {
	kazeCreateResult()
}

func cmdControllerCreateStash() {
	kazeCreateStash()
}
