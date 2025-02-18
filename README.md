The screenshots are the captures from Wireshark. Note the difference in the actual data and the content-length value in the request. quic-go/http3 answers with 200 OK.

# How to replicate

 1. Complie main.go
 2. Run server
 3. Setup wireshark to capture HTTP/3 and HTTP/2 secrets (see below) and start capturing
 4. Edit attacker.py to your desired content-length and data, then run it
 5. Observe traffic in wireshark

# Commands to run services


 - attacker: `python3 attacker.py` (you may need to install qh3, `python3 -m pip install qh3`)
 - compile main.go: `go mod init go_server` `go mod tidy` `go build -o server`
 - run server: `sudo ./server`

# Setup Wireshark

 1. Edit attacker.py PATH_TO_TLS_SECRECTS (example "/home/USER/keys.log")
 2. Same path into Wireshark -> Edit -> Preferences -> (Left) Protocols -> TLS -> (Pre)-Master-Secret log filname
