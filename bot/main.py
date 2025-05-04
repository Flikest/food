import os
from dotenv import load_dotenv
import telebot

load_dotenv()
user_data = {} #dict for users {'chat_id': '_user_status'}
bot = telebot.TeleBot(os.getenv('BOT_TOKEN'))

@bot.message_handler(commands=['start'])
def welcome(message):
    chat_id = message.chat.id
    #TODO: Check if user in database
    #True
    bot.send_message(chat_id, 'Здравствуйте, отправьте название блюда, которое вы хотели бы попробовать.')
    #user_data[chat_id] = '_is_writing_dish'
    #False
    bot.send_message(chat_id, 'Здравствуйте, это бот для поиска где и с кем вы можете поесть ту еду, которую вы выбрали.')
    bot.send_message(chat_id, 'Давайте создадим вам анкету.')
    bot.send_message(chat_id, 'Введите свой пол:')
    user_date[char_id] = '_is_writing_sex'

@bot.message_handler(func = lambda message: True):
def handle_message(message):
    chat_id = message.chat.id
    if user_data[chat_id] == '_is_writing_sex':
        #TODO: add the sex in sql tablet

        bot.send_message(chat_id, 'Введите свое имя:')
        user_data[chat_id] = '_is_writing_name'
    if user_data[chat_id] == '_is_writing_name':
        #TODO: add the name in sql tablet

        bot.send_message(chat_id, 'Введите свой возраст:')
        user_data[chat_id] = '_is_writing_age'
    if user_data[chat_id] == '_is_writing_age':
        #TODO: add the age in sql tablet

        bot.send_message(chat_id, 'Введите описание для своей анкеты:')
        user_data[chat_id] = '_is_writing_description'
    if user_data[chat_id] == '_is_writing_description':
        #TODO: add the description in sql tablet

        bot.send_message(chat_id, 'Поздравляю, вы были зарегестрированы, теперь напишите команду /start еще раз.')
    if user_data[chat_id] = '_is_writing_dish':
        #TODO: Make api request and processing an api response

        #Show restorance (mb in button to make them clickable) and other users
        user_data[chat_id] = '_is_selecting_restaurant'





bot.infinity_polling()

