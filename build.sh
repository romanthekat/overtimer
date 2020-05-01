GOOS=linux GOARCH=amd64 go build -o overtimer_linux romangaranin.dev/overtimer && \
GOOS=darwin GOARCH=amd64 go build -o overtimer_mac romangaranin.dev/overtimer && \
GOOS=windows GOARCH=amd64 go build -o overtimer_win romangaranin.dev/overtimer