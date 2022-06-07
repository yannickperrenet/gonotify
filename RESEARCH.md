* Understanding DBUS: https://www.freedesktop.org/wiki/Software/dbus/

DBUS
* How about creating a small script that opens a new tab in my browser
* https://web.archive.org/web/20200522193008/http://0pointer.net/blog/the-new-sd-bus-api-of-systemd.html
    * Set up Wireshark to monitor D-Bus traffic
* I should probably spend some time to better understand systemd given that services sit on the
  system's bus
    * As an example, what is the system journal?
    * Probably a good source: https://wiki.archlinux.org/title/Systemd

## Sending notification through DBUS
The [DBUS protocol for
notifications](https://specifications.freedesktop.org/notification-spec/notification-spec-latest.html)
and [source code](https://github.com/GNOME/libnotify).

From the `notify-send` manpage ([source
code](https://github.com/GNOME/libnotify/blob/master/tools/notify-send.c)):

>   With notify-send you can send desktop notifications to the user via a notification daemon from
>   the command line.

What I think this means is that `libnotify` creates a service on the (user session) bus, aka the
"notification daemon". With `notify-send` you can then send messages to that daemon.

Example of using `busctl` to send messages to the daemon directly instead of through `notify-send`.
```
busctl --user call org.freedesktop.Notifications /org/freedesktop/Notifications org.dunstproject.cmd0 NotificationShow

busctl --user call \
  org.freedesktop.Notifications \
  /org/freedesktop/Notifications \
  org.freedesktop.Notifications \
  Notify \
  susssasa\{sv\}i \
  "my-app" 0 "" "A summary" "Some body" 0 0 5000
```
Meaning of `susssasa\{sv\}i` (which refers to the function signature):
* `s` a string
* `u` unsigned integer
* `as` array of strings. Which is specified through e.g. `2 "string" "string"`, so `0` meaning no
  array
* `i` signed integer
* [complete specification](https://dbus.freedesktop.org/doc/dbus-specification.html)

So then what does `dunst` do? It can be explained when trying to run `dunst` when another `dunst`
daemon is already running:
```
❯ dunst
WARNING: BadAccess (attempt to access private resource denied)
WARNING: BadAccess (attempt to access private resource denied)
WARNING: Unable to grab key 'mod1+shift+n'.
WARNING: No icon found in path: 'dialog-information'
CRITICAL: Cannot acquire 'org.freedesktop.Notifications': Name is acquired by 'dunst' with PID '2254'.
```
It seems like `dunst` will take over the `org.freedesktop.Notifications` name in the bus and thus
will become the daemon that handles notifications.

Indeed in `/usr/share/dbus-1/services/org.knopwob.dunst.service` (where DBUS looks for all its
defined services) there is a service for dunst:
```
[D-BUS Service]
Name=org.freedesktop.Notifications
Exec=/usr/bin/dunst
SystemdService=dunst.service
```

Great example of creating a new systemd service that sits on the bus:
https://github.com/sezanzeb/systemd-pydbus-example

And how systemd services get on the bus:
https://stackoverflow.com/questions/31702465/how-to-define-a-d-bus-activated-systemd-service
