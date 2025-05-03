import telebot
import os
import requests
from dotenv import load_dotenv


token = os.getenv("BOT_TOKEN")

bot=telebot.TeleBot(token)

@bot.message_handler(commands=["start"])
def start_message():
    bot.send_message(message.chat_id, "привет! 👋" \
    "это бот для поиска еды в кафе и ресторанах твоего города")
    requests("http://localhost:5000/user/")