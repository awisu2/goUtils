package time

import (
	"github.com/awisu2/goUtils/cmd/time/sleep"
	"github.com/spf13/cobra"
)

var params = struct {
}{}

// コマンド作成
var Cmd = &cobra.Command{
	Use:   "time",
	Short: "時間処理",
	Args:  cobra.ArbitraryArgs, // 引数設定(ArbitraryArgs: なんでもOK)
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
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

	// 配下コマンドの追加
	Cmd.AddCommand(sleep.Cmd)
}
