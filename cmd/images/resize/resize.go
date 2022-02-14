package resize

import (
	"github.com/awisu2/goUtils/images"
	"github.com/spf13/cobra"
)

var params = struct {
	input   string
	out     string
	width   int
	height  int
	fit     int
	quality int
}{}

// コマンド作成
var Cmd = &cobra.Command{
	Use:   "resize",
	Short: "resize image",
	Args:  cobra.ArbitraryArgs, // 引数設定(ArbitraryArgs: なんでもOK)
	Run: func(cmd *cobra.Command, args []string) {
		img, err := images.Open(params.input)
		if err != nil {
			panic(err)
		}

		resizedImg, err := images.Resize(img, &images.Size{
			X: params.width,
			Y: params.height,
		}, &images.ResizeOption{
			Fit: images.Fit(params.fit),
		})
		if err != nil {
			panic(err)
		}

		err = images.Save(resizedImg, &images.SaveOption{
			Path:    params.out,
			Quality: params.quality,
		})
		if err != nil {
			panic(err)
		}
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
	flags.StringVarP(&params.input, "input", "i", "", "image output")

	// output path
	flags.StringVarP(&params.out, "output", "o", "", "image output")

	// size
	flags.IntVarP(&params.width, "width", "x", 0, "resize width")
	flags.IntVarP(&params.height, "height", "y", 0, "resize height")

	// jpeg quality
	flags.IntVarP(&params.quality, "quality", "q", 70, "quality of jpeg (0~100)")

	// fit type
	flags.IntVarP(&params.fit, "fit", "f", 0, "fit type (0: fill, 1: Contain, 2: Cover)")

	// 必須指定
	requireds := []string{"input", "output", "width", "height"}
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
