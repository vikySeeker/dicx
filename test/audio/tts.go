package main

import htgotts "github.com/hegedustibor/htgo-tts"

func main() {
	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	speech.Speak("Threat hunting is a proactive and iterative process,")
}
