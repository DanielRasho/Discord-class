{{if eq (.Reaction.Emoji.Name) "1️⃣"}}

{{sendDM (cembed "title" "**Comando:** `duda`" "description" "Útil si tienes una **Duda o Pregunta** que compartir con tu maestra. Esta es su sintaxis:\n```yaml\nduda \"pregunta\" $ \"Comentario (Opcional)\"\n\nEjemplos:\nduda ¿Le gusto mi tarea? $ Me esforzé muchísimo.\nduda ¿Cuales son los temas del examen?\nduda Que ejercicios son? $ Pares o impares.\n```\n\n**Nota:**\n>No olvides escribir `duda` al **Inicio**  y `$` entre tu **Pregunta** y el **Desarrollo** de está.\n>Revisa el historial de dudas por si tu pregunta ya fue respondida, y se paciente.\n>Este comando solo se puede usar cada `24h` y solo en canales de `dudas`, así que haz preguntas específicas." "color" 4093106 )}}

 

{{else if eq (.Reaction.Emoji.Name) "2️⃣"}}

{{sendDM (cembed "title" "**Comando:** `respuesta`" "description" "Útil para **maestros** que quieren __**responder una duda**__ escriba este comando en el chat, __**la Sintaxis es:**__\n```yaml\nrespuesta <ID> <Texto de respuesta>\n\nEjemplo:\nrespuesta 727713590458581032 Solo ejercicios pares \nrespuesta 728751113150398515 Es resumen de cap. 3\nrespuesta 728460971630133349 Examen para 16 de Mayo\n```\n**Nota:**\n>El `ID` es ese **número largo** al final de cada mensaje de duda; poner ID inválidos dará `ERROR`.\n>También puede usar este comando para editar sus respuestas, es la misma sintaxis, usa el `ID` de la pregunta que quieres editar." "color" 4093106)}}

 

{{else if eq (.Reaction.Emoji.Name) "3️⃣"}}

{{sendDM (cembed "title" "**Comando:** `editar`" "description" "Escribiste una pregunta pero te arrepentiste de algo y quieres corregirlo? Puedes hacerlo con este comando, Sintaxis:\n```yaml\nActualmente este comando no ha sido creado, ya que solo es una demo.\n```" "color" 4093106)}}

 

{{else if eq (.Reaction.Emoji.Name) "4️⃣"}}

{{sendDM (cembed "title" "**Otros Comandos...**" "description" "Contamos con estos comandos extras: \n\n> `ayuda`\n> despliega un mensaje con instrucciones para responder dudas (solo maestros)." "color" 4093106)}}

{{end}}
