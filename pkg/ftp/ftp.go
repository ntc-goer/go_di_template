package ftp

import (
	"crypto/tls"
	"fmt"
	"github.com/jlaffaye/ftp"
	"github.com/sirupsen/logrus"
	"go_di_template/config"
	"go_di_template/internal/util"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type FTP struct {
	Config *config.Config
	Client *ftp.ServerConn
}

func NewFTP(cfg *config.Config) *FTP {
	// Connect to the FTP server
	return &FTP{
		Config: cfg,
		Client: nil,
	}
}

func (f *FTP) Connect() error {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	ftpConnectionString := fmt.Sprintf("%s:%d", f.Config.FTP.Server, f.Config.FTP.Port)
	conn, err := ftp.Dial(ftpConnectionString, ftp.DialWithExplicitTLS(tlsConfig))
	if err != nil {
		log.Fatal("Failed to connect to FTP server:", err)
		return err
	}

	// Login to the FTP server
	err = conn.Login(f.Config.FTP.UserName, f.Config.FTP.Password)
	if err != nil {
		log.Fatal("Failed to login to FTP server:", err)
		return err
	}
	f.Client = conn
	return nil
}

func (f *FTP) Disconnect() {
	if f.Client != nil {
		f.Client.Quit()
		f.Client = nil
	}
}

func (f *FTP) List(dir string) ([]*ftp.Entry, error) {
	entries, err := f.Client.List(dir)
	if err != nil {
		log.Fatal("Failed to list files:", err)
		return nil, err
	}
	return entries, nil
}

func (f *FTP) DownloadFile(remoteFilePath, localDirPath string) error {
	// Check LocalDir Exist
	if !util.FolderExists(localDirPath) {
		err := os.Mkdir(localDirPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create local dir: %v", err)
		}
	}
	// Retrieve the remote file
	_, fileName := filepath.Split(remoteFilePath)
	// Create the local file
	localFilePath := filepath.Join(localDirPath, fileName)
	localFile, err := os.Create(localFilePath)
	if err != nil {
		return fmt.Errorf("failed to create local file: %v", err)
	}
	defer localFile.Close()

	resp, err := f.Client.Retr(remoteFilePath)
	if err != nil {
		return fmt.Errorf("failed to retrieve remote file: %v", err)
	}
	defer resp.Close()
	logrus.Infof("File downloaded successfully: %s", fileName)
	_, err = io.Copy(localFile, resp)
	if err != nil {
		return fmt.Errorf("failed to copy file content: %v", err)
	}
	return nil
}

func (f *FTP) DownloadDir(remoteDirPath, localDirPath string) error {
	// Check LocalDir Exist
	if !util.FolderExists(localDirPath) {
		err := os.Mkdir(localDirPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create local dir: %v", err)
		}
	}

	// List the contents of the remote folder
	entries, err := f.List(remoteDirPath)
	if err != nil {
		return fmt.Errorf("failed to list remote folder: %v", err)
	}

	// Download each file in the folder
	var wg sync.WaitGroup

	for _, entry := range entries {
		wg.Add(1)
		go func(et ftp.Entry) {
			defer wg.Done()
			if et.Type == ftp.EntryTypeFile {
				remoteFilePath := fmt.Sprintf("%s/%s", remoteDirPath, et.Name)
				err = f.DownloadFile(remoteFilePath, localDirPath)
				if err != nil {
					logrus.Errorf("DownloadFile %s error %s", et.Name, err)
				} else {
					logrus.Errorf("DownloadFile %s success", et.Name)
				}
			}
		}(*entry)
	}
	wg.Wait()
	fmt.Printf("Folder downloaded successfully: %s\n", remoteDirPath)
	return nil
}
