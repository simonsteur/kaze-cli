# kaze-cli

kaze-cli is a cli management tool for sensu. it provides a way to easily interact with the sensu api which gives you more time to build up your monitoring solution. 

# download
* [osx 64bit](https://github.com/simonsteur/kaze-cli/releases/download/1.0/kaze-cli-darwin-amd64.zip)
* [freebsd 64bit](https://github.com/simonsteur/kaze-cli/releases/download/1.0/kaze-cli-freebsd-amd64.zip)
* [linux 64bit](https://github.com/simonsteur/kaze-cli/releases/download/1.0/kaze-cli-linux-amd64.zip)
* [windows 64bit](https://github.com/simonsteur/kaze-cli/releases/download/1.0/kaze-cli-win-amd64.zip)


# configuration
To operate kaze-cli needs to know your sensu-api endpoint. When running kaze-cli for the first time use the following command: 
```
kaze configure -address <address> -port <port> 
```
# how to use
```
kaze-cli is command line interface tool for sensu operations

Usage:
  kaze [command] [options]

Commands:
  configure        configure kaze-cli
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

# license 

Copyright 2017 Simon Steur

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.