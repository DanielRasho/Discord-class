{{$embed := cembed "title" (joinStr " " "**Enhorabuena**" .User.Username "!!!") "description" "La **1.era Fase** de la verificación se realizo con éxito <a:blobwave:480810966939074560> <a:blobbounce:480810958621769756>" "color" 15573803}}

{{sendDM $embed}}

{{sleep 2.5}}

{{$recordatorio := sendMessageRetID nil (cembed "title" "Atención 📢 !!!" "description" "Todo salio bien , ahora **dirígete a:**\n\n👉<#723196268830261299>\n\nRecuerda que puedes hacerlo cliqueando el texto azul de arriba 👆 o en la barra de canales a la izquierda 👈." "color" 15573803)}}

{{deleteMessage nil $recordatorio 120}}

{{addRoleID 724697434621477065}}

{{removeRoleID 724697438635425954}}
