# isucon9-prep

## memo

### go

- GOROOT `/opt/go`
- PATH `/opt/go/bin`

### kataribe

- binary `/usr/local/bin/kataribe`
- config `/etc/kataribe.toml`

| コマンド | 結果 |
| --- | --- |
| `kataru` | 統計を表示する |
| `katarazu` | logを全削除する |

### netdata

- `/opt/netdata`

**最終計測前に停止すること**

listening on `http://localhost:19999/`

### notify_slack

- binary `/usr/local/bin/notify_slack`
- config `/etc/notify_slack.toml`

Example
```
uname -a | notify_slack
cat /proc/cpuinfo | notify_slack -snippet
```
