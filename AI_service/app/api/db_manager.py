from app.api.models import Request, Response
from app.api.db import requests, database
from app.api.AI_func import AI_func


async def add_item(request: Request):
    cur_request = request.request
    cur_response = AI_func(cur_request)
    query = requests.insert().values(request=cur_request, response=cur_response)
    
    return await database.execute(query=query), cur_request, cur_response

async def get_item(id):
    query = requests.select().where(requests.c.id==id)
    return await database.fetch_one(query=query)