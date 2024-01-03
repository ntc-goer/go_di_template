package ftp

import (
	"github.com/stretchr/testify/assert"
	"go_di_template/config"
	"testing"
)

var cfg = &config.Config{
	FTP: config.FTP{
		Server:   "jpro2.jpo.or.jp",
		UserName: "vj_kinkan@receive",
		Port:     21,
		Password: "03-3294-7652",
	},
}

func TestFTP_List(t *testing.T) {
	ftp := NewFTP(cfg)
	err := ftp.Connect()
	defer ftp.Disconnect()
	assert.Nil(t, err)

	entries, err := ftp.List("/")
	assert.Nil(t, err)
	assert.Greater(t, len(entries), 0)
}

func TestFTP_DownloadFile(t *testing.T) {
	ftp := NewFTP(cfg)
	err := ftp.Connect()
	defer ftp.Disconnect()
	assert.Nil(t, err)

	err = ftp.DownloadFile("/1/9784010114995.jpg", "data")
	assert.Nil(t, err)
}

func TestFTP_DownloadDir(t *testing.T) {
	ftp := NewFTP(cfg)
	err := ftp.Connect()
	defer ftp.Disconnect()
	assert.Nil(t, err)

	err = ftp.DownloadDir("/1", "data")
	assert.Nil(t, err)
}
