module example.com/filesrv/server

go 1.17

replace example.com/filesrv/server/serverlib => ./serverlib/

require example.com/filesrv/server/serverlib v0.0.0-00010101000000-000000000000

require github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
