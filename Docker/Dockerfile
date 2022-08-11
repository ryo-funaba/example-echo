FROM golang:1.19.0-alpine3.16

#パッケージをインストール
RUN apk update && \
    apk --no-cache add \
    curl=7.83.1-r2 \
    git=2.36.2-r0 \
    make=4.3-r0 \
    npm=8.10.0-r0 \
    tzdata=2022a-r0 \
    vim=8.2.5000-r0 \
    zsh=5.8.1-r4

# Serverless Frameworkをインストール
RUN npm install -g serverless@3.21.0

# hadolint ignore=DL3059
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

ENV USER developer
ENV HOME /home/$USER

#非rootユーザーを使用
RUN addgroup -S $USER && \
    adduser -S -G $USER $USER && \
    chown -R $USER:$USER $HOME

USER $USER

WORKDIR $HOME

COPY .zshrc .

# プラグインをインストール
RUN git clone https://github.com/zsh-users/zsh-autosuggestions ~/.zsh/zsh-autosuggestions && \
    git clone https://github.com/zsh-users/zsh-completions.git ~/.zsh/zsh-completions && \
    git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ~/.zsh/zsh-syntax-highlighting && \
    curl -o ~/.zsh/git-prompt.sh https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh && \
    curl -o _git https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.zsh && \
    curl -o ~/.zsh/git-completion.bash https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.bash

CMD ["/bin/zsh"]
