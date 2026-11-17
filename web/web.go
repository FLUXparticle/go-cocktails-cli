package web

import "embed"

// Files contains the browser client and its templates.
//
//go:embed *.html app.js
var Files embed.FS
