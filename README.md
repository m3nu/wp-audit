# Wordpress Version Scanner

If you host a lot of Wordpress sites, you probably want to make sure they have all required updates.

This little Go tool will scan a list of Wordpress sites in parallel and print the corresponding version.

## Installation
Use the pre-built binaries from the *Release* page or with Go installed:

`$ go get -uv github.com/m3nu/wp-audit`

## Usage
First add all your WordPress domains in a textfile, omitting the protocol. See `domains.txt` for an example. Lines starting with `#` are ignored.

Then run `wp-audit` by passing the textfile. It will look for a file called `domains.txt` by default.

```
$ wp-audit --domain-list domains.txt
Site www.saiaja.com.br is on version 5.1
Site www.thetruthseeker.co.uk is on version 5.1
Site murviel-info-beziers.com is on version 5.1
```

## License
(C) Manuel Riel, 2019

This project is licensed under the terms of the MIT license.
