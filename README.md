# gonotify

`notify-send` alternative writting in Go.

Should be just a few lines of code so you can send notifications through DBUS communication.

Sending a notification using `busctl`:
```sh
busctl --user call \
  org.freedesktop.Notifications \   # peer
  /org/freedesktop/Notifications \  # object
  org.freedesktop.Notifications \   # interface
  Notify \                          # method
  susssasa\{sv\}i \                 # signature
  "my-app" 0 "" "A summary" "Some body" 0 0 5000
```

[Specification for
notifications](https://specifications.freedesktop.org/notification-spec/notification-spec-latest.html)

## Resources

* http://0pointer.net/blog/the-new-sd-bus-api-of-systemd.html
