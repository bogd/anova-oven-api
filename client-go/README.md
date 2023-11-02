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


### Option 2 - Building your own binaries



