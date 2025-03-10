package settings

import (
	"gopkg.in/yaml.v3"
	"os"
)

var cfgPath = "manifest/config/config.yml"

type server struct {
	Address          string `yaml:"address"`
	UploadDir        string `yaml:"upload_dir"`
	CoverFileMaxSize int64  `yaml:"cover_file_max_size"`
	ImageFileMaxSize int64  `yaml:"image_file_max_size"`
}

type mysql struct {
	Link     string `yaml:"link"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Dbname   string `yaml:"dbname"`
}

type database struct {
	Mysql mysql `yaml:"mysql"`
}

type setting struct {
	Server   server   `yaml:"server"`
	Database database `yaml:"database"`
}

var Settings = &setting{}

func init() {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(data, &Settings); err != nil {
		panic(err)
	}
}
