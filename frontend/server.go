package frontend

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/spudtrooper/lyft/api"
	"github.com/spudtrooper/lyft/handlers"
	"github.com/spudtrooper/minimalcli/handler"
)

func ListenAndServe(ctx context.Context, client *api.Client, port int, host string) error {
	var hostPort string
	if host == "localhost" {
		hostPort = fmt.Sprintf("http://localhost:%d", port)
	} else {
		hostPort = fmt.Sprintf("https://%s", host)
	}

	locs := handlers.SourceLocations
	handlers := handlers.CreateHandlers(client)
	mux := http.NewServeMux()
	handler.Init(mux)
	if err := handler.AddHandlers(ctx, mux, handlers,
		handler.AddHandlersPrefix("api"),
		handler.AddHandlersKey("lyft"),
		handler.AddHandlersIndexTitle("unofficial lyft API"),
		handler.AddHandlersFooterHTML(`Details: <a target="_" href="//github.com/spudtrooper/lyft">github.com/spudtrooper/lyft</a>`),
		handler.AddHandlersSourceLinks(true),
		handler.AddHandlersSerializedSourceLocations(locs),
	); err != nil {
		return err
	}
	mux.Handle("/", http.RedirectHandler("/api", http.StatusSeeOther))

	log.Printf("listening on %s", hostPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		return err
	}

	return nil
}
