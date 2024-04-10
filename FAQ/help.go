{{deleteTrigger 0.5}}

{{- $error := sendMessageRetID nil (cembed "title" "**Necesita ayuda?**" "description" "Para __**responder una duda**__ escriba este comando en el chat, __**la Sintaxis es:**__\n```yaml\nrespuesta <ID> <Texto de respuesta>\n\nEjemplo:\nrespuesta 727713590458581032 Solo ejercicios pares \n```\n**Nota:**\n>El `ID` es el **número largo** al final de cada mensaje; poner ID inválidos dará `ERROR`.\n>Tambien puede usar este comando para editar respuestas, solo use el mismo ID." "color" 16768256) -}}

{{deleteMessage nil $error 120}}
