# RSS-Aggregator

## Introduction

I created a Go program to aggregate [RSS](https://en.wikipedia.org/wiki/RSS) feeds. Websites can post updates to their content using RSS feeds. This project allows you to follow your preferred podcasts, news sites, blogs, and more! It's a web server that lets users access:

* Add RSS feeds to be collected
* Follow and unfollow RSS feeds that other users have added
* Fetch all of the latest posts from the RSS feeds they follow

## Setup

Before we dive into the project, let's make sure you have everything you'll need on your machine.

1. An editor. I use [VS code](https://code.visualstudio.com/), you can use whatever you like.
2. A command line. I use [WSL 2](https://docs.microsoft.com/en-us/windows/wsl/install).
3. The latest [Go toolchain](https://golang.org/doc/install).
4. If you're in VS Code, I recommend the official [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go).
5. An HTTP client. I use [Postman](https://www.postman.com/downloads/), but you can use whatever you like.
