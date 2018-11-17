# lllog [![Build Status](https://travis-ci.org/LorenzoLeonardini/lllog.svg?branch=master)](https://travis-ci.org/LorenzoLeonardini/lllog)

Simple and powerful logger for Golang.

![example screen](https://lorenzoleonardini.altervista.org/github/lllog/screen1.png)

**lllog** offers

- different color based on the log level (using [fatih/color](https://github.com/fatih/color) library) + color customization
- custom date and time format to timestamp your logs
- format strings support
- logging to file
- json formatting
- goroutines safety thanks to `sync.Mutex`

### Basic logging

Creating a logger is freaking easy. Every logger requires a name to make easier to identify different loggers

```golang
logger := lllog.New("Name")
```

From now on the logging process is extremely intuitive

```golang
// Notice how the log ends with a \n char, since it is not appended by default
logger.Log("Normal log message\n")
logger.Warn("Different log levels are provided\n")
logger.Err("Nothing very special to say\n")
logger.Fatal("AAAAAAAAAAAAAA\n")

// lllog supports format strings
a := 2
b := 3
logger.Log("The result of a * b is %d\n", a * b)
logger.Log("This is %s!\n", "awesome")
```

### Timestamps

![time screen](https://lorenzoleonardini.altervista.org/github/lllog/screen2.png)

**lllog** can add a time and date header, you just need to add a go time format

```golang
logger.SetFormat("02 January 2006 - 15:04")
```

### Custom color scheme

Since each of us has different taste in color you can edit the color for each log level. (You clearly must choose a `color.Attribute` from the [fatih/color](https://github.com/fatih/color) library)

```golang
logger.SetErrColor(color.FgHiGreen)
```

### Log to file

You can decide to log to a file. Unfortunately you cannot specify the file name, which is in the `yyyy-mm-dd.log` format, but you can specify the path

```golang
logger.LogToFile("logs")
```

Sometimes, if you are logging to a file, it could be useful not to log to the console too. You can enable and disable console logging with

```golang
logger.WriteToConsole(false)
```

### Log formatter

When logging to a file you can decide to use two different formatters. The default one is the console formatter, which prints the exact copy of what you see in the console. **lllog** also features a json formatter which gives you an output like the following, easy and nice to be read by some other program:

```json
{"Level":"Fatal","Msg":"This is sick!\n","Timestamp":1542384480,"Format":"02 January 2006 - 15:04","Name":"File logger"}
```

To change the log formatter:

```golang
logger.SetOutputFormatter(lllog.ConsoleFormatter) // default formatter
logger.SetOutputFormatter(lllog.JSONFormatter)    // json formatter
```

Every log function also returns a string with the log outputted by the chosen log formatter, if you need it for something

```golang
log := logger.Info("This is %s and returned\n", "evaluated")
```
