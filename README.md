# PrettyPrinter
Pretty Printer package used for colorful console output support for Golang

## Usage

Pretty Printer simpyfies writing colorful output to command line providing simple API which is more readable than:
```Go
	fmt.Printf("\x1b[47m\x1b[31;1mRed title: \x1b[0m%s: \n\x1b[32;1mBody: \x1b[0m%s\x1b[0m\n", "Sample title param", "Sample body param")
```

With this API you can do:
```Go
	fmt.Printf(prettyPrinter.New().
		BgYellow().BoldRed("Red title: ").Text("%s: \n").
		BoldGreen("Body: ").Text("%s\n").
		Format("Sample title param", "Sample body param").String())
```

Every text element (colorful or not) will reset background color

Sample above will have following output:

![Alt text](https://dl.dropboxusercontent.com/u/4557819/pretty_printer_sample.png "Sample output screenshot")
