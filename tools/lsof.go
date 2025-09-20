package tools

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"unicode/utf8"
)

const FILE_ACCESS_MODE = 'a'
const PROCESS_COMMAND_NAME = 'c'
const FILE_STRUCTURE_SHARE_COUNT = 'C'
const FILE_DEVICE_CHARACTER_CODE = 'd'
const FILE_DEVICE_NUMBER = 'D'
const FILE_DESCRIPTOR = 'f'
const FILE_STRUCTURE_ADDRESS = 'F'
const FILE_FLAGS = 'G'
const PROCESS_GROUP = 'g'
const INODE = 'i'
const TASK_ID = 'K'
const LINK_COUNT = 'k'
const FILE_LOCK_STATUS = 'l'
const PROCESS_LOGIN_NAME = 'L'
const MARKER = 'm'
const TASK_COMMAND_NAME = 'M'
const FILE_NAME = 'n'
const NODE_ID = 'N'
const FILE_OFFSET = 'o'
const PROCESS_ID = 'p'
const PROTOCOL = 'P'
const DEVICE_NUMBER = 'r'
const PARENT_PROCESS = 'R'
const FILE_SIZE = 's'
const FILE_STREAM_ID = 'S'
const FILE_TYPE = 't'
const TCP_IP_INFO = 'T'
const PROCESS_USER_ID = 'u'

func Lsof(pid int) []map[rune]string {

	cmd := exec.Command("lsof", "-p", fmt.Sprint(pid), "-F")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	return parseFiles(stdout)
}

func parseFiles(reader io.Reader) []map[rune]string {
	scanner := bufio.NewScanner(reader)

	files := []map[rune]string{}
	currentFile := map[rune]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			c, lineDetail := trimFirstRune(line)
			switch c {
			case FILE_DESCRIPTOR:
				files = append(files, currentFile)
				currentFile = map[rune]string{}
				currentFile[c] = lineDetail

			default:
				currentFile[c] = lineDetail
			}
		}
	}

	return files
}

func trimFirstRune(s string) (rune, string) {
	r, i := utf8.DecodeRuneInString(s)
	return r, s[i:]
}
