package create

import (
	"github.com/awisu2/goUtils/images"
	"github.com/spf13/cobra"
)

var params = struct {
	width  int
	height int
	out    string
}{}

// コマンド作成
var Cmd = &cobra.Command{
	Use:   "create",
	Short: "画像生成",
	Args:  cobra.ArbitraryArgs, // 引数設定(ArbitraryArgs: なんでもOK)
	Run: func(cmd *cobra.Command, args []string) {
		images.Create(&images.CreateOption{
			Size: images.Size{
				Width:  params.width,
				Height: params.height,
			},
			SaveOption: images.SaveOption{
				Path:    params.out,
				Format:  images.Jpg,
				Quality: 70,
			},
			Color: images.RgbaBlack,
		})

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

	flags.StringVarP(&params.out, "output", "o", "", "image output")
	flags.IntVarP(&params.width, "width", "w", 0, "image width")
	flags.IntVarP(&params.height, "height", "y", 0, "image height")

	// 必須指定
	requireds := []string{"output", "width", "height"}
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
