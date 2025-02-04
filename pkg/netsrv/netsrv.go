package netsrv

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"unicode"
	"web_crawler/pkg/index"
)

func handler(conn net.Conn, index *index.InvIndex) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		i := 0
		for ; i < len(buf) && unicode.IsLetter(rune(buf[i])); i++ {
		}
		query := string(buf[:i])
		if query == "e" {
			return
		}

		query = strings.ToLower(query)
		queryRes := index.GetDocuments(query)

		for _, doc := range queryRes {
			conn.Write([]byte(doc.Title + string("\n\r")))
			if err != nil {
				return
			}
		}
		conn.Write([]byte{0})

	}

}

func ListenAndServe(port string, index *index.InvIndex) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection accepted")
		go handler(conn, index)
	}
}
