module github.com/iranika/gomob/cmd/gomob-server

go 1.14

replace github.com/iranika/gomob => ../../../gomob

require (
	cloud.google.com/go v0.75.0 // indirect
	cloud.google.com/go/firestore v1.4.0 // indirect
	github.com/andybalholm/cascadia v1.2.0 // indirect
	github.com/iranika/gomob v0.0.0-20200904052240-91de24715903
	github.com/labstack/echo/v4 v4.1.17
	golang.org/x/mod v0.4.1 // indirect
	golang.org/x/oauth2 v0.0.0-20210113205817-d3ed898aa8a3 // indirect
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/tools v0.1.0 // indirect
	google.golang.org/genproto v0.0.0-20210113195801-ae06605f4595 // indirect
	google.golang.org/grpc v1.35.0 // indirect
)
