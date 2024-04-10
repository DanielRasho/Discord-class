{{/* Variables */}}

{{$desarrollo := ""}}

{{$titulo := "" }}

{{$info := .StrippedMsg}}

{{$avatar := .User.AvatarURL "256"}}

{{$nombre := ""}}

{{/* Verificación de si hay un apodo o no*/}}

{{if .Member.Nick}}{{$nombre = .Member.Nick}}	{{else}}{{$nombre = .User.Username}}{{end}} 

	

{{if not (hasRoleID 728601033302999164)}}

	{{addRoleID 728601033302999164}}

	{{removeRoleID 728601033302999164 10}}

	{{if $info}} {{/*Comprobación de existencia de argumentos*/}}

		{{deleteTrigger 0.7}}

		{{if (reFind "[$]+" $info)}} {{/*comprobación de 2 argumentos*/}}	

			{{$msID := sendMessageRetID nil (cembed "description" "cargando...")}}

		

			{{/*Variables para el Embed*/}}

			{{$titulo = index (split $info "$") 0}}

			{{$desarrollo = joinStr "" (index (split $info "$") 1) "\n\nUsuario:" .User.Mention}}

	

			{{$embed:= sdict 

			"author" (sdict "name" $nombre "icon_url" $avatar ) 

			"color" 14695195

			"Title" $titulo

			"description" $desarrollo

			"footer" (sdict "text" (toString $msID) )

			}}

 

			{{ editMessage nil $msID (cembed $embed)  }}

 

		{{else}}

			{{$titulo = $info}}

 

			{{$msID := sendMessageRetID nil (cembed "description" "cargando...")}}

			{{$embed:= sdict 

			"author" (sdict "name" $nombre "icon_url" $avatar ) 

			"color" 14695195

			"Title" $titulo

			"description" (joinStr "" "Usuario:" .User.Mention)

			"footer" (sdict "text" (toString $msID) )

			}}

			{{ editMessage nil $msID (cembed $embed) }}

		{{end}}

 

	{{else}}

	{{deleteTrigger 0.7}}{{$error := sendMessageRetID nil (cembed "title" "**ERROR:**" "description" "Necesitas escribir una pregunta." "color" 	16768256)}}{{deleteMessage nil $error 15}}

	{{end}}

{{else}}

```

ERROR: No puedes usar este comando ahora, espera 3 H entre cada uso.

Pero puedes usar los comandos: "editar" y "eliminar"

```

{{deleteResponse 25}}

{{end}}
