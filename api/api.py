import os
from dotenv import load_dotenv
import requests
load_dotenv()
MAP_TOKEN = os.getenv('MAPS_API')

def find_places(location,food):
    parametrs =  {
            'q': food,
            'location':location,
            'radius': 10000,
            'sort': 'distance',
            'key': MAP_TOKEN,
            'type': 'branch'
            }

    response = requests.get('https://catalog.api.2gis.com/3.0/items', params = parametrs)
    response = response.json()
    return response

def create_user(id, username, avatar_id, blank, rating):
    params = {
            'id' : id,
            'use' : username,
            'avatar' : avatar_id,
            'description' : blank,
            'rating' : rating,
        }
    response = requests.post('http://localhost:5000/user/', params = params)

def get_user_by_id(id):
    response = requests.get(f'http://localhost:5000/user/{id}')
    response = response.json()
    return response

def update_user(id, username, avatar_id, blank, rating):
    params = {
            'id' : id,
            'use' : username,
            'avatar' : avatar_id,
            'description' : blank,
            'rating' : rating,
        }
    response = requests.post('http://localhost:5000/user/', params = params)

def delete_user(id):
    response = requests.delete(f'http://localhost:5000/user/{id}')

def create_group(rest_id):
    params = {'id' : rest_id}
    response = requests.post('http://localhost:5000/room/', params = params)

def join_group(id, rest_id):
    params = {
            'id' : rest_id,
            'user_id' : id,
            }
    response = requests.post('http://localhost:5000/join/', params = params)

def get_all_users_from_group(rest_id):
    response = requests.get(f'http://localhost:5000/room/{rest_id}')

def get_all_group():
    response = requests.get('http://localhost:5000/room/')

def leave(rest_id, id):
    params = {
            'id' : rest_id,
            'user_id' : id,
            }
    response = request.post('http://localhost:5000/room/', params = params)

def delete_room(rest_id):
    response = request.post(f'http://localhost:5000/room/{rest_id}')

#def update_raiting(id, operation):
#    params = {
##            'user_id' : id
#            'operation' : operation
#            }
#    response = request.patch('http://localhost:5000/raiting')

