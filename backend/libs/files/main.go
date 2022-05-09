package files

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func ReplaceKeyInFile(src string, newStr string, key string) error {

	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	output := bytes.Replace(input, []byte(key), []byte(newStr), -1)

	if err = ioutil.WriteFile(src, output, 0666); err != nil {
		return err
	}

	return nil
}

func CopyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = file(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

func AppendEndFile(path string, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		return err
	}

	return nil
}

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

func RemoveFile(path string) error {
    err := os.Remove(path)
    if err != nil {
		return err
    }
	return nil
}

// func AddToTextToFileLine(path string, line int, text string) error {
// 	file, err := os.Open(path)
//     if err != nil {
//         return err
//     }

//     scanner := bufio.NewScanner(file)
// 	line_num := 1
// 	content := ""
//     for scanner.Scan() {
//         text += scanner.Text() + "\n"
// 		if line_num == line {
// 			content += text
// 		}
// 		line_num ++
//     }
	
//     if err := scanner.Err(); err != nil {
//         return err
//     }

// 	file.Close()
// 	RemoveFile(path)

// 	if err := AppendEndFile(path, content); err != nil {
// 		return err
// 	}

// 	return nil
// }

func file(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}