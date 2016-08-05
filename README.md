# Wifi_Hack 
[![Build Status](https://travis-ci.org/travis-ci/travis-web.svg?branch=master)](https://travis-ci.org/travis-ci/travis-web)

To turn of and on your wifi after a certain time period. Aims to solve the macbook and mac wifi problem.

Problem Statement
----------------

After facing the mac wifi resatart problems as stated in many forums I decided to make a Go script that runs in the background, automating the wifi restart process.

Requirements
-----------
* Golang - http://golang.org
* Git

Installation
-----------
Setup a git repository and then pull this.

Usage
-----
Move into the repository and run the program by typing in -
```
go run wifi_hack.go
```
Functionality
-----------
The program will run in the background (do not close the terminal window) and automatically turn off and on your wifi every 20 minutes. Future versions will detect packet drops to activate the wifi reset process.

License
------
MIT






