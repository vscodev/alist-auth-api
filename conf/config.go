package conf

import (
	"github.com/caarlos0/env/v11"
)

var Conf = &Config{}

type AliyunDrive struct {
	ClientID     string `json:"client_id,omitempty" env:"CLIENT_ID"`
	ClientSecret string `json:"client_secret,omitempty" env:"CLIENT_SECRET"`
}

type BaiduNetdisk struct {
	ClientID     string `json:"client_id,omitempty" env:"CLIENT_ID"`
	ClientSecret string `json:"client_secret,omitempty" env:"CLIENT_SECRET"`
}

type Disk115 struct {
	ClientID     string `json:"client_id,omitempty" env:"CLIENT_ID"`
	ClientSecret string `json:"client_secret,omitempty" env:"CLIENT_SECRET"`
}

type Drive123 struct {
	ClientID     string `json:"client_id,omitempty" env:"CLIENT_ID"`
	ClientSecret string `json:"client_secret,omitempty" env:"CLIENT_SECRET"`
}

type Onedrive struct {
	ClientID     string `json:"client_id,omitempty" env:"CLIENT_ID"`
	ClientSecret string `json:"client_secret,omitempty" env:"CLIENT_SECRET"`
}

type GoogleDrive struct {
	ClientID     string `json:"client_id,omitempty" env:"CLIENT_ID"`
	ClientSecret string `json:"client_secret,omitempty" env:"CLIENT_SECRET"`
}

type Dropbox struct {
	ClientID     string `json:"client_id,omitempty" env:"CLIENT_ID"`
	ClientSecret string `json:"client_secret,omitempty" env:"CLIENT_SECRET"`
}

type Box struct {
	ClientID     string `json:"client_id,omitempty" env:"CLIENT_ID"`
	ClientSecret string `json:"client_secret,omitempty" env:"CLIENT_SECRET"`
}

type YandexDisk struct {
	ClientID     string `json:"client_id,omitempty" env:"CLIENT_ID"`
	ClientSecret string `json:"client_secret,omitempty" env:"CLIENT_SECRET"`
}

type PCloud struct {
	ClientID     string `json:"client_id,omitempty" env:"CLIENT_ID"`
	ClientSecret string `json:"client_secret,omitempty" env:"CLIENT_SECRET"`
}

type Config struct {
	AliyunDrive  AliyunDrive  `json:"aliyun_drive" envPrefix:"ALIYUN_DRIVE_"`
	BaiduNetdisk BaiduNetdisk `json:"baidu_netdisk" envPrefix:"BAIDU_NETDISK_"`
	Disk115      Disk115      `json:"disk_115" envPrefix:"DISK_115_"`
	Drive123     Drive123     `json:"drive_123" envPrefix:"DRIVE_123_"`
	Onedrive     Onedrive     `json:"onedrive" envPrefix:"ONEDRIVE_"`
	GoogleDrive  GoogleDrive  `json:"google_drive" envPrefix:"GOOGLE_DRIVE_"`
	Dropbox      Dropbox      `json:"dropbox" envPrefix:"DROPBOX_"`
	Box          Box          `json:"box" envPrefix:"BOX_"`
	YandexDisk   YandexDisk   `json:"yandex_disk" envPrefix:"YANDEX_DISK_"`
	PCloud       PCloud       `json:"pcloud" envPrefix:"PCLOUD_"`
}

func (c Config) Public() Config {
	return Config{
		AliyunDrive: AliyunDrive{
			ClientID: c.AliyunDrive.ClientID,
		},
		BaiduNetdisk: BaiduNetdisk{
			ClientID: c.BaiduNetdisk.ClientID,
		},
		Disk115: Disk115{
			ClientID: c.Disk115.ClientID,
		},
		Drive123: Drive123{
			ClientID: c.Drive123.ClientID,
		},
		Onedrive: Onedrive{
			ClientID: c.Onedrive.ClientID,
		},
		GoogleDrive: GoogleDrive{
			ClientID: c.GoogleDrive.ClientID,
		},
		Dropbox: Dropbox{
			ClientID: c.Dropbox.ClientID,
		},
		Box: Box{
			ClientID: c.Box.ClientID,
		},
		YandexDisk: YandexDisk{
			ClientID: c.YandexDisk.ClientID,
		},
		PCloud: PCloud{
			ClientID: c.PCloud.ClientID,
		},
	}
}

func init() {
	if err := env.Parse(Conf); err != nil {
		panic(err)
	}
}
