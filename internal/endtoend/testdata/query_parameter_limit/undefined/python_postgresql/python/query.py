# Code generated by sqlc. DO NOT EDIT.
# versions:
#   sqlc v1.15.0
# source: query.sql
import dataclasses

import sqlalchemy
import sqlalchemy.ext.asyncio

from querytest import models


DELETE_BAR_BY_ID = """-- name: delete_bar_by_id \\:execrows
DELETE FROM bar WHERE id = :p1
"""


DELETE_BAR_BY_ID_AND_NAME = """-- name: delete_bar_by_id_and_name \\:execrows
DELETE FROM bar WHERE id = :p1 AND name = :p2
"""


@dataclasses.dataclass()
class DeleteBarByIDAndNameParams:
    id: int
    name: str


class Querier:
    def __init__(self, conn: sqlalchemy.engine.Connection):
        self._conn = conn

    def delete_bar_by_id(self, *, id: int) -> int:
        result = self._conn.execute(sqlalchemy.text(DELETE_BAR_BY_ID), {"p1": id})
        return result.rowcount

    def delete_bar_by_id_and_name(self, arg: DeleteBarByIDAndNameParams) -> int:
        result = self._conn.execute(sqlalchemy.text(DELETE_BAR_BY_ID_AND_NAME), {"p1": arg.id, "p2": arg.name})
        return result.rowcount


class AsyncQuerier:
    def __init__(self, conn: sqlalchemy.ext.asyncio.AsyncConnection):
        self._conn = conn

    async def delete_bar_by_id(self, *, id: int) -> int:
        result = await self._conn.execute(sqlalchemy.text(DELETE_BAR_BY_ID), {"p1": id})
        return result.rowcount

    async def delete_bar_by_id_and_name(self, arg: DeleteBarByIDAndNameParams) -> int:
        result = await self._conn.execute(sqlalchemy.text(DELETE_BAR_BY_ID_AND_NAME), {"p1": arg.id, "p2": arg.name})
        return result.rowcount
