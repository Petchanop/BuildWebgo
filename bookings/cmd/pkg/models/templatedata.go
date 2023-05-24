/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   templatedata.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/05/22 14:54:36 by npiya-is          #+#    #+#             */
/*   Updated: 2023/05/22 14:54:57 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
