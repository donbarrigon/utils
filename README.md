# Utils - Herramientas para parceros y desarrolladores

La idea es estandarizar y agilizar la forma de hacer APIs y aplicaciones web con Go. Esta colección de paquetes ahorra tiempo y código repetitivo, estableciendo una estructura consistente inspirada en Laravel. simplemente lo importás y listo, a trabajar en lo que realmente importa.

Esta librería está inspirada en la forma como Laravel hace las cosas pero adaptada para Go.

**Utils** fue diseñado específicamente para funcionar con el starter kit [New](https://github.com/donbarrigon/new), igual tambien funciona de forma independiente si la quieres usar sola o solo algun paquete en especifico.

## Instalación
```bash
go get github.com/donbarrigon/utils
```

---

## Paquetes Disponibles

### 🔐 Auth

Maneja toda la autenticación de tu aplicación. JWT, passwords, sesiones y toda esa vuelta de seguridad que siempre te toca hacer.

**Ejemplo de uso**
```go
import "github.com/donbarrigon/utils/auth"

// Generá tokens, maneja sessiones de usuarios
```

---

### ⚙️ Config

Para cargar y manejar toda la configuración de tu app. Variables de entorno, archivos de config, y esas cosas que siempre hay que configurar. Súper útil para no tener valores quemados en el código.

**Ejemplo de uso**
```go
import "github.com/donbarrigon/utils/config"

// Cargá tu .env y usá las variables donde las necesitás
```

---

### 🗄️ DB

Todo lo relacionado con bases de datos. Conexiones, pools, y las funciones básicas para no tener que estar escribiendo SQL a mano todo el tiempo. Compatible con MongoDB y lo que necesités.

**Ejemplo de uso**
```go
import "github.com/donbarrigon/utils/db"

// Conectate a tu base de datos y usala facil
```

---

### 🎯 Handler

Helpers para manejar los HTTP handlers de forma más fácil. Respuestas JSON, errores, middleware y toda esa vuelta de las APIs REST.

**Ejemplo de uso**
```go
import "github.com/donbarrigon/utils/handler"

// Crea el context para los controladores
```
---

### ❌ Herror

Manejo de errores personalizado y descriptivo, inspirado en `createError` de Nuxt. Va un pokito más allá del paquete estándar `errors` de Go, ofreciendo errores más detallados y fáciles de debuggear. Se llama `herror` porque Go ya tiene ocupados todos los nombres que yo quería y `httperror` me sonaba muy largo y confuso.

**Ejemplo de uso**
```go
import "github.com/donbarrigon/utils/herror"

// Manejá y crea errores estandarizados
```

---

### 🌐 Lang

Internacionalización y localización para tu app. Si querés que tu aplicación hable en varios idiomas, este paquete te resuelve esa vuelta sin tanto drama.

**Ejemplo de uso**
```go
import "github.com/donbarrigon/utils/lang"

// Traducí mensajes y manejá varios idiomas automaticamente
```

---

### 📝 Logs

Sistema de logging bien organizado. Para que sepás qué está pasando en tu app, debuggear más fácil y tener logs que sirvan de verdad, no puro `fmt.Println` por ahí regados.

**Ejemplo de uso**
```go
import "github.com/donbarrigon/utils/logs"

// Loguea todo lo que necesitás, con niveles y estructura
```

---

### 🔨 QB (Query Builder)

Un query builder que te facilita la vida para construir queries simplificando la horrible sintaxis del driver de mongodb en go.

**Ejemplo de uso**
```go
import "github.com/donbarrigon/utils/qb"

// Construí tus queries sin tanto dolor de cabeza
```

---
### 🌐 Server

Inicia un servidor HTTP o HTTPS con soporte para certificados automáticos mediante Let's Encrypt (autocert). Te olvidás de configurar manualmente certificados SSL y toda esa vuelta complicada del HTTPS.

**Instalación**
```bash
go get github.com/donbarrigon/utils/server
```

---

**Ejemplo de uso**
```go
import "github.com/donbarrigon/utils/server"

// Levantá tu servidor HTTP o HTTPS
```

---

### 🔤 Str

Helpers para manipular strings. Esto solo es para uso interno de otros paquetes si quieres miralo pero no hay nada interesante.

**Ejemplo de uso**
```go
import "github.com/donbarrigon/utils/str"

// Esto fuera de los paquetes que lo usan no tiene ninguna utilidad.
```

---

### ✅ Validate

Validación de datos de entrada. Forms, json o msgpack que necesitás validar antes de procesar la información.

**Ejemplo de uso**
```go
import "github.com/donbarrigon/utils/validate"

// Validá los datos que te mandan antes de hacer algo con ellos
```

---

## Dependencias Externas

Esta librería usa algunos paquetes externos que son una chimba:

- **[MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)** - Para conectarte a MongoDB sin complicaciones
- **[msgpack/v5](https://github.com/vmihailenco/msgpack)** - Serialización eficiente de datos, ideal para cachés y comunicación entre servicios
- **crypto** (paquete estándar de Golang) - Para todo lo relacionado con seguridad, cifrado y hashing

## Contribuciones

¿Tenés alguna idea bacana o encontraste un bug? ¡Mandá tu PR o abrí un issue! Parce todo aporte es bienvenido.

## Licencia

GNU AFFERO GENERAL PUBLIC LICENSE Version 3, 19 November 2007

## Créditos

Hecho con ❤️ por donbarrigon en Medellín, Antioquia 🇨🇴

---

**¿Dudas o problemas?** Abrí un issue en el repo aver que puedo hacer.
