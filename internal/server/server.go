package server

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/lucastomic/web-generator-service/internal/input"
	"github.com/lucastomic/web-generator-service/internal/logging"
	"github.com/lucastomic/web-generator-service/internal/types"
	webprocessor "github.com/lucastomic/web-generator-service/internal/webProcessor"
)

type Server struct {
	listenAddr  string
	service     webprocessor.WebProcessor
	inputParser input.Reader
	logging     logging.Logger
}

func New(
	listenAddr string,
	service webprocessor.WebProcessor,
	inputParser input.Reader,
	logging logging.Logger,
) Server {
	return Server{
		listenAddr,
		service,
		inputParser,
		logging,
	}
}

func (s *Server) Run() {
	http.HandleFunc("/", s.makeHTTPHandlerFunc(s.handleGenerationReq))
	s.logging.Info(fmt.Sprintf("Service running in %s", s.listenAddr))
	http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) makeHTTPHandlerFunc(apiFn types.APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(10000000))
	return func(w http.ResponseWriter, r *http.Request) {
		err := apiFn(ctx, w, r)
		if err != nil {
			s.logging.Request(ctx, r, http.StatusBadRequest)
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func (s Server) handleGenerationReq(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) error {
	pageData, err := s.inputParser.RetrieveInput(*req)
	if err != nil {
		return err
	}
	err = s.service.Process(ctx, pageData)
	if err != nil {
		return err
	}
	writeJSON(writer, http.StatusOK, map[string]any{"message": "Web generated successfully"})
	return nil
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
