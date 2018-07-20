__NOT PRODUCTION READY__

# BotManager
BotManager is a simplistic barebones manager for bots, primarily Discord for now. Currently it provides a server and desktop cli client that connect via gRPC. 

Make the service.pb.go file via `make proto`. Then build and run the server `make server`. Afterwards build and start your client `make client`.