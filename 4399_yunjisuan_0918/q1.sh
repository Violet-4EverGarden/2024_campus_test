#!/bin/bash

# 要监控的ip
ip_add=("1.1.1.1" "2.2.2.2")
# 发送告警信息给指定用户
notify_user="user01"

# 初始化告警时间
last_alert_time=0

while true; do
    for ip in "${ip_add[@]}"; do
        # 使用ping命令测试对应ip主机存活状态
        if ping -c 1 -W "$ip" > /dev/null 2 >&1; then
            echo "[$(date +'%Y-%m-%d %H:%M:%S')] $ip is alive."
        else
            # 检查是否需要发送告警信息
            cur_time=$(date +%s)
            time_diff=$((cur_time - last_alert_time))
            # 如果距离上次告警时间超过1h则发送新的告警信息
            if ((time_diff >= 3600)); then
                msg="Host $ip is down at $(date +'%Y-%m-%d %H:%M:%S')."
                notify "$notify_user" "$msg"
                last_alert_time=$cur_time
            fi
        fi
    done

    sleep 30
done