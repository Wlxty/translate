package server
import(
	"net/http"
	"log"
	"translateapp/internal/languages"
)

type taskServer struct {
	store *languages.Repository
}

func NewTaskServer() *taskServer {
	repository := languages.New()
	return &taskServer{store: &repository}
}
