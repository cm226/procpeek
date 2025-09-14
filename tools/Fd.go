package tools

import (
	"fmt"
	"log"
	"os"
	"path"
)

type FD struct {
	FDName string
}

type File struct {
	FD
	Name string
}

type Socket struct {
	FD
	Src  string
	Dest string
	Port string
}

type Other struct {
	FdType string
}

type FDs struct {
	Files   []File
	Sockets []Socket
	Other   []Other
}

func Fd(pid int) func() FDs {
	return func() FDs {
		dirs, err := os.ReadDir(fmt.Sprintf("/proc/%d/fd", pid))

		if err != nil {
			log.Fatal(err)
		}

		var fds FDs
		for _, dir := range dirs {

			filepath := fmt.Sprintf("/proc/%d/fd/%s", pid, dir.Name())
			fileInfo, err := os.Stat(filepath)
			if err != nil {
				log.Fatal(err)
			}
			fdTypes(filepath, fileInfo, &fds)
		}
		return fds
	}
}

func fdTypes(filepath string, fileInfo os.FileInfo, fds *FDs) {

	mode := fileInfo.Mode()

	switch {
	case mode.IsRegular():
		linkSrc, err := os.Readlink(filepath)
		if err != nil {
			log.Fatal(err)
		}
		f := File{
			Name: linkSrc,
			FD: FD{
				FDName: path.Base(filepath),
			},
		}
		fds.Files = append(fds.Files, f)
	case mode.IsDir():
		break
	case mode&os.ModeSymlink != 0:
		fds.Other = append(fds.Other, Other{
			FdType: "link",
		})
	case mode&os.ModeNamedPipe != 0:
		fds.Other = append(fds.Other, Other{
			FdType: "pipe",
		})
	case mode&os.ModeSocket != 0:
		fds.Other = append(fds.Other, Other{
			FdType: "socket",
		})
	case mode&os.ModeDevice != 0 && mode&os.ModeCharDevice != 0:
		fds.Other = append(fds.Other, Other{
			FdType: "device",
		})
	case mode&os.ModeDevice != 0:
		fds.Other = append(fds.Other, Other{
			FdType: "device",
		})
	default:
		fds.Other = append(fds.Other, Other{
			FdType: "unknown",
		})
	}
}
