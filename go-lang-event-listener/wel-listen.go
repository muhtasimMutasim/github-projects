package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"

	"github.com/wel-api/logutils"
)

const (
	CONN_HOST           = "0.0.0.0"
	CONN_PORT           = "1525"
	CONN_TYPE           = "tcp"
	FILESinDIRthresHold = 20
)

func mainOld() {

	logFile, err := os.OpenFile("/home/ubuntu/wel-api/go-logs/wel.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, ":"+CONN_PORT)
	checkError("Error Listening: ", err)

	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		// option to implement log output on both Stdout and to log file
		mw := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(mw)

		// Listen for an incoming connection.
		conn, err := l.Accept()
		checkError("Error Accepting: ", err)

		// Handle connections in a new goroutine.
		go handleMess(conn)

	}
}

func handleMess(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	for {
		fmt.Println("Received Command")
		cmd, err := rw.ReadString('\n')

		if err != nil {
			fmt.Println("Client left.")
			conn.Close()
			return
		}
		cmd = strings.Trim(cmd, "\n ")
		log.Println(cmd)
	}

}

func checkDirs() string {
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

func checkError(error_message string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %s", error_message, err.Error())
		os.Exit(1)
	}
}

func main() {
	logFile, err := os.OpenFile("/home/ubuntu/wel-api/go-logs/wel.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	aTest := ""
	fmt.Println(aTest)
	// logutils.IsDirEmpty("/home/ubuntu/wel-api/go-logs")
	logutils.WriteToEventLogFile(aTest)

}

/*

second process
	check each directory:
		5 miniutes  // checks each directory every 5 minutes for updates.

		check.script:
			- delta or updates within a file or directory
			- sending new data to s3
			- delete file that was parsed
			ex:
				- wel-1.log
					* gets new data
					* pushes to s3
					* clears data pushed

				updated wel-1.log file will be 0 mbs


golang process
	data/
		1/
			wel-1.log <- 0 mbs data
			wel-2.log <- 3 mbs data
		2/
			wel-1.log <- 3 mbs data
			wel-2.log <- 3 mbs data
		3/
			wel-1.log <- 3 mbs data
			wel-2.log <- 3 mbs data



*/
