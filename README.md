# QIF Parser

A Go library for parsing Quicken Interchange Format (QIF) files.

This is a work in progress, docs will be provided when a first version is complete.

## Episub fork

This is a quick and dirty fork of dmjones/qif library to add support for `!Account` metadata.

An additional function is added to the reader. `ReadAccountMetadata` should be run before trying to `Read` if `!Account` is expected in the file.
