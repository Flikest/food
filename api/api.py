import os
from dotenv import load_dotenv
import requests
load_dotenv()
MAP_TOKEN = os.getenv('MAPS_API')

def find_places(location,food):
    parametrs =  {
            'q': food,
            'location':location,
            'radius': 1000,
            'sort': 'distance',
            'key': MAP_TOKEN,
            'point':location,
            }

    responce = requests.get('https://catalog.api.2gis.com/3.0/items', params = parametrs)
    return responce
