# Opera
> 💃 MRPIS implementation for osu!

Opera is a program to send currently playing osu! maps/songs over DBus.  
It implements the MPRIS specification and acts as a "music player" to
utilities and tools that read data about MPRIS players.

As a good example, an osu! rich presence with [Clematis](https://github.com/TorchedSammy/Clematis)
;)

# Install
Opera requires running [Gosumemory](https://github.com/l3lackShark/gosumemory) to get info.

## Compiling
Run `go install github.com/TorchedSammy/Opera`, or if you prefer to manually build:  
```
git clone https://github.com/TorchedSammy/Opera
cd Opera
go get
go build
```

# Usage
Running Opera standalone will work, and anything which gets info from DBus
for MPRIS will pick up on what Opera sends.

# License
MIT

