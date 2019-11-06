# Krabbel

[![golang](https://img.shields.io/badge/Language-Go-green.svg)](https://golang.org)
[![GoDoc](https://godoc.org/github.com/mwat56/krabbel?status.svg)](https://godoc.org/github.com/mwat56/krabbel)
[![Go Report](https://goreportcard.com/badge/github.com/mwat56/krabbel)](https://goreportcard.com/report/github.com/mwat56/krabbel)
[![Issues](https://img.shields.io/github/issues/mwat56/krabbel.svg)](https://github.com/mwat56/krabbel/issues?q=is%3Aopen+is%3Aissue)
[![Size](https://img.shields.io/github/repo-size/mwat56/krabbel.svg)](https://github.com/mwat56/krabbel/)
[![Tag](https://img.shields.io/github/tag/mwat56/krabbel.svg)](https://github.com/mwat56/krabbel/tags)
[![View examples](https://img.shields.io/badge/learn%20by-examples-0077b3.svg)](https://github.com/mwat56/krabbel/blob/master/app/krabbel.go)
[![License](https://img.shields.io/github/mwat56/krabbel.svg)](https://github.com/mwat56/krabbel/blob/master/LICENSE)

- [Krabbel](#krabbel)
	- [Purpose](#purpose)
	- [Installation](#installation)
	- [Usage](#usage)
	- [Licence](#licence)

## Purpose

When writing web applications (server and/or client) debugging, refactoring and testing brings you only so far.
One of the things that are not easy to mock or fake is a certain workload that shows you how your application is behaving under load.
This is were `krabbel` comes in: it's basically just a web-crawler that tries to get all local links within an URL as soon as possible.
Its sole purpose is to produce some kind of stress test that shows how well your web-server reacts under load.

## Installation

You can use `Go` to install this package for you:

	go get -u github.com/mwat56/krabbel

## Usage

First you've to compile the main file

	go build app/krabble.go

you can run it by calling it with the start URL to use, e.g.

	./krabbel http://127.0.0.1:8080/

Depending on the number of linked pages it might run a few seconds while printing out the respective page processed and finally showing a line like

	2019/10/28 23:37:41 checked 3422 pages in 5.9901556s

The actual number of pages shown and the time used will change depending on the load of the computer you use to run the tool and the load of the server tested.
Things like routing details and network latency will take their time as well.
In other words: This is _not_ a benchmarking tool.

> Please _note_ that you should use this tool only with web-servers/-pages that you're personally responsible for.
> Do _not_ use this tool with servers you don't own – that's not only impolite but also _illegal_ in some countries.

## Licence

        Copyright © 2019 M.Watermann, 10247 Berlin, Germany
                        All rights reserved
                    EMail : <support@mwat.de>

> This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.
>
> This software is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
>
> You should have received a copy of the GNU General Public License along with this program. If not, see the [GNU General Public License](http://www.gnu.org/licenses/gpl.html) for details.
