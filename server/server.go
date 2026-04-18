package server

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/donbarrigon/utils/config"
	"github.com/donbarrigon/utils/db"
	"github.com/donbarrigon/utils/handler"
	"github.com/donbarrigon/utils/logs"
	"golang.org/x/crypto/acme/autocert"
)

func Start(h *handler.Router) {
	config.LoadEnv()
	db.InitMongoDB()
	var tlsConfig *tls.Config

	if config.ServerHttpsEnabled {
		manager := &autocert.Manager{
			Cache:      autocert.DirCache(config.ServerHttpsCertPath), // Carpeta donde guarda los certs
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(strings.Split(config.ServerHostWhiteList, ",")...),
		}
		tlsConfig = manager.TLSConfig()
		go func() {
			// Redirigir todo lo que llegue por HTTP a HTTPS
			if er := http.ListenAndServe(":80", manager.HTTPHandler(nil)); er != nil {
				logs.Error("🔴💥 No se inicio el server en el puerto 80: %s", er.Error())
			}
		}()
	}

	server := &http.Server{
		Addr:         ":" + config.ServerPort,
		Handler:      h,
		TLSConfig:    tlsConfig,
		ReadTimeout:  time.Duration(config.ServerReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.ServerWriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(config.ServerReadTimeout+config.ServerWriteTimeout) * time.Second,
	}

	logs.Info(`🚀 Server running on %s 
  ____   ___  ____  ____  ___  ___  _   _ ____   ___
 / ___| / _ \|  _ \|  _ \|_ _|| __|| \ | |  _ \ / _ \
| |    | | | | |_) | |_) || ||||__ |  \| | | | | | | |
| |___ | |_| |  _ <|  _ < | ||||__ | |\  | |_| | |_| |
 \____(_)___/|_| \_\_| \_\___||___||_| \_|____/ \___/
`, config.AppURL)

	// funciona en dev pero en produccion es feo.
	// espera la señal en segundo plano, el bun run dev lo reinicia pero el main se termina y no salen los mensajes
	go HttpServerGracefulShutdown(server)
	time.Sleep(100 * time.Millisecond) // para que salga el mensaje de corriendo.

	if config.ServerHttpsEnabled {
		if e := server.ListenAndServeTLS("", ""); e != nil && e != http.ErrServerClosed {
			logs.Error("💥 Could not start server tls: %s", e.Error())
		}
	} else {
		if e := server.ListenAndServe(); e != nil && e != http.ErrServerClosed {
			logs.Error("💥 Could not start server: %s", e.Error())
		}
	}

	// funciona mejor en produccion en dev no.
	// server en segundo plano espera la señal de cierre, salen los mensajes pero el bun run dev no lo reinicia
	// go func() {
	// 	if Env.SERVER_HTTPS_ENABLED {
	// 		if err := server.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
	// 			PrintError("💥 Could not start server: :error", Entry{"error", err.Error()})
	// 		}
	// 	} else {
	// 		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 			PrintError("💥 Could not start server: :error", Entry{"error", err.Error()})
	// 		}
	// 	}
	// }()
	// HttpServerGracefulShutdown(server)
}

// maneja el apagado graceful del servidor
func HttpServerGracefulShutdown(server *http.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Espera por la señal de terminación
	<-sigChan
	logs.Info("⏻ Initiating controlled server shutdown...")

	// se cierra la conexion con mono db
	if e := db.CloseMongoDB(); e != nil {
		logs.Warning("💥 Error closing connection to MongoDB %s", e.Error())
	} else {
		logs.Info("🔌 Connection to MongoDB successfully closed")
	}

	// Crea un contexto con timeout para el shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//se cierra el servidor HTTP para que no acepte nuevas conexiones
	if e := server.Shutdown(ctx); e != nil {
		logs.Warning("⏻ Server forced to close: %s", e.Error())
	} else {
		logs.Info("⏻ HTTP server stopped successfully")
	}

	logs.Info("💀 Apagado controlado completado")
}
