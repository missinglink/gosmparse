package gosmparse

import (
	"log"
	"os"
	"path/filepath"
)

// AutoloadIndex - automatically load index file if one if available
func (d *Decoder) AutoloadIndex() {
	idxPath, _ := filepath.Abs(d.r.Name() + ".idx")
	if _, err := os.Stat(idxPath); err == nil {
		if nil == d.Index {
			log.Println("autoload idx:", idxPath)
			d.Index = &BlobIndex{}
			d.Index.ReadFromFile(idxPath)
		}
	}
}

// AutoSaveIndex - automatically save index file if feature is enabled
func (d *Decoder) AutoSaveIndex() {
	idxPath, _ := filepath.Abs(d.r.Name() + ".idx")
	log.Println("autosave idx:", idxPath)
	d.Index.WriteToFile(idxPath)
}
