//go:build windows

package gamemeta

import "github.com/bi-zone/go-fileversion"

func readMeta(exePath string) (GameMetaInfo, error) {
	// 声明
	result := GameMetaInfo{}
	// 读取
	f, err := fileversion.New(exePath)
	if err != nil {
		return result, err
	}
	// 赋值
	result.ExecutableFileVersion = f.FileVersion()
	result.ExecutableProductVersion = f.ProductVersion()

	// 返回
	return result, nil
}
