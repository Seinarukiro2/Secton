#!/bin/bash

# Переменные для базы данных
DB_NAME="secton_db"
DB_USER="secton_user"
DB_PASSWORD="4464"

# Проверка, установлен ли PostgreSQL
if ! command -v psql > /dev/null; then
    echo "PostgreSQL не установлен. Устанавливаем..."

    # Для Ubuntu/Debian
    if [ -f /etc/debian_version ]; then
        sudo apt update
        sudo apt install -y postgresql postgresql-contrib
    # Для macOS через Homebrew
    elif [ "$(uname)" == "Darwin" ]; then
        brew install postgresql
    else
        echo "Эта система не поддерживается скриптом."
        exit 1
    fi
else
    echo "PostgreSQL уже установлен."
fi

# Запуск PostgreSQL сервиса
echo "Запуск PostgreSQL сервиса..."
if [ -f /etc/debian_version ]; then
    sudo service postgresql start
elif [ "$(uname)" == "Darwin" ]; then
    brew services start postgresql@14
fi

# Подключение к базе данных postgres и создание пользователя и базы данных
echo "Создание пользователя $DB_USER и базы данных $DB_NAME..."

psql -v ON_ERROR_STOP=1 -d postgres <<-EOSQL
    CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';
    CREATE DATABASE $DB_NAME;
    GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;
EOSQL

# Проверка успешности выполнения команд
if [ $? -eq 0 ]; then
    echo "Пользователь и база данных успешно созданы."
else
    echo "Произошла ошибка при создании пользователя или базы данных."
    exit 1
fi
