package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"sync"
	"time"

	"strconv"

	"github.com/mattn/go-isatty"
	"github.com/valyala/fasttemplate"

	"github.com/labstack/gommon/color"
)

type (
	// Logger defines logger information.
	Logger struct {
		prefix     string
		level      Lvl
		output     io.Writer
		template   *fasttemplate.Template
		levels     []string
		color      *color.Color
		bufferPool sync.Pool
		mutex      sync.Mutex
		timeFormat string
	}

	// Lvl defines log level.
	Lvl uint8

	// JSON defines json format.
	JSON map[string]interface{}
)

const (
	DEBUG Lvl = iota + 1
	INFO
	WARN
	ERROR
	FATAL
	PANIC
	OFF
)

var (
	// global defines default logger.
	global = New("-")
	// defaultHeader is default log header.
	defaultHeader = `{"time":"${time}","level":"${level}","prefix":"${prefix}",` +
		`"file":"${short_file}","line":"${line}"}`
)

// New creates a new logger.
func New(prefix string) (l *Logger) {
	l = &Logger{
		level:    INFO,
		prefix:   prefix,
		template: l.newTemplate(defaultHeader),
		color:    color.New(),
		bufferPool: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 256))
			},
		},
		timeFormat: time.RFC3339,
	}
	l.initLevels()
	l.SetOutput(output())
	return
}

// initLevels initializes log level.
func (l *Logger) initLevels() {
	l.levels = []string{
		"-",
		l.color.Blue("DEBUG"),
		l.color.Green("INFO"),
		l.color.Yellow("WARN"),
		l.color.Red("ERROR"),
		l.color.Red("FATAL"),
		l.color.Red("PANIC"),
	}
}

// newTemplate creates a new logger template.
func (l *Logger) newTemplate(format string) *fasttemplate.Template {
	return fasttemplate.New(format, "${", "}")
}

// DisableColor disables colorized output.
func (l *Logger) DisableColor() {
	l.color.Disable()
	l.initLevels()
}

// EnableColor enables colorized output.
func (l *Logger) EnableColor() {
	l.color.Enable()
	l.initLevels()
}

// Prefix gets the output prefix for the standard logger.
func (l *Logger) Prefix() string {
	return l.prefix
}

// SetPrefix sets the output prefix for the standard logger.
func (l *Logger) SetPrefix(p string) {
	l.prefix = p
}

// Level gets log level.
func (l *Logger) Level() Lvl {
	return l.level
}

// SetLevel sets log level.
func (l *Logger) SetLevel(v Lvl) {
	l.level = v
}

// Output gets the output destination for the standard logger.
func (l *Logger) Output() io.Writer {
	return l.output
}

// SetOutput sets the output destination for the standard logger.
func (l *Logger) SetOutput(w io.Writer) {
	l.output = w
	if w, ok := w.(*os.File); !ok || !isatty.IsTerminal(w.Fd()) {
		l.DisableColor()
	}
}

// Color gets log color setting.
func (l *Logger) Color() *color.Color {
	return l.color
}

// SetHeader sets a logger header.
func (l *Logger) SetHeader(h string) {
	l.template = l.newTemplate(h)
}

// Print calls Output to print to the standard logger.
func (l *Logger) Print(i ...interface{}) {
	l.log(0, "", i...)
	// fmt.Fprintln(l.output, i...)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, args ...interface{}) {
	l.log(0, format, args...)
}

// Printj calls Output to print to the standard logger.
func (l *Logger) Printj(j JSON) {
	l.log(0, "json", j)
}

// Debug calls Output to print debug message to the standard logger.
func (l *Logger) Debug(i ...interface{}) {
	l.log(DEBUG, "", i...)
}

// Debugf calls Output to print debug message to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

// Debugj calls Output to print debug message to the standard logger.
func (l *Logger) Debugj(j JSON) {
	l.log(DEBUG, "json", j)
}

// Info calls Output to print info message to the standard logger.
func (l *Logger) Info(i ...interface{}) {
	l.log(INFO, "", i...)
}

// Infof calls Output to print info message to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

// Infoj calls Output to print info message to the standard logger.
func (l *Logger) Infoj(j JSON) {
	l.log(INFO, "json", j)
}

// Warn calls Output to print warning message to the standard logger.
func (l *Logger) Warn(i ...interface{}) {
	l.log(WARN, "", i...)
}

// Warnf calls Output to print warning message to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.log(WARN, format, args...)
}

// Warnj calls Output to print warning message to the standard logger.
func (l *Logger) Warnj(j JSON) {
	l.log(WARN, "json", j)
}

// Error calls Output to print error message to the standard logger.
func (l *Logger) Error(i ...interface{}) {
	l.log(ERROR, "", i...)
}

// Errorf calls Output to print error message to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

// Errorj calls Output to print error message to the standard logger.
func (l *Logger) Errorj(j JSON) {
	l.log(ERROR, "json", j)
}

// Fatal calls Output to print fatal message to the standard logger.
func (l *Logger) Fatal(i ...interface{}) {
	l.log(FATAL, "", i...)
	os.Exit(1)
}

// Fatalf calls Output to print fatal message to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.log(FATAL, format, args...)
	os.Exit(1)
}

// Fatalj calls Output to print fatal message to the standard logger.
func (l *Logger) Fatalj(j JSON) {
	l.log(FATAL, "json", j)
	os.Exit(1)
}

