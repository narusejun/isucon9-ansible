# isucon9-prep

## memo

### Linux User

githubアカウントと同名。githubに登録済みの鍵でログインできる。

### MySQL User

githubアカウントと同名。パスワードはアカウント名と同じ文字列。
localhostからしか接続できない。SSHポートフォワードを使うか、SSHしてから`mysql`する。

### deploy

`/home/kiritan/Makefile` を使う。
`sudo -i -u kiritan`して、`make`でビルド＋アプリ再起動する。

↑これはslack botが叩いてくれる(_TODO_)ので、直接叩くのはトラブル時のみ。

### systemctl/jounalctl

一般ユーザでも勝手にsudoになるaliasが存在。
`sc` → systemctl、`jc` → jounalctlのショートカットも。

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

listening on `http://localhost:19999/`

_最終計測前に必ず停止する_

### dstat

netdataがトラブったとき用に入れてある

### notify_slack

- binary `/usr/local/bin/notify_slack`
- config `/etc/notify_slack.toml`

Example
```
uname -a | notify_slack
cat /proc/cpuinfo | notify_slack -snippet
```
