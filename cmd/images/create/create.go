package create

import (
	"image/color"

	"github.com/awisu2/goUtils/images"
	"github.com/spf13/cobra"
)

var params = struct {
	width  int
	height int
	out    string
	color  struct {
		red   uint8
		green uint8
		blue  uint8
		alpha uint8
	}
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
			Color: color.RGBA{
				R: params.color.red,
				G: params.color.green,
				B: params.color.blue,
				A: params.color.alpha,
			},
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

	// output path
	flags.StringVarP(&params.out, "output", "o", "", "image output")

	// size
	flags.IntVarP(&params.width, "width", "w", 0, "image width")
	flags.IntVarP(&params.height, "height", "y", 0, "image height")

	// color
	flags.Uint8VarP(&params.color.red, "red", "r", 0, "color red")
	flags.Uint8VarP(&params.color.green, "green", "g", 0, "color green")
	flags.Uint8VarP(&params.color.blue, "blue", "b", 0, "color blue")
	flags.Uint8VarP(&params.color.alpha, "alpha", "a", 255, "color alpha")

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
