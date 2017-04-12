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
		kazeDelete(clientsapi, name, deleteCheckName)
	}
	if event {
		kazeDelete(eventsapi, name, deleteCheckName)
	}
	if result {
		kazeDelete(resultsapi, name, deleteCheckName)
	}
	if aggregate {
		kazeDelete(aggregatesapi, name, deleteCheckName)
	}
	if stash {
		kazeDelete(stashesapi, name, deleteCheckName)
	}
}

func cmdControllerSilence() {
	if client {
		kazeSilence(name)
	}
}

func cmdControllerSilenceClear() {
	kazeClear(name)
}

func cmdControllerCheck() {
	kazeCheck(name)
}
