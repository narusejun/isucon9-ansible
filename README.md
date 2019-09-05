# isucon9-prep

チーム NaruseJun

## reminder

以下は最終計測前に必ず停止すること。
上から重要。

- hq.sysad.netのPortForwarding系
	- 🚨🚨🚨🚨忘れると不正行為になる可能性アリ🚨🚨🚨🚨🚨
	- SSHでつないでいるため
- Deploy Bot
	- 🚨🚨🚨🚨忘れると不正行為になる可能性アリ🚨🚨🚨🚨🚨
	- SSHにつなぎにいくため
- netdata
	- `systemctl disable netdata && systemctl stop netdata`
- slow_query_log
	- /etc/mysql/mysql.conf.d/zz-isucon.cnf `slow_query_log = OFF`
- nginx log
	- /etc/nginx/nginx.conf `access_log off;`
- pprof
	- コードをいじってはがす

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

### アプリ

https://github.com/kaz/isucon9-app

リポジトリのトップに実行ファイル`app`を生成するMakefileを置くこと。後述のデプロイスクリプトはこれをキックする。
systemdで起動される。サービス名は`app.service`である。必ず`9000/tcp`でLISTENすること。

### deploy

`/home/kiritan/deploy.sh` を使う。
`sudo -i -u kiritan`して、`./deploy.sh`でビルド＋アプリ再起動する。

※ これは下記のbotが叩いてくれるので、直接叩くのはトラブル時のみ。

### deploy bot

Slackの`#deploy`にいる。

isucon9-appリポジトリのmasterが更新されると、「デプロイしますか？」的な質問を飛ばす → そこからデプロイできる。
`@kiritan deploy [commit_id]` とか `@kiritan deploy origin/branch-name` とかで任意のタイミングでのdeployも可。

`@kiritan target 1,3` とするとデプロイ先を変更することができる。（この場合isu1とisu3にデプロイする。）

### MySQL User

githubアカウントと同名。パスワードはアカウント名と同じ文字列。
localhostからしか接続できない。SSHポートフォワードを使うか、SSHしてから`mysql`する。

### phpMyAdmin

↓競技用インスタンス3台に接続しているphpMyAdmin

http://pma.hq.sysad.net/

slowlogは `mysql.slow_log` テーブルに格納されている。

### myprofiler

localのDBにつなぎに行くaliasが設定されているので、DBパスワードなどは気にしなくてOK。
その他のパラメータ → https://github.com/KLab/myprofiler

オススメ `myprofiler | notify_slack`

### systemctl/jounalctl

一般ユーザでも勝手にsudoになるaliasが存在。
`sc` → systemctl、`jc` → jounalctlのショートカットも。

### go

`1.13` がインストールされている。

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
