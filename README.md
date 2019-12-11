## BendoBot

[![Go Report Card](https://goreportcard.com/badge/github.com/Bendodroid/BendoBot)](https://goreportcard.com/report/github.com/Bendodroid/BendoBot)

This is a rewrite in Go of my old Discord bot, which was written in Java.
I wrote the old version in 11th grade together with @Yurifag based on a series of Youtube tutorials by @zekroTJA.
After some time it turned into an unmaintainable mess and I took it offline.
This time I try to avoid that.

#### Installation

There are efforts to use Docker to deploy the bot, but it's currently not officially supported.

```shell script
# Clone the repo
$ git clone https://github.com/Bendodroid/BendoBot.git && cd BendoBot
# Compile
$ go build
# Run
$ ./BendoBot
```

To work, the bot needs a Config.json in your current working directory.
An example is provided in this repo.
Of course you need your own Discord API token, which you can get via the Discord Developer Portal.

#### License

Will be chosen when version 0.1 is released.
