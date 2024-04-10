{{$mensaje := getMessage nil 727961705589506130}}

{{$embed := index $mensaje.Embeds 0}}

{{$IDusuario := reFind "<@[0-9]+>"  $embed.Description}}

{{$texto := joinStr "" "Esto es una prueba " $IDusuario}}

{{sendMessage nil (cembed "description" $texto)}}
