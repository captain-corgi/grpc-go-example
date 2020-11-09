package internal

import routeguide "github.com/captain-corgi/grpc-go-example/internal/svc/route_guide"

//StartServices run all services
func StartServices() {
	//Start basic service
	routeguide.Bootstrap()

	// //Start authentication service
	// auth.Bootstrap()
	// //Start user service
	// user.Bootstrap()
	// //Start role service
	// role.Bootstrap()
	// //Start role permission service
	// permission.Bootstrap()
}
