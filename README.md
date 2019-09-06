# isucon9-prep

チーム NaruseJun

## 開発機にインストールしておくべきリスト

- go (1.13)
- ansible (2.7.12)
	- 多分2.8でも大丈夫だと思う！

## 最初にやることリスト

- レギュレーションと当日マニュアルを読み込む
	- 得点計算周りをよく理解する。例年と大きく異なる場合は、施策のヒントになっているかも。
	- 外部連携APIの仕様も重要。何ができるかをよく把握しておく。
	- インスタンスのサイズ（メモリ、CPU）等も確認して共有。
- DNSの設定
	- _環境に関するメモ → ドメイン_ を参照
- `~/ssh/config` の用意（推奨）
	- ```
	  Host *
	  	User (githubのユーザ名)
	  	TCPKeepAlive yes
	  	ForwardAgent yes
	  	ServerAliveInterval 60
	  Host isu1
	  	HostName isu1.sysad.net
	  Host isu2
	  	HostName isu2.sysad.net
	  Host isu3
	  	HostName isu3.sysad.net
	  Host hq
	  	HostName hq.sysad.net
	  ```
- Pythonをリモートに入れる（一応）
	- Pythonが入ってないとansibleが動かない。python実装が存在するので、入っているはずだけど。
- ログインユーザを作る
	- `ansible-playbook playbooks/all.yml -t common.users`
	- その後のAnsibleも`-t`付きで実行することを推奨。環境がぶっ壊れたときのリカバリを早めるため。
		- 最初は`-l`もつけるべきかもしれない。
- ソースコードと静的ファイル類をlocalへ持っていき、appリポジトリに上げる
	- `tar zcvf ~/code.tar.gz /home/isucon/(ほげ)`
	- `post_slack code.tar.gz` (推奨)
		- fallback: `rsync isu1:code.tar.gz .` or `scp isu1:code.tar.gz .`
	- ソースコードと、静的ファイル類。例年はisuconユーザのホームディレクトリにおいてあった。
- DBをダンプしてlocalに持ってきておく
	- `mysqldump -u (だれか) -p (DB名) > dump.sql` && `gzip dump.sql`
	- 最初から初期データのダンプがおいてある場合もある。要確認。

## 最終計測前チェックリスト

上から順に重要。

- hq.sysad.netのPortForwarding系を止める
	- 🚨🚨🚨忘れると不正行為になる可能性アリ🚨🚨🚨
	- SSHでつないでいるため
- Deploy Botを止める
	- 🚨🚨🚨忘れると不正行為になる可能性アリ🚨🚨🚨
	- SSHにつなぎにいくため
- 再起動後、アプリが完動しているかを調べる
	- 再起動時に `/etc/hosts` が書き換わる場合があるので、要チェック
- netdataを止める
	- `systemctl disable netdata && systemctl stop netdata`
- slow_query_logを止める
	- /etc/mysql/mysql.conf.d/zz-isucon.cnf `slow_query_log = OFF`
- nginx logを止める
	- /etc/nginx/nginx.conf `access_log off;`
- pprofを止める
	- コードをいじってはがす

## 環境に関するメモ

### ドメイン

IP直打ちでの(SSH|HTTP)アクセスではどれが何台目か混乱しがちなので、以下のドメインでアクセスする。
ただし、SSHについては別途ssh_configも用意することを推奨。

- isu1.sysad.net
- isu2.sysad.net
- isu3.sysad.net

これらは、開始直後に手動で設定する。
各インスタンスの `/etc/hosts` は、内部接続を利用するように設定されているので、このドメインを使って内部の通信も行って良い。

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

DBのオペレーションをするやつ。一応負荷状況も見れる。
slowlogはテーブルに書くようになっているので、phpMyAdminで`mysql.slow_log`を確認すれば良い。

↓競技用インスタンス3台に接続しているphpMyAdmin

http://pma.hq.sysad.net/

### myprofiler

DBに飛んできてるクエリを集計するやつ。

localのDBにつなぎに行くaliasが設定されているので、DBパスワードなどは気にしなくてOK。
その他のパラメータ → https://github.com/KLab/myprofiler

オススメ `myprofiler | notify_slack`

### systemctl/jounalctl

一般ユーザでも勝手にsudoになるaliasが入れてあるので、一般ユーザでもふつうに叩ける。
ほか、以下のようなショートカットもあり。

| ショートカット | 展開後 | 備考 |
| --- | --- | --- |
| `sc` | `systemctl` | |
| `sce` | `systemctl status` | Examineのイメージ |
| `scs` | `systemctl start` | Start |
| `sck` | `systemctl stop` | Killのイメージ |
| `scr` | `systemctl restart` | Restart |
| `scl` | `systemctl reload` | reLoadのイメージ |
| `jc` | `jounalctl` | |
| `jcf` | `jounalctl -f -u` | `jcf nginx` とかいう感じで使う |
| `jcn` | `jounalctl -n 100 -u` | |
| `jcnn` | `jounalctl -n 1000 -u` | nが長いイメージ |

### go

`1.13` がインストールされている。

- GOROOT `/opt/go`
- PATH `/opt/go/bin`

### kataribe

nginxのログファイルを集計して表示するやつ。

- config `/etc/kataribe.toml`

| コマンド | 結果 |
| --- | --- |
| `kataru` | 統計を表示する |
| `katarazu` | logを全削除する |

オススメは `kataru | slack_notify -snippet` でSlackに飛ばしてから見る。

### netdata

ホストの負荷状況を確認するやつ（GUI）

各ホストの 19999/tcp でLISTENしている。
以下のURLでもアクセスできる。（推奨）

- http://isu1-netdata.hq.sysad.net/
- http://isu2-netdata.hq.sysad.net/
- http://isu3-netdata.hq.sysad.net/

### dstat

ホストの負荷状況を確認するやつ（CUI）
netdataがトラブったとき用に入れてある。

### notify_slack

テキストをslackに投げるやつ。
`#stdout`に上がる。

- config `/etc/notify_slack.toml`

Example
```
uname -a | notify_slack
cat /proc/cpuinfo | notify_slack -snippet
```

### upload_slack

ファイルをslackにアップロードするやつ。
`#stdout`に上がる。

Example
```
upload_slack ./some_file
```

### HTTPフォワード

以下のURLは各ホストの `8080/tcp` につながるようになっている。
L7でフォワードしているので、HTTP限定。

- http://isu1.hq.sysad.net/
- http://isu2.hq.sysad.net/
- http://isu3.hq.sysad.net/

`python -m SimpleHTTPServer 8080` とか `go tool pprof -http 0:8080` で使うことを想定。
