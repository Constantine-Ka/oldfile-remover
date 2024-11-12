# oldfile-remover
Приложение для удаления файлов и папок. Проблема-в проекте не github.com/Constantine-Ka/ocr_jatcrm (репозиторий - private) не успел реализовать удаление файлов, после окончания работы с ними..

# Флаги
## folder
Работать с папками или файлами. Если ```true```,  флаг ```fileException``` не будет учитываться
```json
{
  "arg":"folder",
  "type": "bool",
  "default": false,
  "usage": "-folder"
} 
```
## date
Дата, до которой требуется действие. По умолчанию текущая дата. При непустом значении ```durationYear```,```durationMonth```,```durationWeek```,```durationDays```,```durationHours``` флаги не будут учитываться 
```json
{
  "arg":"date",
  "type": "string",
  "default": "time.Now()",
  "usage": "-date='21.02.2024'"
} 
```
## workdir
Папка, для которой будет выполняться программа
```json
{
  "arg":"workdir",
  "type": "string",
  "default": "false",
  "usage": "-workdir='/var/www/testdeleter/'"
} 
```
## fileException
Расширение файлов, которые будут удаляться.
```json
{
  "arg":"fileException",
  "type": "string",
  "default": "",
  "usage": "-fileException=png"
} 
```
---
Как давно были созданы файлы?
## durationYear
```json
{
  "arg":"durationYear",
  "type": "integer",
  "default": 0,
  "usage": "-durationYear=1"
} 
```
## durationMonth
```json
{
  "arg":"durationMonth",
  "type": "integer",
  "default": 0,
  "usage": "-durationMonth=2"
} 
```
## durationWeek
```json
{
  "arg":"durationWeek",
  "type": "integer",
  "default": 0,
  "usage": "-durationWeek=3"
} 
```
## durationDays
```json
{
  "arg":"durationDays",
  "type": "integer",
  "default": 0,
  "usage": "-durationDays=5"
} 
```
## durationHours
```json
{
  "arg":"durationHours",
  "type": "integer",
  "default": 0,
  "usage": "-durationHours=12"
} 
```

