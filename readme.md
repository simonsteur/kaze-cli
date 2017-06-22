# kaze-cli

kaze-cli is a cli management tool for sensu. it provides a way to easily interact with the sensu api which gives you more time to build up your monitoring solution. 

# configuration
To operate kaze-cli needs to know your sensu-api endpoint. When running kaze-cli for the first time use the following command: 
```
kaze -configure -address <address> -port <port> 
```
# how to use
```
kaze-cli is command line interface tool for sensu operations

Usage:
  kaze [command] [options]

Commands:
  list             list objects
  create-client    creates a proxy client
  create-result    creates a check result
  create-stash     creates a stash
  delete           delete clients, results, stashes
  clear-silence    clear a silence entry
  check            request to schedule a check
  resolve          resolve a check result
  help             print help text


for help use: kaze [command] -help
```

# the name

sensu refers to the japense foldable fan, the dashboard for sensu is called uchiwa which is a japenese rigid fan. kaze is japense for wind. It seemed appropriate. 

