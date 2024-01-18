# Anova Precision Oven API

Want to skip my rants? Go directly to [the API documentation](./docs/README.md)

## Introduction

Before we start - a big FU to Anova! After [many years](https://community.anovaculinary.com/t/api-in-2021) of people in the community asking for an open API, or some way, any way, of interacting with their products in an open way, they still will not provide that. In fact, they repeatedly took active action to prevent people from doing this (changing authentication mechanisms, undocumented [API changes](https://anovaculinary.com/pages/software-updates), or [telling people that "they are working on a public API"](https://github.com/bmedicke/anova.py/issues/1) and instead implementing certificate pinning, etc). 

In fact, no - let me rephrase this. This is not directed to all the people at Anova. I am sure they have plenty of good people, working to create good products. And they do make quality products to prove that. This is just for the few people in charge who just don't understand that allowing customers to actually interact with their products in an open way would in no way hurt the company or their market. In fact, it would actually be beneficial for everyone, making possible new use cases.

Most people asking for an API are (just like me) huge fans of Anova and their products, looking for a way to actually extend the functionality of those products. Too bad Anova doesn't seem to understand that...

Rant over, back to the docs.

## Disclaimer

> [!WARNING]
> 
> This project is in no way affiliated with Anova. Anything you find here is just based on whatever public information I was able to gather. And if the past is any guide, things can (and will) change based on Anova's whims. 
> 
> Use this at your own risk.

## Acknowledgements 

Some useful links I used when putting this together:
* [API decoding](https://mcolyer.github.io/anova-oven-api/#introduction) by mcolyer. Unfortunately, this is for the v1 API, which was deprecated by Anova in August 2022
* [Anova oven forwarder](https://github.com/huangyq23/anova-oven-forwarder) by huangyq23 . The project that got me started in the right direction for the v2 API. Also, a great project for anyone looking to collect data from the oven, and even get a nice Grafana dashboard with the various oven parameters. Unfortunately, I was unable to get in touch with the creator - he/she does not respond to issues, and the discord invite in the repo does not work :(
* [The discussion](https://community.home-assistant.io/t/anova-precision-oven/541722) that got this project started. Hopefully, in the future, the information I post here will help someone write an actual Home Assistant integration for the oven
* [The request](https://community.anovaculinary.com/t/api-in-2021) for a public API. Maybe one day this will be available, and my repo will become obsolete. Just don't hold your breath...

## Using the API

The detailed documentation is [here](./docs/README.md)

A Postman Workspace is available [here](https://www.postman.com/bowery/workspace/anova/overview). Follow the directions for forking it to a private workspace so you do not expose your tokens.

