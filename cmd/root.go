/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
	"golang.design/x/clipboard"

	"github.com/dishonoreded/cbhelper/lib"
)

const (
	cbHelperFlagsDecode = "decode"
	cbHelperFlagsEncode = "encode"
	cbHelperFlagsOutput = "output"
	cbHelperFlagsFormat = "format"

	cbHelperFlagsOutputClipboard = "clipboard"
	cbHelperFlagsOutputStdout    = "stdout"

	cbHelperFlagsFormatRaw    = "raw"
	cbHelperFlagsFormatPretty = "pretty-json"
)

type cbHelperFlagsOutputOption string
type cbHelperFlagsOutputFunc func([]byte)
type OutputKeyToFuncMap map[cbHelperFlagsOutputOption]cbHelperFlagsOutputFunc

var cbHelperFlagsOutputFuncMap = OutputKeyToFuncMap{
	"stdout": func(bytes []byte) {
		_, err := os.Stdout.Write(bytes)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = os.Stdout.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	},
	"clipboard": func(bytes []byte) {
		clipboard.Write(clipboard.FmtText, bytes)
	}}

func (m OutputKeyToFuncMap) Options() []string {
	var options []string
	for option := range cbHelperFlagsOutputFuncMap {
		options = append(options, string(option))
	}
	return options
}

func (m OutputKeyToFuncMap) RetrieveExecution(option string) interface{} {
	f, ok := cbHelperFlagsOutputFuncMap[cbHelperFlagsOutputOption(option)]
	if !ok {
		return nil
	}
	return f
}

type cbHelperFlagsFormatOption string
type cbHelperFlagsFormatFunc func([]byte) []byte
type OptionKeyToFuncMap map[cbHelperFlagsFormatOption]cbHelperFlagsFormatFunc

var cbHelperFlagsFormatFuncMap = OptionKeyToFuncMap{
	"raw": func(bytes []byte) []byte {
		return bytes
	},
	"json": lib.FormatJSON,
	"pretty-json": func(bytes []byte) []byte {
		return pretty.Color(pretty.PrettyOptions(bytes, pretty.DefaultOptions), nil)
	}}

func (m OptionKeyToFuncMap) Options() []string {
	var options []string
	for option := range cbHelperFlagsFormatFuncMap {
		options = append(options, string(option))
	}
	return options
}

func (m OptionKeyToFuncMap) RetrieveExecution(option string) interface{} {
	f, ok := cbHelperFlagsFormatFuncMap[cbHelperFlagsFormatOption(option)]
	if !ok {
		return nil
	}
	return f
}

type cbHelperFlags struct {
	Decode string
	Encode string
	Output string
	Format string
}

type bytesHandlers func([]byte) ([]byte, error)

var cbhFlags cbHelperFlags

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "cbHelper",
	Aliases: []string{"cbh"},
	Short:   "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		raw := clipboard.Read(clipboard.FmtText)
		if cmd.Flags().Lookup(cbHelperFlagsDecode).Value.String() != "" {
			raw = cbHelperDecoder(raw)
		} else if cmd.Flags().Lookup(cbHelperFlagsEncode).Value.String() != "" {
			raw = cbHelperEncoder(raw)
		}

		formatFunc := cbHelperFlagsFormatFuncMap.RetrieveExecution(cmd.Flags().Lookup(cbHelperFlagsFormat).Value.String())
		if formatFunc == nil {
			cmd.PrintErr("unknown format option")
			return
		}
		raw = formatFunc.(cbHelperFlagsFormatFunc)(raw)

		outputFunc := cbHelperFlagsOutputFuncMap.RetrieveExecution(cmd.Flags().Lookup(cbHelperFlagsOutput).Value.String())
		if outputFunc == nil {
			cmd.PrintErr("unknown output option")
			return
		}
		outputFunc.(cbHelperFlagsOutputFunc)(raw)
	},

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	err := clipboard.Init()
	if err != nil {
		log.Fatal(err)
	}
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cbHelper.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&cbhFlags.Decode, cbHelperFlagsDecode, "d", "", "")
	rootCmd.Flags().StringVarP(&cbhFlags.Encode, cbHelperFlagsEncode, "e", "", "")
	rootCmd.MarkFlagsMutuallyExclusive(cbHelperFlagsDecode, cbHelperFlagsEncode)
	rootCmd.Flags().StringVarP(&cbhFlags.Output, cbHelperFlagsOutput, "o", "clipboard", `clipboard/stdout`)
	rootCmd.Flags().StringVarP(&cbhFlags.Format, cbHelperFlagsFormat, "f", "raw",
		strings.Join(cbHelperFlagsFormatFuncMap.Options(), "/"))
}

var decoderFuncMap = map[byte]bytesHandlers{
	'b': lib.Base64Decode,
	'z': lib.UnGzip,
	'u': lib.URLUnescape,
	'y': lib.BytesArrayDecode,
}

var encoderFuncMap = map[byte]bytesHandlers{
	'b': lib.Base64Encode,
	'z': lib.Gzip,
	'u': lib.URLEscape,
}

func cbHelperDecoder(raw []byte) []byte {
	var decodeExecutions []bytesHandlers
	for _, b := range []byte(cbhFlags.Decode) {
		decodeExecutions = append(decodeExecutions, decoderFuncMap[b])
	}

	var err error
	for _, f := range decodeExecutions {
		raw, err = f(raw)
		if err != nil {
			log.Fatal(err)
		}
	}
	return raw
}

func cbHelperEncoder(raw []byte) []byte {
	var encodeExecutions []bytesHandlers
	for _, b := range []byte(cbhFlags.Encode) {
		encodeExecutions = append(encodeExecutions, encoderFuncMap[b])
	}

	var err error
	for _, f := range encodeExecutions {
		raw, err = f(raw)
		if err != nil {
			log.Fatal(err)
		}
	}
	return raw
}
