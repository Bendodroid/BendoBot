# BendoBot

[![Go Report Card](https://goreportcard.com/badge/github.com/Bendodroid/BendoBot)](https://goreportcard.com/report/github.com/Bendodroid/BendoBot)

This is a rewrite in Go of my old Discord bot, which was written in Java.
I wrote the old version in 11th grade together with @Yurifag based on a series of Youtube tutorials by @zekroTJA.
After some time it turned into an unmaintainable mess and I took it offline.
This time I try to avoid that.

## Installation

### Classic

```shell script
# Clone the repo
git clone https://github.com/Bendodroid/BendoBot.git && cd BendoBot
# Compile
go build
# Run
./BendoBot
```

To work, the bot needs a `Config.json` in your current working directory.
An example is provided in this repo.
Of course you need your own Discord API token, which you can get via the Discord Developer Portal.

### Docker

#### Step 1: Obtaining the Docker Image

Currently there's no Docker image available via Docker Hub or the like, so you need to build one yourself:

```shell script
# Clone the repo
git clone https://github.com/Bendodroid/BendoBot.git && cd BendoBot
# Build the Docker image and name it bendobot
docker image build -t bendobot .
```

#### Step 2: Creating a Docker Volume for the Configuration File

In this step you create a Docker volume with a configuration file inside it.
This volume will then be mounted inside the Docker container.
This way we can easily use a configuration file inside a Docker container.

First create the Docker volume:

```shell script
# Create a Docker volume named BendoBotData.
docker volume create BendoBotData
```

Get the mountpoint of the Docker volume:

```shell script
docker volume inspect BendoBotData | grep Mountpoint
```

This command should give you the path of the mountpoint.

Now create a `Config.json` as described [above](#classic) and move it into the Docker volume using the mountpoint you just obtained.

Here's an example of how that can be done:

```shell script
# Copy the example Config.json (Config.example.json) to Config.json
cp Config.example.json Config.json
# Edit the Config.json to replace the placeholder values
vim Config.json
# Move the Config.json into the Docker volume
mv Config.json /path/of/Docker/volume/mountpoint
```

#### Step 3: Running the Bot

Now you can easily run the Bot by using the following command:

```shell script
docker run -d --rm --name BendoBot --mount source=BendoBotData,target=/app/config_volume bendobot
```

You can get BendoBots output using this command:

```shell script
docker container logs BendoBot
```

And you can stop the container and bot using this command:

```shell script
docker container stop BendoBot
```

## License

Will be chosen when version 0.1 is released.
