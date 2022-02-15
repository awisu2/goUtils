package output

import (
	"github.com/awisu2/goUtils/output"
	"github.com/spf13/cobra"
)

var params = struct {
	stdout string
	stderr string
}{}

// コマンド作成
var Cmd = &cobra.Command{
	Use:   "output",
	Short: "output the set value to stdout and stderr",
	Args:  cobra.ArbitraryArgs, // 引数設定(ArbitraryArgs: なんでもOK)
	Run: func(cmd *cobra.Command, args []string) {
		output.Output(params.stdout, params.stderr)
	},
}

func initConfig() {
	// viper(config値/ファイルに強いmodule)などを実行
}

// go 標準関数: 実行時に処理
func init() {
	// 引数が解析されたあと、Runの実行前に実行されるイベント設定
	cobra.OnInitialize(initConfig)

	// 引数設定
	// flags, pFlags := Cmd.Flags(), Cmd.PersistentFlags()
	flags := Cmd.Flags()

	// std
	flags.StringVarP(&params.stdout, "out", "o", "", "string of standard out")
	flags.StringVarP(&params.stderr, "err", "e", "", "string of standard error")

	// 必須指定
	requireds := []string{}
	persistentRequireds := []string{}
	for _, required := range requireds {
		Cmd.MarkFlagRequired(required)
	}
	for _, required := range persistentRequireds {
		Cmd.MarkPersistentFlagRequired(required)
	}

	// 配下コマンドの追加
	// Cmd.AddCommand(subCmd)
}
