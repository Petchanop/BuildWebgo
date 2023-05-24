/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   routes.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/05/22 15:10:08 by npiya-is          #+#    #+#             */
/*   Updated: 2023/05/24 14:42:55 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"BuildWebgo/cmd/pkg/config"
	"BuildWebgo/cmd/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	//pat package routing
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.HandlerHttp))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.AboutHttp))

	//chir package routing
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.HandlerHttp)
	mux.Get("/about", handlers.Repo.AboutHttp)
	return mux
}
