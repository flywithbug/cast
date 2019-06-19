package file_manager

import (
	"os"
	"errors"
	"strings"
)



func WriteFileString(path ,content string,cover bool)error  {
	if !cover{
		ex,_ := PathExists(path)
		if ex {
			return errors.New("PathExists")
		}
	}
	fil,err := os.Create(path)
	if err != nil {
		err = CreatePath(path)
		if err != nil {
			return err
		}
		return WriteFileString(path,content,cover)
	}
	defer fil.Close()
	fil.WriteString(content)
	return nil
}

func WriteFile(path string,content []byte,cover bool)error  {
	if !cover{
		ex,_ := PathExists(path)
		if ex {
			return errors.New("PathExists")
		}
	}
	fil,err := os.Create(path)
	if err != nil {
		err = CreatePath(path)
		if err != nil {
			return err
		}
		return WriteFile(path,content,cover)
	}
	defer fil.Close()
	fil.Write(content)
	return nil
}

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

func createPath(filePath string)error  {
	b,err := PathExists(filePath)
	if err != nil{
		return err
	}
	if !b{
		err := os.Mkdir(filePath,os.ModePerm)
		return err
	}
	return nil
}

func CreatePath(path string)error  {
	list := strings.Split(path,"/")
	for i := 1;i < len(list);i ++ {
		path := strings.Join(list[:i],"/") + "/"
		b,err := PathExists(path)
		if err != nil{
			return err
		}
		if !b {
			err = createPath(path)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func RemoveAll(path string)error  {
	return os.RemoveAll(path)
}