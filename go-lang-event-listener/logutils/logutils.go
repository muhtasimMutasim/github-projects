package logutils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func checkDirs(FILESinDIRthresHold int) string {
	// Function will check directories for file threshold

	var dirPaths [3]string

	dirPaths[0] = "/home/ubuntu/data/1"
	dirPaths[1] = "/home/ubuntu/data/2"
	// dirPaths[2] = "/home/ubuntu/data/3"

	for i := 0; i < len(dirPaths); i++ {

		currentDirPath := dirPaths[i]
		currentDirectory, _ := ioutil.ReadDir(currentDirPath)

		if len(currentDirectory) != FILESinDIRthresHold {
			//return dirOneFiles[len(dirOneFiles)]
			fmt.Printf("\n\nDirectory has not hit threshold\n\n")
			return currentDirPath
		} else if len(currentDirectory) == FILESinDIRthresHold {
			fmt.Printf("\n\nDirectory IS AT THRESHOLD\n\n")
			continue
		}
	}
	return ""
}

func IsDirEmpty(name string) (bool, error) {
	// Will Check if Directory is Empty.
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// read in ONLY one file
	_, err = f.Readdir(1)

	// if Error equals "EOF" means there is no directory
	if err == io.EOF {
		return true, nil
	}
	// returns false if there are files within directory
	return false, err
}

func WriteToEventLogFile(logData string) {
	// Returns a list of the current valid directories to write to
	numberOfFilesThreshold := 20
	fileSizeThreshold := 1205597
	targetDirectory := checkDirs(numberOfFilesThreshold)
	var existingFile string = ""
	var newEventLogFileName string

	fmt.Printf("Target Directory: %v \n\n", targetDirectory)

	dirValue, err := IsDirEmpty(targetDirectory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error with checking if Directory is empty %v", err.Error())
		os.Exit(1)
		//return false, err
	}

	// Check size of string

	if dirValue {
		// create new file for log ingestion
		currentTimeStamp := time.Now().Format("20060102EST150405.00") //  2006 01 02 EST 15 04 05 .00
		newEventLogFileName = targetDirectory + "/" + currentTimeStamp + ".log"
		fmt.Printf("Target PATH: %v \n\n", newEventLogFileName)

		f, err := os.OpenFile(newEventLogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(logData); err != nil {
			log.Println(err)
		}
	}

	if !dirValue {
		// Get list of Files from directory
		files, err := ioutil.ReadDir(targetDirectory)
		checkError("Error reading target directory", err)

		for _, f := range files {
			fileName := f.Name()
			fileSize := f.Size()
			fmt.Printf("\nFile Name: %v  Size: %v\n\n", fileName, fileSize)
			if fileSize < int64(fileSizeThreshold) {
				existingFile = targetDirectory + "/" + fileName
				break
			}
		}
		if existingFile == "" {
			/* if the target file's size is bellow the threshold but writing
			new data will make the file size exponentially greater than threshold */

		} else {
			f, err := os.OpenFile(existingFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(logData); err != nil {
				log.Println(err)
			}
		}

	}

	//return true, nil
}

func checkError(error_message string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %s", error_message, err.Error())
		os.Exit(1)
	}
}
