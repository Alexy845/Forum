package _const

import "os"

var TemplatesDir = os.Getenv("TEMPLATES_DIR")
var HtmlDir = os.Getenv("HTML_DIR")
var Database = os.Getenv("DATABASE_DIR")
