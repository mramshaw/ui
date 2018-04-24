# andlabs / ui

[![Go Report Card](https://goreportcard.com/badge/github.com/mramshaw/ui?style=flat-square)](https://goreportcard.com/report/github.com/mramshaw/ui)
[![GoDoc](https://godoc.org/github.com/mramshaw/ui?status.svg)](https://godoc.org/github.com/mramshaw/ui)
[![GitHub release](https://img.shields.io/github/release/mramshaw/ui.svg?style=flat-square)](https://github.com/mramshaw/ui/releases)

[Just a mickey-mouse demo I wrote up for the Golang 'Code and Chat' Meetup.]

Find the code, etc. for `andlabs/ui` at this [link](https://github.com/andlabs/ui).

There is a nice [Getting Started](https://github.com/andlabs/ui/wiki/Getting-Started) page.
Be sure to read this, as it is pretty good.

## Prerequisites

Golang (Go) installed. Find the [Download Go](https://golang.org/) link on the Golang home page.

## Installation

1. Verify that you are on a [supported platform](https://github.com/andlabs/ui#update-5-june-2016-you-can-finally-go-get-this-package). It looks like __OS/X__, __linux__ and __Windows__ are supported.

2. Download the latest version:

    ```
    $ go get -u github.com/andlabs/ui
    ```

3. Dependencies: if the previous command results in something like:

    ```
    $ go get -u github.com/andlabs/ui
    # pkg-config --cflags gtk+-3.0
    Package gtk+-3.0 was not found in the pkg-config search path.
    Perhaps you should add the directory containing `gtk+-3.0.pc'
    to the PKG_CONFIG_PATH environment variable
    No package 'gtk+-3.0' found
    pkg-config: exit status 1
    $
    ```

    Then install any missing dependencies (`lib-gtk-3-dev` plus extras in my case).

    Try the installation again (which may require installing more dependencies):

        $ go get -u github.com/andlabs/ui
        $

## Running

To simply run the example (i.e. no ___build___):

    $ go run main.go

Type something in the textbox and press 'Greet' to see what happens.

Close the window to terminate.

## Building

To build the example:

    $ go build main.go

This should produce a binary: `main`

To run the binary:

    $ ./main

Type something in the textbox and press 'Greet' to see what happens.

Close the window to terminate.

## To Do

* [ ] Figure what makes a box element ___stretchy___
* [x] Figure out how to get the selected items text from a __combobox__
* [ ] Figure out how to get the selected items text from a __radiobox__
