# Tibia Scrapper

This project is a scrapper for the tibia's game website,
since it doesn't have a public api, all data has to be fetched via scrapping.

For now, it only looks for characters in bazaar that have a Ferumbras hat equipped,
meaning that the character has the "Alumni" achievement.

## Starting 🚀

_This instructions will guide you on how to run the project successfully_


### Prerequisites for running 📋

_Dependencies (minimum tested versions)_

* Docker version 20.10.7
* docker-compose version 1.29.2
* Make

### Set up 🔧

_Spin up the scrapper container_

```makefile
make start
```

The `make start` command is going to spin up a container that will automatically run the binary generated.
Once the process is finished it will create a file with all the characters found to the outputs folder,
using json format.

### Useful commands ℹ️
Display the status of the container to know if it's running
```makefile
make status
```

Display logs of scrapper container
```makefile
make logs
```

Stop running scrapper container
```makefile
make stop
```

## Built with 🛠️

* [Golang](https://golang.org/) - The language
* [Colly](http://go-colly.org/) - The scrapping framework
* [Logrus](https://github.com/sirupsen/logrus) - The logging implementation

## Authors ✒️

* **Norberto Gómez** - *Owner* - [norbertogomez](https://github.com/norbertogomez)

## TODOs

- [ ] Create dynamic filters (level, price, skills, etc.)
- [ ] Create configuration files
- [ ] Write unit tests
- [ ] Extend scrapping to collect world information (pvp config, green/yellow, battleye status)
- [ ] Scrap server deaths on open-pvp+ servers