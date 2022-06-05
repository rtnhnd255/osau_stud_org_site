.PHONY: build
build:
	go build -o .\bin\app.exe .

.PHONY:run
run:
	.\bin\app.exe --servc=".\bin\server-config.json" --storc=".\bin\storage-config.json"

.PHONY: buru
buru:
	make build_win
	make run_win
