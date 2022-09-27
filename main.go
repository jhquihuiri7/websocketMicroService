package main
import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)




func main() {
	router := mux.NewRouter()




	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	router.HandleFunc("/ws",func (w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(conn.RemoteAddr())
	}).Methods("GET","OPTIONS")
	//router.HandleFunc("/socket", Socket).Methods("GET","OPTIONS")

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)
}

