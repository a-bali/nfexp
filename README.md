# nfexp - Netflow explorer

## Objective
Nfexp is a frontend for [nfdump](https://github.com/phaag/nfdump) to parse network flows and present various statistics and reports on a web page.

Currently in very early phase with the following features:
* Show chart of flows and traffic for a given time period
* Perform ad hoc queries and show output from nfdump
* Parse CSV output from nfdump and look up IP addresses in the output (on server side)
* Save queries as reports and show reports for a given time period

## Building and usage
Nfexp is a Go application that builds to a single binary that includes and serves all web assets. For more details, please refer to the included `Makefile`.

Nfexp expects some command line parameters for usage, run `nfexp -h` for more details.

Usage is not yet very intuitive and probably quite a few bugs are also present. The manual page of nfdump is recommended to understand how to define queries/reports.

## License
Nfexp is licensed under GPL 3.0.

