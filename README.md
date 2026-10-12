## Ledger
Ledger is a logging library designed around what I believe a logger should do.
I created it because logging is an important part of any project and I was genuinely interested in learning more. It's still a work in progress, and I intend to improve it over time.
Additionally, I intend to use this in my other projects from now on, so if you're here from one of those... Hey, hope you're doing great!

## Features
- **Custom Writers**: Ledger accepts any writer that implements Go's `io.Writer` interface.
- **Error Handling**: Ledger returns errors from failed operations and provides its own exported errors for conditions that occur within Ledger itself. These can all be checked using `errors.Is` if necessary.
- **Severity Filtering**: Entries below the minimum severity are discarded (by default, everything is recorded).
- **Timestamps**: Every entry includes a timestamp in RFC 3339 formatting.

## Usage
A minimal example of creating a Ledger instance and recording an entry:

```
	// Create a writer that implements io.Writer
	writer := os.Stdout

	// Initialize a new Ledger
	scribe, err := ledger.New(writer)

    // set the minimum severity to whatever you'd like (optional)
    //err = scribe.SetMinimumSeverity(ledger.Warn)

	// some example records
	err = scribe.Record(ledger.Info, "Ledger setup complete!")       // typical program steps
	err = scribe.Record(ledger.Warn, "This is a warning.")           // strange state reached
	err = scribe.Record(ledger.Error, "And this is an Error!")       // problem reached
	err = scribe.Record(ledger.Fatal, "Finish Him! (MK mentioned?)") // unrecoverable state

	// note: the program can still run after a fatal log.
	// Fatal simply means that the program can no longer recover and function as intended
```

For a complete working example, including imports and error handling, please view [USAGE.md](USAGE.md).

## Severity Levels
Ledger currently splits entries into four severity levels: `Info`, `Warn`, `Error`, and `Fatal`. Below is how I suggest they are to be used (although it is not required to follow my convention).

- **Info**: Routine events, such as starting a process or initializing components.
- **Warn**: Unexpected or unusual conditions that have not caused any significant problems.
- **Error**: Problems that affect normal operation but are recoverable.
- **Fatal**: Serious issues that cause the application to no longer operate as intended. This does not necessarily mean the process must terminate immediately, simply that it can no longer function as intended.

## Testing
To run the test suite for this project, clone the repository, enter its directory, and run `go test`
