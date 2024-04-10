{{deleteTrigger 0}}

{{$var := cembed "title" "**HOLA DE NUEVO 😅 !!!**" "description" "Reacciona con los botones abajo de este mensaje 👇 según tu sección para obtener tu identificador, hazlo honestamente **NO habrá vuelta atrás.**\n\n🧪**- Ciencias y Letras**\n🎨**- Diseño Gráfico**\n\nCon ello se te dará acceso a los canales de tus materias. Las verás a la izquierda, en la barra de canales 👈 (si ocurre algún error comunícalo al correo del colegio)" "color" 4093106}}

{{$var2 := sendMessageRetID nil $var}}{{addMessageReactions nil $var2 "🧪" "🎨"}}
