# isucon9-prep

チーム NaruseJun

## memo

### ドメイン

IP直打ちでの(SSH|HTTP)アクセスではどれが何台目か混乱しがちなので、以下のドメインでアクセスする。
ただし、SSHについては別途ssh_configも用意することを推奨。

- isu1.sysad.net
- isu2.sysad.net
- isu3.sysad.net

これらは、開始直後に手動で設定する。

### Linux User

githubアカウントと同名。githubに登録済みの鍵でSSHログインできる。

### deploy

`/home/kiritan/Makefile` を使う。
`sudo -i -u kiritan`して、`make`でビルド＋アプリ再起動する。

↑これは下記のbot↓が叩いてくれるので、直接叩くのはトラブル時のみ。

### deploy bot

Slackの`#deploy`にいる。

isucon9-appリポジトリのmasterが更新されると、「デプロイしますか？」的な質問を飛ばす → そこからデプロイできる。
`@kiritan deploy [commit_id]` とか `@kiritan deploy origin/branch-name` とかで任意のタイミングでのdeployも可。

### MySQL User

githubアカウントと同名。パスワードはアカウント名と同じ文字列。
localhostからしか接続できない。SSHポートフォワードを使うか、SSHしてから`mysql`する。

### phpMyAdmin

↓競技用インスタンス3台に接続しているphpMyAdmin

http://pma.hq.sysad.net/

slowlogは `mysql.slow_log` テーブルに格納されている。

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

オススメは `kataru | slack_notify -snippet` でSlackに飛ばしてから見る。

### netdata

- `/opt/netdata`

各ホストの 19999/tcp でLISTENしている。
以下のURLでもアクセスできる。

- http://isu1-netdata.hq.sysad.net/
- http://isu2-netdata.hq.sysad.net/
- http://isu3-netdata.hq.sysad.net/

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

### HTTPフォワード

以下のURLは各ホストの `8080/tcp` につながるようになっている。
L7でフォワードしているので、HTTP限定。

- http://isu1.hq.sysad.net/
- http://isu2.hq.sysad.net/
- http://isu3.hq.sysad.net/

`python -m SimpleHTTPServer 8080` とか `go tool pprof -http 0:8080` で使うことを想定。
