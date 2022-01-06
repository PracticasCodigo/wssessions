package ws 

import (
    "fmt"
    "log"
    "net/http"
	"github.com/gorilla/websocket"
	"encoding/json"
)

var Sessions = make(map[int]*websocket.Conn)

type LoginRequest struct {
    Client_id int `json:"client_id"`
	Type string `json:"type"`
	Message string `json:"message"`
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func saveSession(client_id *int,conn *websocket.Conn){
	Sessions[*client_id] = conn

	response := "se ha guardado correctamente la conexion"

    if err := conn.WriteMessage(websocket.TextMessage, []byte(response)); err != nil {
		log.Println(err)
            return
	}

}


func sendMessage(client_id *int,message *string){
	conn :=Sessions[*client_id]

    if err := conn.WriteMessage(websocket.TextMessage, []byte(*message)); err != nil {
		log.Println(err)
            return
	}
}



func reader(conn *websocket.Conn) {
    for {
    // read in a message
        _, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
		fmt.Println("Message type : " , string(p))
		// print out that message for clarity
        var result LoginRequest
		if err := json.Unmarshal(p, &result); err != nil {   // Parse []byte to go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}
		// aca va el switch que filtra todos los mensajes por tipos

		fmt.Println("Typo de mensaje " , result)

		switch result.Type {
		case "login":
			log.Println("Es tipo login")
			saveSession(&result.Client_id,conn)
			break
		case "message":
			log.Println("es de tipo mensaje")
			sendMessage(&result.Client_id,&result.Message)
		}
	
    }
}


func WsEndpoint(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }

    // upgrade this connection to a WebSocket
    // connection
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
    }
	reader(ws)

}

