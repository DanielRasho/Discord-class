{{sleep 4}}

{{deleteMessage 724368722743263234 ((dbGet 0 "1").Value) 1}} {{sleep 4}}

{{$a := sendMessageRetID nil ">>> Si tienes una **Duda o Pregunta** que compartir con tu maestra usa el siguiente **comando:**\n```yaml\nduda \"pregunta\" $ \"Comentario (Opcional)\"\n\nEjemplos:\nduda ¿Le gusto mi tarea? $ Me esforze muchísimo.\nduda ¿Cuales son los temas del examen?\n```\n**Nota:**\n>No olvides escribir `duda` al **Inicio**  y `$` entre tu **Pregunta** y el **Desarrollo** de está.\n>Revisa el historial de dudas por si tu pregunta ya fue respondida, y se paciente."}}

{{dbSet 0 "1" (toString $a)}}
