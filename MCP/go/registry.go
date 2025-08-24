package main

import (
	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	tools_notes "github.com/input-api/mcp-server/tools/notes"
	tools_api "github.com/input-api/mcp-server/tools/api"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_notes.CreateGet_notes_idTool(cfg),
		tools_notes.CreatePut_notes_idTool(cfg),
		tools_notes.CreateDelete_notes_idTool(cfg),
		tools_api.CreateGetTool(cfg),
		tools_api.CreateGet_apiTool(cfg),
		tools_notes.CreateGet_notesTool(cfg),
		tools_notes.CreatePost_notesTool(cfg),
	}
}
