alias systemctl="sudo systemctl"
alias jounalctl="sudo journalctl"

alias sc="sudo systemctl"
alias sce="sudo systemctl status"
alias scs="sudo systemctl start"
alias sck="sudo systemctl stop"
alias scr="sudo systemctl restart"
alias scl="sudo systemctl reload"
alias scon="sudo systemctl enable"
alias scof="sudo systemctl disable"

alias jc="sudo journalctl"
alias jcf="sudo journalctl -f -u"
alias jcn="sudo journalctl -n 100 -u"
alias jcnn="sudo journalctl -n 1000 -u"

alias ppf="sudo /opt/go/bin/go tool pprof -http :8080"
alias pph="sudo /opt/go/bin/go tool pprof -http :8080 http://localhost/debug/pprof/profile"
