package crypto

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	addressGroupName          = "address"
	operationGroupName        = "operation"
	fromCoinGroupName         = "from_coin"
	toCoinGroupName           = "to_coin"
	fromToNumberCoinGroupName = "from_to_number"
	amountCoinGroupName       = "amount_coin"
	amountNumberGroupName     = "amount_number"
	feePercentGroupName       = "fee_percent"
	feeAmountGroupName        = "fee_amount"
	feeCurrencyGroupName      = "fee_currency"
	fixedFeeGroupName         = "fixed_fee"
	datetimeGroupName         = "datetime"
	dateTimeFormat            = "02/Jan/2006:15:04:05 -0700"
)

var errInvalidLogFormat = errors.New("invalid log format")

// NewFile wraps an os.File creating a special apache common log format regex
// and adding useful helper functions such as seekLine and search for easier working with log files
func NewFile(file *os.File) *File {
	// start
	regExString := `^`
	// 0xeeaFf5e4B8B488303A9F1db36edbB9d73b38dFcf - crypto address
	regExString += fmt.Sprintf(`(?P<%s>\S+) `, addressGroupName)
	// BUY | SELL | CONVERT | WITHDRAW - operation type
	regExString += fmt.Sprintf(`(?P<%s>\S+) `, operationGroupName)
	// BTC/USD:26782.60 - COIN/COIN:price | COIN/COIN:amount
	regExString += fmt.Sprintf(`((?P<%s>\S+)\/((?P<%s>\S+):(?P<%s>\S+))) `, fromCoinGroupName, toCoinGroupName, fromToNumberCoinGroupName)
	// USD:13967.95 - COIN:amount
	regExString += fmt.Sprintf(`((?P<%s>\S+):(?P<%s>\S+)) `, amountCoinGroupName, amountNumberGroupName)
	// 2%(0.02 USD) | 0% | 15USD - fee
	regExString += fmt.Sprintf(`((?P<%s>\d%%)*(\((?P<%s>\S+)\s(?P<%s>\S+)\))*(?P<%s>\d+\S+)*) `, feePercentGroupName, feeAmountGroupName, feeCurrencyGroupName, fixedFeeGroupName)
	// 03/13/2022 11:36:51 +0000 - date time
	regExString += fmt.Sprintf(`(?P<%s>\d{2}\/\d{2}\/\d{4} \d{2}:\d{2}:\d{2} \+\d{4})`, datetimeGroupName)
	// end
	regExString += `$`
	return &File{
		File:  file,
		regEx: regexp.MustCompile(regExString),
	}
}

// File represents a wrapped structure around os.File
// providing additional constructs and helpers for working with log files
type File struct {
	*os.File
	regEx *regexp.Regexp
}

// IndexTime applies a binary search on a log file looking for
// the offset of the log that is withing lookup time (that took place within the last T time).
// offset >= 0 -> means an actual log line to begin reading logs at was found
// offset == -1 -> all the logs inside the log file are older than the lookup time T
func (file *File) IndexTime(lookupTime time.Time) (int64, error) {
	var top, bottom, pos, prevPos, offset, prevOffset int64
	scanLines := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanLines(data, atEOF)
		prevPos = pos
		pos += int64(advance)
		return
	}

	stat, err := os.Stat(file.Name())
	if err != nil {
		return -1, err
	}
	bottom = stat.Size()
	var prevLogTime time.Time
	for top <= bottom {
		// define the middle relative to the top and bottom positions
		middle := top + (bottom-top)/2
		// seek the file at the middle
		_, err := file.Seek(middle, io.SeekStart)
		if err != nil {
			return -1, err
		}
		// reposition the middle to the beginning of the current line
		offset, err = file.seekLine(0, io.SeekCurrent)
		if err != nil {
			return -1, err
		}

		// scan 1 line and parse 1 log line
		scanner := bufio.NewScanner(file)
		scanner.Split(scanLines)
		scanner.Scan()
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			// we'll consider empty line an EOF
			break
		}

		logTime, err := file.parseLogTime(line)
		if err != nil {
			return -1, err
		}

		if lookupTime.Sub(logTime) > 0 {
			// the starting log is way down (relative to the middle)
			// move down the top
			top = offset + (pos - prevPos)
		} else if prevLogTime.Sub(logTime) < 0 {
			// the starting log is way up (relative to the middle)
			// move up the bottom
			bottom = offset - (pos - prevPos)
		} else if lookupTime.Sub(prevLogTime) < 0 && offset != top {
			if lookupTime.Minute() == logTime.Minute() {
				return offset - (pos - prevPos), nil
			}
			return top, nil
		}

		if offset == top {
			if lookupTime.Minute() == logTime.Minute() || top == 0 {
				return top, nil
			}
			return offset - (pos - prevPos), nil
		}
		if offset == bottom {
			if lookupTime.Minute() > logTime.Minute() {
				return top, nil
			}
			return bottom, nil
		}
		if top == bottom && top == stat.Size() {
			return -1, nil
		}

		prevLogTime = logTime
		prevOffset = offset
	}

	if lookupTime.Unix() == prevLogTime.Unix() {
		return prevOffset, nil
	}

	return -1, nil
}

