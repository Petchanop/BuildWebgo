/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/05/15 15:01:18 by npiya-is          #+#    #+#             */
/*   Updated: 2023/05/22 15:55:20 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"BuildWebgo/cmd/pkg/config"
	"BuildWebgo/cmd/pkg/handlers"
	"BuildWebgo/cmd/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create templat cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	// http.HandleFunc("/", handler.Repo.HandlerHttp)
	// http.HandleFunc("/about", handler.Repo.AboutHttp)

	fmt.Printf("Starting server on port %s\n", portNumber)
	// if err := http.ListenAndServe(portNumber, nil); err != nil {
	// 	log.Fatal(err)
	// }
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
