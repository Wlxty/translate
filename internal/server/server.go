package server
import(
	"translateapp/internal/languagesstore"
)
type taskServer struct {
	store *languagesstore.LanguagesStore
}

func NewTaskServer() *taskServer {
	store := languagesstore.New()
	return &taskServer{store: &store}
}

