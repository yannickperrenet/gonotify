# gonotify

An example of sending a desktop notification through D-Bus, written in Go.

```sh
go build
./gonotify
```

## `notify-send`

Another program to send desktop notifications is `notify-send`. From its manpage ([source
code](https://github.com/GNOME/libnotify/blob/master/tools/notify-send.c)):

>   With notify-send you can send desktop notifications to the user via a notification daemon from
>   the command line.

What I think this means is that `libnotify` ([source code](https://github.com/GNOME/libnotify))
creates a service on the (user session) bus, aka the "notification daemon". With `notify-send` you
can then send messages to that daemon to handle.

## `busctl`

Let's use `busctl` to send a message to the notification daemon instead of using the `notify-send`
client:

```sh
busctl --user call \
  org.freedesktop.Notifications \   # peer
  /org/freedesktop/Notifications \  # object
  org.freedesktop.Notifications \   # interface
  Notify \                          # method
  susssasa\{sv\}i \                 # signature
  "my-app" 0 "" "A summary" "Some body" 0 0 5000
```

In the example (D-Bus specific terms incoming) you can see how we first have to specify the `peer`
(this can be a client or service) that exposes the notifications `object`. On this object we can then
call a specific `method` by specifying its `interface` and `signature`. The meaning of the signature
is as follows:
* `s` a string
* `u` unsigned integer
* `as` array of strings. Which is specified through e.g. `2 "string" "string"`, so `0` meaning no
  array
* `i` signed integer

## The `dunst` notification daemon

So then how can it be possible to use `dunst` as you notification daemon if there already is a
service listening under the name `org.freedesktop.Notifications`? To add to the question, note that
`notify-send` will still work with `dunst` without having to change anything.

A first step to answering this question is to run `dunst` when another `dunst` daemon is already
running:
```
❯ dunst
WARNING: BadAccess (attempt to access private resource denied)
WARNING: BadAccess (attempt to access private resource denied)
WARNING: Unable to grab key 'mod1+shift+n'.
WARNING: No icon found in path: 'dialog-information'
CRITICAL: Cannot acquire 'org.freedesktop.Notifications': Name is acquired by 'dunst' with PID '2254'.
```
From the warnings above we can infer that `dunst` will put itself under the
`org.freedesktop.Notifications` on the bus and thus will be the daemon that handles the
notifications (given that it is the only possible daemon that can handle notifications).

Indeed in `/usr/share/dbus-1/services/org.knopwob.dunst.service` (where DBUS looks for all its
defined services) there is a service for dunst under the `org.freedesktop.Notifications` name:
```
[D-BUS Service]
Name=org.freedesktop.Notifications
Exec=/usr/bin/dunst
SystemdService=dunst.service
```

## Resources

* [Introduction to D-Bus](https://www.freedesktop.org/wiki/IntroductionToDBus/)
* [Pid Eins blog on `busctl` and D-Bus](http://0pointer.net/blog/the-new-sd-bus-api-of-systemd.html)
* [D-Bus specification](https://dbus.freedesktop.org/doc/dbus-specification.html)
* [Notifications
  specification](https://specifications.freedesktop.org/notification-spec/notification-spec-latest.html)
* [Stackoverflow - how a systemd service gets on the
  bus](https://stackoverflow.com/questions/31702465/how-to-define-a-d-bus-activated-systemd-service)
* [Python example of creating systemd service that sits on the
  bus](https://github.com/sezanzeb/systemd-pydbus-example)
