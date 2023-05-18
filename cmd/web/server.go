/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/05/15 15:01:18 by npiya-is          #+#    #+#             */
/*   Updated: 2023/05/17 15:54:45 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"BuildWebgo/cmd/pkg/config"
	handler "BuildWebgo/cmd/pkg/handlers"
	render "BuildWebgo/cmd/pkg/render"
	"fmt"
	"log"
	"net/http"
)

var portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create templat cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)
	render.NewTemplates(&app)
	http.HandleFunc("/", handler.Repo.HandlerHttp)
	http.HandleFunc("/about", handler.Repo.AboutHttp)
	fmt.Printf("Starting server on port %s\n", portNumber)
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Fatal(err)
	}
}
