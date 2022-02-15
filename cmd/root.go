package cmd

import (
	"github.com/awisu2/goUtils/cmd/images"
	"github.com/awisu2/goUtils/cmd/output"
	"github.com/awisu2/goUtils/cmd/time"
	"github.com/spf13/cobra"
)

var params = struct {
}{}

// コマンド作成(一番topのコマンドは(通常)rootCmdと呼ぶ)
var Cmd = &cobra.Command{
	Short: "go 言語でのutil環境",
	Args:  cobra.ArbitraryArgs, // 引数設定(ArbitraryArgs: なんでもOK)
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
	Cmd.AddCommand(images.Cmd)
	Cmd.AddCommand(time.Cmd)
	Cmd.AddCommand(output.Cmd)
}

// 参考実装では、cmd.Execute()で実行できるようにしている
func Execute() error {
	return Cmd.Execute()
}
