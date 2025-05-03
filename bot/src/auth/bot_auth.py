import requests


def auth(
            self,
            id: int,
            use: str,
            avatar: str, 
            description: str
    ):
    response = requests.post("http://localhost:5000/user/")
    print(response.json)
    return response.json