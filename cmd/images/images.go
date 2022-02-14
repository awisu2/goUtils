package images

import (
	"github.com/awisu2/goUtils/cmd/images/create"
	"github.com/awisu2/goUtils/cmd/images/resize"
	"github.com/spf13/cobra"
)

var params = struct {
}{}

// コマンド作成
var Cmd = &cobra.Command{
	Use:   "images",
	Short: "各種画像関係処理",
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
	Cmd.AddCommand(create.Cmd)
	Cmd.AddCommand(resize.Cmd)
}
