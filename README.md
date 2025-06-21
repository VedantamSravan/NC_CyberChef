# 🧪 CyberChef Dashboard Server

This project sets up a secure HTTPS server in Go to serve a lightweight AngularJS-based dashboard that embeds the CyberChef tool. It includes a simple navigation sidebar to access CyberChef and a status panel.

---

## 📁 Project Structure

```
project/
├── main.go                  # Go HTTPS server with embedded CyberChef
├── cert.pem                 # TLS certificate (self-signed or CA)
├── key.pem                  # TLS private key
├── index.html               # Main AngularJS app with sidebar and routing
├── static/
│   └── controller.js        # AngularJS controller for routing
└── cyberchef/               # Embedded CyberChef application
```

---

## 🚀 Features

- Serves CyberChef securely over HTTPS
- Uses Go's `embed` package to bundle the CyberChef frontend
- AngularJS UI with routes for:
  - CyberChef tool
  - Status panel
- Easy navigation via sidebar

---

## 🔧 Requirements

- Go 1.16+ (for `embed`)
- `cert.pem` and `key.pem` in the root directory

---

## 🔐 TLS Setup (Self-Signed Certificate)

To generate a self-signed certificate:

```bash
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
```

---

## 🛠️ Build & Run

```bash
go run main.go
```

Then open your browser to:

- **Dashboard**: [https://localhost:8443](https://localhost:8443)
- **CyberChef**: [https://localhost:8443/#/cyberchef](https://localhost:8443/#/cyberchef)
- **Status Panel**: [https://localhost:8443/#/status](https://localhost:8443/#/status)

---

## 📜 Example AngularJS App Behavior

### Sidebar Navigation

- **CyberChef**:  
  Loads `/cyberchef/index.html` in an iframe

- **Status**:  
  Displays basic HTTPS and availability message

### Routes (`controller.js`)

```js
$routeProvider
  .when("/cyberchef", {
    template: '<iframe src="/cyberchef/" title="CyberChef" style="width:100%;height:90vh;border:none;"></iframe>'
  })
  .when("/status", {
    template: '<h2>Server Status</h2><p>CyberChef is available and server is running securely over HTTPS.</p>'
  })
  .otherwise({
    redirectTo: "/cyberchef"
  });
```

---

## 📦 Embedding Notes

The CyberChef directory is embedded using Go's `embed.FS`:

```go
//go:embed cyberchef/*
var cyberchefFS embed.FS
```

---

## 📃 License

MIT or public domain — this project is intended for demonstration or internal tooling use.

---

## 👤 Author

Built by [Your Name]. Contributions are welcome!
