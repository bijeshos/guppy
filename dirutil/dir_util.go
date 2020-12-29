package dirutil

import (
	"github.com/bijeshos/guppy/arrayutil"
	"os"
	"path/filepath"
)

//ReadFiles to read from dir
func ReadFiles(srcDir string, ignoreList []string) ([]string, error) {

	var files []string
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && arrayutil.IsPresent(ignoreList, info.Name()) {
			return filepath.SkipDir
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return files, nil
}

func ReadDirs(srcDir string, ignoreList []string) ([]string, error) {

	var dirs []string
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && arrayutil.IsPresent(ignoreList, info.Name()) {

			return filepath.SkipDir
		}
		if info.IsDir() && info.Name() != srcDir {
			dirs = append(dirs, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dirs, nil
}

//ReadAll to read and dir from source dir
/*func ReadAll(srcDir string, ignoreList []string) ([]string, error) {
	zap.S().Infow("reading files", "from", srcDir)
	var files []string
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && arrayutil.IsPresent(ignoreList, info.Name()) {
			zap.S().Infow("ignoring dir", "dir name", info.Name())
			return filepath.SkipDir
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}*/

//MkDirAll to create directories
func MkDirAll(targetDir string) error {
	_, err := os.Stat(targetDir)
	if err != nil {
		if os.IsNotExist(err) {

			err := os.MkdirAll(targetDir, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

//IsExist to check if the directory exists
func IsExist(dir string) (bool, error) {
	//check whether directory exists
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

//IsSame to check if both directories are same
func IsSame(srcDir string, targetDir string) (bool, error) {

	srcInfo, srcErr := os.Stat(srcDir)
	if srcErr != nil {
		return false, srcErr
	}
	targetInfo, targetErr := os.Stat(targetDir)
	if targetErr != nil {
		if os.IsNotExist(targetErr) {
			return false, nil
		}
		return false, targetErr
	}
	// check if source and target dirs are same
	if os.SameFile(srcInfo, targetInfo) {
		return true, nil
	} else {
		return false, nil
	}
}
