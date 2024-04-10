{{$args := parseArgs 0 "" (carg "int" "ID del mensaje") (carg "string" "respuesta")}}

 

{{/*Variables*/}}

{{$msgID := $args.Get 0}}

{{$respuesta := $args.Get 1}}

 

{{/*Chequeo de variables*/}}

{{if ($msgID)}}	

	{{if ($respuesta)}}

	{{$mensajeOriginal := getMessage nil $msgID}}

	{{$EmbedOriginal := index $mensajeOriginal.Embeds 0}}

	{{$serverID := 723196078706524171}}

	{{$URL:= joinStr "" "https://discordapp.com/channels/" $serverID "/" $mensajeOriginal.ChannelID "/" $mensajeOriginal.ID}}

 

		{{if ne ($EmbedOriginal.Color) 4831309}} {{/*Edición Msg si este no ha sido respondido (en rojo)*/}}

 

			{{$descripcion := reReplace "Usuario:<@[0-9]+>" $EmbedOriginal.Description ""}}

			{{$respuesta = joinStr "" $respuesta "\n\nUsuario:" (reFind "<@[0-9]+>" $EmbedOriginal.Description)}}{{/*Encontramos la mencion de usuario en descripción*/}}

		

			{{/*Creación del nuevo mensaje*/}}

			{{$EmbedRespuesta := cembed 

			"author" (sdict "name" $EmbedOriginal.Author.Name "icon_url" $EmbedOriginal.Author.IconURL)

			"Title" $EmbedOriginal.Title

			"color" 4831309

			"description" $descripcion

			"fields" (cslice (sdict "name" "Respuesta:" "value" $respuesta))

			"footer" (sdict "text" $EmbedOriginal.Footer.Text)}}

 

			{{deleteTrigger 0.7}}

			{{editMessage nil $msgID $EmbedRespuesta}}

 

 

			{{/*Mensaje de notificación*/}}	

			{{$usuario := reFind "<@[0-9]+>" $EmbedOriginal.Description}}

			{{$embedNoti := cembed "title" (joinStr "" "*" $EmbedOriginal.Title "*") "color" 51587 "description" (joinStr "" "Ha sido respondida <a:blobwave:480810966939074560> \n" "**[Link a tu pregunta](" $URL ")**" ) }}

			{{sendMessage 727663143307968532 (complexMessage "content" $usuario "embed" $embedNoti)}}

 

		{{else}}

			

 

			{{$descripcion := $EmbedOriginal.Description}}

			{{$IDuser := (index $EmbedOriginal.Fields 0).Value}}

			{{$respuesta = joinStr "" $respuesta "\n\nUsuario:" (reFind "<@[0-9]+>" $IDuser)}}{{/*Buscamos la mención a usuario en Field*/}}

		

			{{/*Creación del nuevo mensaje*/}}

			{{$EmbedRespuesta := cembed 

			"author" (sdict "name" $EmbedOriginal.Author.Name "icon_url" $EmbedOriginal.Author.IconURL)

			"Title" $EmbedOriginal.Title

			"color" 4831309

			"description" $descripcion

			"fields" (cslice (sdict "name" "Respuesta:" "value" $respuesta))

			"footer" (sdict "text" $EmbedOriginal.Footer.Text)}}

 

			{{deleteTrigger 0.7}}

			{{editMessage nil $msgID $EmbedRespuesta}}

 

 

			{{/*Mensaje de notificación*/}}	

			{{$usuario := reFind "<@[0-9]+>" ((index $EmbedOriginal.Fields 0).Value)}}

			{{$embedNoti := cembed "title" (joinStr "" "*" $EmbedOriginal.Title "*") "color" 13400552 "description" (joinStr "" "Atención!!! la respuesta a tu pregunta fue editada <a:100:480810944616988673>\n" "**[Link a tu pregunta](" $URL ")**") }}

			{{sendMessage 727663143307968532 (complexMessage "content" $usuario "embed" $embedNoti)}}

		{{end}}					

	

	{{$exito := sendMessageRetID nil ">>> Su mensaje fue respondido/editado exitosamente :ok_hand: !!!"}}

	{{deleteMessage nil $exito 8}}

 

	{{else}}

		{{deleteTrigger 0.7}}

		{{$error := sendMessageRetID nil (cembed "title" "**ERROR:**" "description" "Necesita __**Añadir Una Respuesta**__, la Sintaxis es:\n```yaml\nrespuesta <ID> <Texto de respuesta>\n\nEjemplo:\nrespuesta 727713590458581032 Solo ejercicios pares \n```" "color" 16768256)}}

		{{deleteMessage nil $error 45}}

	{{end}}

 

{{else}}

{{deleteTrigger 0.7}}

{{- $error := sendMessageRetID nil (cembed "title" "**ERROR:**" "description" "Necesita __**Añadir el ID y**__ de algún mensaje y __**Una Respuesta**__, la Sintaxis es:\n```yaml\nrespuesta <ID> <Texto de respuesta>\n\nEjemplo:\nrespuesta 727713590458581032 Solo ejercicios pares \n```" "color" 16768256) -}}

{{deleteMessage nil $error 45}}

{{end}}
