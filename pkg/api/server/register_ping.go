package server

import (
	"net/http"

	"github.com/containers/libpod/pkg/api/handlers"
	"github.com/gorilla/mux"
)

func (s *APIServer) registerPingHandlers(r *mux.Router) error {

	r.Handle("/_ping", s.APIHandler(handlers.Ping)).Methods(http.MethodGet)
	r.Handle("/_ping", s.APIHandler(handlers.Ping)).Methods(http.MethodHead)

	// swagger:operation GET /libpod/_ping libpod libpodPingGet
	// ---
	//   summary: Ping service
	//   description: |
	//     Return protocol information in response headers.
	//     `HEAD /libpod/_ping` is also supported.
	//     `/_ping` is available for compatibility with other engines.
	//   tags:
	//   - system (compat)
	//   - system
	//   produces:
	//   - text/plain
	//   responses:
	//     200:
	//       description: Success
	//       schema:
	//         description: OK
	//         type: string
	//         example: "OK"
	//       headers:
	//         API-Version:
	//           type: string
	//           description: Max compatibility API Version the server supports
	//         BuildKit-Version:
	//           type: string
	//           description: Default version of docker image builder
	//         Docker-Experimental:
	//           type: boolean
	//           description: If the server is running with experimental mode enabled, always true
	//         Cache-Control:
	//           type: string
	//           description: always no-cache
	//         Pragma:
	//           type: string
	//           description: always no-cache
	//         Libpod-API-Version:
	//           type: string
	//           description: |
	//             Max Podman API Version the server supports.
	//             Available if service is backed by Podman, therefore may be used to
	//             determine if talking to Podman engine or another engine
	//         Libpod-Buildha-Version:
	//           type: string
	//           description: |
	//             Default version of libpod image builder.
	//               Available if service is backed by Podman, therefore may be used to
	//               determine if talking to Podman engine or another engine
	//     500:
	//       $ref: "#/responses/InternalError"
	r.Handle("/libpod/_ping", s.APIHandler(handlers.Ping)).Methods(http.MethodGet)
	r.Handle("/libpod/_ping", s.APIHandler(handlers.Ping)).Methods(http.MethodHead)
	return nil
}
