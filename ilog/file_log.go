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

func (self *FileLog) Debug(v ...interface{}){
	if self.level <= Debug{
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

func (self *FileLog) Debugf(format string,v ...interface{}){
	if self.level <= Info{
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

func (self *FileLog) Info(v ...interface{}){
	if self.level <= Info{
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

func (self *FileLog) Infof(format string,v ...interface{}){
	if self.level <= Info{
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

func (self *FileLog) Error(v ...interface{}){
	if self.level <= Error{
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

func (self *FileLog) Errorf(format string,v ...interface{}){
	if self.level <= Error{
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

func (self *FileLog) Warn(v ...interface{}){
	if self.level <= Warn{
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

func (self *FileLog) Warnf(format string,v ...interface{}){
	if self.level <= Warn{
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

func (self *FileLog) Fatal(v ...interface{}){
	if self.level <= Fatal{
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

func (self *FileLog) Fatalf(format string,v ...interface{}){
	if self.level <= Fatal{
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