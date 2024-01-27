package server

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/lucastomic/web-generator/web-generator/internal/generator"
	"github.com/lucastomic/web-generator/web-generator/internal/input"
	"github.com/lucastomic/web-generator/web-generator/internal/logging"
	"github.com/lucastomic/web-generator/web-generator/internal/types"
)

type Server struct {
	listenAddr  string
	service     generator.Generator
	inputParser input.Reader
	logging     logging.Logger
}

func New(
	listenAddr string,
	service generator.Generator,
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
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleGenerationReq))
	http.ListenAndServe(s.listenAddr, nil)
}

func makeHTTPHandlerFunc(apiFn types.APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(10000000))

	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
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
	paths, err := s.service.GenerateAndGetPaths(ctx, pageData)
	if err != nil {
		return err
	}
	for path := range paths {
		_ = path
		// send file to infraestructure service
	}
	writeJSON(writer, 200, nil)
	return nil
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
