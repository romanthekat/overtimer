GOOS=linux GOARCH=amd64 go build -o overtimer_linux github.com/EvilKhaosKat/overtimer && \
GOOS=darwin GOARCH=amd64 go build -o overtimer_mac github.com/EvilKhaosKat/overtimer && \
GOOS=windows GOARCH=amd64 go build -o overtimer_win github.com/EvilKhaosKat/overtimer