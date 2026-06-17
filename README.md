# SdalGIABot (недоступен, находится в разработке)

Telegram-бот для получения результатов экзаменов с нескольких источников ([ИКСОРА](https://sdr.ixora.ru/login), [НИМРО](http://nscm.ru/), [CheckEGE](https://checkege.rustest.ru/)) с автоматическим обновлением.

~Доступен для использования: [SdalGIABot](https://t.me/SdalGIABot)~

---

### TODO
- [ ] Внедрение логирования и ротации логов
- [ ] Приём сообщений из Telegram
- [ ] Приём сообщений из VK
- [ ] Получение результатов с ИКСОРА
- [ ] Получение результатов с НИМРО
- [ ] Получение результатов с CheckEGE
- [ ] Выборочное обновление результатов (репрезентативная выборка по предметам и источникам)
- [ ] Покрытие кода тестами

---

### Setup & Run

- Клонирование репозитория:
```bash
git clone git@github.com:mrzlkvvv/SdalGIABot.git && cd SdalGIABot
```

- Копирование и редактирование файлов конфигурации:
```bash
# Не забудьте отредактировать скопированные файлы!
cp ./config/bot.env.example      ./config/bot.env
cp ./config/pgadmin.env.example  ./config/pgadmin.env
cp ./config/postgres.env.example ./config/postgres.env
```

- Сборка и запуск:
```bash
make up-prod
```
