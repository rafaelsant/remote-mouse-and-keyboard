package queues

import (
	"encoding/json"
	"log"

	models "github.com/rafaelsant/remote-mouse-and-keyboard/models"
	services "github.com/rafaelsant/remote-mouse-and-keyboard/services"
)

func Init() {

	mCh, mQu := InitReceiverQueue("mouse-queue")
	mouse := getChannel(mCh, &mQu)
	defer mCh.Close()

	mBCh, mBQu := InitReceiverQueue("mouse-button-queue")
	mouseButton := getChannel(mBCh, &mBQu)
	defer mBCh.Close()

	kCh, kQu := InitReceiverQueue("keyboard-queue")
	keyboard := getChannel(kCh, &kQu)
	defer kCh.Close()

	go func() {
		for {
			select {
			case m := <-mouse:
				if len(m.Body) > 0 {
					log.Println(string(m.Body))
					var body models.BodyMouseResponse
					err := json.Unmarshal(m.Body, &body)
					if err != nil {
						log.Print(err)
					}
					services.MoveMouse(&body)
					log.Printf("Received a message: %s", m.Body)
				}
			case k := <-keyboard:
				if len(k.Body) > 0 {
					var body models.BodyKeyboardResponse
					err := json.Unmarshal(k.Body, &body)
					if err != nil {
						log.Print(err)
					}
					services.SendKey(&body)
					log.Printf("Received a message: %s", k.Body)
				}
			case mb := <-mouseButton:
				if len(mb.Body) > 0 {
					var body models.BodyMouseClick
					err := json.Unmarshal(mb.Body, &body)
					if err != nil {
						log.Print(err)
					}
					services.MouseClick(&body)
					log.Printf("Received a message: %s", mb.Body)
				}
			}
		}
	}()
	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