// Panic calls Output to print panic message to the standard logger.
func (l *Logger) Panic(i ...interface{}) {
	l.log(PANIC, "", i...)
	panic(fmt.Sprint(i...))
}

// Panicf calls Output to print panic message to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.log(PANIC, format, args...)
	panic(fmt.Sprintf(format, args))
}

// Panicj calls Output to print panic message to the standard logger.
func (l *Logger) Panicj(j JSON) {
	l.log(PANIC, "json", j)
	panic(j)
}

// DisableColor disables colorized output.
func DisableColor() {
	global.DisableColor()
}

// EnableColor enables colorized output.
func EnableColor() {
	global.EnableColor()
}

// Prefix gets the output prefix for the standard logger.
func Prefix() string {
	return global.Prefix()
}

// SetPrefix sets the output prefix for the standard logger.
func SetPrefix(p string) {
	global.SetPrefix(p)
}

// Level gets log level.
func Level() Lvl {
	return global.Level()
}

// SetLevel sets log level.
func SetLevel(v Lvl) {
	global.SetLevel(v)
}

// Output gets the output destination for the standard logger.
func Output() io.Writer {
	return global.Output()
}

// SetOutput sets the output destination for the standard logger.
func SetOutput(w io.Writer) {
	global.SetOutput(w)
}

// SetHeader sets a logger header.
func SetHeader(h string) {
	global.SetHeader(h)
}

// Print calls Output to print to the standard logger.
func Print(i ...interface{}) {
	global.Print(i...)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, args ...interface{}) {
	global.Printf(format, args...)
}

// Printj calls Output to print to the standard logger.
func Printj(j JSON) {
	global.Printj(j)
}

// Debug calls Output to print debug message to the standard logger.
func Debug(i ...interface{}) {
	global.Debug(i...)
}

// Debugf calls Output to print debug message to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, args ...interface{}) {
	global.Debugf(format, args...)
}

// Debugj calls Output to print debug message to the standard logger.
func Debugj(j JSON) {
	global.Debugj(j)
}

// Info calls Output to print info message to the standard logger.
func Info(i ...interface{}) {
	global.Info(i...)
}

// Infof calls Output to print info message to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, args ...interface{}) {
	global.Infof(format, args...)
}

// Infoj calls Output to print info message to the standard logger.
func Infoj(j JSON) {
	global.Infoj(j)
}

// Warn calls Output to print warning message to the standard logger.
func Warn(i ...interface{}) {
	global.Warn(i...)
}

// Warnf calls Output to print warning message to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Warnf(format string, args ...interface{}) {
	global.Warnf(format, args...)
}

// Warnj calls Output to print warning message to the standard logger.
func Warnj(j JSON) {
	global.Warnj(j)
}

// Error calls Output to print error message to the standard logger.
func Error(i ...interface{}) {
	global.Error(i...)
}

// Errorf calls Output to print error message to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, args ...interface{}) {
	global.Errorf(format, args...)
}

// Errorj calls Output to print error message to the standard logger.
func Errorj(j JSON) {
	global.Errorj(j)
}

// Fatal calls Output to print fatal message to the standard logger.
func Fatal(i ...interface{}) {
	global.Fatal(i...)
}

// Fatalf calls Output to print fatal message to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Fatalf(format string, args ...interface{}) {
	global.Fatalf(format, args...)
}

// Fatalj calls Output to print fatal message to the standard logger.
func Fatalj(j JSON) {
	global.Fatalj(j)
}

// Panic calls Output to print panic message to the standard logger.
func Panic(i ...interface{}) {
	global.Panic(i...)
}

// Panicf calls Output to print panic message to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Panicf(format string, args ...interface{}) {
	global.Panicf(format, args...)
}

// Panicj calls Output to print panic message to the standard logger.
func Panicj(j JSON) {
	global.Panicj(j)
}

func (l *Logger) log(v Lvl, format string, args ...interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	buf := l.bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer l.bufferPool.Put(buf)
	_, file, line, _ := runtime.Caller(3)

	if v >= l.level || v == 0 {
		var message string
		if format == "" {
			message = fmt.Sprint(args...)
		} else if format == "json" {
			b, err := json.Marshal(args[0])
			if err != nil {
				panic(err)
			}
			message = string(b)
		} else {
			message = fmt.Sprintf(format, args...)
		}

		_, err := l.template.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
			switch tag {
			case "time":
				return w.Write([]byte(time.Now().Format(l.timeFormat)))
			case "level":
				return w.Write([]byte(l.levels[v]))
			case "prefix":
				return w.Write([]byte(l.prefix))
			case "long_file":
				return w.Write([]byte(file))
			case "short_file":
				return w.Write([]byte(path.Base(file)))
			case "line":
				return w.Write([]byte(strconv.Itoa(line)))
			}
			return 0, nil
		})

		if err == nil {
			s := buf.String()
			i := buf.Len() - 1
			if s[i] == '}' {
				// JSON header
				buf.Truncate(i)
				buf.WriteByte(',')
				if format == "json" {
					buf.WriteString(message[1:])
				} else {
					buf.WriteString(fmt.Sprintf(`"message":"%s"}`, message))
				}
			} else {
				// Text header
				buf.WriteByte(' ')
				buf.WriteString(message)
			}
			buf.WriteByte('\n')
			l.output.Write(buf.Bytes())
		}
	}
}
