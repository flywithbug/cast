package file_manager

import "os"


func ReadAll(name string)([]byte,error)  {
	fb,err := os.Open(name)
	if err != nil {
		return nil,err
	}
	defer fb.Close()
	fileInfo ,err := fb.Stat()
	if err != nil {
		return nil,err
	}
	buffer := make([]byte,fileInfo.Size())
	_,err = fb.Read(buffer)
	if err != nil {
		return nil,err
	}
	return buffer,nil
}