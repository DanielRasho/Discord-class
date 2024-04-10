{{$ccll := "🧪"}}

{{$dg := "🎨"}}

{{$emote := .Reaction.Emoji.Name}}

 

{{if eq $ccll $emote}}

	{{addRoleID 724364482616492162}}

	{{addRoleID 728095944561655908}}

	{{removeRoleID 724697434621477065}}

	{{sendDM (cembed "color" 3583304 "description" "<a:whoop6:448846256295378945> Haz recibido el rol de **Estudiante CCLL**")}}

	{{$mensaje := sendMessageRetID nil ">>> Perfecto!!!\nTodo salió bien, dirigete a 👉 <#724392000044007515> y escribe tu nombre completo y clave."}}

	{{deleteMessage nil $mensaje 45}}

{{else if eq $dg $emote}}

	{{addRoleID 724367077284249680}}

	{{addRoleID 728095944561655908}}

	{{removeRoleID 724697434621477065}}

	{{sendDM (cembed "color" 3583304 "description" "<a:whoop6:448846256295378945> Haz recibido el rol de **Estudiante DG**")}}

	{{sleep 2}}

	{{$mensaje := sendMessageRetID nil ">>> Perfecto!!!\nTodo salió bien, dirigete a 👉 <#724392000044007515> y escribe tu **nombre completo** y **clave.**"}}

	{{deleteMessage nil $mensaje 45}}

{{end}}
