from __future__ import annotations
from enum import Enum

from pydantic import BaseModel


class Bmi(str, Enum):
    UNDERWEIGHT = "underweight"
    NORMAL = "normal"
    OVERWEIGHT = "overweight"
    VERY_OVERWEIGHT = "very_overweight"

    @classmethod
    def from_bmi(cls, bmi: float) -> Bmi:
        """Constructor given a numeric bmi"""
        if bmi < 18.5:
            return Bmi.UNDERWEIGHT
        elif bmi < 24.9:
            return Bmi.NORMAL
        elif bmi < 29.9:
            return Bmi.OVERWEIGHT
        return Bmi.VERY_OVERWEIGHT


class Output(BaseModel):
    id: int
    name: str
    age: int
    bmi: Bmi

