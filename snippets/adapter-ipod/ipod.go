package main

import "fmt"

type ArchivoDeMusica interface {
	Reproducir() string
}

// Mp3
type FileMp3 struct {
	name string
}

func NewFileMp3(name string) *FileMp3 {
	return &FileMp3{name: name}
}

func (f *FileMp3) PlayMp3() string {
	return "Reproduciendo archivo MP3. Nombre: " + f.name
}

// Ogg
type FileOgg struct {
	name string
}

func NewFileOgg(name string) *FileOgg {
	return &FileOgg{name: name}
}

func (f *FileOgg) PlayOgg() string {
	return "Reproduciendo archivo Ogg. Nombre: " + f.name
}

type AdapterMp3 struct {
	archivo *FileMp3
}

func NewAdapterMp3(mp3 *FileMp3) *AdapterMp3 {
	return &AdapterMp3{archivo: mp3}
}

func (amp3 *AdapterMp3) Reproducir() string {
	return amp3.archivo.PlayMp3()
}

type AdapterOgg struct {
	archivo *FileOgg
}

func NewAdapterOgg(ogg *FileOgg) *AdapterOgg {
	return &AdapterOgg{archivo: ogg}
}

func (aogg *AdapterOgg) Reproducir() string {
	return aogg.archivo.PlayOgg()
}

type Ipod struct {
	lista []ArchivoDeMusica
}

func NewIpod() *Ipod {
	return &Ipod{}
}

func (ipod *Ipod) Add(pista ArchivoDeMusica) {
	ipod.lista = append(ipod.lista, pista)
}

func (ipod *Ipod) Reproducir() {
	for _, p := range ipod.lista {
		fmt.Println(p.Reproducir())
	}
}