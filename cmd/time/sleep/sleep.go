package sleep

import (
	"github.com/awisu2/goUtils/cmd/images/create"
	"github.com/awisu2/goUtils/times"
	"github.com/spf13/cobra"
)

var params = struct {
	time int64
}{}

// コマンド作成
var Cmd = &cobra.Command{
	Use:   "sleep",
	Short: "一定時間待機",
	Args:  cobra.ArbitraryArgs, // 引数設定(ArbitraryArgs: なんでもOK)
	Run: func(cmd *cobra.Command, args []string) {
		times.SleepMilliSecond(int(params.time))
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
	flags, _ := Cmd.Flags(), Cmd.PersistentFlags()

	// output path
	flags.Int64VarP(&params.time, "time", "t", 0, "sleep time(millisecond)")

	// 配下コマンドの追加
	Cmd.AddCommand(create.Cmd)
}
