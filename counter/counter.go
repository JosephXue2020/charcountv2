package counter

import (
	"log"
	"projects/office"
	"runtime"
	"strconv"
	"sync"
)

// CountDir counts all the files in the directory
func CountDir(direc string) [][]string {
	var result [][]string

	fInfo, err := office.GetPathInfo(direc)
	if err != nil {
		log.Println(err)
		return result
	}

	for _, inSli := range fInfo {
		inSliRet := []string{}
		inSliRet = append(inSliRet, inSli...)
		ext := inSli[1]
		if ext != ".docx" {
			inSliRet = append(inSliRet, "", "非.docx格式")
		} else {
			str, err := office.ReadDocx(inSli[2])
			if err != nil {
				inSliRet = append(inSliRet, "", "文件不能正常读取")
			} else {
				charNum := len([]rune(str))
				charNumStr := strconv.Itoa(charNum)
				inSliRet = append(inSliRet, charNumStr, "")
			}
		}
		result = append(result, inSliRet)
	}
	return result
}

var wg sync.WaitGroup

// CountDirMultiThread counts all the files in the directory by multithread
func CountDirMultiThread(direc string) [][]string {
	var result [][]string

	// Get file info
	fInfo, err := office.GetPathInfo(direc)
	if err != nil {
		log.Println(err)
		return result
	}

	// Count by multithread
	fNum := len(fInfo)

	taskChan := make(chan []string, fNum)
	for _, item := range fInfo {
		taskChan <- item
	}
	close(taskChan)

	resultChan := make(chan []string, fNum)

	threadNum := runtime.NumCPU() * 2
	wg.Add(threadNum)

	for i := 0; i < threadNum; i++ {
		go countFile(taskChan, resultChan)
	}

	wg.Wait()

	close(resultChan)

	// Get result from channel
	for r := range resultChan {
		result = append(result, r)
	}
	return result
}

func countFile(task chan []string, result chan<- []string) {
	for {
		tsk, ok := <-task
		if !ok {
			wg.Done() // closure
			break
		}

		inSliRet := []string{}
		inSliRet = append(inSliRet, tsk...)
		ext := tsk[1]
		if ext != ".docx" {
			inSliRet = append(inSliRet, "", "非.docx格式")
		} else {
			str, err := office.ReadDocx(tsk[2])
			if err != nil {
				inSliRet = append(inSliRet, "", "文件不能正常读取")
			} else {
				charNum := len([]rune(str))
				charNumStr := strconv.Itoa(charNum)
				inSliRet = append(inSliRet, charNumStr, "")
			}
		}
		result <- inSliRet
	}
}
