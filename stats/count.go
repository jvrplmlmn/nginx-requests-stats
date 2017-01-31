package stats

import (
	"io"
	"os"
	"strconv"
	"time"

	"github.com/satyrius/gonx"
)

// Counter24h is a wrapper around 'counter' that counts the requests
// in the last 24 hours.
func Counter24h(parser *gonx.Parser, logFile string) (int, error) {
	return counter(parser, logFile, 1)
}

// counter reads a given logFile, and counts the requests that happened
// in the time range between now and the last N days (using now as a time reference).
// The filter uses the 'combined_log' format to parse the log entries.
func counter(parser *gonx.Parser, logFile string, days int) (int, error) {

	// Read given file
	var logReader io.Reader
	file, err := os.Open(logFile)
	if err != nil {
		return 0, err
	}
	logReader = file
	defer file.Close()

	// Get yesterday's and current timestamp in UTC format
	yesterday := time.Now().AddDate(0, 0, -1).UTC()
	now := time.Now().UTC()

	// Make a chain of reducers.
	// Filter based on the date range.
	// Count the entries that pass the filter.
	reducer := gonx.NewChain(
		&gonx.Datetime{
			Field:  "time_local",
			Format: "02/Jan/2006:15:04:05 -0700", // combined_log format
			Start:  yesterday,
			End:    now,
		},
		&gonx.Count{})

	// Process the file with the given chain of reducers (filter + count).
	output := gonx.MapReduce(logReader, parser, reducer)
	for res := range output {
		value, err := res.Field("count")
		if err != nil {
			return 0, err
		}
		return strconv.Atoi(value)
	}
	return 0, nil
}
