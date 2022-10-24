package controller

import (
	"fmt"
	"image/png"
	_"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/image/webp"
)


const p = ".webp$"
var reg = regexp.MustCompile(p)

func WebpTrans() {
	//新增資料夾
	var newDir = "D:/comic1-trans/"
	os.Mkdir(newDir, 0777)

	//圖片所在的資料夾
	var folderDir = "D:/comic1"
	//取得圖檔清單
	fileInfoList, err := os.ReadDir(folderDir)
	fmt.Println(fileInfoList)
	if err != nil {
		fmt.Println(err)
	}

	// fnameChan = make(chan string, len(fileInfoList))
	//for loop 每份圖檔
	for _, fileInfo := range fileInfoList {
		//取得檔名
		var fileName = fileInfo.Name()
		fmt.Println(fileName)
		
		//判斷特定檔名的檔案，如果非特定檔名則跳過該檔案
		isMatch := reg.MatchString(fileName)
		if !isMatch {
			continue
		}
		//取得副檔名
		var fileExt = filepath.Ext(fileName)
		fmt.Println(fileExt)
		//去掉 filename由後往前且符合 fileExt 字串的字元 (去掉.副檔名)
		var fileNameOnly = strings.TrimSuffix(fileName, fileExt)
		fmt.Println(fileNameOnly)

		//webp檔案路徑
		webpFilename := folderDir + "/" + fileName
		fmt.Println(webpFilename)
		f0, err := os.Open(webpFilename)
		fmt.Println(f0)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f0.Close()
		//decode webp檔
		img0, err := webp.Decode(f0)
		if err != nil {
			fmt.Println(err)
			fmt.Println(err)
			return
		}
		
		pngFile, err := os.Create(newDir+ fileNameOnly + ".png")
		if err != nil {
			fmt.Println(err)
		}
		err = png.Encode(pngFile, img0)
		if err != nil {
			fmt.Println(err)
		}
		
	}
}