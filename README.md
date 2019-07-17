any-os-syslog
=========

This repository provides a very simple `gsyslog` package. The point of this
package is to allow safe importing of syslog without introducing cross-compilation
issues. The stdlib `log/syslog` cannot be imported on Windows systems, and without
conditional compilation this adds complications.

Instead, `gsyslog` provides a very simple wrapper around `log/syslog`. By default, it just prints out log messages on a non Linux or OSX system. You can change the destination of log messages (a file, a database, other logging system, etc) by setting custom writer, that implements `io.Writer`. You just need to use `SetOtherWriter(io.Writer)` method. (It does nothing on a Linux or OSX system).
