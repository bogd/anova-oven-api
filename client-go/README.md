# API Client Sample

## Overview

A (very basic) implementation of a client for the Anova Precision Oven API.

> [!WARNING]
> 
> This project is in no way affiliated with Anova. Anything you find here is just based on whatever public information I was able to gather. And if the past is any guide, things can (and will) change based on Anova's whims. 
> 
> Use this at your own risk.

## Prerequisites

A Linux machine, running Docker. It should also work on Windows, but I have only tested using Linux.

## Usage

### Getting the Credentials

You will need the `refreshToken` and `apiKey` for your installation. See [the main README file](../README.md#getting-the-refresh-token) for details on how to get them.

Once you have them, copy `creds_anova.yml.sample` to `creds_anova.yml`:

```
cp creds_anova.yml.sample creds_anova.yml
```

Edit the file, placing your `refreshToken` and `apiKey` in the correct places.

### Option 1 - Using a pre-built Docker image
The pre-built image (for `x86_64`, the only architecture I am using) is published on dockerhub as `bogd/anova-oven-api` 

Make sure you have the `creds_anova.yml` file in the current directory, and run:

```
docker container run -v $(pwd)/creds_anova.yml:/creds_anova.yml bogd/anova-oven-api
```

You can append command-line arguments as needed (see `anova-oven-api --help` for the available options):

```
docker container run -v $(pwd)/creds_anova.yml:/creds_anova.yml bogd/anova-oven-api -showjson
```


### Option 2 - Building your own binaries

From the `client-go` directory:

```
make run
```

This will build the application and launch it (with the default settings). It should start decoding `EVENT_APO_STATE` messages, and print some of the info in the messages:

```
Reading credentials from file
Creating an access token
Preparing WS connection
Connecting to websocket
Oven with ID |0123456789abcdef|: mode idle, current temperature (dry) 31.12, timer 00:00
Oven with ID |0123456789abcdef|: mode idle, current temperature (dry) 31.38, timer 00:00
```

### Command-line arguments

```
anova-oven-api --help
```
