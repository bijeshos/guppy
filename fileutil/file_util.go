package fileutil

import (
	"bufio"
	gdu "github.com/bijeshos/guppy/dirutil"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//CreateFile creates a file in the target directory, including creating missing directories in the path
func CreateFile(targetDir, fileName string) error {
	dirErr := gdu.MkDirAll(targetDir)
	if dirErr != nil {
		return dirErr
	}
	_, err := os.Create(filepath.Join(targetDir, fileName))
	if err != nil {
		return err
	}
	return nil
}

//CopyFile copies a file from source path to target path
func CopyFile(srcPath, targetPath string, forceReplace bool) error {
	//open source file
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	//create sub directories at target if needed
	targetSubDir := filepath.Dir(targetPath)
	err = gdu.MkDirAll(targetSubDir)
	if err != nil {
		return err
	}

	//open target file
	target, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer target.Close()
	proceedCopy := false

	if forceReplace {
		proceedCopy = true
	} else {
		isSame, _ := isSameMetadata(srcPath, targetPath)
		proceedCopy = !isSame
	}
	if proceedCopy {
		//perform copying
		_, err = io.Copy(target, src)
		if err != nil {
			return err
		}
	}
	return nil
}

//MoveFile moves a file from source path to target path
func MoveFile(srcPath, targetPath string, forceReplace bool) error {

	//create sub directories at target if needed
	targetSubDir := filepath.Dir(targetPath)
	err := gdu.MkDirAll(targetSubDir)
	if err != nil {
		return err
	}

	proceedMove := false

	if forceReplace {
		proceedMove = true
	} else {
		isSame, _ := isSameMetadata(srcPath, targetPath)
		proceedMove = !isSame
	}
	if proceedMove {
		//perform move
		err := os.Rename(srcPath, targetPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func isSameMetadata(srcFilePath string, targetFilePath string) (bool, error) {
	srcFileInfo, srcErr := os.Stat(srcFilePath)
	targetFileInfo, targetErr := os.Stat(targetFilePath)

	if srcErr != nil {
		return false, srcErr
	}

	if targetErr != nil {
		if os.IsNotExist(targetErr) {
			return false, targetErr
		}
	}

	if strings.Compare(srcFileInfo.Name(), targetFileInfo.Name()) == 0 && srcFileInfo.Size() == targetFileInfo.Size() {
		return true, nil
	}
	return false, nil

}

//ReadFileContent read content of the input file to a string array
func ReadFileContent(inputFilePath string) ([]string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
