package main

import (
	"gopkg.in/gographics/imagick.v2/imagick"
	"os"
	"strconv"
)

func main() {
	ConvertPdfToImage("test.pdf", 800, 1212, 200, 85)
}

//ConvertPdfToImage 转换pdf为图片格式
//@resolution:扫描精度
//@CompressionQuality:图片质量: 1~100
func ConvertPdfToImage(bookname string, pageWidth uint, pageHeight uint, resolution float64, compressionQuality uint) (err error) {

	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	//defer mw.Clear()
	defer mw.Destroy()

	if err := mw.SetResolution(resolution, resolution); err != nil {
		println("扫描精度设置失败")
		return err
	}

	if err := mw.ReadImage(bookname); err != nil {
		println("文件读取失败")
		return err
	}

	var pages = int(mw.GetNumberImages())
	println("页数:", pages)

	//裁剪会使页数增加
	addPages := 0
	path := ""
	for i := 0; i < pages; i++ {
		mw.SetIteratorIndex(i) // This being the page offset

		//压平图像，去掉alpha通道，防止JPG中的alpha变黑,用在ReadImage之后
		if err := mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_FLATTEN); err != nil {
			println("图片")
			return err
		}

		mw.SetImageFormat("jpg")
		mw.SetImageCompression(imagick.COMPRESSION_JPEG)
		mw.SetImageCompressionQuality(compressionQuality)

		//如果width>height ,就裁剪成两张
		pWidth := mw.GetImageWidth()
		pHeight := mw.GetImageHeight()

		//需要裁剪
		if pWidth > pHeight {

			//mw.ResizeImage(pageWidth*2, pageHeight, imagick.FILTER_UNDEFINED, 1.0)
			mw.ThumbnailImage(pageWidth*2, pageHeight)

			tempImage := mw.GetImageFromMagickWand()
			leftMw := imagick.NewMagickWandFromImage(tempImage) //由于返回的是指针,需要提前初始化,不然写完左半业tempImage就变了

			//左半页
			mw.CropImage(pageWidth, pageHeight, 0, 0)
			path = "./book/page" + strconv.Itoa(i+addPages) + ".jpg"
			mw.WriteImage(path)

			//右半页
			leftMw.SetImageFormat("jpg")
			leftMw.SetImageCompression(imagick.COMPRESSION_JPEG)
			leftMw.SetImageCompressionQuality(compressionQuality)
			leftMw.CropImage(pageWidth, pageHeight, int(pageWidth), 0)
			addPages++
			path = "./book/page" + strconv.Itoa(i+addPages) + ".jpg"
			leftMw.WriteImage(path)
			leftMw.Destroy()

		} else {

			mw.ThumbnailImage(pageWidth, pageHeight)
			path = "./book/page" + strconv.Itoa(i+addPages) + ".jpg"
			mw.WriteImage(path)

		}

	}

	println("转换完毕!")
	os.Exit(0) //模拟退出程序,删掉!
	return nil
}
