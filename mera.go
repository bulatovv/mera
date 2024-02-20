package main

import (
	"encoding/json"
	"log"
	"net"
	"github.com/influxdata/go-syslog/v3/rfc3164"
)

type logEntry struct {
    IP string `json:"ip"`
    URL string `json:"url"`
    Host string `json:"host"`
    Referrer string `json:"referrer"`
    UserAgent string `json:"user_agent"`
    Time string `json:"time"`
}

func main() {
    pc, err := net.ListenPacket("udp", ":1234")
    if err != nil {
        log.Fatal(err)
    }
    defer pc.Close()
    
    p := rfc3164.NewParser()


    for {
        buf := make([]byte, 2048)
        n, _, err := pc.ReadFrom(buf)
        if err != nil {
            log.Fatal(err)
        }

        msg, err := p.Parse(buf[:n])
        if err != nil {
            log.Fatal(err)
        }

        rfc3164 := msg.(*rfc3164.SyslogMessage)

        log.Println(*rfc3164.Message)

        var entry logEntry
        err = json.Unmarshal(
            []byte(*rfc3164.Message),
            &entry,
        )
        if err != nil {
            log.Fatal(err)
        }


    }

    

}
