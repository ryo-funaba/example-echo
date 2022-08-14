# 新たにインストールしたコマンドを即認識させる
zstyle ":completion:*:commands" rehash 1

# 補完で小文字でも大文字にマッチさせる
zstyle ':completion:*' matcher-list 'm:{a-z}={A-Z}'

# git-completionの設定
fpath=(~/.zsh $fpath)
zstyle ':completion:*:*:git:*' script ~/.zsh/git-completion.bash
autoload -Uz compinit && compinit

# git-promptの設定
source ~/.zsh/git-prompt.sh
GIT_PS1_SHOWDIRTYSTATE=false
GIT_PS1_SHOWSTASHSTATE=false
GIT_PS1_SHOWUNTRACKEDFILES=false
GIT_PS1_SHOWUPSTREAM=false

# zsh-autosuggestionsの設定
source ~/.zsh/zsh-autosuggestions/zsh-autosuggestions.zsh
ZSH_AUTOSUGGEST_HIGHLIGHT_STYLE="fg=blue,bg=black,bold,underline"

# zsh-completionsの設定
fpath=(~/.zsh/zsh-completions/src $fpath)

# zsh-syntax-highlightingの設定
source ~/.zsh/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh

# 色の設定
autoload -Uz colors && colors

# 補完機能の設定
autoload -Uz compinit && compinit

# cdコマンドを省略して、ディレクトリ名のみの入力で移動
setopt auto_cd

# コマンドミスを修正
setopt correct

# 同じコマンドを履歴に残さない
setopt hist_ignore_all_dups

# 直前と同じコマンドの場合は履歴に追加しない
setopt hist_ignore_dups

# スペースから始まるコマンド行は履歴に残さない
setopt hist_ignore_space

# ヒストリに保存するときに余分なスペースを削除する
setopt hist_reduce_blanks

# ビープ音の停止
setopt no_beep

# ビープ音の停止(補完時)
setopt nolistbeep

# 他のターミナルと履歴を共有
setopt share_history

# 履歴の保存に関する設定
HISTFILE=~/.zsh_history
HISTSIZE=10000
SAVEHIST=10000

# コマンド実行結果のあとに空行を挿入
add_newline() {
  if [[ -z $PS1_NEWLINE_LOGIN ]]; then
    PS1_NEWLINE_LOGIN=true
  else
    printf '\n'
  fi
}

# プロンプトの設定
set_prompt() {
  PROMPT="%F{cyan}(Docker)%f%F{green}:%~%f %F{red}$(__git_ps1)%f"$'\n'"$ "
}

precmd() {
  add_newline
  set_prompt
}

export LANG=ja_JP.UTF-8

# エイリアス
alias ..='cd ..'
alias ...='cd ../..'
alias ....='cd ../../..'
alias ll="ls -l"
alias lla="ls -la"
alias szh='source ~/.zshrc'
alias vzh='vim ~/.zshrc'
