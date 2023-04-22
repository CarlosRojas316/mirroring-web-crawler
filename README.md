# Implement a recursive, mirroring web crawler

The crawler should be a command-line tool that accepts a starting URL and a destination directory.
The crawler will then download the page at the URL, save it in the destination directory, 
and then recursively proceed to any valid links in this page.
A valid link is the value of an <u>href</u> attribute in an <u><`a`></u> tag the resolves to https://start.url/abc/foo and https://start.url/abc/foo/bar are valid URLs, but ones that resolve to https://another.domain/ or to https://start.url/baz are <i>not</i> valid URLs, and
should be skipped.

Additionally, the crawler should:
- Correctly handle being interrupted by Ctrl-C
- Perform work in parallel where reasonable
- Support resume functionality by checking the destination directory for downloaded pages and skip downloading and processing where not necessary
- Provide “happy-path” test coverage


# 

Usage:

```shell
go run main.go <--url url> <--path destination>

> go run main.go --url http://www.w3school.com/cpp --path websites
```

Developed using Golang version 1.20


Main features:
- File processing to guarantee that links are relative to base directory
- Handle Ctrl-C
  > Not top priority feature while no concurrency implemented
- Concurrency 
  > Mutex to handle set visited url
