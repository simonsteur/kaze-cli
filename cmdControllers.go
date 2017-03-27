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

func cmdControllerDelete() {
	if client {
		kazeDelete(clientsapi, name)
	}
	if event {
		kazeDelete(eventsapi, name)
	}
	if result {
		kazeDelete(resultsapi, name)
	}
	if aggregate {
		kazeDelete(aggregatesapi, name)
	}
	if stash {
		kazeDelete(stashesapi, name)
	}
}

func cmdControllerSilence() {
	if silenceList {
		kazeList(silencedapi, name)
	}
	if silenceClear {
		if client {
			kazeClear(name)
		}
	}
}
