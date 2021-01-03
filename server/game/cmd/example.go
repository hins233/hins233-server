package main

import (
	"html/template"
)



var index = template.Must(template.New("").Parse(`
<html>
<body>
<h1>Welcome to WebSocket test server!</h1>
<h4>Ready to Autobahn!</h4>
<a href="/report">Reports</a>
</body>
</html>
`))
