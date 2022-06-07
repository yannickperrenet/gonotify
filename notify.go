// Example notification through DBUS org.freedesktop.Notifications
// Turned out to be pretty much a copy of:
// https://github.com/godbus/dbus/blob/v5.1.0/_examples/notification.go
package main

import (
    "github.com/godbus/dbus/v5"
)

func main() {
    conn, err := dbus.ConnectSessionBus()
    if err != nil {
		panic(err)
	}
    defer conn.Close()

    // Get the Object identified by path "/org/freedesktop/Notifications"
    // that is defined on peer (could be client or service) "org.freedesktop.Notifications"
    obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")

    // Call method ".Notify"
    // on interface "org.freedesktop.Notifications"
    // which has signature "susssasa{sv}i"
    // and returns "u"
    // Specification: https://specifications.freedesktop.org/notification-spec/notification-spec-latest.html
    call := obj.Call(
        "org.freedesktop.Notifications.Notify",
        // Flags that godbus can pass with the D-Bus message
        0,
        // Optional name of application sending the notification
        "gonotify",
        // ID of notification it replaces, or 0 for new notification
        uint(0),
        // Optional program icon
        "",
        // Summary of the notification
        "Warning",
        // body
        "Something terrible happened",
        // Actions
        make([]string, 0),
        // Hints
        map[string]dbus.Variant{},
        // Expire timeout in milliseconds
        5000,
    )

    if call.Err != nil {
        panic(call.Err)
    }
}
