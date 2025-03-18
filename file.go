package helpers

import (
	"os"
	"path/filepath"

	log "github.com/catalystgo/logger/cli"
)

type SaveFileOpt struct {
	override bool // true - override existing files (default is false)
}

func SaveFile(file string, data []byte, cfg *SaveFileOpt) error {
	if len(data) == 0 {
		log.Warnf("no data to write in file (%s) therefore skipping", file)
		return nil
	}

	_, err := os.Stat(file)
	if err == nil {
		if cfg.override {
			log.Warnf("override file (%s)", file)
		} else {
			log.Warnf("skip file (%s) => already exist", file)
			return nil
		}
	}

	err = os.MkdirAll(filepath.Dir(file), os.ModePerm)
	if err != nil {
		log.Errorf("mkdir file (%s) => %v", file, err)
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		log.Errorf("create file (%s) => %v", file, err)
		return err
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Errorf("close file (%s) => %v", file, err)
		}
	}()

	_, err = f.Write(data)
	if err != nil {
		log.Errorf("write file (%s) => %v", file, err)
		return err
	}

	log.Infof("createad %s", file)

	return nil
}
