# pro-bot

<h1 align="center">🤖 Pro Bot</h1>
<p align="center">
    <em>A Telegram Bot to Play Around With The Community Telegram Channels</em>
</p>

<p align="center">
    <a href="https://app.netlify.com/sites/pro-bot/deploys">
        <img src="https://api.netlify.com/api/v1/badges/eb4b8a29-181d-4d48-980e-ebce39edafb8/deploy-status">
    </a>
    <a href="https://hits.seeyoufarm.com">
        <img src="https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2FTechProber%2Fpro-bot&count_bg=%235322B2&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false"/>
    </a>
    <a href="https://img.shields.io/tokei/lines/github/TechProber/pro-bot?color=orange">
        <img src="https://img.shields.io/tokei/lines/github/TechProber/pro-bot?color=orange" alt="lines">
    </a>
    <a href="https://hub.docker.com/repository/docker/hikariai/">
        <img src="https://img.shields.io/badge/docker-v20.10.7-blue" alt="Version">
    </a>
    <a href="https://kubernetes.io/">
        <img src="https://img.shields.io/badge/kubernetes-v1.23-navy.svg" alt="Kubernetes"/>
    </a>
    <a href="https://github.com/TechProber/pro-bot">
        <img src="https://img.shields.io/github/last-commit/TechProber/pro-bot" alt="lastcommit"/>
    </a>
    <a href="https://github.com/TechProber/pro-bot/blob/master/LICENSE">
        <img src="https://img.shields.io/github/license/TechProber/pro-bot?color=critical" alt="License"/>
    </a>
</p>

## Introduction

CopyRight 2022 TechProber. All rights reserved.

Maintainer: [ Kevin Yu (@yqlbu) ](https://github.com/yqlbu)

This repo serves to provide the end-users with a way to interact with community telegram channels. Feel free to contribute.

## Pre-commit Hooks

This repo uses the [pre-commit framework](https://github.com/pre-commit/pre-commit-hooks) to set up hooks. To use it, run the following commands:

```bash
# Install with pip
pip install pre-commit
# Install with Homebrew
brew install pre-commit
# Install pre-commit under root directory
pre-commit install
```

## Community

- [Hikariai_AI Telegram Group](https://t.me/hikariai_channel)
- [unRAID Telegram Group](https://t.me/unraid_zh)
- [TechProber Telegram Group](https://t.me/joinchat/7AG3aEQ5I00wY2Q5)
- [TechProber Discord Channel](https://discord.gg/se4RtrB473)

## Contributors

Special thanks go to all [ contributors ](https://github.com/TechProber/pro-bot/graphs/contributors). If you would like to contribute, please see the [instructions](https://github.com/TechProber/pro-bot/blob/master/docs/contribute.md).

[Pro Bot](https://github.com/TechProber/pro-bot) is inspired and introduced in 2022. Currently, it is maintained by [ Kevin Yu (@yqlbu) ](https://github.com/yqlbu)

## Reference

- [Telegram Bot API](https://core.telegram.org/bots/api)
- [Sending a message to a Telegram channel the easy way](https://medium.com/javarevisited/sending-a-message-to-a-telegram-channel-the-easy-way-eb0a0b32968)
- [Can a Telegram bot read messages of channel](https://stackoverflow.com/questions/42672034/can-a-telegram-bot-read-messages-of-channel)
- [https://stackoverflow.com/questions/52765833/why-i-cant-send-sticker-by-its-id](https://stackoverflow.com/questions/52765833/why-i-cant-send-sticker-by-its-id)
- [Telegram, getting file_id for existing sticker](https://stackoverflow.com/questions/34355648/telegram-getting-file-id-for-existing-sticker)
- [How to set up push notifications in your Telegram bot](https://www.freecodecamp.org/news/telegram-push-notifications-58477e71b2c2/)
- [https://lornajane.net/posts/2020/a-first-netlify-function-in-golang](https://lornajane.net/posts/2020/a-first-netlify-function-in-golang)

## Development Notes

<details><summary>Sample API Call</summary>

### Send Text

POST `https://api.telegram.org/bot<token>/sendMessage?chat_id=<id>&text=<text_string>`

```bash
curl -X POST https://api.telegram.org/bot<token>/sendMessage?chat_id=<id>&text=<text_string>
```

### Send Sticker

POST `https://api.telegram.org/bot<token>/sendSticker?chat_id=<id>&sticker=<sticker_id>`

```bash
curl -X POST https://api.telegram.org/bot<token>/sendSticker?chat_id=<id>&sticker=<sticker_id>
```

### Get Update

GET `https://api.telegram.org/bot{bot_token}/getUpdates`

```bash
curl -X GET https://api.telegram.org/bot{bot_token}/getUpdates
```

</details>

## License

[MIT (C) TechProber](https://github.com/TechProber/pro-bot/blob/master/LICENSE)
