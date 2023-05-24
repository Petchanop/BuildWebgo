/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   routes.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/05/22 15:10:08 by npiya-is          #+#    #+#             */
/*   Updated: 2023/05/24 18:25:03 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net/http"

	"github.com/Petchanop/bookings/cmd/pkg/config"
	"github.com/Petchanop/bookings/cmd/pkg/handlers"

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
