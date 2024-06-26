package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/luevano/mangal/config"
	"github.com/luevano/mangal/meta"
	"github.com/luevano/mangal/provider/loader"
	"github.com/luevano/mangal/script"
	"github.com/luevano/mangal/script/lib"
	"github.com/luevano/mangal/util/afs"
	"github.com/spf13/cobra"
)

var scriptArgs = script.Args{}

func init() {
	subcommands = append(subcommands, scriptCmd)
	setDefaultModeShort(scriptCmd)
	// To shorten the statements a bit
	f := scriptCmd.Flags()
	lOpts := loader.Options{}

	f.StringVarP(&scriptArgs.File, "file", "f", "", "Read script from file")
	f.StringVarP(&scriptArgs.String, "string", "s", "", "Read script from script")
	f.BoolVarP(&scriptArgs.Stdin, "stdin", "i", false, "Read script from stdin")
	f.StringVarP(&scriptArgs.Provider, "provider", "p", "", "Load provider by tag")
	f.StringToStringVarP(&scriptArgs.Variables, "vars", "v", nil, "Variables to set in the `Vars` table")
	setupLoaderOptions(f, &lOpts)
	scriptArgs.LoaderOptions = &lOpts

	scriptCmd.MarkPersistentFlagRequired("provider")
	scriptCmd.MarkPersistentFlagRequired("vars")
	scriptCmd.MarkFlagsOneRequired("file", "string", "stdin")
	scriptCmd.MarkFlagsMutuallyExclusive("file", "string", "stdin")
	scriptCmd.RegisterFlagCompletionFunc("provider", completionProviderIDs)
}

var scriptCmd = &cobra.Command{
	Use:     config.ModeScript.String(),
	Short:   "Script, useful for custom process with Lua",
	GroupID: groupMode,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		var reader io.Reader

		switch {
		case cmd.Flag("file").Changed:
			file, err := afs.Afero.OpenFile(
				scriptArgs.File,
				os.O_RDONLY,
				config.Config.Download.ModeFile.Get(),
			)
			if err != nil {
				errorf(cmd, err.Error())
			}

			defer file.Close()

			reader = file
		case cmd.Flag("string").Changed:
			reader = strings.NewReader(scriptArgs.String)
		case cmd.Flag("stdin").Changed:
			reader = os.Stdin
		}

		if err := script.Run(context.Background(), scriptArgs, reader); err != nil {
			errorf(cmd, err.Error())
		}
	},
}

func init() {
	scriptCmd.AddCommand(scriptDocCmd)
}

var scriptDocCmd = &cobra.Command{
	Use:   "doc",
	Short: "Generate documentation for the `mangal` lua library",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		// TODO: use Sprintf?
		filename := fmt.Sprint(meta.AppName, ".lua")
		err := afs.Afero.WriteFile(filename, []byte(lib.LuaDoc()), config.Config.Download.ModeFile.Get())
		if err != nil {
			errorf(cmd, "Error writting library specs: %s", err.Error())
		}
		successf(cmd, "Library specs written to %s", filename)
	},
}
