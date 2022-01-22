GOOS=linux GOARCH=amd64 go build -o overtimer_linux github.com/romanthekat/overtimer && \
GOOS=darwin GOARCH=amd64 go build -o overtimer_mac github.com/romanthekat/overtimer && \
GOOS=windows GOARCH=amd64 go build -o overtimer_win github.com/romanthekat/overtimer