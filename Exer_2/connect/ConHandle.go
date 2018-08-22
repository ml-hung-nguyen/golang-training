package connect

import (
	"encoding/json"
	"example/Exer_2/model"
	"net/http"
)

type ConnectHandle struct {
	UC ConnectUseCaseInterface
}

func (h *ConnectHandle) ConnectHandle(w http.ResponseWriter, r *http.Request) {
	err := h.UC.ConnectUC()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.MessageResponse{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.MessageResponse{Message: "Connected Database"})
}

func NewConHandle(uc *ConnectUseCase) *ConnectHandle {
	return &ConnectHandle{
		UC: uc,
	}
}
