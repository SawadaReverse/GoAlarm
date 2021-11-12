package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	var SetTime string
	var nowTime string
	fmt.Println("Enter time (HH:mm:ss)")
	fmt.Scanf("%s", &SetTime)

	for {
		nowTime = time.Now().Format("15:04:05")
		if nowTime == SetTime {
			fmt.Println("It's time !!!")

			// 以下、 https://qiita.com/usk81/items/8590172a23bb71e21329 より引用
			// ファイル名だけ変えた
			f, err := os.Open("sound.mp3")
			if err != nil {
				log.Fatal(err)
			}
			st, format, err := mp3.Decode(f)
			if err != nil {
				log.Fatal(err)
			}
			defer st.Close()

			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

			done := make(chan bool)
			speaker.Play(beep.Seq(st, beep.Callback(func() {
				done <- true
			})))
			<-done

			return
		}

	}
}