// seekLine resets the cursor for N lines relative to whence, back to the beginning (seek back)
// lines: 0 ->  means seek back (till new line) for the current line
// lines > 0 -> means seek back that many lines
func (file *File) seekLine(lines int64, whence int) (int64, error) {
	const bufferSize = 32 * 1024 // 32KB
	buf := make([]byte, bufferSize)
	bufLen := 0
	lines = int64(math.Abs(float64(lines)))
	seekBack := lines < 1
	lineCount := int64(0)

	// seekBack ignores the first match lines == 0
	// then goes to the beginning of the current line
	if seekBack {
		lineCount = -1
	}

	pos, err := file.Seek(0, whence)
	left := pos
	offset := int64(bufferSize * -1)
	for b := 1; ; b++ {
		if seekBack {
			// on seekBack 2nd buffer onward needs to seek
			// past what was just read plus another buffer size
			if b == 2 {
				offset *= 2
			}

			// if next seekBack will pass beginning of file
			// buffer is 0 to unread position
			if pos+offset <= 0 {
				buf = make([]byte, left)
				left = 0
				pos, err = file.Seek(0, io.SeekStart)
			} else {
				left = left - bufferSize
				pos, err = file.Seek(offset, io.SeekCurrent)
			}
		}
		if err != nil {
			break
		}

		bufLen, err = file.Read(buf)
		if err != nil {
			return file.Seek(0, io.SeekEnd)
		}
		for i := 0; i < bufLen; i++ {
			idx := i
			if seekBack {
				idx = bufLen - i - 1
			}
			if buf[idx] == '\n' {
				lineCount++
			}
			if lineCount == lines {
				if seekBack {
					return file.Seek(int64(i)*-1, io.SeekCurrent)
				}
				return file.Seek(int64(bufLen*-1+i+1), io.SeekCurrent)
			}
		}
		if seekBack && left == 0 {
			return file.Seek(0, io.SeekStart)
		}
	}

	return pos, err
}

// parseLogTime parses a given apache common log line and attempts to convert it into time.Time
// example of apache common log line:
// 127.0.0.1 user-identifier frank [04/Mar/2022:05:30:00 +0000] "GET /api/endpoint HTTP/1.0" 500 123
func (file *File) parseLogTime(l string) (time.Time, error) {
	matches := file.regEx.FindStringSubmatch(l)
	if len(matches) == 0 {
		return time.Time{}, fmt.Errorf("line '%s': %w", l, errInvalidLogFormat)
	}

	var dateTime string
	for i, name := range file.regEx.SubexpNames() {
		if name == datetimeGroupName {
			dateTime = matches[i]
			break
		}
	}
	if dateTime == "" {
		return time.Time{}, fmt.Errorf("invalid date: %w", errInvalidLogFormat)
	}

	t, err := time.Parse(dateTimeFormat, dateTime)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
