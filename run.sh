pkill Xephyr
go build
Xephyr :1 -br -ac -once -br -reset -terminate -screen 1600x900 1>/dev/null 2>/dev/null 0>/dev/null &
sleep 0.5
DISPLAY=:1 ./salientwm

