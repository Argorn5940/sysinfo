package root_test

import (
	"os"
	"testing"

	"github.com/Argorn5940/sysinfo/cmd/root" // rootパッケージをインポート
)

func TestRootCmd(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{args: []string{"cmd", "all"}, expected: ""},
		{args: []string{"cmd", "cpu"}, expected: ""},
		{args: []string{"cmd", "disk"}, expected: ""},
		{args: []string{"cmd", "memory"}, expected: ""},
		{args: []string{"cmd", "network"}, expected: ""},
	}
	for _, test := range tests {
		t.Run(test.args[1], func(t *testing.T) {
			//コマンドライン引数を設定
			os.Args = test.args
			//コマンドを実行
			err := root.RootCmd.Execute() // root.RootCmdを使用
			if err != nil {
				t.Fatalf("コマンド実行中にエラーが発生しました: %v", err)
			}

		})
	}
}
