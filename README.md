# XGOFIND
The goal of the **xgofind** project is to create a flexible and efficient HTTP interceptor proxy that enables real-time inspection, modification, and manipulation of HTTP and HTTPS traffic. This tool primarily aims to assist security 
professionals, developers, and researchers in tasks such as:

## How to use

```bash
$ xgofind --port <PORT> --dest <DOMAIN>
```

```bash
$ ./xgofind --help

XGOFIND - http/https interceptor
==========================================
Version: dev | Jabes Eduardo @yhk0

Use:
  proxy [opções]
Options:
  -dest string
        Destination URL where the proxy will redirect requests.
  -help
        Exibe a ajuda do programa
  -log
        Enables logging of intercepted requests
  -port string
        Port for the proxy to listen
```

#### parameters:

```bash
--help
--port
--dest
--log
```

## Instalation:

```bash
$ git clone https://github.com/yhk0/xgofind.git
$ make buid-[YOUR OS]
$ cd bin
```

### Security Testing (Pentesting):

Inspect and modify requests and responses to identify vulnerabilities in web applications.
Simulate attacks like header injections, payload manipulation, and authentication testing.

### Application Development and Debugging:

Monitor traffic between clients and servers to understand API and web application behavior.
Test request modifications without altering client or server code.

### Customizable Flexibility:

Enable the extension of the proxy with new features, such as detailed logging, header injection, or custom traffic manipulation rules.

---------------------------
The xgofind project is designed to be simple, modular, and open-source, allowing users to adapt the tool to their specific needs. It also serves as an educational project for exploring Go’s capabilities in building network-based applications.
