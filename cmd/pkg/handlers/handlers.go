/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   handlers.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/05/16 14:57:18 by npiya-is          #+#    #+#             */
/*   Updated: 2023/05/17 15:50:24 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package handlers

import (
	"BuildWebgo/cmd/pkg/config"
	render "BuildWebgo/cmd/pkg/render"
	"net/http"
)

//Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) AboutHttp(w http.ResponseWriter, req *http.Request) {
	render.RenderTemplates(w, "about.page.tmpl")
}

func (m *Repository) HandlerHttp(w http.ResponseWriter, req *http.Request) {
	render.RenderTemplates(w, "home.page.tmpl")
}
