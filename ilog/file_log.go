package ilog

import(
	"os"
	"sync"
    "strings"

	"time"
	"log"
	"path/filepath"
)

var(
	file_daliy = "daliy"
	file_single = "single"
	file_write_mutex *sync.Mutex
)



type FileLog struct{
	level int
	mode string
}

func newFileLog(log_level int) (*FileLog){
	fileLog := new(FileLog)
	fileLog.level = log_level
	file_write_mutex = new(sync.Mutex)
	return fileLog
}

func (self *FileLog)getFileLogName()(string, error){
	filename := "log_" + time.Now().Format("2006-01-02") + ".log"
	path, err := filepath.Abs("storage/logs/" + filename)
	if err != nil{
        return "", err
	}
	path = strings.Replace(path, "\\", "/", -1)
    path = strings.Replace(path, "\\\\", "/", -1)
    return path, nil
}

func (self *FileLog)getWriter()(*os.File, error){
	filename, err := self.getFileLogName()
	if err != nil{
        return nil, err
	}
	writer, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil{
        return nil, err
	}
	return writer, nil
}

func (self *FileLog) debug(v ...interface{}){
	if self.level <= Level_Debug{
		go func(){
			file_write_mutex.Lock()
			defer file_write_mutex.Unlock()
			writer, err := self.getWriter()
			if err != nil{
				return
			}
			defer writer.Close()
			logger := log.New(writer, "[Debug]", log.LstdFlags)
			logger.Print(v...)
		}()
	}
}

func (self *FileLog) debugf(format string,v ...interface{}){
	if self.level <= Level_Debug{
		go func(){
			file_write_mutex.Lock()
			defer file_write_mutex.Unlock()
			writer, err := self.getWriter()
			if err != nil{
				return
			}
			defer writer.Close()
			logger := log.New(writer, "[Debug]", log.LstdFlags)
			logger.Printf(format, v...)
		}()
	}
}

func (self *FileLog) info(v ...interface{}){
	if self.level <= Level_Info{
		go func(){
			file_write_mutex.Lock()
			defer file_write_mutex.Unlock()
			writer, err := self.getWriter()
			if err != nil{
				return
			}
			defer writer.Close()
			logger := log.New(writer, "[Info]", log.LstdFlags)
			logger.Println(v...)
		}()
	}
}

func (self *FileLog) infof(format string,v ...interface{}){
	if self.level <= Level_Info{
		go func(){
			file_write_mutex.Lock()
			defer file_write_mutex.Unlock()
			writer, err := self.getWriter()
			if err != nil{
				return
			}
			defer writer.Close()
			logger := log.New(writer, "[Info]", log.LstdFlags)
			logger.Printf(format, v...)
		}()
	}
}

func (self *FileLog) error(v ...interface{}){
	if self.level <= Level_Error{
		go func(){
			file_write_mutex.Lock()
			defer file_write_mutex.Unlock()
			writer, err := self.getWriter()
			if err != nil{
				return
			}
			defer writer.Close()
			logger := log.New(writer, "[Error]", log.LstdFlags)
			logger.Println(v...)
		}()
	}
}

func (self *FileLog) errorf(format string,v ...interface{}){
	if self.level <= Level_Error{
		go func(){
			file_write_mutex.Lock()
			defer file_write_mutex.Unlock()
			writer, err := self.getWriter()
			if err != nil{
				return
			}
			defer writer.Close()
			logger := log.New(writer, "[Error]", log.LstdFlags)
			logger.Printf(format, v...)
		}()
	}
}

func (self *FileLog) warn(v ...interface{}){
	if self.level <= Level_Warn{
		go func(){
			file_write_mutex.Lock()
			defer file_write_mutex.Unlock()
			writer, err := self.getWriter()
			if err != nil{
				return
			}
			defer writer.Close()
			logger := log.New(writer, "[Warn]", log.LstdFlags)
			logger.Println(v...)
		}()
	}
}

func (self *FileLog) warnf(format string,v ...interface{}){
	if self.level <= Level_Warn{
		go func(){
			file_write_mutex.Lock()
			defer file_write_mutex.Unlock()
			writer, err := self.getWriter()
			if err != nil{
				return
			}
			defer writer.Close()
			logger := log.New(writer, "[Warn]", log.LstdFlags)
			logger.Printf(format, v...)
		}()
	}
}

func (self *FileLog) fatal(v ...interface{}){
	if self.level <= Level_Fatal{
		go func(){
			file_write_mutex.Lock()
			defer file_write_mutex.Unlock()
			writer, err := self.getWriter()
			if err != nil{
				return
			}
			defer writer.Close()
			logger := log.New(writer, "[Fatal]", log.LstdFlags)
			logger.Fatalln(v...)
		}()
	}
}

func (self *FileLog) fatalf(format string,v ...interface{}){
	if self.level <= Level_Fatal{
		go func(){
			file_write_mutex.Lock()
			defer file_write_mutex.Unlock()
			writer, err := self.getWriter()
			if err != nil{
				return
			}
			defer writer.Close()
			logger := log.New(writer, "[Fatal]", log.LstdFlags)
			logger.Fatalf(format, v...)
		}()
	}
}