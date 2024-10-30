//go:build !windows

package common

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"syscall"
)

func Umask(mask int) int {
	return syscall.Umask(mask)
}

func CheckDirWritePermission(dir string) error {
	fileInfo, err := os.Stat(dir)
	if err != nil {
		return fmt.Errorf("cannot access directory: %w", err)
	}

	// Confirm that the path is a directory
	if !fileInfo.IsDir() {
		return fmt.Errorf("specified path is not a directory")
	}
	return checkWritePerm(fileInfo)
}

func CheckFileWritePermission(file string) error {
	fileInfo, err := os.Stat(file)
	if os.IsNotExist(err) {
		return CheckDirWritePermission(filepath.Dir(file))
	} else if err != nil {
		return fmt.Errorf("cannot access file: %w", err)
	}
	if fileInfo.IsDir() {
		return fmt.Errorf("specified path is a directory")
	}
	return checkWritePerm(fileInfo)
}

func checkWritePerm(fileInfo os.FileInfo) error {
	// Get file's system-specific data to access UID and GID
	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return fmt.Errorf("cannot retrieve file information")
	}

	// Get the current user's UID and GID
	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("cannot get current user: %w", err)
	}

	if currentUser.Uid == "0" {
		return nil
	}

	currentUID := fmt.Sprint(stat.Uid)
	currentGID := fmt.Sprint(stat.Gid)

	// Convert UID and GID to the appropriate format
	isOwner := currentUID == currentUser.Uid
	isGroupMember := currentGID == currentUser.Gid

	// Get the mode and check permissions
	mode := fileInfo.Mode()

	// Check if the owner has write permission and the user is the owner
	if isOwner && mode&0200 != 0 {
		return nil // Owner has write permissions
	}

	// Check if the group has write permission and the user is in the group
	if isGroupMember && mode&0020 != 0 {
		return nil // Group has write permissions
	}

	// Check if others have write permission
	if mode&0002 != 0 {
		return nil // Others have write permissions
	}

	return fmt.Errorf("permission denied to write to %s", fileInfo.Name())
}
