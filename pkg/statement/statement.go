package statement

import (
	"net/http"
	"strconv"

	"github.com/Gilmardealcantara/rinha/pkg/utils"
)

func GetStatement(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	_, err := strconv.Atoi(idPath)
	if err != nil {
		utils.WriteErrorJson(w, err, 400)
		return
	}	

}
