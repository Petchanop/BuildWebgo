# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2023/05/16 15:13:05 by npiya-is          #+#    #+#              #
#    Updated: 2023/05/16 17:46:32 by npiya-is         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

NAME= testserv

CC= go

CFLAGS=

DIR_WEB= cmd/web/

DIR_HANDLERS= cmd/pkg/handlers/

DIR_RENDER= cmd/pkg/render/

WEB= server.go \

PKG_HANDLERS= handlers.go \

PKG_RENDER= render.go \

SRCS = $(addprefix $(DIR_WEB), $(WEB)) \
	$(addprefix $(DIR_HANDLERS), $(PKG_HANDLERS)) \
	$(addprefix $(DIR_RENDER), $(PKG_RENDER)) \

run:
	$(CC) run $(SRCS)

test:
	$(CC) test $(SRCS) 