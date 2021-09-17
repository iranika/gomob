
Invoke-RestMethod -Uri "http://localhost:1323/dlsitesq" -Method Post -InFile "$PSScriptRoot/product-sample.json" -ContentType "application/json"
