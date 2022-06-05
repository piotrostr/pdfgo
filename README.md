# PDF Merger

It is super annoying when you have to pay for merging of your pdf's and the
websites are flooded with ads. It takes a long time too smh.

This is a very simple app that just merges pdfs.

```bash
go install github.com/piotrostr/pdfgo@v0.0.3
pdfgo serve
```

This will run a server on `localhost:8080` that serves as a dropbox for pdf files.

It is also possible to run stuff with the cli.

```cli
merge pdfs, lookup eth txs

Usage:
  pdfgo [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  merge       merges the pdfs in the given directory
  serve       run server
  tx          gets the tx by hash
  version     prints the version

Flags:
  -h, --help   help for pdfgo

Use "pdfgo [command] --help" for more information about a command.
```
