package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"syscall"

	"cigol.com/mini-paas/common/utils"
)

func printMap(fileMap map[string][]string, writer io.WriteCloser) {
	for name, dirs := range fileMap {
		if len(dirs) > 1 {
			//			fmt.Println(name, dirs)

			writer.Write([]byte(name + "\t" + strconv.Itoa(len(dirs)) + "\r\n"))
			for _, dir := range dirs {
				writer.Write([]byte("\t" + dir + "\r\n"))
			}
			writer.Write([]byte("================================================\r\n"))
		}
	}
}

func list(m map[string][]string, exts *utils.Set, dir string) error {
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range fs {
		if f.IsDir() {
			list(m, exts, dir+"/"+f.Name())
		} else {
			name := f.Name()
			lowerName := strings.ToLower(name)
			i := strings.LastIndex(lowerName, ".")
			if i >= 0 {
				lowerName = lowerName[i:]
			}
			if exts.Contains(lowerName) {
				dirs, ok := m[name]
				//			d := dir + "/" + name
				if ok {
					dirs = append(dirs, dir)
				} else {
					dirs = []string{dir}
				}
				m[name] = dirs
			}
		}
	}
	return nil
}

func GetLogicalDrives() []string {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")
	n, _, _ := GetLogicalDrives.Call()
	s := strconv.FormatInt(int64(n), 2)
	var drives_all = []string{"A:", "B:", "C:", "D:", "E:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P:", "Q:", "R:", "S:", "T:", "U:", "V:", "W:", "X:", "Y:", "Z:"}
	temp := drives_all[0:len(s)]
	var d []string
	for i, v := range s {
		if v == 49 {
			l := len(s) - i - 1
			d = append(d, temp[l])
		}
	}
	var drives []string
	for i, v := range d {
		//		fmt.Println(i, v)
		drives = append(drives[i:], append([]string{v}, drives[:i]...)...)
	}
	return drives
}

func main() {
	//	var m map[string][]string = map[string][]string{"1": []string{"1.1", "1.2"}, "2": []string{"2.1", "2.2"}}
	//	printMap(m, os.Stdout)
	var m map[string][]string = map[string][]string{}
	//	var exts []string = []string{".js", ".css", ".html", ".doc", ".mp3", ".pdf", ".rmvb", ".mkv", ".mp4"}
	exts := utils.NewSet(".doc", ".mp3", ".pdf", ".rmvb", ".mkv", ".mp4")

	name := "c:/duplicate-files.txt"
	f, err := os.Create(name)
	if err != nil {
		fmt.Println("failed to create file", name)
		fmt.Println(err)
		return
	}
	defer f.Close()

	//	disks := GetLogicalDrives()
	//	for _, disk := range disks {
	//		list(m, disk+"/")
	//		fmt.Println(disk)
	//		list(m, "E:/DTLFolder")
	//	}
	list(m, exts, "d:/")
	list(m, exts, "e:/")

	//	dir := []string{"abc"}
	//	dir = append(dir, "def", "ghi")
	//	fmt.Println(dir)

	printMap(m, f)
}
