package api

import (
	"encoding/json"
	"net/http"

	"github.com/AbhishekCS3459/goGraph/internal/graph"
)

type Graph struct {
	Edges map[int][]int `json:"edges"`
}
type RequestBody struct {
	Graph     Graph `json:"graph`
	StartNode int   `json:"startNode"`
	EndNode   int   `json:"endNode"`
}

func FindPathsHandler(w http.ResponseWriter, r *http.Request) {
	var req RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invallid request body", http.StatusBadRequest)
	}

	// validate input
	if len(req.Graph.Edges) == 0 {
		http.Error(w, "Empty Graph", http.StatusBadRequest)
	}

	if _, ok := req.Graph.Edges[req.StartNode]; !ok {
		http.Error(w, "Start Node not found", http.StatusBadRequest)
	}
	if _, ok := req.Graph.Edges[req.EndNode]; !ok {
		http.Error(w, "End NOde not found", http.StatusBadRequest)
	}
	allPath := [][]int{}

	graph.DFS(req.StartNode, req.EndNode, req.Graph.Edges, []int{}, &allPath)

	if len(allPath) == 0 {
		http.Error(w, "No path found between start and end node", http.StatusNotFound)
		return
	}
	resp := map[string]interface{}{
		"paths": allPath}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}


// an health check handler
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am alive"))
}
