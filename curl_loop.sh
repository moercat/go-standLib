#!/bin/bash

# 持续执行 curl 请求的脚本
# 目标URL: http://127.0.0.1:19930/ip
# 包含特定的头部信息

# 设置变量
TARGET_URL="http://127.0.0.1:19930/ip"
SERVER_IP="222.184.96.155"
BIZ_TYPE="fusion-biz"
BIZ_PARAMS="fusion-biz|key#1"
FORWARDED_FOR="222.184.96.155"
URL_HASHKEY="/static/xxx111.txt"
SOURCE_URL="http://lf9-webcast-cdn-tos-ncdn.bytegecko.com/static/1M.txt"

# 统计变量
REQUEST_COUNT=0
SUCCESS_COUNT=0
ERROR_COUNT=0

# 默认参数
SLEEP_TIME=1  # 请求间隔时间(秒)
MAX_REQUESTS=0  # 最大请求数，0 表示无限循环

# 显示使用帮助
show_help() {
    echo "Usage: $0 [OPTIONS]"
    echo "Options:"
    echo "  -i INTERVAL   请求间隔时间(秒，支持小数)，默认为 1"
    echo "  -n COUNT      最大请求数，0 表示无限循环，默认为 0"
    echo "  -q            静默模式，只输出错误信息"
    echo "  -h            显示此帮助信息"
    echo ""
    echo "Example:"
    echo "  $0 -i 2 -n 100       # 每2秒发送一次请求，共发送100次"
    echo "  $0 -i 0.06           # 每0.06秒发送一次请求，大约每分钟1000次"
    echo "  $0                   # 每1秒发送一次请求，无限循环"
}

# 解析命令行参数
QUIET_MODE=false
while getopts "i:n:qh" opt; do
    case $opt in
        i)
            SLEEP_TIME=$OPTARG
            ;;
        n)
            MAX_REQUESTS=$OPTARG
            ;;
        q)
            QUIET_MODE=true
            ;;
        h)
            show_help
            exit 0
            ;;
        \?)
            echo "Invalid option: -$OPTARG" >&2
            show_help
            exit 1
            ;;
    esac
done

# 检查是否支持高精度睡眠
use_high_precision_sleep=false
if command -v perl >/dev/null 2>&1; then
    use_high_precision_sleep=true
elif [[ "$OSTYPE" == "linux-gnu"* ]] && sleep 0.1 2>/dev/null; then
    use_high_precision_sleep=true
fi

# 发送信号处理程序 - 捕获 Ctrl+C 信号
trap 'echo -e "\n脚本已停止。请求统计 - 总数: $REQUEST_COUNT, 成功: $SUCCESS_COUNT, 错误: $ERROR_COUNT"; exit 0' INT TERM

# 主循环
echo "开始发送 curl 请求到 $TARGET_URL"
echo "请求间隔: ${SLEEP_TIME}秒"
if [ $MAX_REQUESTS -gt 0 ]; then
    echo "最大请求数: $MAX_REQUESTS"
else
    echo "模式: 无限循环 (按 Ctrl+C 停止)"
fi
echo "----------------------------------------"

while true; do
    # 检查是否达到最大请求数
    if [ $MAX_REQUESTS -gt 0 ] && [ $REQUEST_COUNT -ge $MAX_REQUESTS ]; then
        echo "已达到最大请求数: $MAX_REQUESTS"
        break
    fi

    # 增加请求计数器
    REQUEST_COUNT=$((REQUEST_COUNT+1))

    # 执行 curl 命令并捕获输出和状态码
    response=$(curl -H "X-Server-Ip:$SERVER_IP" \
                -H "X-Biz-Type:$BIZ_TYPE" \
                -H "X-Biz-Params:$BIZ_PARAMS" \
                -H "X-Forwarded-For:$FORWARDED_FOR" \
                -H "X-Url-Hashkey: $URL_HASHKEY" \
                -H "X-Url:$SOURCE_URL" \
                "$TARGET_URL" \
                -vvv 2>&1)
    curl_status=$?

    # 统计请求结果
    if [ $curl_status -eq 0 ]; then
        SUCCESS_COUNT=$((SUCCESS_COUNT+1))
        if [ "$QUIET_MODE" = false ]; then
            echo "[$REQUEST_COUNT] 请求成功 $(date '+%Y-%m-%d %H:%M:%S'):"
            echo "$response"
            echo "----------------------------------------"
        fi
    else
        ERROR_COUNT=$((ERROR_COUNT+1))
        # 总是输出错误信息，即使在静默模式下
        echo "[$REQUEST_COUNT] 请求失败 $(date '+%Y-%m-%d %H:%M:%S') (状态码: $curl_status):"
        echo "$response"
        echo "----------------------------------------"
    fi

    # 检查是否需要停止
    if [ $MAX_REQUESTS -gt 0 ] && [ $REQUEST_COUNT -ge $MAX_REQUESTS ]; then
        break
    fi

    # 等待指定的时间间隔，支持小数秒
    if (( $(echo "$SLEEP_TIME > 0" | bc -l 2>/dev/null || echo 0) )); then
        if [ "$use_high_precision_sleep" = true ]; then
            # 如果是 Linux 系统，使用内置的 sleep 命令支持小数
            if [[ "$OSTYPE" == "linux-gnu"* ]]; then
                sleep "$SLEEP_TIME"s
            else
                # 使用 Perl 实现高精度睡眠 (适用于 macOS)
                perl -e "select(undef, undef, undef, $SLEEP_TIME);"
            fi
        else
            # 如果不支持高精度睡眠，则只使用整数秒
            sleep "${SLEEP_TIME%.*}"
        fi
    fi
done

# 输出最终统计信息
echo "脚本执行完成。"
echo "请求统计 - 总数: $REQUEST_COUNT, 成功: $SUCCESS_COUNT, 错误: $ERROR_COUNT"