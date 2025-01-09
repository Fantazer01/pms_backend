from fastapi import APIRouter, HTTPException

from app.api.models import Request, Response
from app.api import db_manager

ai = APIRouter()

@ai.post('/', response_model=Response, status_code=201)
async def create_response(request: Request):

    id, request, responce = await db_manager.add_item(request)

    response = {
        'id': id,
        'request': request,
        'response': responce
    }

    return response

@ai.get('/{id}/', response_model=Response)
async def get_item(id: int):
    response = await db_manager.get_item(id)
    if not response:
        raise HTTPException(status_code=404, detail="item not found")
    return response