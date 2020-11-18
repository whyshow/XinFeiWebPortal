package utils

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/nfnt/resize"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"hash"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// ClientIP 通过结合使用 ClientPublicIP和 ClientIP 来获取 IP 地址。
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0]); ip != "" {
		return ip
	}
	if ip := strings.TrimSpace(r.Header.Get("X-Real-Ip")); ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return "0.0.0.0"
}

//配置INFO日志
func Init_info_log() *logs.BeeLogger {
	log := logs.NewLogger(10000) // 创建一个日志记录器，参数为缓冲区的大小
	// 设置配置文件
	jsonConfig := `{
        "filename" : "logs/admin.txt",
         "maxlines" : 10000000,
         "daily" : false,
         "maxdays" : 90
    }`
	log.SetLogger("file", jsonConfig) // 设置日志记录方式：本地文件记录
	//log.EnableFuncCallDepth(true)     // 输出log时能显示输出文件名和行号（非必须）
	return log
}

//生成随机数 n位随机数字 用于表id   //缺点：如果太快可能会造成生成的数字会重复。
func Random(a int) string {
	num := 10
	for i := 1; i < a; i++ {
		num = num * 10
	}
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(rd.Intn(num)) //好像只有这样将int 类型转换成字符串类型才不乱码。
}

//判断文件夹是否存在，如果不存在则创建
func IsDir(path string) string {
	if _, err := os.Stat(path); err != nil {
		os.MkdirAll(path, os.ModePerm)
		return path
	}
	return path
}

// 判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 生成MD5
func MD5(str string) string {
	var h hash.Hash
	if h == nil {
		h := md5.New()
		h.Write([]byte(str))
		return hex.EncodeToString(h.Sum(nil))
	} else {
		h.Write([]byte(str))
		return hex.EncodeToString(h.Sum(nil))
	}
}

// 遍历所有文件夹内的文件  返回字符串数组
func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := fi.Name()
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return s, err
			}
		} else {
			fullName := fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

//解压操作 参数 一：要解压的文件所处于的位置及名称。参数二：解压后存放的位置
//已解决 非国产解压软件解压乱码问题
func DeCompress(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()
	var decodeName string
	for _, f := range zipReader.File {
		if f.Flags == 0 {
			//如果标致位是0  则是默认的本地编码   默认为gbk
			i := bytes.NewReader([]byte(f.Name))
			decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
			content, _ := ioutil.ReadAll(decoder)
			decodeName = string(content)
		} else {
			//如果标志为是 1 << 11也就是 2048  则是utf-8编码
			decodeName = f.Name
		}

		fpath := filepath.Join(destDir, decodeName)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//获得路径
func GetDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

//字符串截取
func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

// 生成压缩后的缩略图
// 通过路径获得图片文件（路径中包括文件名） -> 判断后缀名
// -> 根据不同的后缀名来压缩指定大小的图片 -> 将图片安装原图片路径和名称覆盖式保存下来。
func Thumbnail(imgFilePge string, width uint, height uint) error {
	//用到了第三方依赖包   go get github.com/nfnt/resize 或者 import "github.com/nfnt/resize"
	file, err := os.Open(imgFilePge)
	if err != nil {
		return err
	}
	//得到格式
	_, imgtype, _ := image.Decode(file)
	file.Close()
	//判断格式
	if imgtype == "png" {
		file, _ := os.Open(imgFilePge)
		img, err := png.Decode(file)
		if err != nil {
			return err
		}
		file.Close()
		//缩放指定的大小
		m := resize.Resize(width, height, img, resize.NearestNeighbor)
		//创建新的图片
		out, err := os.Create(imgFilePge)
		if err != nil {
			return err
		}
		defer out.Close()
		// 写入新的图片文件
		png.Encode(out, m)
	} else if imgtype == "jpg" || imgtype == "jpeg" {
		file, _ := os.Open(imgFilePge)
		img, err := jpeg.Decode(file)
		if err != nil {
			return err
		}
		file.Close()
		//缩放指定的大小
		m := resize.Resize(width, height, img, resize.NearestNeighbor)
		//创建新的图片
		out, err := os.Create(imgFilePge)
		if err != nil {
			return err
		}
		defer out.Close()
		// 写入新的图片文件
		jpeg.Encode(out, m, nil)
	}

	return nil
}

/**
 *  有BUG，第四页有问题
 */
//分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func Paginator(page, prepage int, nums int64) map[string]interface{} {
	var firstpage int //前一页地址
	var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		if totalpages == page {
			lastpage = page
		} else {
			lastpage = page + 1
		}
	}

	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages //总页数
	paginatorMap["firstpage"] = firstpage   //上一页
	paginatorMap["lastpage"] = lastpage     //下一页
	paginatorMap["currpage"] = page         //当前页
	return paginatorMap
}

// 转毫秒
func Millisecond(namo time.Duration) string {
	ms1 := namo / 1e6
	if ms1 > 0 {
		ms2 := namo % 1e6
		return strconv.Itoa(int(ms1)) + "." + strconv.Itoa(int(ms2)) + "ms"
	} else {
		ms1 := namo / 1e3
		ms2 := namo % 1e3
		return strconv.Itoa(int(ms1)) + "." + strconv.Itoa(int(ms2)) + "µs"
	}
	return "未知错误"
}
