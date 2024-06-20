package main

func main() {
	laCumbita := NewFileMp3("La cumbita")
	mentiroso := NewFileOgg("Mentiroso")
	laCumbitaAdaptada := NewAdapterMp3(laCumbita)
	mentirosoAdaptado := NewAdapterOgg(mentiroso)
	ipod := NewIpod()
	ipod.Add(laCumbitaAdaptada)
	ipod.Add(mentirosoAdaptado)
	ipod.Reproducir()
}