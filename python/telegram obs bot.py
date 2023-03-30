# https://stackoverflow.com/questions/45558984/how-to-make-telegram-bot-dynamic-keyboardbutton-in-python-every-button-on-one-ro

import ast, os
import time
import telebot
from obswebsocket import obsws, requests
import dotenv

dotenv.load_dotenv()

bot = telebot.TeleBot(os.getenv('BOT_TOKEN'))
stringList = {"Name": "John", "Language": "Python", "API": "pyTelegramBotAPI"}
crossIcon = u"\u274C"

def makeKeyboard():
	markup = telebot.types.InlineKeyboardMarkup()

	for key, value in stringList.items():
		markup.add(telebot.types.InlineKeyboardButton(text=value,
											  callback_data="['value', '" + value + "', '" + key + "']"),
		telebot.types.InlineKeyboardButton(text=crossIcon,
								   callback_data="['key', '" + key + "']"))

	return markup

@bot.message_handler(commands=['test'])
def handle_command_adminwindow(message):
	bot.send_message(chat_id=message.chat.id,
					 text="Here are the values of stringList",
					 reply_markup=makeKeyboard(),
					 parse_mode='HTML')
	

@bot.callback_query_handler(func=lambda call: True)
def handle_query(call):

	if (call.data.startswith("['value'")):
		print(f"call.data : {call.data} , type : {type(call.data)}")
		print(f"ast.literal_eval(call.data) : {ast.literal_eval(call.data)} , type : {type(ast.literal_eval(call.data))}")
		valueFromCallBack = ast.literal_eval(call.data)[1]
		keyFromCallBack = ast.literal_eval(call.data)[2]
		bot.answer_callback_query(callback_query_id=call.id,
							  show_alert=True,
							  text="You Clicked " + valueFromCallBack + " and key is " + keyFromCallBack)
		client = obsws("localhost", 4455, "123456")
		client.connect()
		print(client.call(requests.GetVersion()).getObsVersion())
		client.call(requests.SetCurrentProgramScene(sceneName="Logo"))
		client.disconnect()


	if (call.data.startswith("['key'")):
		keyFromCallBack = ast.literal_eval(call.data)[1]
		del stringList[keyFromCallBack]
		bot.edit_message_text(chat_id=call.message.chat.id,
							  text="Here are the values of stringList",
							  message_id=call.message.message_id,
							  reply_markup=makeKeyboard(),
							  parse_mode='HTML')
while True:
	try:
		bot.polling(none_stop=True, interval=0, timeout=0)
	except:
		time.sleep(10)